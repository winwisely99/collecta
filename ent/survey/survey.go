// Code generated by entc, DO NOT EDIT.

package survey

import (
	"time"
)

const (
	// Label holds the string label denoting the survey type in the database.
	Label = "survey"
	// FieldID holds the string denoting the id field in the database.
	FieldID              = "id"               // FieldTags holds the string denoting the tags vertex property in the database.
	FieldTags            = "tags"             // FieldLastInteraction holds the string denoting the lastinteraction vertex property in the database.
	FieldLastInteraction = "last_interaction" // FieldDueDate holds the string denoting the duedate vertex property in the database.
	FieldDueDate         = "due_date"         // FieldTitle holds the string denoting the title vertex property in the database.
	FieldTitle           = "title"            // FieldDescription holds the string denoting the description vertex property in the database.
	FieldDescription     = "description"      // FieldMetadata holds the string denoting the metadata vertex property in the database.
	FieldMetadata        = "metadata"         // FieldDone holds the string denoting the done vertex property in the database.
	FieldDone            = "done"             // FieldIsPublic holds the string denoting the ispublic vertex property in the database.
	FieldIsPublic        = "is_public"

	// EdgeFlow holds the string denoting the flow edge name in mutations.
	EdgeFlow = "flow"
	// EdgeFor holds the string denoting the for edge name in mutations.
	EdgeFor = "for"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"

	// Table holds the table name of the survey in the database.
	Table = "surveys"
	// FlowTable is the table the holds the flow relation/edge.
	FlowTable = "flows"
	// FlowInverseTable is the table name for the Flow entity.
	// It exists in this package in order to avoid circular dependency with the "flow" package.
	FlowInverseTable = "flows"
	// FlowColumn is the table column denoting the flow relation/edge.
	FlowColumn = "survey_flow"
	// ForTable is the table the holds the for relation/edge.
	ForTable = "surveys"
	// ForInverseTable is the table name for the Person entity.
	// It exists in this package in order to avoid circular dependency with the "person" package.
	ForInverseTable = "persons"
	// ForColumn is the table column denoting the for relation/edge.
	ForColumn = "person_surveys"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "surveys"
	// OwnerInverseTable is the table name for the Domain entity.
	// It exists in this package in order to avoid circular dependency with the "domain" package.
	OwnerInverseTable = "domains"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "domain_surveys"
)

// Columns holds all SQL columns for survey fields.
var Columns = []string{
	FieldID,
	FieldTags,
	FieldLastInteraction,
	FieldDueDate,
	FieldTitle,
	FieldDescription,
	FieldMetadata,
	FieldDone,
	FieldIsPublic,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Survey type.
var ForeignKeys = []string{
	"domain_surveys",
	"person_surveys",
}

var (
	// DefaultDueDate holds the default value on creation for the dueDate field.
	DefaultDueDate func() time.Time
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DefaultDone holds the default value on creation for the done field.
	DefaultDone bool
	// DefaultIsPublic holds the default value on creation for the isPublic field.
	DefaultIsPublic bool
)
