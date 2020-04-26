// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/minskylab/collecta/ent/predicate"
	"github.com/minskylab/collecta/ent/short"
)

// ShortDelete is the builder for deleting a Short entity.
type ShortDelete struct {
	config
	predicates []predicate.Short
}

// Where adds a new predicate to the delete builder.
func (sd *ShortDelete) Where(ps ...predicate.Short) *ShortDelete {
	sd.predicates = append(sd.predicates, ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *ShortDelete) Exec(ctx context.Context) (int, error) {
	return sd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *ShortDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *ShortDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: short.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: short.FieldID,
			},
		},
	}
	if ps := sd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
}

// ShortDeleteOne is the builder for deleting a single Short entity.
type ShortDeleteOne struct {
	sd *ShortDelete
}

// Exec executes the deletion query.
func (sdo *ShortDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{short.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *ShortDeleteOne) ExecX(ctx context.Context) {
	sdo.sd.ExecX(ctx)
}