// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/minskylab/collecta/ent/datum"
	"github.com/minskylab/collecta/ent/predicate"
)

// DatumDelete is the builder for deleting a Datum entity.
type DatumDelete struct {
	config
	hooks      []Hook
	mutation   *DatumMutation
	predicates []predicate.Datum
}

// Where adds a new predicate to the delete builder.
func (dd *DatumDelete) Where(ps ...predicate.Datum) *DatumDelete {
	dd.predicates = append(dd.predicates, ps...)
	return dd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dd *DatumDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dd.hooks) == 0 {
		affected, err = dd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DatumMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dd.mutation = mutation
			affected, err = dd.sqlExec(ctx)
			return affected, err
		})
		for i := len(dd.hooks) - 1; i >= 0; i-- {
			mut = dd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (dd *DatumDelete) ExecX(ctx context.Context) int {
	n, err := dd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dd *DatumDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: datum.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: datum.FieldID,
			},
		},
	}
	if ps := dd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, dd.driver, _spec)
}

// DatumDeleteOne is the builder for deleting a single Datum entity.
type DatumDeleteOne struct {
	dd *DatumDelete
}

// Exec executes the deletion query.
func (ddo *DatumDeleteOne) Exec(ctx context.Context) error {
	n, err := ddo.dd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{datum.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ddo *DatumDeleteOne) ExecX(ctx context.Context) {
	ddo.dd.ExecX(ctx)
}