package auth

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func (collectaAuth *CollectaAuth) startGoogleAuth(callbackURL string) {
	clientID := viper.GetString("google.clientID")
	secretKey := viper.GetString("google.secret")

	goth.UseProviders(
		google.New(clientID, secretKey, callbackURL),
	)

	r := collectaAuth.mainRouter.Group("/auth")

	r.GET("/google", func(c *gin.Context) {
		modQuery := c.Request.URL.Query()
		modQuery.Add("provider", "google")
		c.Request.URL.RawQuery = modQuery.Encode()
		if gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
			log.Info(gothUser)
			spew.Fdump(c.Writer, gothUser)
		} else {
			log.WithError(err).Info("trying new google auth")
			gothic.BeginAuthHandler(c.Writer, c.Request)
		}
	})

	r.GET("/google/callback", func(c *gin.Context) {
		modQuery := c.Request.URL.Query()
		modQuery.Add("provider", "google")
		c.Request.URL.RawQuery = modQuery.Encode()

		user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
		if err != nil {
			log.WithError(err).Info("fail at complete google user auth")
			return
		}
		// TODO: Verify the correctness of the google auth token
		log.Info(user)
		jwtToken, err := collectaAuth.ingressWithGoogle(c, user)
		if err != nil {
			log.WithError(err).Info("error at try to ingress with user")
			return
		}

		log.Info(jwtToken)
		c.Header("Collecta-Token", jwtToken)

		redirect, err := collectaAuth.matchGoogleUserWithCollectaDomain(c, user)
		if err != nil {
			log.WithError(err).Info("error at try to create a redirect link")
			c.Redirect(http.StatusPermanentRedirect, "/")
			return
		}

		c.Redirect(http.StatusPermanentRedirect, redirect)
	})

}
