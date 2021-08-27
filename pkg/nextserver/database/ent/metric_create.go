// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/metric"
)

// MetricCreate is the builder for creating a Metric entity.
type MetricCreate struct {
	config
	mutation *MetricMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mc *MetricCreate) SetCreatedAt(t time.Time) *MetricCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MetricCreate) SetNillableCreatedAt(t *time.Time) *MetricCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MetricCreate) SetUpdatedAt(t time.Time) *MetricCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MetricCreate) SetNillableUpdatedAt(t *time.Time) *MetricCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetDeletedAt sets the "deleted_at" field.
func (mc *MetricCreate) SetDeletedAt(t time.Time) *MetricCreate {
	mc.mutation.SetDeletedAt(t)
	return mc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mc *MetricCreate) SetNillableDeletedAt(t *time.Time) *MetricCreate {
	if t != nil {
		mc.SetDeletedAt(*t)
	}
	return mc
}

// SetTs sets the "ts" field.
func (mc *MetricCreate) SetTs(t time.Time) *MetricCreate {
	mc.mutation.SetTs(t)
	return mc
}

// SetValue sets the "value" field.
func (mc *MetricCreate) SetValue(f float64) *MetricCreate {
	mc.mutation.SetValue(f)
	return mc
}

// SetEndpointId sets the "endpointId" field.
func (mc *MetricCreate) SetEndpointId(u uint) *MetricCreate {
	mc.mutation.SetEndpointId(u)
	return mc
}

// SetTypeId sets the "typeId" field.
func (mc *MetricCreate) SetTypeId(u uint) *MetricCreate {
	mc.mutation.SetTypeId(u)
	return mc
}

// SetNameId sets the "nameId" field.
func (mc *MetricCreate) SetNameId(u uint) *MetricCreate {
	mc.mutation.SetNameId(u)
	return mc
}

// SetLabelId sets the "labelId" field.
func (mc *MetricCreate) SetLabelId(u uint) *MetricCreate {
	mc.mutation.SetLabelId(u)
	return mc
}

// SetClusterId sets the "clusterId" field.
func (mc *MetricCreate) SetClusterId(u uint) *MetricCreate {
	mc.mutation.SetClusterId(u)
	return mc
}

// SetNodeId sets the "nodeId" field.
func (mc *MetricCreate) SetNodeId(u uint) *MetricCreate {
	mc.mutation.SetNodeId(u)
	return mc
}

// SetProcesId sets the "procesId" field.
func (mc *MetricCreate) SetProcesId(u uint) *MetricCreate {
	mc.mutation.SetProcesId(u)
	return mc
}

// SetContainerId sets the "containerId" field.
func (mc *MetricCreate) SetContainerId(u uint) *MetricCreate {
	mc.mutation.SetContainerId(u)
	return mc
}

// Mutation returns the MetricMutation object of the builder.
func (mc *MetricCreate) Mutation() *MetricMutation {
	return mc.mutation
}

// Save creates the Metric in the database.
func (mc *MetricCreate) Save(ctx context.Context) (*Metric, error) {
	var (
		err  error
		node *Metric
	)
	mc.defaults()
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MetricMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MetricCreate) SaveX(ctx context.Context) *Metric {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MetricCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MetricCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MetricCreate) defaults() {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := metric.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := metric.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mc.mutation.DeletedAt(); !ok {
		v := metric.DefaultDeletedAt()
		mc.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MetricCreate) check() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := mc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "deleted_at"`)}
	}
	if _, ok := mc.mutation.Ts(); !ok {
		return &ValidationError{Name: "ts", err: errors.New(`ent: missing required field "ts"`)}
	}
	if _, ok := mc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "value"`)}
	}
	if _, ok := mc.mutation.EndpointId(); !ok {
		return &ValidationError{Name: "endpointId", err: errors.New(`ent: missing required field "endpointId"`)}
	}
	if _, ok := mc.mutation.TypeId(); !ok {
		return &ValidationError{Name: "typeId", err: errors.New(`ent: missing required field "typeId"`)}
	}
	if _, ok := mc.mutation.NameId(); !ok {
		return &ValidationError{Name: "nameId", err: errors.New(`ent: missing required field "nameId"`)}
	}
	if _, ok := mc.mutation.LabelId(); !ok {
		return &ValidationError{Name: "labelId", err: errors.New(`ent: missing required field "labelId"`)}
	}
	if _, ok := mc.mutation.ClusterId(); !ok {
		return &ValidationError{Name: "clusterId", err: errors.New(`ent: missing required field "clusterId"`)}
	}
	if _, ok := mc.mutation.NodeId(); !ok {
		return &ValidationError{Name: "nodeId", err: errors.New(`ent: missing required field "nodeId"`)}
	}
	if _, ok := mc.mutation.ProcesId(); !ok {
		return &ValidationError{Name: "procesId", err: errors.New(`ent: missing required field "procesId"`)}
	}
	if _, ok := mc.mutation.ContainerId(); !ok {
		return &ValidationError{Name: "containerId", err: errors.New(`ent: missing required field "containerId"`)}
	}
	return nil
}

func (mc *MetricCreate) sqlSave(ctx context.Context) (*Metric, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mc *MetricCreate) createSpec() (*Metric, *sqlgraph.CreateSpec) {
	var (
		_node = &Metric{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: metric.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: metric.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: metric.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: metric.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: metric.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := mc.mutation.Ts(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: metric.FieldTs,
		})
		_node.Ts = value
	}
	if value, ok := mc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: metric.FieldValue,
		})
		_node.Value = value
	}
	if value, ok := mc.mutation.EndpointId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldEndpointId,
		})
		_node.EndpointId = value
	}
	if value, ok := mc.mutation.TypeId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldTypeId,
		})
		_node.TypeId = value
	}
	if value, ok := mc.mutation.NameId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldNameId,
		})
		_node.NameId = value
	}
	if value, ok := mc.mutation.LabelId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldLabelId,
		})
		_node.LabelId = value
	}
	if value, ok := mc.mutation.ClusterId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldClusterId,
		})
		_node.ClusterId = value
	}
	if value, ok := mc.mutation.NodeId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldNodeId,
		})
		_node.NodeId = value
	}
	if value, ok := mc.mutation.ProcesId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldProcesId,
		})
		_node.ProcesId = value
	}
	if value, ok := mc.mutation.ContainerId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldContainerId,
		})
		_node.ContainerId = value
	}
	return _node, _spec
}

// MetricCreateBulk is the builder for creating many Metric entities in bulk.
type MetricCreateBulk struct {
	config
	builders []*MetricCreate
}

// Save creates the Metric entities in the database.
func (mcb *MetricCreateBulk) Save(ctx context.Context) ([]*Metric, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Metric, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MetricMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MetricCreateBulk) SaveX(ctx context.Context) []*Metric {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MetricCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MetricCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}