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
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/event"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/metric"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/metricname"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/metrictype"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/predicate"
)

// MetricNameQuery is the builder for querying MetricName entities.
type MetricNameQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.MetricName
	// eager-loading edges.
	withMetrics *MetricQuery
	withEvents  *EventQuery
	withOwners  *MetricTypeQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MetricNameQuery builder.
func (mnq *MetricNameQuery) Where(ps ...predicate.MetricName) *MetricNameQuery {
	mnq.predicates = append(mnq.predicates, ps...)
	return mnq
}

// Limit adds a limit step to the query.
func (mnq *MetricNameQuery) Limit(limit int) *MetricNameQuery {
	mnq.limit = &limit
	return mnq
}

// Offset adds an offset step to the query.
func (mnq *MetricNameQuery) Offset(offset int) *MetricNameQuery {
	mnq.offset = &offset
	return mnq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mnq *MetricNameQuery) Unique(unique bool) *MetricNameQuery {
	mnq.unique = &unique
	return mnq
}

// Order adds an order step to the query.
func (mnq *MetricNameQuery) Order(o ...OrderFunc) *MetricNameQuery {
	mnq.order = append(mnq.order, o...)
	return mnq
}

// QueryMetrics chains the current query on the "metrics" edge.
func (mnq *MetricNameQuery) QueryMetrics() *MetricQuery {
	query := &MetricQuery{config: mnq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mnq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mnq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(metricname.Table, metricname.FieldID, selector),
			sqlgraph.To(metric.Table, metric.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, metricname.MetricsTable, metricname.MetricsColumn),
		)
		fromU = sqlgraph.SetNeighbors(mnq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEvents chains the current query on the "events" edge.
func (mnq *MetricNameQuery) QueryEvents() *EventQuery {
	query := &EventQuery{config: mnq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mnq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mnq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(metricname.Table, metricname.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, metricname.EventsTable, metricname.EventsColumn),
		)
		fromU = sqlgraph.SetNeighbors(mnq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOwners chains the current query on the "owners" edge.
func (mnq *MetricNameQuery) QueryOwners() *MetricTypeQuery {
	query := &MetricTypeQuery{config: mnq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mnq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mnq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(metricname.Table, metricname.FieldID, selector),
			sqlgraph.To(metrictype.Table, metrictype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, metricname.OwnersTable, metricname.OwnersColumn),
		)
		fromU = sqlgraph.SetNeighbors(mnq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MetricName entity from the query.
// Returns a *NotFoundError when no MetricName was found.
func (mnq *MetricNameQuery) First(ctx context.Context) (*MetricName, error) {
	nodes, err := mnq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{metricname.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mnq *MetricNameQuery) FirstX(ctx context.Context) *MetricName {
	node, err := mnq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MetricName ID from the query.
// Returns a *NotFoundError when no MetricName ID was found.
func (mnq *MetricNameQuery) FirstID(ctx context.Context) (id uint, err error) {
	var ids []uint
	if ids, err = mnq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{metricname.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mnq *MetricNameQuery) FirstIDX(ctx context.Context) uint {
	id, err := mnq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MetricName entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one MetricName entity is not found.
// Returns a *NotFoundError when no MetricName entities are found.
func (mnq *MetricNameQuery) Only(ctx context.Context) (*MetricName, error) {
	nodes, err := mnq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{metricname.Label}
	default:
		return nil, &NotSingularError{metricname.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mnq *MetricNameQuery) OnlyX(ctx context.Context) *MetricName {
	node, err := mnq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MetricName ID in the query.
// Returns a *NotSingularError when exactly one MetricName ID is not found.
// Returns a *NotFoundError when no entities are found.
func (mnq *MetricNameQuery) OnlyID(ctx context.Context) (id uint, err error) {
	var ids []uint
	if ids, err = mnq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{metricname.Label}
	default:
		err = &NotSingularError{metricname.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mnq *MetricNameQuery) OnlyIDX(ctx context.Context) uint {
	id, err := mnq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MetricNames.
func (mnq *MetricNameQuery) All(ctx context.Context) ([]*MetricName, error) {
	if err := mnq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mnq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mnq *MetricNameQuery) AllX(ctx context.Context) []*MetricName {
	nodes, err := mnq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MetricName IDs.
func (mnq *MetricNameQuery) IDs(ctx context.Context) ([]uint, error) {
	var ids []uint
	if err := mnq.Select(metricname.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mnq *MetricNameQuery) IDsX(ctx context.Context) []uint {
	ids, err := mnq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mnq *MetricNameQuery) Count(ctx context.Context) (int, error) {
	if err := mnq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mnq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mnq *MetricNameQuery) CountX(ctx context.Context) int {
	count, err := mnq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mnq *MetricNameQuery) Exist(ctx context.Context) (bool, error) {
	if err := mnq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mnq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mnq *MetricNameQuery) ExistX(ctx context.Context) bool {
	exist, err := mnq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MetricNameQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mnq *MetricNameQuery) Clone() *MetricNameQuery {
	if mnq == nil {
		return nil
	}
	return &MetricNameQuery{
		config:      mnq.config,
		limit:       mnq.limit,
		offset:      mnq.offset,
		order:       append([]OrderFunc{}, mnq.order...),
		predicates:  append([]predicate.MetricName{}, mnq.predicates...),
		withMetrics: mnq.withMetrics.Clone(),
		withEvents:  mnq.withEvents.Clone(),
		withOwners:  mnq.withOwners.Clone(),
		// clone intermediate query.
		sql:  mnq.sql.Clone(),
		path: mnq.path,
	}
}

// WithMetrics tells the query-builder to eager-load the nodes that are connected to
// the "metrics" edge. The optional arguments are used to configure the query builder of the edge.
func (mnq *MetricNameQuery) WithMetrics(opts ...func(*MetricQuery)) *MetricNameQuery {
	query := &MetricQuery{config: mnq.config}
	for _, opt := range opts {
		opt(query)
	}
	mnq.withMetrics = query
	return mnq
}

// WithEvents tells the query-builder to eager-load the nodes that are connected to
// the "events" edge. The optional arguments are used to configure the query builder of the edge.
func (mnq *MetricNameQuery) WithEvents(opts ...func(*EventQuery)) *MetricNameQuery {
	query := &EventQuery{config: mnq.config}
	for _, opt := range opts {
		opt(query)
	}
	mnq.withEvents = query
	return mnq
}

// WithOwners tells the query-builder to eager-load the nodes that are connected to
// the "owners" edge. The optional arguments are used to configure the query builder of the edge.
func (mnq *MetricNameQuery) WithOwners(opts ...func(*MetricTypeQuery)) *MetricNameQuery {
	query := &MetricTypeQuery{config: mnq.config}
	for _, opt := range opts {
		opt(query)
	}
	mnq.withOwners = query
	return mnq
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
//	client.MetricName.Query().
//		GroupBy(metricname.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (mnq *MetricNameQuery) GroupBy(field string, fields ...string) *MetricNameGroupBy {
	group := &MetricNameGroupBy{config: mnq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mnq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mnq.sqlQuery(ctx), nil
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
//	client.MetricName.Query().
//		Select(metricname.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (mnq *MetricNameQuery) Select(fields ...string) *MetricNameSelect {
	mnq.fields = append(mnq.fields, fields...)
	return &MetricNameSelect{MetricNameQuery: mnq}
}

func (mnq *MetricNameQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mnq.fields {
		if !metricname.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mnq.path != nil {
		prev, err := mnq.path(ctx)
		if err != nil {
			return err
		}
		mnq.sql = prev
	}
	return nil
}

func (mnq *MetricNameQuery) sqlAll(ctx context.Context) ([]*MetricName, error) {
	var (
		nodes       = []*MetricName{}
		withFKs     = mnq.withFKs
		_spec       = mnq.querySpec()
		loadedTypes = [3]bool{
			mnq.withMetrics != nil,
			mnq.withEvents != nil,
			mnq.withOwners != nil,
		}
	)
	if mnq.withOwners != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, metricname.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &MetricName{config: mnq.config}
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
	if err := sqlgraph.QueryNodes(ctx, mnq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := mnq.withMetrics; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uint]*MetricName)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Metrics = []*Metric{}
		}
		query.withFKs = true
		query.Where(predicate.Metric(func(s *sql.Selector) {
			s.Where(sql.InValues(metricname.MetricsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.metric_name_metrics
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "metric_name_metrics" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "metric_name_metrics" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Metrics = append(node.Edges.Metrics, n)
		}
	}

	if query := mnq.withEvents; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uint]*MetricName)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Events = []*Event{}
		}
		query.withFKs = true
		query.Where(predicate.Event(func(s *sql.Selector) {
			s.Where(sql.InValues(metricname.EventsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.metric_name_events
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "metric_name_events" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "metric_name_events" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Events = append(node.Edges.Events, n)
		}
	}

	if query := mnq.withOwners; query != nil {
		ids := make([]uint, 0, len(nodes))
		nodeids := make(map[uint][]*MetricName)
		for i := range nodes {
			if nodes[i].metric_type_metric_names == nil {
				continue
			}
			fk := *nodes[i].metric_type_metric_names
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(metrictype.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "metric_type_metric_names" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Owners = n
			}
		}
	}

	return nodes, nil
}

func (mnq *MetricNameQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mnq.querySpec()
	return sqlgraph.CountNodes(ctx, mnq.driver, _spec)
}

func (mnq *MetricNameQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mnq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (mnq *MetricNameQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   metricname.Table,
			Columns: metricname.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: metricname.FieldID,
			},
		},
		From:   mnq.sql,
		Unique: true,
	}
	if unique := mnq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mnq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, metricname.FieldID)
		for i := range fields {
			if fields[i] != metricname.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mnq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mnq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mnq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mnq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mnq *MetricNameQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mnq.driver.Dialect())
	t1 := builder.Table(metricname.Table)
	columns := mnq.fields
	if len(columns) == 0 {
		columns = metricname.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mnq.sql != nil {
		selector = mnq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range mnq.predicates {
		p(selector)
	}
	for _, p := range mnq.order {
		p(selector)
	}
	if offset := mnq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mnq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MetricNameGroupBy is the group-by builder for MetricName entities.
type MetricNameGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mngb *MetricNameGroupBy) Aggregate(fns ...AggregateFunc) *MetricNameGroupBy {
	mngb.fns = append(mngb.fns, fns...)
	return mngb
}

// Scan applies the group-by query and scans the result into the given value.
func (mngb *MetricNameGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mngb.path(ctx)
	if err != nil {
		return err
	}
	mngb.sql = query
	return mngb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mngb *MetricNameGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := mngb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (mngb *MetricNameGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(mngb.fields) > 1 {
		return nil, errors.New("ent: MetricNameGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := mngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mngb *MetricNameGroupBy) StringsX(ctx context.Context) []string {
	v, err := mngb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mngb *MetricNameGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mngb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metricname.Label}
	default:
		err = fmt.Errorf("ent: MetricNameGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mngb *MetricNameGroupBy) StringX(ctx context.Context) string {
	v, err := mngb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (mngb *MetricNameGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(mngb.fields) > 1 {
		return nil, errors.New("ent: MetricNameGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := mngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mngb *MetricNameGroupBy) IntsX(ctx context.Context) []int {
	v, err := mngb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mngb *MetricNameGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mngb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metricname.Label}
	default:
		err = fmt.Errorf("ent: MetricNameGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mngb *MetricNameGroupBy) IntX(ctx context.Context) int {
	v, err := mngb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (mngb *MetricNameGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(mngb.fields) > 1 {
		return nil, errors.New("ent: MetricNameGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := mngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mngb *MetricNameGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := mngb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mngb *MetricNameGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mngb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metricname.Label}
	default:
		err = fmt.Errorf("ent: MetricNameGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mngb *MetricNameGroupBy) Float64X(ctx context.Context) float64 {
	v, err := mngb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (mngb *MetricNameGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(mngb.fields) > 1 {
		return nil, errors.New("ent: MetricNameGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := mngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mngb *MetricNameGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := mngb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mngb *MetricNameGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mngb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metricname.Label}
	default:
		err = fmt.Errorf("ent: MetricNameGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mngb *MetricNameGroupBy) BoolX(ctx context.Context) bool {
	v, err := mngb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mngb *MetricNameGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mngb.fields {
		if !metricname.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mngb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mngb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mngb *MetricNameGroupBy) sqlQuery() *sql.Selector {
	selector := mngb.sql.Select()
	aggregation := make([]string, 0, len(mngb.fns))
	for _, fn := range mngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(mngb.fields)+len(mngb.fns))
		for _, f := range mngb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(mngb.fields...)...)
}

// MetricNameSelect is the builder for selecting fields of MetricName entities.
type MetricNameSelect struct {
	*MetricNameQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (mns *MetricNameSelect) Scan(ctx context.Context, v interface{}) error {
	if err := mns.prepareQuery(ctx); err != nil {
		return err
	}
	mns.sql = mns.MetricNameQuery.sqlQuery(ctx)
	return mns.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mns *MetricNameSelect) ScanX(ctx context.Context, v interface{}) {
	if err := mns.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (mns *MetricNameSelect) Strings(ctx context.Context) ([]string, error) {
	if len(mns.fields) > 1 {
		return nil, errors.New("ent: MetricNameSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := mns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mns *MetricNameSelect) StringsX(ctx context.Context) []string {
	v, err := mns.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (mns *MetricNameSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mns.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metricname.Label}
	default:
		err = fmt.Errorf("ent: MetricNameSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mns *MetricNameSelect) StringX(ctx context.Context) string {
	v, err := mns.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (mns *MetricNameSelect) Ints(ctx context.Context) ([]int, error) {
	if len(mns.fields) > 1 {
		return nil, errors.New("ent: MetricNameSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := mns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mns *MetricNameSelect) IntsX(ctx context.Context) []int {
	v, err := mns.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (mns *MetricNameSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mns.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metricname.Label}
	default:
		err = fmt.Errorf("ent: MetricNameSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mns *MetricNameSelect) IntX(ctx context.Context) int {
	v, err := mns.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (mns *MetricNameSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(mns.fields) > 1 {
		return nil, errors.New("ent: MetricNameSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := mns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mns *MetricNameSelect) Float64sX(ctx context.Context) []float64 {
	v, err := mns.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (mns *MetricNameSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mns.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metricname.Label}
	default:
		err = fmt.Errorf("ent: MetricNameSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mns *MetricNameSelect) Float64X(ctx context.Context) float64 {
	v, err := mns.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (mns *MetricNameSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(mns.fields) > 1 {
		return nil, errors.New("ent: MetricNameSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := mns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mns *MetricNameSelect) BoolsX(ctx context.Context) []bool {
	v, err := mns.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (mns *MetricNameSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mns.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metricname.Label}
	default:
		err = fmt.Errorf("ent: MetricNameSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mns *MetricNameSelect) BoolX(ctx context.Context) bool {
	v, err := mns.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mns *MetricNameSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mns.sql.Query()
	if err := mns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
