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
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/metriclabel"
)

// MetricLabelCreate is the builder for creating a MetricLabel entity.
type MetricLabelCreate struct {
	config
	mutation *MetricLabelMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mlc *MetricLabelCreate) SetCreatedAt(t time.Time) *MetricLabelCreate {
	mlc.mutation.SetCreatedAt(t)
	return mlc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mlc *MetricLabelCreate) SetNillableCreatedAt(t *time.Time) *MetricLabelCreate {
	if t != nil {
		mlc.SetCreatedAt(*t)
	}
	return mlc
}

// SetUpdatedAt sets the "updated_at" field.
func (mlc *MetricLabelCreate) SetUpdatedAt(t time.Time) *MetricLabelCreate {
	mlc.mutation.SetUpdatedAt(t)
	return mlc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mlc *MetricLabelCreate) SetNillableUpdatedAt(t *time.Time) *MetricLabelCreate {
	if t != nil {
		mlc.SetUpdatedAt(*t)
	}
	return mlc
}

// SetDeletedAt sets the "deleted_at" field.
func (mlc *MetricLabelCreate) SetDeletedAt(t time.Time) *MetricLabelCreate {
	mlc.mutation.SetDeletedAt(t)
	return mlc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mlc *MetricLabelCreate) SetNillableDeletedAt(t *time.Time) *MetricLabelCreate {
	if t != nil {
		mlc.SetDeletedAt(*t)
	}
	return mlc
}

// SetLabel sets the "label" field.
func (mlc *MetricLabelCreate) SetLabel(s string) *MetricLabelCreate {
	mlc.mutation.SetLabel(s)
	return mlc
}

// SetID sets the "id" field.
func (mlc *MetricLabelCreate) SetID(u uint) *MetricLabelCreate {
	mlc.mutation.SetID(u)
	return mlc
}

// AddMetricIDs adds the "metrics" edge to the Metric entity by IDs.
func (mlc *MetricLabelCreate) AddMetricIDs(ids ...int) *MetricLabelCreate {
	mlc.mutation.AddMetricIDs(ids...)
	return mlc
}

// AddMetrics adds the "metrics" edges to the Metric entity.
func (mlc *MetricLabelCreate) AddMetrics(m ...*Metric) *MetricLabelCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mlc.AddMetricIDs(ids...)
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (mlc *MetricLabelCreate) AddEventIDs(ids ...int) *MetricLabelCreate {
	mlc.mutation.AddEventIDs(ids...)
	return mlc
}

// AddEvents adds the "events" edges to the Event entity.
func (mlc *MetricLabelCreate) AddEvents(e ...*Event) *MetricLabelCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return mlc.AddEventIDs(ids...)
}

// Mutation returns the MetricLabelMutation object of the builder.
func (mlc *MetricLabelCreate) Mutation() *MetricLabelMutation {
	return mlc.mutation
}

// Save creates the MetricLabel in the database.
func (mlc *MetricLabelCreate) Save(ctx context.Context) (*MetricLabel, error) {
	var (
		err  error
		node *MetricLabel
	)
	mlc.defaults()
	if len(mlc.hooks) == 0 {
		if err = mlc.check(); err != nil {
			return nil, err
		}
		node, err = mlc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MetricLabelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mlc.check(); err != nil {
				return nil, err
			}
			mlc.mutation = mutation
			if node, err = mlc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mlc.hooks) - 1; i >= 0; i-- {
			if mlc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mlc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mlc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mlc *MetricLabelCreate) SaveX(ctx context.Context) *MetricLabel {
	v, err := mlc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mlc *MetricLabelCreate) Exec(ctx context.Context) error {
	_, err := mlc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mlc *MetricLabelCreate) ExecX(ctx context.Context) {
	if err := mlc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mlc *MetricLabelCreate) defaults() {
	if _, ok := mlc.mutation.CreatedAt(); !ok {
		v := metriclabel.DefaultCreatedAt()
		mlc.mutation.SetCreatedAt(v)
	}
	if _, ok := mlc.mutation.UpdatedAt(); !ok {
		v := metriclabel.DefaultUpdatedAt()
		mlc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mlc.mutation.DeletedAt(); !ok {
		v := metriclabel.DefaultDeletedAt()
		mlc.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mlc *MetricLabelCreate) check() error {
	if _, ok := mlc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := mlc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := mlc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "deleted_at"`)}
	}
	if _, ok := mlc.mutation.Label(); !ok {
		return &ValidationError{Name: "label", err: errors.New(`ent: missing required field "label"`)}
	}
	if v, ok := mlc.mutation.Label(); ok {
		if err := metriclabel.LabelValidator(v); err != nil {
			return &ValidationError{Name: "label", err: fmt.Errorf(`ent: validator failed for field "label": %w`, err)}
		}
	}
	return nil
}

func (mlc *MetricLabelCreate) sqlSave(ctx context.Context) (*MetricLabel, error) {
	_node, _spec := mlc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mlc.driver, _spec); err != nil {
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

func (mlc *MetricLabelCreate) createSpec() (*MetricLabel, *sqlgraph.CreateSpec) {
	var (
		_node = &MetricLabel{config: mlc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: metriclabel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: metriclabel.FieldID,
			},
		}
	)
	if id, ok := mlc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mlc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: metriclabel.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := mlc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: metriclabel.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := mlc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: metriclabel.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := mlc.mutation.Label(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: metriclabel.FieldLabel,
		})
		_node.Label = value
	}
	if nodes := mlc.mutation.MetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metriclabel.MetricsTable,
			Columns: []string{metriclabel.MetricsColumn},
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
	if nodes := mlc.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metriclabel.EventsTable,
			Columns: []string{metriclabel.EventsColumn},
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
	return _node, _spec
}

// MetricLabelCreateBulk is the builder for creating many MetricLabel entities in bulk.
type MetricLabelCreateBulk struct {
	config
	builders []*MetricLabelCreate
}

// Save creates the MetricLabel entities in the database.
func (mlcb *MetricLabelCreateBulk) Save(ctx context.Context) ([]*MetricLabel, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mlcb.builders))
	nodes := make([]*MetricLabel, len(mlcb.builders))
	mutators := make([]Mutator, len(mlcb.builders))
	for i := range mlcb.builders {
		func(i int, root context.Context) {
			builder := mlcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MetricLabelMutation)
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
					_, err = mutators[i+1].Mutate(root, mlcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mlcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mlcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mlcb *MetricLabelCreateBulk) SaveX(ctx context.Context) []*MetricLabel {
	v, err := mlcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mlcb *MetricLabelCreateBulk) Exec(ctx context.Context) error {
	_, err := mlcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mlcb *MetricLabelCreateBulk) ExecX(ctx context.Context) {
	if err := mlcb.Exec(ctx); err != nil {
		panic(err)
	}
}