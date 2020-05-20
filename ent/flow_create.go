// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/minskylab/collecta/ent/flow"
	"github.com/minskylab/collecta/ent/question"
	"github.com/minskylab/collecta/ent/survey"
	"github.com/minskylab/collecta/uuid"
)

// FlowCreate is the builder for creating a Flow entity.
type FlowCreate struct {
	config
	mutation *FlowMutation
	hooks    []Hook
}

// SetState sets the state field.
func (fc *FlowCreate) SetState(u uuid.UUID) *FlowCreate {
	fc.mutation.SetState(u)
	return fc
}

// SetStateTable sets the stateTable field.
func (fc *FlowCreate) SetStateTable(s string) *FlowCreate {
	fc.mutation.SetStateTable(s)
	return fc
}

// SetInitialState sets the initialState field.
func (fc *FlowCreate) SetInitialState(u uuid.UUID) *FlowCreate {
	fc.mutation.SetInitialState(u)
	return fc
}

// SetTerminationState sets the terminationState field.
func (fc *FlowCreate) SetTerminationState(u uuid.UUID) *FlowCreate {
	fc.mutation.SetTerminationState(u)
	return fc
}

// SetPastState sets the pastState field.
func (fc *FlowCreate) SetPastState(u uuid.UUID) *FlowCreate {
	fc.mutation.SetPastState(u)
	return fc
}

// SetInputs sets the inputs field.
func (fc *FlowCreate) SetInputs(s []string) *FlowCreate {
	fc.mutation.SetInputs(s)
	return fc
}

// SetID sets the id field.
func (fc *FlowCreate) SetID(u uuid.UUID) *FlowCreate {
	fc.mutation.SetID(u)
	return fc
}

// SetSurveyID sets the survey edge to Survey by id.
func (fc *FlowCreate) SetSurveyID(id uuid.UUID) *FlowCreate {
	fc.mutation.SetSurveyID(id)
	return fc
}

// SetNillableSurveyID sets the survey edge to Survey by id if the given value is not nil.
func (fc *FlowCreate) SetNillableSurveyID(id *uuid.UUID) *FlowCreate {
	if id != nil {
		fc = fc.SetSurveyID(*id)
	}
	return fc
}

// SetSurvey sets the survey edge to Survey.
func (fc *FlowCreate) SetSurvey(s *Survey) *FlowCreate {
	return fc.SetSurveyID(s.ID)
}

// AddQuestionIDs adds the questions edge to Question by ids.
func (fc *FlowCreate) AddQuestionIDs(ids ...uuid.UUID) *FlowCreate {
	fc.mutation.AddQuestionIDs(ids...)
	return fc
}

// AddQuestions adds the questions edges to Question.
func (fc *FlowCreate) AddQuestions(q ...*Question) *FlowCreate {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return fc.AddQuestionIDs(ids...)
}

// Save creates the Flow in the database.
func (fc *FlowCreate) Save(ctx context.Context) (*Flow, error) {
	if _, ok := fc.mutation.State(); !ok {
		return nil, errors.New("ent: missing required field \"state\"")
	}
	if _, ok := fc.mutation.StateTable(); !ok {
		return nil, errors.New("ent: missing required field \"stateTable\"")
	}
	if v, ok := fc.mutation.StateTable(); ok {
		if err := flow.StateTableValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"stateTable\": %v", err)
		}
	}
	if _, ok := fc.mutation.InitialState(); !ok {
		return nil, errors.New("ent: missing required field \"initialState\"")
	}
	if _, ok := fc.mutation.TerminationState(); !ok {
		return nil, errors.New("ent: missing required field \"terminationState\"")
	}
	if len(fc.mutation.QuestionsIDs()) == 0 {
		return nil, errors.New("ent: missing required edge \"questions\"")
	}
	var (
		err  error
		node *Flow
	)
	if len(fc.hooks) == 0 {
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fc.mutation = mutation
			node, err = fc.sqlSave(ctx)
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			mut = fc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FlowCreate) SaveX(ctx context.Context) *Flow {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fc *FlowCreate) sqlSave(ctx context.Context) (*Flow, error) {
	var (
		f     = &Flow{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: flow.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: flow.FieldID,
			},
		}
	)
	if id, ok := fc.mutation.ID(); ok {
		f.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fc.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: flow.FieldState,
		})
		f.State = value
	}
	if value, ok := fc.mutation.StateTable(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flow.FieldStateTable,
		})
		f.StateTable = value
	}
	if value, ok := fc.mutation.InitialState(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: flow.FieldInitialState,
		})
		f.InitialState = value
	}
	if value, ok := fc.mutation.TerminationState(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: flow.FieldTerminationState,
		})
		f.TerminationState = value
	}
	if value, ok := fc.mutation.PastState(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: flow.FieldPastState,
		})
		f.PastState = value
	}
	if value, ok := fc.mutation.Inputs(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: flow.FieldInputs,
		})
		f.Inputs = value
	}
	if nodes := fc.mutation.SurveyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   flow.SurveyTable,
			Columns: []string{flow.SurveyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: survey.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.QuestionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flow.QuestionsTable,
			Columns: []string{flow.QuestionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: question.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return f, nil
}
