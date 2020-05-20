// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/minskylab/collecta/ent/flow"
	"github.com/minskylab/collecta/ent/input"
	"github.com/minskylab/collecta/ent/question"
	"github.com/minskylab/collecta/uuid"
)

// Question is the model entity for the Question schema.
type Question struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Hash holds the value of the "hash" field.
	Hash string `json:"hash,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Validator holds the value of the "validator" field.
	Validator string `json:"validator,omitempty"`
	// Anonymous holds the value of the "anonymous" field.
	Anonymous bool `json:"anonymous,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the QuestionQuery when eager-loading is set.
	Edges          QuestionEdges `json:"edges"`
	flow_questions *uuid.UUID
}

// QuestionEdges holds the relations/edges for other nodes in the graph.
type QuestionEdges struct {
	// Answers holds the value of the answers edge.
	Answers []*Answer
	// Input holds the value of the input edge.
	Input *Input
	// Flow holds the value of the flow edge.
	Flow *Flow
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// AnswersOrErr returns the Answers value or an error if the edge
// was not loaded in eager-loading.
func (e QuestionEdges) AnswersOrErr() ([]*Answer, error) {
	if e.loadedTypes[0] {
		return e.Answers, nil
	}
	return nil, &NotLoadedError{edge: "answers"}
}

// InputOrErr returns the Input value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e QuestionEdges) InputOrErr() (*Input, error) {
	if e.loadedTypes[1] {
		if e.Input == nil {
			// The edge input was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: input.Label}
		}
		return e.Input, nil
	}
	return nil, &NotLoadedError{edge: "input"}
}

// FlowOrErr returns the Flow value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e QuestionEdges) FlowOrErr() (*Flow, error) {
	if e.loadedTypes[2] {
		if e.Flow == nil {
			// The edge flow was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: flow.Label}
		}
		return e.Flow, nil
	}
	return nil, &NotLoadedError{edge: "flow"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Question) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&sql.NullString{}, // hash
		&sql.NullString{}, // title
		&sql.NullString{}, // description
		&[]byte{},         // metadata
		&sql.NullString{}, // validator
		&sql.NullBool{},   // anonymous
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Question) fkValues() []interface{} {
	return []interface{}{
		&uuid.UUID{}, // flow_questions
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Question fields.
func (q *Question) assignValues(values ...interface{}) error {
	if m, n := len(values), len(question.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		q.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field hash", values[0])
	} else if value.Valid {
		q.Hash = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field title", values[1])
	} else if value.Valid {
		q.Title = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field description", values[2])
	} else if value.Valid {
		q.Description = value.String
	}

	if value, ok := values[3].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field metadata", values[3])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &q.Metadata); err != nil {
			return fmt.Errorf("unmarshal field metadata: %v", err)
		}
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field validator", values[4])
	} else if value.Valid {
		q.Validator = value.String
	}
	if value, ok := values[5].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field anonymous", values[5])
	} else if value.Valid {
		q.Anonymous = value.Bool
	}
	values = values[6:]
	if len(values) == len(question.ForeignKeys) {
		if value, ok := values[0].(*uuid.UUID); !ok {
			return fmt.Errorf("unexpected type %T for field flow_questions", values[0])
		} else if value != nil {
			q.flow_questions = value
		}
	}
	return nil
}

// QueryAnswers queries the answers edge of the Question.
func (q *Question) QueryAnswers() *AnswerQuery {
	return (&QuestionClient{config: q.config}).QueryAnswers(q)
}

// QueryInput queries the input edge of the Question.
func (q *Question) QueryInput() *InputQuery {
	return (&QuestionClient{config: q.config}).QueryInput(q)
}

// QueryFlow queries the flow edge of the Question.
func (q *Question) QueryFlow() *FlowQuery {
	return (&QuestionClient{config: q.config}).QueryFlow(q)
}

// Update returns a builder for updating this Question.
// Note that, you need to call Question.Unwrap() before calling this method, if this Question
// was returned from a transaction, and the transaction was committed or rolled back.
func (q *Question) Update() *QuestionUpdateOne {
	return (&QuestionClient{config: q.config}).UpdateOne(q)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (q *Question) Unwrap() *Question {
	tx, ok := q.config.driver.(*txDriver)
	if !ok {
		panic("ent: Question is not a transactional entity")
	}
	q.config.driver = tx.drv
	return q
}

// String implements the fmt.Stringer.
func (q *Question) String() string {
	var builder strings.Builder
	builder.WriteString("Question(")
	builder.WriteString(fmt.Sprintf("id=%v", q.ID))
	builder.WriteString(", hash=")
	builder.WriteString(q.Hash)
	builder.WriteString(", title=")
	builder.WriteString(q.Title)
	builder.WriteString(", description=")
	builder.WriteString(q.Description)
	builder.WriteString(", metadata=")
	builder.WriteString(fmt.Sprintf("%v", q.Metadata))
	builder.WriteString(", validator=")
	builder.WriteString(q.Validator)
	builder.WriteString(", anonymous=")
	builder.WriteString(fmt.Sprintf("%v", q.Anonymous))
	builder.WriteByte(')')
	return builder.String()
}

// Questions is a parsable slice of Question.
type Questions []*Question

func (q Questions) config(cfg config) {
	for _i := range q {
		q[_i].config = cfg
	}
}
