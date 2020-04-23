package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"github.com/minskylab/collecta/auth"
	"github.com/minskylab/collecta/db"
)

type Resolver struct {
	Auth *auth.Auth
	DB *db.DB
}
