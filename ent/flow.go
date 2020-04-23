// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/minskylab/collecta/ent/flow"
)

// Flow is the model entity for the Flow schema.
type Flow struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// State holds the value of the "state" field.
	State uuid.UUID `json:"state,omitempty"`
	// StateTable holds the value of the "stateTable" field.
	StateTable string `json:"stateTable,omitempty"`
	// PastState holds the value of the "pastState" field.
	PastState uuid.UUID `json:"pastState,omitempty"`
	// Inputs holds the value of the "inputs" field.
	Inputs []string `json:"inputs,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FlowQuery when eager-loading is set.
	Edges FlowEdges `json:"edges"`
}

// FlowEdges holds the relations/edges for other nodes in the graph.
type FlowEdges struct {
	// Questions holds the value of the questions edge.
	Questions []*Question
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// QuestionsOrErr returns the Questions value or an error if the edge
// was not loaded in eager-loading.
func (e FlowEdges) QuestionsOrErr() ([]*Question, error) {
	if e.loadedTypes[0] {
		return e.Questions, nil
	}
	return nil, &NotLoadedError{edge: "questions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Flow) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&uuid.UUID{},      // state
		&sql.NullString{}, // stateTable
		&uuid.UUID{},      // pastState
		&[]byte{},         // inputs
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Flow fields.
func (f *Flow) assignValues(values ...interface{}) error {
	if m, n := len(values), len(flow.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		f.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field state", values[0])
	} else if value != nil {
		f.State = *value
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field stateTable", values[1])
	} else if value.Valid {
		f.StateTable = value.String
	}
	if value, ok := values[2].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field pastState", values[2])
	} else if value != nil {
		f.PastState = *value
	}

	if value, ok := values[3].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field inputs", values[3])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &f.Inputs); err != nil {
			return fmt.Errorf("unmarshal field inputs: %v", err)
		}
	}
	return nil
}

// QueryQuestions queries the questions edge of the Flow.
func (f *Flow) QueryQuestions() *QuestionQuery {
	return (&FlowClient{config: f.config}).QueryQuestions(f)
}

// Update returns a builder for updating this Flow.
// Note that, you need to call Flow.Unwrap() before calling this method, if this Flow
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Flow) Update() *FlowUpdateOne {
	return (&FlowClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (f *Flow) Unwrap() *Flow {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Flow is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Flow) String() string {
	var builder strings.Builder
	builder.WriteString("Flow(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteString(", state=")
	builder.WriteString(fmt.Sprintf("%v", f.State))
	builder.WriteString(", stateTable=")
	builder.WriteString(f.StateTable)
	builder.WriteString(", pastState=")
	builder.WriteString(fmt.Sprintf("%v", f.PastState))
	builder.WriteString(", inputs=")
	builder.WriteString(fmt.Sprintf("%v", f.Inputs))
	builder.WriteByte(')')
	return builder.String()
}

// Flows is a parsable slice of Flow.
type Flows []*Flow

func (f Flows) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}
