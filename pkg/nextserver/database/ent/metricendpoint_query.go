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
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/metricendpoint"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/predicate"
)

// MetricEndpointQuery is the builder for querying MetricEndpoint entities.
type MetricEndpointQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.MetricEndpoint
	// eager-loading edges.
	withMetrics *MetricQuery
	withEvents  *EventQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MetricEndpointQuery builder.
func (meq *MetricEndpointQuery) Where(ps ...predicate.MetricEndpoint) *MetricEndpointQuery {
	meq.predicates = append(meq.predicates, ps...)
	return meq
}

// Limit adds a limit step to the query.
func (meq *MetricEndpointQuery) Limit(limit int) *MetricEndpointQuery {
	meq.limit = &limit
	return meq
}

// Offset adds an offset step to the query.
func (meq *MetricEndpointQuery) Offset(offset int) *MetricEndpointQuery {
	meq.offset = &offset
	return meq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (meq *MetricEndpointQuery) Unique(unique bool) *MetricEndpointQuery {
	meq.unique = &unique
	return meq
}

// Order adds an order step to the query.
func (meq *MetricEndpointQuery) Order(o ...OrderFunc) *MetricEndpointQuery {
	meq.order = append(meq.order, o...)
	return meq
}

// QueryMetrics chains the current query on the "metrics" edge.
func (meq *MetricEndpointQuery) QueryMetrics() *MetricQuery {
	query := &MetricQuery{config: meq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := meq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := meq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(metricendpoint.Table, metricendpoint.FieldID, selector),
			sqlgraph.To(metric.Table, metric.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, metricendpoint.MetricsTable, metricendpoint.MetricsColumn),
		)
		fromU = sqlgraph.SetNeighbors(meq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEvents chains the current query on the "events" edge.
func (meq *MetricEndpointQuery) QueryEvents() *EventQuery {
	query := &EventQuery{config: meq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := meq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := meq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(metricendpoint.Table, metricendpoint.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, metricendpoint.EventsTable, metricendpoint.EventsColumn),
		)
		fromU = sqlgraph.SetNeighbors(meq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MetricEndpoint entity from the query.
// Returns a *NotFoundError when no MetricEndpoint was found.
func (meq *MetricEndpointQuery) First(ctx context.Context) (*MetricEndpoint, error) {
	nodes, err := meq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{metricendpoint.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (meq *MetricEndpointQuery) FirstX(ctx context.Context) *MetricEndpoint {
	node, err := meq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MetricEndpoint ID from the query.
// Returns a *NotFoundError when no MetricEndpoint ID was found.
func (meq *MetricEndpointQuery) FirstID(ctx context.Context) (id uint, err error) {
	var ids []uint
	if ids, err = meq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{metricendpoint.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (meq *MetricEndpointQuery) FirstIDX(ctx context.Context) uint {
	id, err := meq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MetricEndpoint entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one MetricEndpoint entity is not found.
// Returns a *NotFoundError when no MetricEndpoint entities are found.
func (meq *MetricEndpointQuery) Only(ctx context.Context) (*MetricEndpoint, error) {
	nodes, err := meq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{metricendpoint.Label}
	default:
		return nil, &NotSingularError{metricendpoint.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (meq *MetricEndpointQuery) OnlyX(ctx context.Context) *MetricEndpoint {
	node, err := meq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MetricEndpoint ID in the query.
// Returns a *NotSingularError when exactly one MetricEndpoint ID is not found.
// Returns a *NotFoundError when no entities are found.
func (meq *MetricEndpointQuery) OnlyID(ctx context.Context) (id uint, err error) {
	var ids []uint
	if ids, err = meq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{metricendpoint.Label}
	default:
		err = &NotSingularError{metricendpoint.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (meq *MetricEndpointQuery) OnlyIDX(ctx context.Context) uint {
	id, err := meq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MetricEndpoints.
func (meq *MetricEndpointQuery) All(ctx context.Context) ([]*MetricEndpoint, error) {
	if err := meq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return meq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (meq *MetricEndpointQuery) AllX(ctx context.Context) []*MetricEndpoint {
	nodes, err := meq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MetricEndpoint IDs.
func (meq *MetricEndpointQuery) IDs(ctx context.Context) ([]uint, error) {
	var ids []uint
	if err := meq.Select(metricendpoint.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (meq *MetricEndpointQuery) IDsX(ctx context.Context) []uint {
	ids, err := meq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (meq *MetricEndpointQuery) Count(ctx context.Context) (int, error) {
	if err := meq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return meq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (meq *MetricEndpointQuery) CountX(ctx context.Context) int {
	count, err := meq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (meq *MetricEndpointQuery) Exist(ctx context.Context) (bool, error) {
	if err := meq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return meq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (meq *MetricEndpointQuery) ExistX(ctx context.Context) bool {
	exist, err := meq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MetricEndpointQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (meq *MetricEndpointQuery) Clone() *MetricEndpointQuery {
	if meq == nil {
		return nil
	}
	return &MetricEndpointQuery{
		config:      meq.config,
		limit:       meq.limit,
		offset:      meq.offset,
		order:       append([]OrderFunc{}, meq.order...),
		predicates:  append([]predicate.MetricEndpoint{}, meq.predicates...),
		withMetrics: meq.withMetrics.Clone(),
		withEvents:  meq.withEvents.Clone(),
		// clone intermediate query.
		sql:  meq.sql.Clone(),
		path: meq.path,
	}
}

// WithMetrics tells the query-builder to eager-load the nodes that are connected to
// the "metrics" edge. The optional arguments are used to configure the query builder of the edge.
func (meq *MetricEndpointQuery) WithMetrics(opts ...func(*MetricQuery)) *MetricEndpointQuery {
	query := &MetricQuery{config: meq.config}
	for _, opt := range opts {
		opt(query)
	}
	meq.withMetrics = query
	return meq
}

// WithEvents tells the query-builder to eager-load the nodes that are connected to
// the "events" edge. The optional arguments are used to configure the query builder of the edge.
func (meq *MetricEndpointQuery) WithEvents(opts ...func(*EventQuery)) *MetricEndpointQuery {
	query := &EventQuery{config: meq.config}
	for _, opt := range opts {
		opt(query)
	}
	meq.withEvents = query
	return meq
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
//	client.MetricEndpoint.Query().
//		GroupBy(metricendpoint.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (meq *MetricEndpointQuery) GroupBy(field string, fields ...string) *MetricEndpointGroupBy {
	group := &MetricEndpointGroupBy{config: meq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := meq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return meq.sqlQuery(ctx), nil
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
//	client.MetricEndpoint.Query().
//		Select(metricendpoint.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (meq *MetricEndpointQuery) Select(fields ...string) *MetricEndpointSelect {
	meq.fields = append(meq.fields, fields...)
	return &MetricEndpointSelect{MetricEndpointQuery: meq}
}

func (meq *MetricEndpointQuery) prepareQuery(ctx context.Context) error {
	for _, f := range meq.fields {
		if !metricendpoint.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if meq.path != nil {
		prev, err := meq.path(ctx)
		if err != nil {
			return err
		}
		meq.sql = prev
	}
	return nil
}

func (meq *MetricEndpointQuery) sqlAll(ctx context.Context) ([]*MetricEndpoint, error) {
	var (
		nodes       = []*MetricEndpoint{}
		_spec       = meq.querySpec()
		loadedTypes = [2]bool{
			meq.withMetrics != nil,
			meq.withEvents != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &MetricEndpoint{config: meq.config}
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
	if err := sqlgraph.QueryNodes(ctx, meq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := meq.withMetrics; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uint]*MetricEndpoint)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Metrics = []*Metric{}
		}
		query.withFKs = true
		query.Where(predicate.Metric(func(s *sql.Selector) {
			s.Where(sql.InValues(metricendpoint.MetricsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.metric_endpoint_metrics
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "metric_endpoint_metrics" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "metric_endpoint_metrics" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Metrics = append(node.Edges.Metrics, n)
		}
	}

	if query := meq.withEvents; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uint]*MetricEndpoint)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Events = []*Event{}
		}
		query.withFKs = true
		query.Where(predicate.Event(func(s *sql.Selector) {
			s.Where(sql.InValues(metricendpoint.EventsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.metric_endpoint_events
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "metric_endpoint_events" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "metric_endpoint_events" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Events = append(node.Edges.Events, n)
		}
	}

	return nodes, nil
}

func (meq *MetricEndpointQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := meq.querySpec()
	return sqlgraph.CountNodes(ctx, meq.driver, _spec)
}

func (meq *MetricEndpointQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := meq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (meq *MetricEndpointQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   metricendpoint.Table,
			Columns: metricendpoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: metricendpoint.FieldID,
			},
		},
		From:   meq.sql,
		Unique: true,
	}
	if unique := meq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := meq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, metricendpoint.FieldID)
		for i := range fields {
			if fields[i] != metricendpoint.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := meq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := meq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := meq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := meq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (meq *MetricEndpointQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(meq.driver.Dialect())
	t1 := builder.Table(metricendpoint.Table)
	columns := meq.fields
	if len(columns) == 0 {
		columns = metricendpoint.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if meq.sql != nil {
		selector = meq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range meq.predicates {
		p(selector)
	}
	for _, p := range meq.order {
		p(selector)
	}
	if offset := meq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := meq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MetricEndpointGroupBy is the group-by builder for MetricEndpoint entities.
type MetricEndpointGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (megb *MetricEndpointGroupBy) Aggregate(fns ...AggregateFunc) *MetricEndpointGroupBy {
	megb.fns = append(megb.fns, fns...)
	return megb
}

// Scan applies the group-by query and scans the result into the given value.
func (megb *MetricEndpointGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := megb.path(ctx)
	if err != nil {
		return err
	}
	megb.sql = query
	return megb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (megb *MetricEndpointGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := megb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (megb *MetricEndpointGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(megb.fields) > 1 {
		return nil, errors.New("ent: MetricEndpointGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := megb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (megb *MetricEndpointGroupBy) StringsX(ctx context.Context) []string {
	v, err := megb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (megb *MetricEndpointGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = megb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metricendpoint.Label}
	default:
		err = fmt.Errorf("ent: MetricEndpointGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (megb *MetricEndpointGroupBy) StringX(ctx context.Context) string {
	v, err := megb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (megb *MetricEndpointGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(megb.fields) > 1 {
		return nil, errors.New("ent: MetricEndpointGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := megb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (megb *MetricEndpointGroupBy) IntsX(ctx context.Context) []int {
	v, err := megb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (megb *MetricEndpointGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = megb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metricendpoint.Label}
	default:
		err = fmt.Errorf("ent: MetricEndpointGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (megb *MetricEndpointGroupBy) IntX(ctx context.Context) int {
	v, err := megb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (megb *MetricEndpointGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(megb.fields) > 1 {
		return nil, errors.New("ent: MetricEndpointGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := megb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (megb *MetricEndpointGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := megb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (megb *MetricEndpointGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = megb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metricendpoint.Label}
	default:
		err = fmt.Errorf("ent: MetricEndpointGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (megb *MetricEndpointGroupBy) Float64X(ctx context.Context) float64 {
	v, err := megb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (megb *MetricEndpointGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(megb.fields) > 1 {
		return nil, errors.New("ent: MetricEndpointGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := megb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (megb *MetricEndpointGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := megb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (megb *MetricEndpointGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = megb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metricendpoint.Label}
	default:
		err = fmt.Errorf("ent: MetricEndpointGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (megb *MetricEndpointGroupBy) BoolX(ctx context.Context) bool {
	v, err := megb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (megb *MetricEndpointGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range megb.fields {
		if !metricendpoint.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := megb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := megb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (megb *MetricEndpointGroupBy) sqlQuery() *sql.Selector {
	selector := megb.sql.Select()
	aggregation := make([]string, 0, len(megb.fns))
	for _, fn := range megb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(megb.fields)+len(megb.fns))
		for _, f := range megb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(megb.fields...)...)
}

// MetricEndpointSelect is the builder for selecting fields of MetricEndpoint entities.
type MetricEndpointSelect struct {
	*MetricEndpointQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (mes *MetricEndpointSelect) Scan(ctx context.Context, v interface{}) error {
	if err := mes.prepareQuery(ctx); err != nil {
		return err
	}
	mes.sql = mes.MetricEndpointQuery.sqlQuery(ctx)
	return mes.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mes *MetricEndpointSelect) ScanX(ctx context.Context, v interface{}) {
	if err := mes.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (mes *MetricEndpointSelect) Strings(ctx context.Context) ([]string, error) {
	if len(mes.fields) > 1 {
		return nil, errors.New("ent: MetricEndpointSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := mes.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mes *MetricEndpointSelect) StringsX(ctx context.Context) []string {
	v, err := mes.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (mes *MetricEndpointSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mes.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metricendpoint.Label}
	default:
		err = fmt.Errorf("ent: MetricEndpointSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mes *MetricEndpointSelect) StringX(ctx context.Context) string {
	v, err := mes.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (mes *MetricEndpointSelect) Ints(ctx context.Context) ([]int, error) {
	if len(mes.fields) > 1 {
		return nil, errors.New("ent: MetricEndpointSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := mes.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mes *MetricEndpointSelect) IntsX(ctx context.Context) []int {
	v, err := mes.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (mes *MetricEndpointSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mes.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metricendpoint.Label}
	default:
		err = fmt.Errorf("ent: MetricEndpointSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mes *MetricEndpointSelect) IntX(ctx context.Context) int {
	v, err := mes.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (mes *MetricEndpointSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(mes.fields) > 1 {
		return nil, errors.New("ent: MetricEndpointSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := mes.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mes *MetricEndpointSelect) Float64sX(ctx context.Context) []float64 {
	v, err := mes.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (mes *MetricEndpointSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mes.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metricendpoint.Label}
	default:
		err = fmt.Errorf("ent: MetricEndpointSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mes *MetricEndpointSelect) Float64X(ctx context.Context) float64 {
	v, err := mes.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (mes *MetricEndpointSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(mes.fields) > 1 {
		return nil, errors.New("ent: MetricEndpointSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := mes.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mes *MetricEndpointSelect) BoolsX(ctx context.Context) []bool {
	v, err := mes.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (mes *MetricEndpointSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mes.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metricendpoint.Label}
	default:
		err = fmt.Errorf("ent: MetricEndpointSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mes *MetricEndpointSelect) BoolX(ctx context.Context) bool {
	v, err := mes.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mes *MetricEndpointSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mes.sql.Query()
	if err := mes.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
