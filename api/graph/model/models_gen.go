// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Account struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Sub      string `json:"sub"`
	RemoteID string `json:"remoteID"`
	Owner    *User  `json:"owner"`
}

type Answer struct {
	ID        string    `json:"id"`
	At        time.Time `json:"at"`
	Responses []string  `json:"responses"`
	Valid     bool      `json:"valid"`
	Question  *Question `json:"question"`
}

type Contact struct {
	ID          string  `json:"id"`
	Name        *string `json:"name"`
	Value       string  `json:"value"`
	Kind        string  `json:"kind"`
	Principal   bool    `json:"principal"`
	Validated   bool    `json:"validated"`
	FromAccount bool    `json:"fromAccount"`
	Owner       *User   `json:"owner"`
}

type Domain struct {
	ID      string    `json:"id"`
	Tags    []string  `json:"tags"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Domain  string    `json:"domain"`
	Surveys []*Survey `json:"surveys"`
	Users   []*User   `json:"users"`
}

type Flow struct {
	ID         string      `json:"id"`
	State      string      `json:"state"`
	StateTable string      `json:"stateTable"`
	Inputs     []string    `json:"inputs"`
	Questions  []*Question `json:"questions"`
}

type Input struct {
	ID       string                 `json:"id"`
	Kind     string                 `json:"kind"`
	Multiple bool                   `json:"multiple"`
	Defaults []*string              `json:"defaults"`
	Options  map[string]interface{} `json:"options"`
	Question *Question              `json:"question"`
}

type Question struct {
	ID          string                 `json:"id"`
	Hash        string                 `json:"hash"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Anonymous   bool                   `json:"anonymous"`
	Metadata    map[string]interface{} `json:"metadata"`
	Answers     []*Answer              `json:"answers"`
	Input       *Input                 `json:"input"`
	Flow        *Flow                  `json:"flow"`
}

type Survey struct {
	ID              string                 `json:"id"`
	Tags            []string               `json:"tags"`
	LastInteraction time.Time              `json:"lastInteraction"`
	DueDate         time.Time              `json:"dueDate"`
	Title           string                 `json:"title"`
	Description     string                 `json:"description"`
	Metadata        map[string]interface{} `json:"metadata"`
	Done            bool                   `json:"done"`
	Flow            *Flow                  `json:"flow"`
	For             *User                  `json:"for"`
	Owner           *Domain                `json:"owner"`
}

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Username     string    `json:"username"`
	LastActivity time.Time `json:"lastActivity"`
	Accounts     *Account  `json:"accounts"`
	Contacts     *Contact  `json:"contacts"`
	Survey       *Survey   `json:"survey"`
	Domain       *Domain   `json:"domain"`
}
