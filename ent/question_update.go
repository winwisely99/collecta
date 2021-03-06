// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/minskylab/collecta/ent/answer"
	"github.com/minskylab/collecta/ent/flow"
	"github.com/minskylab/collecta/ent/input"
	"github.com/minskylab/collecta/ent/predicate"
	"github.com/minskylab/collecta/ent/question"
	"github.com/minskylab/collecta/uuid"
)

// QuestionUpdate is the builder for updating Question entities.
type QuestionUpdate struct {
	config
	hooks      []Hook
	mutation   *QuestionMutation
	predicates []predicate.Question
}

// Where adds a new predicate for the builder.
func (qu *QuestionUpdate) Where(ps ...predicate.Question) *QuestionUpdate {
	qu.predicates = append(qu.predicates, ps...)
	return qu
}

// SetHash sets the hash field.
func (qu *QuestionUpdate) SetHash(s string) *QuestionUpdate {
	qu.mutation.SetHash(s)
	return qu
}

// SetTitle sets the title field.
func (qu *QuestionUpdate) SetTitle(s string) *QuestionUpdate {
	qu.mutation.SetTitle(s)
	return qu
}

// SetDescription sets the description field.
func (qu *QuestionUpdate) SetDescription(s string) *QuestionUpdate {
	qu.mutation.SetDescription(s)
	return qu
}

// SetMetadata sets the metadata field.
func (qu *QuestionUpdate) SetMetadata(m map[string]interface{}) *QuestionUpdate {
	qu.mutation.SetMetadata(m)
	return qu
}

// ClearMetadata clears the value of metadata.
func (qu *QuestionUpdate) ClearMetadata() *QuestionUpdate {
	qu.mutation.ClearMetadata()
	return qu
}

// SetAnonymous sets the anonymous field.
func (qu *QuestionUpdate) SetAnonymous(b bool) *QuestionUpdate {
	qu.mutation.SetAnonymous(b)
	return qu
}

// SetNillableAnonymous sets the anonymous field if the given value is not nil.
func (qu *QuestionUpdate) SetNillableAnonymous(b *bool) *QuestionUpdate {
	if b != nil {
		qu.SetAnonymous(*b)
	}
	return qu
}

// AddAnswerIDs adds the answers edge to Answer by ids.
func (qu *QuestionUpdate) AddAnswerIDs(ids ...uuid.UUID) *QuestionUpdate {
	qu.mutation.AddAnswerIDs(ids...)
	return qu
}

// AddAnswers adds the answers edges to Answer.
func (qu *QuestionUpdate) AddAnswers(a ...*Answer) *QuestionUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return qu.AddAnswerIDs(ids...)
}

// SetInputID sets the input edge to Input by id.
func (qu *QuestionUpdate) SetInputID(id uuid.UUID) *QuestionUpdate {
	qu.mutation.SetInputID(id)
	return qu
}

// SetNillableInputID sets the input edge to Input by id if the given value is not nil.
func (qu *QuestionUpdate) SetNillableInputID(id *uuid.UUID) *QuestionUpdate {
	if id != nil {
		qu = qu.SetInputID(*id)
	}
	return qu
}

// SetInput sets the input edge to Input.
func (qu *QuestionUpdate) SetInput(i *Input) *QuestionUpdate {
	return qu.SetInputID(i.ID)
}

// SetFlowID sets the flow edge to Flow by id.
func (qu *QuestionUpdate) SetFlowID(id uuid.UUID) *QuestionUpdate {
	qu.mutation.SetFlowID(id)
	return qu
}

// SetNillableFlowID sets the flow edge to Flow by id if the given value is not nil.
func (qu *QuestionUpdate) SetNillableFlowID(id *uuid.UUID) *QuestionUpdate {
	if id != nil {
		qu = qu.SetFlowID(*id)
	}
	return qu
}

// SetFlow sets the flow edge to Flow.
func (qu *QuestionUpdate) SetFlow(f *Flow) *QuestionUpdate {
	return qu.SetFlowID(f.ID)
}

// RemoveAnswerIDs removes the answers edge to Answer by ids.
func (qu *QuestionUpdate) RemoveAnswerIDs(ids ...uuid.UUID) *QuestionUpdate {
	qu.mutation.RemoveAnswerIDs(ids...)
	return qu
}

// RemoveAnswers removes answers edges to Answer.
func (qu *QuestionUpdate) RemoveAnswers(a ...*Answer) *QuestionUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return qu.RemoveAnswerIDs(ids...)
}

