// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/minskylab/collecta/ent/domain"
	"github.com/minskylab/collecta/ent/flow"
	"github.com/minskylab/collecta/ent/person"
	"github.com/minskylab/collecta/ent/predicate"
	"github.com/minskylab/collecta/ent/survey"
	"github.com/minskylab/collecta/uuid"
)

// SurveyUpdate is the builder for updating Survey entities.
type SurveyUpdate struct {
	config
	hooks      []Hook
	mutation   *SurveyMutation
	predicates []predicate.Survey
}

// Where adds a new predicate for the builder.
func (su *SurveyUpdate) Where(ps ...predicate.Survey) *SurveyUpdate {
	su.predicates = append(su.predicates, ps...)
	return su
}

// SetTags sets the tags field.
func (su *SurveyUpdate) SetTags(s []string) *SurveyUpdate {
	su.mutation.SetTags(s)
	return su
}

// SetLastInteraction sets the lastInteraction field.
func (su *SurveyUpdate) SetLastInteraction(t time.Time) *SurveyUpdate {
	su.mutation.SetLastInteraction(t)
	return su
}

// SetTitle sets the title field.
func (su *SurveyUpdate) SetTitle(s string) *SurveyUpdate {
	su.mutation.SetTitle(s)
	return su
}

// SetDescription sets the description field.
func (su *SurveyUpdate) SetDescription(s string) *SurveyUpdate {
	su.mutation.SetDescription(s)
	return su
}

// SetNillableDescription sets the description field if the given value is not nil.
func (su *SurveyUpdate) SetNillableDescription(s *string) *SurveyUpdate {
	if s != nil {
		su.SetDescription(*s)
	}
	return su
}

// ClearDescription clears the value of description.
func (su *SurveyUpdate) ClearDescription() *SurveyUpdate {
	su.mutation.ClearDescription()
	return su
}

// SetMetadata sets the metadata field.
func (su *SurveyUpdate) SetMetadata(m map[string]interface{}) *SurveyUpdate {
	su.mutation.SetMetadata(m)
	return su
}

// ClearMetadata clears the value of metadata.
func (su *SurveyUpdate) ClearMetadata() *SurveyUpdate {
	su.mutation.ClearMetadata()
	return su
}

// SetDone sets the done field.
func (su *SurveyUpdate) SetDone(b bool) *SurveyUpdate {
	su.mutation.SetDone(b)
	return su
}

// SetNillableDone sets the done field if the given value is not nil.
func (su *SurveyUpdate) SetNillableDone(b *bool) *SurveyUpdate {
	if b != nil {
		su.SetDone(*b)
	}
	return su
}

// ClearDone clears the value of done.
func (su *SurveyUpdate) ClearDone() *SurveyUpdate {
	su.mutation.ClearDone()
	return su
}

// SetIsPublic sets the isPublic field.
func (su *SurveyUpdate) SetIsPublic(b bool) *SurveyUpdate {
	su.mutation.SetIsPublic(b)
	return su
}

// SetNillableIsPublic sets the isPublic field if the given value is not nil.
func (su *SurveyUpdate) SetNillableIsPublic(b *bool) *SurveyUpdate {
	if b != nil {
		su.SetIsPublic(*b)
	}
	return su
}

// ClearIsPublic clears the value of isPublic.
func (su *SurveyUpdate) ClearIsPublic() *SurveyUpdate {
	su.mutation.ClearIsPublic()
	return su
}

// SetFlowID sets the flow edge to Flow by id.
func (su *SurveyUpdate) SetFlowID(id uuid.UUID) *SurveyUpdate {
	su.mutation.SetFlowID(id)
	return su
}

// SetFlow sets the flow edge to Flow.
func (su *SurveyUpdate) SetFlow(f *Flow) *SurveyUpdate {
	return su.SetFlowID(f.ID)
}

// SetForID sets the for edge to Person by id.
func (su *SurveyUpdate) SetForID(id uuid.UUID) *SurveyUpdate {
	su.mutation.SetForID(id)
	return su
}

// SetFor sets the for edge to Person.
func (su *SurveyUpdate) SetFor(p *Person) *SurveyUpdate {
	return su.SetForID(p.ID)
}

// SetOwnerID sets the owner edge to Domain by id.
func (su *SurveyUpdate) SetOwnerID(id uuid.UUID) *SurveyUpdate {
	su.mutation.SetOwnerID(id)
	return su
}

// SetNillableOwnerID sets the owner edge to Domain by id if the given value is not nil.
func (su *SurveyUpdate) SetNillableOwnerID(id *uuid.UUID) *SurveyUpdate {
	if id != nil {
		su = su.SetOwnerID(*id)
	}
	return su
}

// SetOwner sets the owner edge to Domain.
func (su *SurveyUpdate) SetOwner(d *Domain) *SurveyUpdate {
	return su.SetOwnerID(d.ID)
}

// ClearFlow clears the flow edge to Flow.
func (su *SurveyUpdate) ClearFlow() *SurveyUpdate {
	su.mutation.ClearFlow()
	return su
}

// ClearFor clears the for edge to Person.
func (su *SurveyUpdate) ClearFor() *SurveyUpdate {
	su.mutation.ClearFor()
	return su
}

