// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
	"github.com/minskylab/collecta/ent/answer"
	"github.com/minskylab/collecta/ent/question"
	"github.com/rs/xid"
)

// AnswerCreate is the builder for creating a Answer entity.
type AnswerCreate struct {
	config
	id        *xid.ID
	at        *time.Time
	responses *[]string
	valid     *bool
	question  map[uuid.UUID]struct{}
}

// SetAt sets the at field.
func (ac *AnswerCreate) SetAt(t time.Time) *AnswerCreate {
	ac.at = &t
	return ac
}

// SetResponses sets the responses field.
func (ac *AnswerCreate) SetResponses(s []string) *AnswerCreate {
	ac.responses = &s
	return ac
}

// SetValid sets the valid field.
func (ac *AnswerCreate) SetValid(b bool) *AnswerCreate {
	ac.valid = &b
	return ac
}

// SetNillableValid sets the valid field if the given value is not nil.
func (ac *AnswerCreate) SetNillableValid(b *bool) *AnswerCreate {
	if b != nil {
		ac.SetValid(*b)
	}
	return ac
}

// SetID sets the id field.
func (ac *AnswerCreate) SetID(x xid.ID) *AnswerCreate {
	ac.id = &x
	return ac
}

// SetQuestionID sets the question edge to Question by id.
func (ac *AnswerCreate) SetQuestionID(id uuid.UUID) *AnswerCreate {
	if ac.question == nil {
		ac.question = make(map[uuid.UUID]struct{})
	}
	ac.question[id] = struct{}{}
	return ac
}

// SetQuestion sets the question edge to Question.
func (ac *AnswerCreate) SetQuestion(q *Question) *AnswerCreate {
	return ac.SetQuestionID(q.ID)
}

// Save creates the Answer in the database.
func (ac *AnswerCreate) Save(ctx context.Context) (*Answer, error) {
	if ac.at == nil {
		return nil, errors.New("ent: missing required field \"at\"")
	}
	if ac.responses == nil {
		return nil, errors.New("ent: missing required field \"responses\"")
	}
	if len(ac.question) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"question\"")
	}
	if ac.question == nil {
		return nil, errors.New("ent: missing required edge \"question\"")
	}
	return ac.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AnswerCreate) SaveX(ctx context.Context) *Answer {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ac *AnswerCreate) sqlSave(ctx context.Context) (*Answer, error) {
	var (
		a     = &Answer{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: answer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: answer.FieldID,
			},
		}
	)
	if value := ac.id; value != nil {
		a.ID = *value
		_spec.ID.Value = *value
	}
	if value := ac.at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: answer.FieldAt,
		})
		a.At = *value
	}
	if value := ac.responses; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  *value,
			Column: answer.FieldResponses,
		})
		a.Responses = *value
	}
	if value := ac.valid; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: answer.FieldValid,
		})
		a.Valid = *value
	}
	if nodes := ac.question; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   answer.QuestionTable,
			Columns: []string{answer.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: question.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return a, nil
}