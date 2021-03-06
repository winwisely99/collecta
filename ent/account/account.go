// Code generated by entc, DO NOT EDIT.

package account

import (
	"fmt"
)

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID       = "id"        // FieldType holds the string denoting the type vertex property in the database.
	FieldType     = "type"      // FieldSub holds the string denoting the sub vertex property in the database.
	FieldSub      = "sub"       // FieldRemoteID holds the string denoting the remoteid vertex property in the database.
	FieldRemoteID = "remote_id" // FieldSecret holds the string denoting the secret vertex property in the database.
	FieldSecret   = "secret"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"

	// Table holds the table name of the account in the database.
	Table = "accounts"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "accounts"
	// OwnerInverseTable is the table name for the Person entity.
	// It exists in this package in order to avoid circular dependency with the "person" package.
	OwnerInverseTable = "persons"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "person_accounts"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldSub,
	FieldRemoteID,
	FieldSecret,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Account type.
var ForeignKeys = []string{
	"person_accounts",
}

var (
	// SubValidator is a validator for the "sub" field. It is called by the builders before save.
	SubValidator func(string) error
)

// Type defines the type for the type enum field.
type Type string

// Type values.
const (
	TypeGoogle    Type = "Google"
	TypeAnonymous Type = "Anonymous"
	TypeEmail     Type = "Email"
)

func (s Type) String() string {
	return string(s)
}

// TypeValidator is a validator for the "_type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeGoogle, TypeAnonymous, TypeEmail:
		return nil
	default:
		return fmt.Errorf("account: invalid enum value for type field: %q", _type)
	}
}