// ClearOwner clears the owner edge to Domain.
func (su *SurveyUpdate) ClearOwner() *SurveyUpdate {
	su.mutation.ClearOwner()
	return su
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (su *SurveyUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := su.mutation.Title(); ok {
		if err := survey.TitleValidator(v); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"title\": %v", err)
		}
	}

	if _, ok := su.mutation.FlowID(); su.mutation.FlowCleared() && !ok {
		return 0, errors.New("ent: clearing a unique edge \"flow\"")
	}

	if _, ok := su.mutation.ForID(); su.mutation.ForCleared() && !ok {
		return 0, errors.New("ent: clearing a unique edge \"for\"")
	}

	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SurveyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SurveyUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SurveyUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SurveyUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SurveyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   survey.Table,
			Columns: survey.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: survey.FieldID,
			},
		},
	}
	if ps := su.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: survey.FieldTags,
		})
	}
	if value, ok := su.mutation.LastInteraction(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: survey.FieldLastInteraction,
		})
	}
	if value, ok := su.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: survey.FieldTitle,
		})
	}
	if value, ok := su.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: survey.FieldDescription,
		})
	}
	if su.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: survey.FieldDescription,
		})
	}
	if value, ok := su.mutation.Metadata(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: survey.FieldMetadata,
		})
	}
	if su.mutation.MetadataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: survey.FieldMetadata,
		})
	}
	if value, ok := su.mutation.Done(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: survey.FieldDone,
		})
	}
	if su.mutation.DoneCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: survey.FieldDone,
		})
	}
	if value, ok := su.mutation.IsPublic(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: survey.FieldIsPublic,
		})
	}
	if su.mutation.IsPublicCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: survey.FieldIsPublic,
		})
	}
	if su.mutation.FlowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   survey.FlowTable,
			Columns: []string{survey.FlowColumn},
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
	if nodes := su.mutation.FlowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   survey.FlowTable,
			Columns: []string{survey.FlowColumn},
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
	if su.mutation.ForCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   survey.ForTable,
			Columns: []string{survey.ForColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: person.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ForIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   survey.ForTable,
			Columns: []string{survey.ForColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: person.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   survey.OwnerTable,
			Columns: []string{survey.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: domain.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   survey.OwnerTable,
			Columns: []string{survey.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: domain.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{survey.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SurveyUpdateOne is the builder for updating a single Survey entity.
type SurveyUpdateOne struct {
	config
	hooks    []Hook
	mutation *SurveyMutation
}

// SetTags sets the tags field.
func (suo *SurveyUpdateOne) SetTags(s []string) *SurveyUpdateOne {
	suo.mutation.SetTags(s)
	return suo
}

// SetLastInteraction sets the lastInteraction field.
func (suo *SurveyUpdateOne) SetLastInteraction(t time.Time) *SurveyUpdateOne {
	suo.mutation.SetLastInteraction(t)
	return suo
}

// SetTitle sets the title field.
func (suo *SurveyUpdateOne) SetTitle(s string) *SurveyUpdateOne {
	suo.mutation.SetTitle(s)
	return suo
}

// SetDescription sets the description field.
func (suo *SurveyUpdateOne) SetDescription(s string) *SurveyUpdateOne {
	suo.mutation.SetDescription(s)
	return suo
}

// SetNillableDescription sets the description field if the given value is not nil.
func (suo *SurveyUpdateOne) SetNillableDescription(s *string) *SurveyUpdateOne {
	if s != nil {
		suo.SetDescription(*s)
	}
	return suo
}

// ClearDescription clears the value of description.
func (suo *SurveyUpdateOne) ClearDescription() *SurveyUpdateOne {
	suo.mutation.ClearDescription()
	return suo
}

// SetMetadata sets the metadata field.
func (suo *SurveyUpdateOne) SetMetadata(m map[string]interface{}) *SurveyUpdateOne {
	suo.mutation.SetMetadata(m)
	return suo
}

// ClearMetadata clears the value of metadata.
func (suo *SurveyUpdateOne) ClearMetadata() *SurveyUpdateOne {
	suo.mutation.ClearMetadata()
	return suo
}

// SetDone sets the done field.
func (suo *SurveyUpdateOne) SetDone(b bool) *SurveyUpdateOne {
	suo.mutation.SetDone(b)
	return suo
}

// SetNillableDone sets the done field if the given value is not nil.
func (suo *SurveyUpdateOne) SetNillableDone(b *bool) *SurveyUpdateOne {
	if b != nil {
		suo.SetDone(*b)
	}
	return suo
}

// ClearDone clears the value of done.
func (suo *SurveyUpdateOne) ClearDone() *SurveyUpdateOne {
	suo.mutation.ClearDone()
	return suo
}

// SetIsPublic sets the isPublic field.
func (suo *SurveyUpdateOne) SetIsPublic(b bool) *SurveyUpdateOne {
	suo.mutation.SetIsPublic(b)
	return suo
}

// SetNillableIsPublic sets the isPublic field if the given value is not nil.
func (suo *SurveyUpdateOne) SetNillableIsPublic(b *bool) *SurveyUpdateOne {
	if b != nil {
		suo.SetIsPublic(*b)
	}
	return suo
}

// ClearIsPublic clears the value of isPublic.
func (suo *SurveyUpdateOne) ClearIsPublic() *SurveyUpdateOne {
	suo.mutation.ClearIsPublic()
	return suo
}

// SetFlowID sets the flow edge to Flow by id.
func (suo *SurveyUpdateOne) SetFlowID(id uuid.UUID) *SurveyUpdateOne {
	suo.mutation.SetFlowID(id)
	return suo
}

// SetFlow sets the flow edge to Flow.
func (suo *SurveyUpdateOne) SetFlow(f *Flow) *SurveyUpdateOne {
	return suo.SetFlowID(f.ID)
}

// SetForID sets the for edge to Person by id.
func (suo *SurveyUpdateOne) SetForID(id uuid.UUID) *SurveyUpdateOne {
	suo.mutation.SetForID(id)
	return suo
}

// SetFor sets the for edge to Person.
func (suo *SurveyUpdateOne) SetFor(p *Person) *SurveyUpdateOne {
	return suo.SetForID(p.ID)
}

// SetOwnerID sets the owner edge to Domain by id.
func (suo *SurveyUpdateOne) SetOwnerID(id uuid.UUID) *SurveyUpdateOne {
	suo.mutation.SetOwnerID(id)
	return suo
}

// SetNillableOwnerID sets the owner edge to Domain by id if the given value is not nil.
func (suo *SurveyUpdateOne) SetNillableOwnerID(id *uuid.UUID) *SurveyUpdateOne {
	if id != nil {
		suo = suo.SetOwnerID(*id)
	}
	return suo
}

// SetOwner sets the owner edge to Domain.
func (suo *SurveyUpdateOne) SetOwner(d *Domain) *SurveyUpdateOne {
	return suo.SetOwnerID(d.ID)
}

// ClearFlow clears the flow edge to Flow.
func (suo *SurveyUpdateOne) ClearFlow() *SurveyUpdateOne {
	suo.mutation.ClearFlow()
	return suo
}

// ClearFor clears the for edge to Person.
func (suo *SurveyUpdateOne) ClearFor() *SurveyUpdateOne {
	suo.mutation.ClearFor()
	return suo
}

// ClearOwner clears the owner edge to Domain.
func (suo *SurveyUpdateOne) ClearOwner() *SurveyUpdateOne {
	suo.mutation.ClearOwner()
	return suo
}

// Save executes the query and returns the updated entity.
func (suo *SurveyUpdateOne) Save(ctx context.Context) (*Survey, error) {
	if v, ok := suo.mutation.Title(); ok {
		if err := survey.TitleValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"title\": %v", err)
		}
	}

	if _, ok := suo.mutation.FlowID(); suo.mutation.FlowCleared() && !ok {
		return nil, errors.New("ent: clearing a unique edge \"flow\"")
	}

	if _, ok := suo.mutation.ForID(); suo.mutation.ForCleared() && !ok {
		return nil, errors.New("ent: clearing a unique edge \"for\"")
	}

	var (
		err  error
		node *Survey
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SurveyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SurveyUpdateOne) SaveX(ctx context.Context) *Survey {
	s, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// Exec executes the query on the entity.
func (suo *SurveyUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SurveyUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SurveyUpdateOne) sqlSave(ctx context.Context) (s *Survey, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   survey.Table,
			Columns: survey.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: survey.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing Survey.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := suo.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: survey.FieldTags,
		})
	}
	if value, ok := suo.mutation.LastInteraction(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: survey.FieldLastInteraction,
		})
	}
	if value, ok := suo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: survey.FieldTitle,
		})
	}
	if value, ok := suo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: survey.FieldDescription,
		})
	}
	if suo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: survey.FieldDescription,
		})
	}
	if value, ok := suo.mutation.Metadata(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: survey.FieldMetadata,
		})
	}
	if suo.mutation.MetadataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: survey.FieldMetadata,
		})
	}
	if value, ok := suo.mutation.Done(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: survey.FieldDone,
		})
	}
	if suo.mutation.DoneCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: survey.FieldDone,
		})
	}
	if value, ok := suo.mutation.IsPublic(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: survey.FieldIsPublic,
		})
	}
	if suo.mutation.IsPublicCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: survey.FieldIsPublic,
		})
	}
	if suo.mutation.FlowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   survey.FlowTable,
			Columns: []string{survey.FlowColumn},
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
	if nodes := suo.mutation.FlowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   survey.FlowTable,
			Columns: []string{survey.FlowColumn},
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
	if suo.mutation.ForCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   survey.ForTable,
			Columns: []string{survey.ForColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: person.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ForIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   survey.ForTable,
			Columns: []string{survey.ForColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: person.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   survey.OwnerTable,
			Columns: []string{survey.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: domain.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   survey.OwnerTable,
			Columns: []string{survey.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: domain.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	s = &Survey{config: suo.config}
	_spec.Assign = s.assignValues
	_spec.ScanValues = s.scanValues()
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{survey.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return s, nil
}
