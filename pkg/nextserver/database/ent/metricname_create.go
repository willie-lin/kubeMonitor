// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/event"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/metric"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/metricname"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/metrictype"
)

// MetricNameCreate is the builder for creating a MetricName entity.
type MetricNameCreate struct {
	config
	mutation *MetricNameMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mnc *MetricNameCreate) SetCreatedAt(t time.Time) *MetricNameCreate {
	mnc.mutation.SetCreatedAt(t)
	return mnc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mnc *MetricNameCreate) SetNillableCreatedAt(t *time.Time) *MetricNameCreate {
	if t != nil {
		mnc.SetCreatedAt(*t)
	}
	return mnc
}

// SetUpdatedAt sets the "updated_at" field.
func (mnc *MetricNameCreate) SetUpdatedAt(t time.Time) *MetricNameCreate {
	mnc.mutation.SetUpdatedAt(t)
	return mnc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mnc *MetricNameCreate) SetNillableUpdatedAt(t *time.Time) *MetricNameCreate {
	if t != nil {
		mnc.SetUpdatedAt(*t)
	}
	return mnc
}

// SetDeletedAt sets the "deleted_at" field.
func (mnc *MetricNameCreate) SetDeletedAt(t time.Time) *MetricNameCreate {
	mnc.mutation.SetDeletedAt(t)
	return mnc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mnc *MetricNameCreate) SetNillableDeletedAt(t *time.Time) *MetricNameCreate {
	if t != nil {
		mnc.SetDeletedAt(*t)
	}
	return mnc
}

// SetName sets the "name" field.
func (mnc *MetricNameCreate) SetName(s string) *MetricNameCreate {
	mnc.mutation.SetName(s)
	return mnc
}

// SetHelp sets the "help" field.
func (mnc *MetricNameCreate) SetHelp(s string) *MetricNameCreate {
	mnc.mutation.SetHelp(s)
	return mnc
}

// SetTypeId sets the "typeId" field.
func (mnc *MetricNameCreate) SetTypeId(u uint) *MetricNameCreate {
	mnc.mutation.SetTypeId(u)
	return mnc
}

// SetID sets the "id" field.
func (mnc *MetricNameCreate) SetID(u uint) *MetricNameCreate {
	mnc.mutation.SetID(u)
	return mnc
}

// AddMetricIDs adds the "metrics" edge to the Metric entity by IDs.
func (mnc *MetricNameCreate) AddMetricIDs(ids ...int) *MetricNameCreate {
	mnc.mutation.AddMetricIDs(ids...)
	return mnc
}

// AddMetrics adds the "metrics" edges to the Metric entity.
func (mnc *MetricNameCreate) AddMetrics(m ...*Metric) *MetricNameCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mnc.AddMetricIDs(ids...)
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (mnc *MetricNameCreate) AddEventIDs(ids ...int) *MetricNameCreate {
	mnc.mutation.AddEventIDs(ids...)
	return mnc
}

// AddEvents adds the "events" edges to the Event entity.
func (mnc *MetricNameCreate) AddEvents(e ...*Event) *MetricNameCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return mnc.AddEventIDs(ids...)
}

// SetOwnersID sets the "owners" edge to the MetricType entity by ID.
func (mnc *MetricNameCreate) SetOwnersID(id uint) *MetricNameCreate {
	mnc.mutation.SetOwnersID(id)
	return mnc
}

// SetNillableOwnersID sets the "owners" edge to the MetricType entity by ID if the given value is not nil.
func (mnc *MetricNameCreate) SetNillableOwnersID(id *uint) *MetricNameCreate {
	if id != nil {
		mnc = mnc.SetOwnersID(*id)
	}
	return mnc
}

// SetOwners sets the "owners" edge to the MetricType entity.
func (mnc *MetricNameCreate) SetOwners(m *MetricType) *MetricNameCreate {
	return mnc.SetOwnersID(m.ID)
}

// Mutation returns the MetricNameMutation object of the builder.
func (mnc *MetricNameCreate) Mutation() *MetricNameMutation {
	return mnc.mutation
}

