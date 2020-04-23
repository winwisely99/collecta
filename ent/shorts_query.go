// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/minskylab/collecta/ent/predicate"
	"github.com/minskylab/collecta/ent/shorts"
)

// ShortsQuery is the builder for querying Shorts entities.
type ShortsQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Shorts
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (sq *ShortsQuery) Where(ps ...predicate.Shorts) *ShortsQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit adds a limit step to the query.
func (sq *ShortsQuery) Limit(limit int) *ShortsQuery {
	sq.limit = &limit
	return sq
}

// Offset adds an offset step to the query.
func (sq *ShortsQuery) Offset(offset int) *ShortsQuery {
	sq.offset = &offset
	return sq
}

// Order adds an order step to the query.
func (sq *ShortsQuery) Order(o ...Order) *ShortsQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// First returns the first Shorts entity in the query. Returns *NotFoundError when no shorts was found.
func (sq *ShortsQuery) First(ctx context.Context) (*Shorts, error) {
	sSlice, err := sq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(sSlice) == 0 {
		return nil, &NotFoundError{shorts.Label}
	}
	return sSlice[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *ShortsQuery) FirstX(ctx context.Context) *Shorts {
	s, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return s
}

// FirstID returns the first Shorts id in the query. Returns *NotFoundError when no id was found.
func (sq *ShortsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{shorts.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (sq *ShortsQuery) FirstXID(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Shorts entity in the query, returns an error if not exactly one entity was returned.
func (sq *ShortsQuery) Only(ctx context.Context) (*Shorts, error) {
	sSlice, err := sq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(sSlice) {
	case 1:
		return sSlice[0], nil
	case 0:
		return nil, &NotFoundError{shorts.Label}
	default:
		return nil, &NotSingularError{shorts.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *ShortsQuery) OnlyX(ctx context.Context) *Shorts {
	s, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// OnlyID returns the only Shorts id in the query, returns an error if not exactly one id was returned.
func (sq *ShortsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{shorts.Label}
	default:
		err = &NotSingularError{shorts.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (sq *ShortsQuery) OnlyXID(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ShortsSlice.
func (sq *ShortsQuery) All(ctx context.Context) ([]*Shorts, error) {
	return sq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sq *ShortsQuery) AllX(ctx context.Context) []*Shorts {
	sSlice, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return sSlice
}

// IDs executes the query and returns a list of Shorts ids.
func (sq *ShortsQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := sq.Select(shorts.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *ShortsQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *ShortsQuery) Count(ctx context.Context) (int, error) {
	return sq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sq *ShortsQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *ShortsQuery) Exist(ctx context.Context) (bool, error) {
	return sq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *ShortsQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *ShortsQuery) Clone() *ShortsQuery {
	return &ShortsQuery{
		config:     sq.config,
		limit:      sq.limit,
		offset:     sq.offset,
		order:      append([]Order{}, sq.order...),
		unique:     append([]string{}, sq.unique...),
		predicates: append([]predicate.Shorts{}, sq.predicates...),
		// clone intermediate query.
		sql: sq.sql.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Key string `json:"key,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Shorts.Query().
//		GroupBy(shorts.FieldKey).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (sq *ShortsQuery) GroupBy(field string, fields ...string) *ShortsGroupBy {
	group := &ShortsGroupBy{config: sq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = sq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Key string `json:"key,omitempty"`
//	}
//
//	client.Shorts.Query().
//		Select(shorts.FieldKey).
//		Scan(ctx, &v)
//
func (sq *ShortsQuery) Select(field string, fields ...string) *ShortsSelect {
	selector := &ShortsSelect{config: sq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = sq.sqlQuery()
	return selector
}

func (sq *ShortsQuery) sqlAll(ctx context.Context) ([]*Shorts, error) {
	var (
		nodes = []*Shorts{}
		_spec = sq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &Shorts{config: sq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (sq *ShortsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *ShortsQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (sq *ShortsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shorts.Table,
			Columns: shorts.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shorts.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *ShortsQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(shorts.Table)
	selector := builder.Select(t1.Columns(shorts.Columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(shorts.Columns...)...)
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ShortsGroupBy is the builder for group-by Shorts entities.
type ShortsGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *ShortsGroupBy) Aggregate(fns ...Aggregate) *ShortsGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the group-by query and scan the result into the given value.
func (sgb *ShortsGroupBy) Scan(ctx context.Context, v interface{}) error {
	return sgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sgb *ShortsGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (sgb *ShortsGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: ShortsGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sgb *ShortsGroupBy) StringsX(ctx context.Context) []string {
	v, err := sgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (sgb *ShortsGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: ShortsGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sgb *ShortsGroupBy) IntsX(ctx context.Context) []int {
	v, err := sgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (sgb *ShortsGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: ShortsGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sgb *ShortsGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (sgb *ShortsGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: ShortsGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sgb *ShortsGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sgb *ShortsGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sgb.sqlQuery().Query()
	if err := sgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sgb *ShortsGroupBy) sqlQuery() *sql.Selector {
	selector := sgb.sql
	columns := make([]string, 0, len(sgb.fields)+len(sgb.fns))
	columns = append(columns, sgb.fields...)
	for _, fn := range sgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(sgb.fields...)
}

// ShortsSelect is the builder for select fields of Shorts entities.
type ShortsSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (ss *ShortsSelect) Scan(ctx context.Context, v interface{}) error {
	return ss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ss *ShortsSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ss *ShortsSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: ShortsSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ss *ShortsSelect) StringsX(ctx context.Context) []string {
	v, err := ss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ss *ShortsSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: ShortsSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ss *ShortsSelect) IntsX(ctx context.Context) []int {
	v, err := ss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ss *ShortsSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: ShortsSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ss *ShortsSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ss *ShortsSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: ShortsSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ss *ShortsSelect) BoolsX(ctx context.Context) []bool {
	v, err := ss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ss *ShortsSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ss.sqlQuery().Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ss *ShortsSelect) sqlQuery() sql.Querier {
	selector := ss.sql
	selector.Select(selector.Columns(ss.fields...)...)
	return selector
}