// ClearInput clears the input edge to Input.
func (qu *QuestionUpdate) ClearInput() *QuestionUpdate {
	qu.mutation.ClearInput()
	return qu
}

// ClearFlow clears the flow edge to Flow.
func (qu *QuestionUpdate) ClearFlow() *QuestionUpdate {
	qu.mutation.ClearFlow()
	return qu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (qu *QuestionUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := qu.mutation.Title(); ok {
		if err := question.TitleValidator(v); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"title\": %v", err)
		}
	}

	var (
		err      error
		affected int
	)
	if len(qu.hooks) == 0 {
		affected, err = qu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QuestionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			qu.mutation = mutation
			affected, err = qu.sqlSave(ctx)
			return affected, err
		})
		for i := len(qu.hooks) - 1; i >= 0; i-- {
			mut = qu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (qu *QuestionUpdate) SaveX(ctx context.Context) int {
	affected, err := qu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (qu *QuestionUpdate) Exec(ctx context.Context) error {
	_, err := qu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qu *QuestionUpdate) ExecX(ctx context.Context) {
	if err := qu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (qu *QuestionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   question.Table,
			Columns: question.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: question.FieldID,
			},
		},
	}
	if ps := qu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := qu.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: question.FieldHash,
		})
	}
	if value, ok := qu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: question.FieldTitle,
		})
	}
	if value, ok := qu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: question.FieldDescription,
		})
	}
	if value, ok := qu.mutation.Metadata(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: question.FieldMetadata,
		})
	}
	if qu.mutation.MetadataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: question.FieldMetadata,
		})
	}
	if qu.mutation.ValidatorCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: question.FieldValidator,
		})
	}
	if value, ok := qu.mutation.Anonymous(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: question.FieldAnonymous,
		})
	}
	if nodes := qu.mutation.RemovedAnswersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.AnswersTable,
			Columns: []string{question.AnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: answer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.AnswersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.AnswersTable,
			Columns: []string{question.AnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: answer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qu.mutation.InputCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   question.InputTable,
			Columns: []string{question.InputColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: input.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.InputIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   question.InputTable,
			Columns: []string{question.InputColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: input.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qu.mutation.FlowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   question.FlowTable,
			Columns: []string{question.FlowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: flow.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.FlowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   question.FlowTable,
			Columns: []string{question.FlowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: flow.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, qu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{question.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// QuestionUpdateOne is the builder for updating a single Question entity.
type QuestionUpdateOne struct {
	config
	hooks    []Hook
	mutation *QuestionMutation
}

// SetHash sets the hash field.
func (quo *QuestionUpdateOne) SetHash(s string) *QuestionUpdateOne {
	quo.mutation.SetHash(s)
	return quo
}

// SetTitle sets the title field.
func (quo *QuestionUpdateOne) SetTitle(s string) *QuestionUpdateOne {
	quo.mutation.SetTitle(s)
	return quo
}

// SetDescription sets the description field.
func (quo *QuestionUpdateOne) SetDescription(s string) *QuestionUpdateOne {
	quo.mutation.SetDescription(s)
	return quo
}

// SetMetadata sets the metadata field.
func (quo *QuestionUpdateOne) SetMetadata(m map[string]interface{}) *QuestionUpdateOne {
	quo.mutation.SetMetadata(m)
	return quo
}

// ClearMetadata clears the value of metadata.
func (quo *QuestionUpdateOne) ClearMetadata() *QuestionUpdateOne {
	quo.mutation.ClearMetadata()
	return quo
}

// SetAnonymous sets the anonymous field.
func (quo *QuestionUpdateOne) SetAnonymous(b bool) *QuestionUpdateOne {
	quo.mutation.SetAnonymous(b)
	return quo
}

// SetNillableAnonymous sets the anonymous field if the given value is not nil.
func (quo *QuestionUpdateOne) SetNillableAnonymous(b *bool) *QuestionUpdateOne {
	if b != nil {
		quo.SetAnonymous(*b)
	}
	return quo
}

// AddAnswerIDs adds the answers edge to Answer by ids.
func (quo *QuestionUpdateOne) AddAnswerIDs(ids ...uuid.UUID) *QuestionUpdateOne {
	quo.mutation.AddAnswerIDs(ids...)
	return quo
}

// AddAnswers adds the answers edges to Answer.
func (quo *QuestionUpdateOne) AddAnswers(a ...*Answer) *QuestionUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return quo.AddAnswerIDs(ids...)
}

// SetInputID sets the input edge to Input by id.
func (quo *QuestionUpdateOne) SetInputID(id uuid.UUID) *QuestionUpdateOne {
	quo.mutation.SetInputID(id)
	return quo
}

// SetNillableInputID sets the input edge to Input by id if the given value is not nil.
func (quo *QuestionUpdateOne) SetNillableInputID(id *uuid.UUID) *QuestionUpdateOne {
	if id != nil {
		quo = quo.SetInputID(*id)
	}
	return quo
}

// SetInput sets the input edge to Input.
func (quo *QuestionUpdateOne) SetInput(i *Input) *QuestionUpdateOne {
	return quo.SetInputID(i.ID)
}

// SetFlowID sets the flow edge to Flow by id.
func (quo *QuestionUpdateOne) SetFlowID(id uuid.UUID) *QuestionUpdateOne {
	quo.mutation.SetFlowID(id)
	return quo
}

// SetNillableFlowID sets the flow edge to Flow by id if the given value is not nil.
func (quo *QuestionUpdateOne) SetNillableFlowID(id *uuid.UUID) *QuestionUpdateOne {
	if id != nil {
		quo = quo.SetFlowID(*id)
	}
	return quo
}

// SetFlow sets the flow edge to Flow.
func (quo *QuestionUpdateOne) SetFlow(f *Flow) *QuestionUpdateOne {
	return quo.SetFlowID(f.ID)
}

// RemoveAnswerIDs removes the answers edge to Answer by ids.
func (quo *QuestionUpdateOne) RemoveAnswerIDs(ids ...uuid.UUID) *QuestionUpdateOne {
	quo.mutation.RemoveAnswerIDs(ids...)
	return quo
}

// RemoveAnswers removes answers edges to Answer.
func (quo *QuestionUpdateOne) RemoveAnswers(a ...*Answer) *QuestionUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return quo.RemoveAnswerIDs(ids...)
}

// ClearInput clears the input edge to Input.
func (quo *QuestionUpdateOne) ClearInput() *QuestionUpdateOne {
	quo.mutation.ClearInput()
	return quo
}

// ClearFlow clears the flow edge to Flow.
func (quo *QuestionUpdateOne) ClearFlow() *QuestionUpdateOne {
	quo.mutation.ClearFlow()
	return quo
}

// Save executes the query and returns the updated entity.
func (quo *QuestionUpdateOne) Save(ctx context.Context) (*Question, error) {
	if v, ok := quo.mutation.Title(); ok {
		if err := question.TitleValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"title\": %v", err)
		}
	}

	var (
		err  error
		node *Question
	)
	if len(quo.hooks) == 0 {
		node, err = quo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QuestionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			quo.mutation = mutation
			node, err = quo.sqlSave(ctx)
			return node, err
		})
		for i := len(quo.hooks) - 1; i >= 0; i-- {
			mut = quo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, quo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (quo *QuestionUpdateOne) SaveX(ctx context.Context) *Question {
	q, err := quo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return q
}

// Exec executes the query on the entity.
func (quo *QuestionUpdateOne) Exec(ctx context.Context) error {
	_, err := quo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (quo *QuestionUpdateOne) ExecX(ctx context.Context) {
	if err := quo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (quo *QuestionUpdateOne) sqlSave(ctx context.Context) (q *Question, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   question.Table,
			Columns: question.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: question.FieldID,
			},
		},
	}
	id, ok := quo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing Question.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := quo.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: question.FieldHash,
		})
	}
	if value, ok := quo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: question.FieldTitle,
		})
	}
	if value, ok := quo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: question.FieldDescription,
		})
	}
	if value, ok := quo.mutation.Metadata(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: question.FieldMetadata,
		})
	}
	if quo.mutation.MetadataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: question.FieldMetadata,
		})
	}
	if quo.mutation.ValidatorCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: question.FieldValidator,
		})
	}
	if value, ok := quo.mutation.Anonymous(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: question.FieldAnonymous,
		})
	}
	if nodes := quo.mutation.RemovedAnswersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.AnswersTable,
			Columns: []string{question.AnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: answer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.AnswersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.AnswersTable,
			Columns: []string{question.AnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: answer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if quo.mutation.InputCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   question.InputTable,
			Columns: []string{question.InputColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: input.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.InputIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   question.InputTable,
			Columns: []string{question.InputColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: input.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if quo.mutation.FlowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   question.FlowTable,
			Columns: []string{question.FlowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: flow.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.FlowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   question.FlowTable,
			Columns: []string{question.FlowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: flow.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	q = &Question{config: quo.config}
	_spec.Assign = q.assignValues
	_spec.ScanValues = q.scanValues()
	if err = sqlgraph.UpdateNode(ctx, quo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{question.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return q, nil
}