// Save creates the MetricName in the database.
func (mnc *MetricNameCreate) Save(ctx context.Context) (*MetricName, error) {
	var (
		err  error
		node *MetricName
	)
	mnc.defaults()
	if len(mnc.hooks) == 0 {
		if err = mnc.check(); err != nil {
			return nil, err
		}
		node, err = mnc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MetricNameMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mnc.check(); err != nil {
				return nil, err
			}
			mnc.mutation = mutation
			if node, err = mnc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mnc.hooks) - 1; i >= 0; i-- {
			if mnc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mnc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mnc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mnc *MetricNameCreate) SaveX(ctx context.Context) *MetricName {
	v, err := mnc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mnc *MetricNameCreate) Exec(ctx context.Context) error {
	_, err := mnc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mnc *MetricNameCreate) ExecX(ctx context.Context) {
	if err := mnc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mnc *MetricNameCreate) defaults() {
	if _, ok := mnc.mutation.CreatedAt(); !ok {
		v := metricname.DefaultCreatedAt()
		mnc.mutation.SetCreatedAt(v)
	}
	if _, ok := mnc.mutation.UpdatedAt(); !ok {
		v := metricname.DefaultUpdatedAt()
		mnc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mnc.mutation.DeletedAt(); !ok {
		v := metricname.DefaultDeletedAt()
		mnc.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mnc *MetricNameCreate) check() error {
	if _, ok := mnc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := mnc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := mnc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "deleted_at"`)}
	}
	if _, ok := mnc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := mnc.mutation.Name(); ok {
		if err := metricname.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := mnc.mutation.Help(); !ok {
		return &ValidationError{Name: "help", err: errors.New(`ent: missing required field "help"`)}
	}
	if _, ok := mnc.mutation.TypeId(); !ok {
		return &ValidationError{Name: "typeId", err: errors.New(`ent: missing required field "typeId"`)}
	}
	return nil
}

func (mnc *MetricNameCreate) sqlSave(ctx context.Context) (*MetricName, error) {
	_node, _spec := mnc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mnc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint(id)
	}
	return _node, nil
}

func (mnc *MetricNameCreate) createSpec() (*MetricName, *sqlgraph.CreateSpec) {
	var (
		_node = &MetricName{config: mnc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: metricname.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: metricname.FieldID,
			},
		}
	)
	if id, ok := mnc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mnc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: metricname.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := mnc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: metricname.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := mnc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: metricname.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := mnc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: metricname.FieldName,
		})
		_node.Name = value
	}
	if value, ok := mnc.mutation.Help(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: metricname.FieldHelp,
		})
		_node.Help = value
	}
	if value, ok := mnc.mutation.TypeId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metricname.FieldTypeId,
		})
		_node.TypeId = value
	}
	if nodes := mnc.mutation.MetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metricname.MetricsTable,
			Columns: []string{metricname.MetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metric.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mnc.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metricname.EventsTable,
			Columns: []string{metricname.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mnc.mutation.OwnersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metricname.OwnersTable,
			Columns: []string{metricname.OwnersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: metrictype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.metric_type_metric_names = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MetricNameCreateBulk is the builder for creating many MetricName entities in bulk.
type MetricNameCreateBulk struct {
	config
	builders []*MetricNameCreate
}

// Save creates the MetricName entities in the database.
func (mncb *MetricNameCreateBulk) Save(ctx context.Context) ([]*MetricName, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mncb.builders))
	nodes := make([]*MetricName, len(mncb.builders))
	mutators := make([]Mutator, len(mncb.builders))
	for i := range mncb.builders {
		func(i int, root context.Context) {
			builder := mncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MetricNameMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mncb *MetricNameCreateBulk) SaveX(ctx context.Context) []*MetricName {
	v, err := mncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mncb *MetricNameCreateBulk) Exec(ctx context.Context) error {
	_, err := mncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mncb *MetricNameCreateBulk) ExecX(ctx context.Context) {
	if err := mncb.Exec(ctx); err != nil {
		panic(err)
	}
}
