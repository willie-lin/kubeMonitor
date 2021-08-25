// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/k8smetric"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/predicate"
)

// K8sMetricQuery is the builder for querying K8sMetric entities.
type K8sMetricQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.K8sMetric
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the K8sMetricQuery builder.
func (kmq *K8sMetricQuery) Where(ps ...predicate.K8sMetric) *K8sMetricQuery {
	kmq.predicates = append(kmq.predicates, ps...)
	return kmq
}

// Limit adds a limit step to the query.
func (kmq *K8sMetricQuery) Limit(limit int) *K8sMetricQuery {
	kmq.limit = &limit
	return kmq
}

// Offset adds an offset step to the query.
func (kmq *K8sMetricQuery) Offset(offset int) *K8sMetricQuery {
	kmq.offset = &offset
	return kmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (kmq *K8sMetricQuery) Unique(unique bool) *K8sMetricQuery {
	kmq.unique = &unique
	return kmq
}

// Order adds an order step to the query.
func (kmq *K8sMetricQuery) Order(o ...OrderFunc) *K8sMetricQuery {
	kmq.order = append(kmq.order, o...)
	return kmq
}

// First returns the first K8sMetric entity from the query.
// Returns a *NotFoundError when no K8sMetric was found.
func (kmq *K8sMetricQuery) First(ctx context.Context) (*K8sMetric, error) {
	nodes, err := kmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{k8smetric.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (kmq *K8sMetricQuery) FirstX(ctx context.Context) *K8sMetric {
	node, err := kmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first K8sMetric ID from the query.
// Returns a *NotFoundError when no K8sMetric ID was found.
func (kmq *K8sMetricQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = kmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{k8smetric.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (kmq *K8sMetricQuery) FirstIDX(ctx context.Context) int {
	id, err := kmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single K8sMetric entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one K8sMetric entity is not found.
// Returns a *NotFoundError when no K8sMetric entities are found.
func (kmq *K8sMetricQuery) Only(ctx context.Context) (*K8sMetric, error) {
	nodes, err := kmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{k8smetric.Label}
	default:
		return nil, &NotSingularError{k8smetric.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (kmq *K8sMetricQuery) OnlyX(ctx context.Context) *K8sMetric {
	node, err := kmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only K8sMetric ID in the query.
// Returns a *NotSingularError when exactly one K8sMetric ID is not found.
// Returns a *NotFoundError when no entities are found.
func (kmq *K8sMetricQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = kmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{k8smetric.Label}
	default:
		err = &NotSingularError{k8smetric.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (kmq *K8sMetricQuery) OnlyIDX(ctx context.Context) int {
	id, err := kmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of K8sMetrics.
func (kmq *K8sMetricQuery) All(ctx context.Context) ([]*K8sMetric, error) {
	if err := kmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return kmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (kmq *K8sMetricQuery) AllX(ctx context.Context) []*K8sMetric {
	nodes, err := kmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of K8sMetric IDs.
func (kmq *K8sMetricQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := kmq.Select(k8smetric.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (kmq *K8sMetricQuery) IDsX(ctx context.Context) []int {
	ids, err := kmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (kmq *K8sMetricQuery) Count(ctx context.Context) (int, error) {
	if err := kmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return kmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (kmq *K8sMetricQuery) CountX(ctx context.Context) int {
	count, err := kmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (kmq *K8sMetricQuery) Exist(ctx context.Context) (bool, error) {
	if err := kmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return kmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (kmq *K8sMetricQuery) ExistX(ctx context.Context) bool {
	exist, err := kmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the K8sMetricQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (kmq *K8sMetricQuery) Clone() *K8sMetricQuery {
	if kmq == nil {
		return nil
	}
	return &K8sMetricQuery{
		config:     kmq.config,
		limit:      kmq.limit,
		offset:     kmq.offset,
		order:      append([]OrderFunc{}, kmq.order...),
		predicates: append([]predicate.K8sMetric{}, kmq.predicates...),
		// clone intermediate query.
		sql:  kmq.sql.Clone(),
		path: kmq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.K8sMetric.Query().
//		GroupBy(k8smetric.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (kmq *K8sMetricQuery) GroupBy(field string, fields ...string) *K8sMetricGroupBy {
	group := &K8sMetricGroupBy{config: kmq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := kmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return kmq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.K8sMetric.Query().
//		Select(k8smetric.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (kmq *K8sMetricQuery) Select(fields ...string) *K8sMetricSelect {
	kmq.fields = append(kmq.fields, fields...)
	return &K8sMetricSelect{K8sMetricQuery: kmq}
}

func (kmq *K8sMetricQuery) prepareQuery(ctx context.Context) error {
	for _, f := range kmq.fields {
		if !k8smetric.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if kmq.path != nil {
		prev, err := kmq.path(ctx)
		if err != nil {
			return err
		}
		kmq.sql = prev
	}
	return nil
}

func (kmq *K8sMetricQuery) sqlAll(ctx context.Context) ([]*K8sMetric, error) {
	var (
		nodes = []*K8sMetric{}
		_spec = kmq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &K8sMetric{config: kmq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, kmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (kmq *K8sMetricQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := kmq.querySpec()
	return sqlgraph.CountNodes(ctx, kmq.driver, _spec)
}

func (kmq *K8sMetricQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := kmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (kmq *K8sMetricQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   k8smetric.Table,
			Columns: k8smetric.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: k8smetric.FieldID,
			},
		},
		From:   kmq.sql,
		Unique: true,
	}
	if unique := kmq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := kmq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, k8smetric.FieldID)
		for i := range fields {
			if fields[i] != k8smetric.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := kmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := kmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := kmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := kmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (kmq *K8sMetricQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(kmq.driver.Dialect())
	t1 := builder.Table(k8smetric.Table)
	columns := kmq.fields
	if len(columns) == 0 {
		columns = k8smetric.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if kmq.sql != nil {
		selector = kmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range kmq.predicates {
		p(selector)
	}
	for _, p := range kmq.order {
		p(selector)
	}
	if offset := kmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := kmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// K8sMetricGroupBy is the group-by builder for K8sMetric entities.
type K8sMetricGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (kmgb *K8sMetricGroupBy) Aggregate(fns ...AggregateFunc) *K8sMetricGroupBy {
	kmgb.fns = append(kmgb.fns, fns...)
	return kmgb
}

// Scan applies the group-by query and scans the result into the given value.
func (kmgb *K8sMetricGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := kmgb.path(ctx)
	if err != nil {
		return err
	}
	kmgb.sql = query
	return kmgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (kmgb *K8sMetricGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := kmgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (kmgb *K8sMetricGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(kmgb.fields) > 1 {
		return nil, errors.New("ent: K8sMetricGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := kmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (kmgb *K8sMetricGroupBy) StringsX(ctx context.Context) []string {
	v, err := kmgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (kmgb *K8sMetricGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = kmgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{k8smetric.Label}
	default:
		err = fmt.Errorf("ent: K8sMetricGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (kmgb *K8sMetricGroupBy) StringX(ctx context.Context) string {
	v, err := kmgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (kmgb *K8sMetricGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(kmgb.fields) > 1 {
		return nil, errors.New("ent: K8sMetricGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := kmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (kmgb *K8sMetricGroupBy) IntsX(ctx context.Context) []int {
	v, err := kmgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (kmgb *K8sMetricGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = kmgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{k8smetric.Label}
	default:
		err = fmt.Errorf("ent: K8sMetricGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (kmgb *K8sMetricGroupBy) IntX(ctx context.Context) int {
	v, err := kmgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (kmgb *K8sMetricGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(kmgb.fields) > 1 {
		return nil, errors.New("ent: K8sMetricGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := kmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (kmgb *K8sMetricGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := kmgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (kmgb *K8sMetricGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = kmgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{k8smetric.Label}
	default:
		err = fmt.Errorf("ent: K8sMetricGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (kmgb *K8sMetricGroupBy) Float64X(ctx context.Context) float64 {
	v, err := kmgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (kmgb *K8sMetricGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(kmgb.fields) > 1 {
		return nil, errors.New("ent: K8sMetricGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := kmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (kmgb *K8sMetricGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := kmgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (kmgb *K8sMetricGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = kmgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{k8smetric.Label}
	default:
		err = fmt.Errorf("ent: K8sMetricGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (kmgb *K8sMetricGroupBy) BoolX(ctx context.Context) bool {
	v, err := kmgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (kmgb *K8sMetricGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range kmgb.fields {
		if !k8smetric.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := kmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := kmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (kmgb *K8sMetricGroupBy) sqlQuery() *sql.Selector {
	selector := kmgb.sql.Select()
	aggregation := make([]string, 0, len(kmgb.fns))
	for _, fn := range kmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(kmgb.fields)+len(kmgb.fns))
		for _, f := range kmgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(kmgb.fields...)...)
}

// K8sMetricSelect is the builder for selecting fields of K8sMetric entities.
type K8sMetricSelect struct {
	*K8sMetricQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (kms *K8sMetricSelect) Scan(ctx context.Context, v interface{}) error {
	if err := kms.prepareQuery(ctx); err != nil {
		return err
	}
	kms.sql = kms.K8sMetricQuery.sqlQuery(ctx)
	return kms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (kms *K8sMetricSelect) ScanX(ctx context.Context, v interface{}) {
	if err := kms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (kms *K8sMetricSelect) Strings(ctx context.Context) ([]string, error) {
	if len(kms.fields) > 1 {
		return nil, errors.New("ent: K8sMetricSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := kms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (kms *K8sMetricSelect) StringsX(ctx context.Context) []string {
	v, err := kms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (kms *K8sMetricSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = kms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{k8smetric.Label}
	default:
		err = fmt.Errorf("ent: K8sMetricSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (kms *K8sMetricSelect) StringX(ctx context.Context) string {
	v, err := kms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (kms *K8sMetricSelect) Ints(ctx context.Context) ([]int, error) {
	if len(kms.fields) > 1 {
		return nil, errors.New("ent: K8sMetricSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := kms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (kms *K8sMetricSelect) IntsX(ctx context.Context) []int {
	v, err := kms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (kms *K8sMetricSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = kms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{k8smetric.Label}
	default:
		err = fmt.Errorf("ent: K8sMetricSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (kms *K8sMetricSelect) IntX(ctx context.Context) int {
	v, err := kms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (kms *K8sMetricSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(kms.fields) > 1 {
		return nil, errors.New("ent: K8sMetricSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := kms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (kms *K8sMetricSelect) Float64sX(ctx context.Context) []float64 {
	v, err := kms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (kms *K8sMetricSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = kms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{k8smetric.Label}
	default:
		err = fmt.Errorf("ent: K8sMetricSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (kms *K8sMetricSelect) Float64X(ctx context.Context) float64 {
	v, err := kms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (kms *K8sMetricSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(kms.fields) > 1 {
		return nil, errors.New("ent: K8sMetricSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := kms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (kms *K8sMetricSelect) BoolsX(ctx context.Context) []bool {
	v, err := kms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (kms *K8sMetricSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = kms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{k8smetric.Label}
	default:
		err = fmt.Errorf("ent: K8sMetricSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (kms *K8sMetricSelect) BoolX(ctx context.Context) bool {
	v, err := kms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (kms *K8sMetricSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := kms.sql.Query()
	if err := kms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
