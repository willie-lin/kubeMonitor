// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/metricname"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/metrictype"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/predicate"
)

// MetricTypeQuery is the builder for querying MetricType entities.
type MetricTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.MetricType
	// eager-loading edges.
	withMetricNames *MetricNameQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MetricTypeQuery builder.
func (mtq *MetricTypeQuery) Where(ps ...predicate.MetricType) *MetricTypeQuery {
	mtq.predicates = append(mtq.predicates, ps...)
	return mtq
}

// Limit adds a limit step to the query.
func (mtq *MetricTypeQuery) Limit(limit int) *MetricTypeQuery {
	mtq.limit = &limit
	return mtq
}

// Offset adds an offset step to the query.
func (mtq *MetricTypeQuery) Offset(offset int) *MetricTypeQuery {
	mtq.offset = &offset
	return mtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mtq *MetricTypeQuery) Unique(unique bool) *MetricTypeQuery {
	mtq.unique = &unique
	return mtq
}

// Order adds an order step to the query.
func (mtq *MetricTypeQuery) Order(o ...OrderFunc) *MetricTypeQuery {
	mtq.order = append(mtq.order, o...)
	return mtq
}

// QueryMetricNames chains the current query on the "metricNames" edge.
func (mtq *MetricTypeQuery) QueryMetricNames() *MetricNameQuery {
	query := &MetricNameQuery{config: mtq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(metrictype.Table, metrictype.FieldID, selector),
			sqlgraph.To(metricname.Table, metricname.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, metrictype.MetricNamesTable, metrictype.MetricNamesColumn),
		)
		fromU = sqlgraph.SetNeighbors(mtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MetricType entity from the query.
// Returns a *NotFoundError when no MetricType was found.
func (mtq *MetricTypeQuery) First(ctx context.Context) (*MetricType, error) {
	nodes, err := mtq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{metrictype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mtq *MetricTypeQuery) FirstX(ctx context.Context) *MetricType {
	node, err := mtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MetricType ID from the query.
// Returns a *NotFoundError when no MetricType ID was found.
func (mtq *MetricTypeQuery) FirstID(ctx context.Context) (id uint, err error) {
	var ids []uint
	if ids, err = mtq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{metrictype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mtq *MetricTypeQuery) FirstIDX(ctx context.Context) uint {
	id, err := mtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MetricType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one MetricType entity is not found.
// Returns a *NotFoundError when no MetricType entities are found.
func (mtq *MetricTypeQuery) Only(ctx context.Context) (*MetricType, error) {
	nodes, err := mtq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{metrictype.Label}
	default:
		return nil, &NotSingularError{metrictype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mtq *MetricTypeQuery) OnlyX(ctx context.Context) *MetricType {
	node, err := mtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MetricType ID in the query.
// Returns a *NotSingularError when exactly one MetricType ID is not found.
// Returns a *NotFoundError when no entities are found.
func (mtq *MetricTypeQuery) OnlyID(ctx context.Context) (id uint, err error) {
	var ids []uint
	if ids, err = mtq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{metrictype.Label}
	default:
		err = &NotSingularError{metrictype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mtq *MetricTypeQuery) OnlyIDX(ctx context.Context) uint {
	id, err := mtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MetricTypes.
func (mtq *MetricTypeQuery) All(ctx context.Context) ([]*MetricType, error) {
	if err := mtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mtq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mtq *MetricTypeQuery) AllX(ctx context.Context) []*MetricType {
	nodes, err := mtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MetricType IDs.
func (mtq *MetricTypeQuery) IDs(ctx context.Context) ([]uint, error) {
	var ids []uint
	if err := mtq.Select(metrictype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mtq *MetricTypeQuery) IDsX(ctx context.Context) []uint {
	ids, err := mtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mtq *MetricTypeQuery) Count(ctx context.Context) (int, error) {
	if err := mtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mtq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mtq *MetricTypeQuery) CountX(ctx context.Context) int {
	count, err := mtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mtq *MetricTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := mtq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mtq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mtq *MetricTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := mtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MetricTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mtq *MetricTypeQuery) Clone() *MetricTypeQuery {
	if mtq == nil {
		return nil
	}
	return &MetricTypeQuery{
		config:          mtq.config,
		limit:           mtq.limit,
		offset:          mtq.offset,
		order:           append([]OrderFunc{}, mtq.order...),
		predicates:      append([]predicate.MetricType{}, mtq.predicates...),
		withMetricNames: mtq.withMetricNames.Clone(),
		// clone intermediate query.
		sql:  mtq.sql.Clone(),
		path: mtq.path,
	}
}

// WithMetricNames tells the query-builder to eager-load the nodes that are connected to
// the "metricNames" edge. The optional arguments are used to configure the query builder of the edge.
func (mtq *MetricTypeQuery) WithMetricNames(opts ...func(*MetricNameQuery)) *MetricTypeQuery {
	query := &MetricNameQuery{config: mtq.config}
	for _, opt := range opts {
		opt(query)
	}
	mtq.withMetricNames = query
	return mtq
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
//	client.MetricType.Query().
//		GroupBy(metrictype.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (mtq *MetricTypeQuery) GroupBy(field string, fields ...string) *MetricTypeGroupBy {
	group := &MetricTypeGroupBy{config: mtq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mtq.sqlQuery(ctx), nil
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
//	client.MetricType.Query().
//		Select(metrictype.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (mtq *MetricTypeQuery) Select(fields ...string) *MetricTypeSelect {
	mtq.fields = append(mtq.fields, fields...)
	return &MetricTypeSelect{MetricTypeQuery: mtq}
}

func (mtq *MetricTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mtq.fields {
		if !metrictype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mtq.path != nil {
		prev, err := mtq.path(ctx)
		if err != nil {
			return err
		}
		mtq.sql = prev
	}
	return nil
}

func (mtq *MetricTypeQuery) sqlAll(ctx context.Context) ([]*MetricType, error) {
	var (
		nodes       = []*MetricType{}
		_spec       = mtq.querySpec()
		loadedTypes = [1]bool{
			mtq.withMetricNames != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &MetricType{config: mtq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, mtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := mtq.withMetricNames; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uint]*MetricType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.MetricNames = []*MetricName{}
		}
		query.withFKs = true
		query.Where(predicate.MetricName(func(s *sql.Selector) {
			s.Where(sql.InValues(metrictype.MetricNamesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.metric_type_metric_names
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "metric_type_metric_names" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "metric_type_metric_names" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.MetricNames = append(node.Edges.MetricNames, n)
		}
	}

	return nodes, nil
}

func (mtq *MetricTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mtq.querySpec()
	return sqlgraph.CountNodes(ctx, mtq.driver, _spec)
}

func (mtq *MetricTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mtq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (mtq *MetricTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   metrictype.Table,
			Columns: metrictype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: metrictype.FieldID,
			},
		},
		From:   mtq.sql,
		Unique: true,
	}
	if unique := mtq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mtq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, metrictype.FieldID)
		for i := range fields {
			if fields[i] != metrictype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mtq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mtq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mtq *MetricTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mtq.driver.Dialect())
	t1 := builder.Table(metrictype.Table)
	columns := mtq.fields
	if len(columns) == 0 {
		columns = metrictype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mtq.sql != nil {
		selector = mtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range mtq.predicates {
		p(selector)
	}
	for _, p := range mtq.order {
		p(selector)
	}
	if offset := mtq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mtq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MetricTypeGroupBy is the group-by builder for MetricType entities.
type MetricTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mtgb *MetricTypeGroupBy) Aggregate(fns ...AggregateFunc) *MetricTypeGroupBy {
	mtgb.fns = append(mtgb.fns, fns...)
	return mtgb
}

// Scan applies the group-by query and scans the result into the given value.
func (mtgb *MetricTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mtgb.path(ctx)
	if err != nil {
		return err
	}
	mtgb.sql = query
	return mtgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mtgb *MetricTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := mtgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (mtgb *MetricTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(mtgb.fields) > 1 {
		return nil, errors.New("ent: MetricTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := mtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mtgb *MetricTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := mtgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mtgb *MetricTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mtgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metrictype.Label}
	default:
		err = fmt.Errorf("ent: MetricTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mtgb *MetricTypeGroupBy) StringX(ctx context.Context) string {
	v, err := mtgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (mtgb *MetricTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(mtgb.fields) > 1 {
		return nil, errors.New("ent: MetricTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := mtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mtgb *MetricTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := mtgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mtgb *MetricTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mtgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metrictype.Label}
	default:
		err = fmt.Errorf("ent: MetricTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mtgb *MetricTypeGroupBy) IntX(ctx context.Context) int {
	v, err := mtgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (mtgb *MetricTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(mtgb.fields) > 1 {
		return nil, errors.New("ent: MetricTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := mtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mtgb *MetricTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := mtgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mtgb *MetricTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mtgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metrictype.Label}
	default:
		err = fmt.Errorf("ent: MetricTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mtgb *MetricTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := mtgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (mtgb *MetricTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(mtgb.fields) > 1 {
		return nil, errors.New("ent: MetricTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := mtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mtgb *MetricTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := mtgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mtgb *MetricTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mtgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metrictype.Label}
	default:
		err = fmt.Errorf("ent: MetricTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mtgb *MetricTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := mtgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mtgb *MetricTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mtgb.fields {
		if !metrictype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mtgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mtgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mtgb *MetricTypeGroupBy) sqlQuery() *sql.Selector {
	selector := mtgb.sql.Select()
	aggregation := make([]string, 0, len(mtgb.fns))
	for _, fn := range mtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(mtgb.fields)+len(mtgb.fns))
		for _, f := range mtgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(mtgb.fields...)...)
}

// MetricTypeSelect is the builder for selecting fields of MetricType entities.
type MetricTypeSelect struct {
	*MetricTypeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (mts *MetricTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := mts.prepareQuery(ctx); err != nil {
		return err
	}
	mts.sql = mts.MetricTypeQuery.sqlQuery(ctx)
	return mts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mts *MetricTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := mts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (mts *MetricTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(mts.fields) > 1 {
		return nil, errors.New("ent: MetricTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := mts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mts *MetricTypeSelect) StringsX(ctx context.Context) []string {
	v, err := mts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (mts *MetricTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metrictype.Label}
	default:
		err = fmt.Errorf("ent: MetricTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mts *MetricTypeSelect) StringX(ctx context.Context) string {
	v, err := mts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (mts *MetricTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(mts.fields) > 1 {
		return nil, errors.New("ent: MetricTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := mts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mts *MetricTypeSelect) IntsX(ctx context.Context) []int {
	v, err := mts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (mts *MetricTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metrictype.Label}
	default:
		err = fmt.Errorf("ent: MetricTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mts *MetricTypeSelect) IntX(ctx context.Context) int {
	v, err := mts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (mts *MetricTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(mts.fields) > 1 {
		return nil, errors.New("ent: MetricTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := mts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mts *MetricTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := mts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (mts *MetricTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metrictype.Label}
	default:
		err = fmt.Errorf("ent: MetricTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mts *MetricTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := mts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (mts *MetricTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(mts.fields) > 1 {
		return nil, errors.New("ent: MetricTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := mts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mts *MetricTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := mts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (mts *MetricTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metrictype.Label}
	default:
		err = fmt.Errorf("ent: MetricTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mts *MetricTypeSelect) BoolX(ctx context.Context) bool {
	v, err := mts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mts *MetricTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mts.sql.Query()
	if err := mts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}