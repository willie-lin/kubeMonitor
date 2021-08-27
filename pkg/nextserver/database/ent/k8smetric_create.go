// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/k8smetric"
)

// K8sMetricCreate is the builder for creating a K8sMetric entity.
type K8sMetricCreate struct {
	config
	mutation *K8sMetricMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (kmc *K8sMetricCreate) SetCreatedAt(t time.Time) *K8sMetricCreate {
	kmc.mutation.SetCreatedAt(t)
	return kmc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (kmc *K8sMetricCreate) SetNillableCreatedAt(t *time.Time) *K8sMetricCreate {
	if t != nil {
		kmc.SetCreatedAt(*t)
	}
	return kmc
}

// SetUpdatedAt sets the "updated_at" field.
func (kmc *K8sMetricCreate) SetUpdatedAt(t time.Time) *K8sMetricCreate {
	kmc.mutation.SetUpdatedAt(t)
	return kmc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (kmc *K8sMetricCreate) SetNillableUpdatedAt(t *time.Time) *K8sMetricCreate {
	if t != nil {
		kmc.SetUpdatedAt(*t)
	}
	return kmc
}

// SetDeletedAt sets the "deleted_at" field.
func (kmc *K8sMetricCreate) SetDeletedAt(t time.Time) *K8sMetricCreate {
	kmc.mutation.SetDeletedAt(t)
	return kmc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (kmc *K8sMetricCreate) SetNillableDeletedAt(t *time.Time) *K8sMetricCreate {
	if t != nil {
		kmc.SetDeletedAt(*t)
	}
	return kmc
}

// SetTs sets the "ts" field.
func (kmc *K8sMetricCreate) SetTs(t time.Time) *K8sMetricCreate {
	kmc.mutation.SetTs(t)
	return kmc
}

// SetValue sets the "value" field.
func (kmc *K8sMetricCreate) SetValue(f float64) *K8sMetricCreate {
	kmc.mutation.SetValue(f)
	return kmc
}

// SetEndpointId sets the "endpointId" field.
func (kmc *K8sMetricCreate) SetEndpointId(u uint) *K8sMetricCreate {
	kmc.mutation.SetEndpointId(u)
	return kmc
}

// SetTypeId sets the "typeId" field.
func (kmc *K8sMetricCreate) SetTypeId(u uint) *K8sMetricCreate {
	kmc.mutation.SetTypeId(u)
	return kmc
}

// SetNameId sets the "nameId" field.
func (kmc *K8sMetricCreate) SetNameId(u uint) *K8sMetricCreate {
	kmc.mutation.SetNameId(u)
	return kmc
}

// SetLabelId sets the "labelId" field.
func (kmc *K8sMetricCreate) SetLabelId(u uint) *K8sMetricCreate {
	kmc.mutation.SetLabelId(u)
	return kmc
}

// SetK8sClusterId sets the "k8sClusterId" field.
func (kmc *K8sMetricCreate) SetK8sClusterId(u uint) *K8sMetricCreate {
	kmc.mutation.SetK8sClusterId(u)
	return kmc
}

// SetK8sNodeId sets the "k8sNodeId" field.
func (kmc *K8sMetricCreate) SetK8sNodeId(u uint) *K8sMetricCreate {
	kmc.mutation.SetK8sNodeId(u)
	return kmc
}

// SetK8sNamespaceId sets the "k8sNamespaceId" field.
func (kmc *K8sMetricCreate) SetK8sNamespaceId(u uint) *K8sMetricCreate {
	kmc.mutation.SetK8sNamespaceId(u)
	return kmc
}

// SetK8sContainerId sets the "k8sContainerId" field.
func (kmc *K8sMetricCreate) SetK8sContainerId(u uint) *K8sMetricCreate {
	kmc.mutation.SetK8sContainerId(u)
	return kmc
}

// Mutation returns the K8sMetricMutation object of the builder.
func (kmc *K8sMetricCreate) Mutation() *K8sMetricMutation {
	return kmc.mutation
}

// Save creates the K8sMetric in the database.
func (kmc *K8sMetricCreate) Save(ctx context.Context) (*K8sMetric, error) {
	var (
		err  error
		node *K8sMetric
	)
	kmc.defaults()
	if len(kmc.hooks) == 0 {
		if err = kmc.check(); err != nil {
			return nil, err
		}
		node, err = kmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*K8sMetricMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kmc.check(); err != nil {
				return nil, err
			}
			kmc.mutation = mutation
			if node, err = kmc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(kmc.hooks) - 1; i >= 0; i-- {
			if kmc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kmc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kmc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (kmc *K8sMetricCreate) SaveX(ctx context.Context) *K8sMetric {
	v, err := kmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kmc *K8sMetricCreate) Exec(ctx context.Context) error {
	_, err := kmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kmc *K8sMetricCreate) ExecX(ctx context.Context) {
	if err := kmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kmc *K8sMetricCreate) defaults() {
	if _, ok := kmc.mutation.CreatedAt(); !ok {
		v := k8smetric.DefaultCreatedAt()
		kmc.mutation.SetCreatedAt(v)
	}
	if _, ok := kmc.mutation.UpdatedAt(); !ok {
		v := k8smetric.DefaultUpdatedAt()
		kmc.mutation.SetUpdatedAt(v)
	}
	if _, ok := kmc.mutation.DeletedAt(); !ok {
		v := k8smetric.DefaultDeletedAt()
		kmc.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kmc *K8sMetricCreate) check() error {
	if _, ok := kmc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := kmc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := kmc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "deleted_at"`)}
	}
	if _, ok := kmc.mutation.Ts(); !ok {
		return &ValidationError{Name: "ts", err: errors.New(`ent: missing required field "ts"`)}
	}
	if _, ok := kmc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "value"`)}
	}
	if _, ok := kmc.mutation.EndpointId(); !ok {
		return &ValidationError{Name: "endpointId", err: errors.New(`ent: missing required field "endpointId"`)}
	}
	if _, ok := kmc.mutation.TypeId(); !ok {
		return &ValidationError{Name: "typeId", err: errors.New(`ent: missing required field "typeId"`)}
	}
	if _, ok := kmc.mutation.NameId(); !ok {
		return &ValidationError{Name: "nameId", err: errors.New(`ent: missing required field "nameId"`)}
	}
	if _, ok := kmc.mutation.LabelId(); !ok {
		return &ValidationError{Name: "labelId", err: errors.New(`ent: missing required field "labelId"`)}
	}
	if _, ok := kmc.mutation.K8sClusterId(); !ok {
		return &ValidationError{Name: "k8sClusterId", err: errors.New(`ent: missing required field "k8sClusterId"`)}
	}
	if _, ok := kmc.mutation.K8sNodeId(); !ok {
		return &ValidationError{Name: "k8sNodeId", err: errors.New(`ent: missing required field "k8sNodeId"`)}
	}
	if _, ok := kmc.mutation.K8sNamespaceId(); !ok {
		return &ValidationError{Name: "k8sNamespaceId", err: errors.New(`ent: missing required field "k8sNamespaceId"`)}
	}
	if _, ok := kmc.mutation.K8sContainerId(); !ok {
		return &ValidationError{Name: "k8sContainerId", err: errors.New(`ent: missing required field "k8sContainerId"`)}
	}
	return nil
}

func (kmc *K8sMetricCreate) sqlSave(ctx context.Context) (*K8sMetric, error) {
	_node, _spec := kmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (kmc *K8sMetricCreate) createSpec() (*K8sMetric, *sqlgraph.CreateSpec) {
	var (
		_node = &K8sMetric{config: kmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: k8smetric.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: k8smetric.FieldID,
			},
		}
	)
	if value, ok := kmc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8smetric.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := kmc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8smetric.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := kmc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8smetric.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := kmc.mutation.Ts(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8smetric.FieldTs,
		})
		_node.Ts = value
	}
	if value, ok := kmc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: k8smetric.FieldValue,
		})
		_node.Value = value
	}
	if value, ok := kmc.mutation.EndpointId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8smetric.FieldEndpointId,
		})
		_node.EndpointId = value
	}
	if value, ok := kmc.mutation.TypeId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8smetric.FieldTypeId,
		})
		_node.TypeId = value
	}
	if value, ok := kmc.mutation.NameId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8smetric.FieldNameId,
		})
		_node.NameId = value
	}
	if value, ok := kmc.mutation.LabelId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8smetric.FieldLabelId,
		})
		_node.LabelId = value
	}
	if value, ok := kmc.mutation.K8sClusterId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8smetric.FieldK8sClusterId,
		})
		_node.K8sClusterId = value
	}
	if value, ok := kmc.mutation.K8sNodeId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8smetric.FieldK8sNodeId,
		})
		_node.K8sNodeId = value
	}
	if value, ok := kmc.mutation.K8sNamespaceId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8smetric.FieldK8sNamespaceId,
		})
		_node.K8sNamespaceId = value
	}
	if value, ok := kmc.mutation.K8sContainerId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8smetric.FieldK8sContainerId,
		})
		_node.K8sContainerId = value
	}
	return _node, _spec
}

// K8sMetricCreateBulk is the builder for creating many K8sMetric entities in bulk.
type K8sMetricCreateBulk struct {
	config
	builders []*K8sMetricCreate
}

// Save creates the K8sMetric entities in the database.
func (kmcb *K8sMetricCreateBulk) Save(ctx context.Context) ([]*K8sMetric, error) {
	specs := make([]*sqlgraph.CreateSpec, len(kmcb.builders))
	nodes := make([]*K8sMetric, len(kmcb.builders))
	mutators := make([]Mutator, len(kmcb.builders))
	for i := range kmcb.builders {
		func(i int, root context.Context) {
			builder := kmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*K8sMetricMutation)
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
					_, err = mutators[i+1].Mutate(root, kmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, kmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kmcb *K8sMetricCreateBulk) SaveX(ctx context.Context) []*K8sMetric {
	v, err := kmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kmcb *K8sMetricCreateBulk) Exec(ctx context.Context) error {
	_, err := kmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kmcb *K8sMetricCreateBulk) ExecX(ctx context.Context) {
	if err := kmcb.Exec(ctx); err != nil {
		panic(err)
	}
}