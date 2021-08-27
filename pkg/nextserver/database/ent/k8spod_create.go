// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/k8spod"
)

// K8sPodCreate is the builder for creating a K8sPod entity.
type K8sPodCreate struct {
	config
	mutation *K8sPodMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (kpc *K8sPodCreate) SetCreatedAt(t time.Time) *K8sPodCreate {
	kpc.mutation.SetCreatedAt(t)
	return kpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (kpc *K8sPodCreate) SetNillableCreatedAt(t *time.Time) *K8sPodCreate {
	if t != nil {
		kpc.SetCreatedAt(*t)
	}
	return kpc
}

// SetUpdatedAt sets the "updated_at" field.
func (kpc *K8sPodCreate) SetUpdatedAt(t time.Time) *K8sPodCreate {
	kpc.mutation.SetUpdatedAt(t)
	return kpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (kpc *K8sPodCreate) SetNillableUpdatedAt(t *time.Time) *K8sPodCreate {
	if t != nil {
		kpc.SetUpdatedAt(*t)
	}
	return kpc
}

// SetDeletedAt sets the "deleted_at" field.
func (kpc *K8sPodCreate) SetDeletedAt(t time.Time) *K8sPodCreate {
	kpc.mutation.SetDeletedAt(t)
	return kpc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (kpc *K8sPodCreate) SetNillableDeletedAt(t *time.Time) *K8sPodCreate {
	if t != nil {
		kpc.SetDeletedAt(*t)
	}
	return kpc
}

// SetName sets the "name" field.
func (kpc *K8sPodCreate) SetName(s string) *K8sPodCreate {
	kpc.mutation.SetName(s)
	return kpc
}

// SetQos sets the "qos" field.
func (kpc *K8sPodCreate) SetQos(s string) *K8sPodCreate {
	kpc.mutation.SetQos(s)
	return kpc
}

// SetK8sClusterId sets the "k8sClusterId" field.
func (kpc *K8sPodCreate) SetK8sClusterId(u uint) *K8sPodCreate {
	kpc.mutation.SetK8sClusterId(u)
	return kpc
}

// SetK8sNamespaceId sets the "k8sNamespaceId" field.
func (kpc *K8sPodCreate) SetK8sNamespaceId(u uint) *K8sPodCreate {
	kpc.mutation.SetK8sNamespaceId(u)
	return kpc
}

// SetK8sObjectId sets the "k8sObjectId" field.
func (kpc *K8sPodCreate) SetK8sObjectId(u uint) *K8sPodCreate {
	kpc.mutation.SetK8sObjectId(u)
	return kpc
}

// SetID sets the "id" field.
func (kpc *K8sPodCreate) SetID(u uint) *K8sPodCreate {
	kpc.mutation.SetID(u)
	return kpc
}

// Mutation returns the K8sPodMutation object of the builder.
func (kpc *K8sPodCreate) Mutation() *K8sPodMutation {
	return kpc.mutation
}

// Save creates the K8sPod in the database.
func (kpc *K8sPodCreate) Save(ctx context.Context) (*K8sPod, error) {
	var (
		err  error
		node *K8sPod
	)
	kpc.defaults()
	if len(kpc.hooks) == 0 {
		if err = kpc.check(); err != nil {
			return nil, err
		}
		node, err = kpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*K8sPodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kpc.check(); err != nil {
				return nil, err
			}
			kpc.mutation = mutation
			if node, err = kpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(kpc.hooks) - 1; i >= 0; i-- {
			if kpc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kpc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kpc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (kpc *K8sPodCreate) SaveX(ctx context.Context) *K8sPod {
	v, err := kpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kpc *K8sPodCreate) Exec(ctx context.Context) error {
	_, err := kpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kpc *K8sPodCreate) ExecX(ctx context.Context) {
	if err := kpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kpc *K8sPodCreate) defaults() {
	if _, ok := kpc.mutation.CreatedAt(); !ok {
		v := k8spod.DefaultCreatedAt()
		kpc.mutation.SetCreatedAt(v)
	}
	if _, ok := kpc.mutation.UpdatedAt(); !ok {
		v := k8spod.DefaultUpdatedAt()
		kpc.mutation.SetUpdatedAt(v)
	}
	if _, ok := kpc.mutation.DeletedAt(); !ok {
		v := k8spod.DefaultDeletedAt()
		kpc.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kpc *K8sPodCreate) check() error {
	if _, ok := kpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := kpc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := kpc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "deleted_at"`)}
	}
	if _, ok := kpc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := kpc.mutation.Name(); ok {
		if err := k8spod.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := kpc.mutation.Qos(); !ok {
		return &ValidationError{Name: "qos", err: errors.New(`ent: missing required field "qos"`)}
	}
	if v, ok := kpc.mutation.Qos(); ok {
		if err := k8spod.QosValidator(v); err != nil {
			return &ValidationError{Name: "qos", err: fmt.Errorf(`ent: validator failed for field "qos": %w`, err)}
		}
	}
	if _, ok := kpc.mutation.K8sClusterId(); !ok {
		return &ValidationError{Name: "k8sClusterId", err: errors.New(`ent: missing required field "k8sClusterId"`)}
	}
	if _, ok := kpc.mutation.K8sNamespaceId(); !ok {
		return &ValidationError{Name: "k8sNamespaceId", err: errors.New(`ent: missing required field "k8sNamespaceId"`)}
	}
	if _, ok := kpc.mutation.K8sObjectId(); !ok {
		return &ValidationError{Name: "k8sObjectId", err: errors.New(`ent: missing required field "k8sObjectId"`)}
	}
	return nil
}

func (kpc *K8sPodCreate) sqlSave(ctx context.Context) (*K8sPod, error) {
	_node, _spec := kpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kpc.driver, _spec); err != nil {
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

func (kpc *K8sPodCreate) createSpec() (*K8sPod, *sqlgraph.CreateSpec) {
	var (
		_node = &K8sPod{config: kpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: k8spod.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: k8spod.FieldID,
			},
		}
	)
	if id, ok := kpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := kpc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8spod.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := kpc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8spod.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := kpc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8spod.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := kpc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: k8spod.FieldName,
		})
		_node.Name = value
	}
	if value, ok := kpc.mutation.Qos(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: k8spod.FieldQos,
		})
		_node.Qos = value
	}
	if value, ok := kpc.mutation.K8sClusterId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8spod.FieldK8sClusterId,
		})
		_node.K8sClusterId = value
	}
	if value, ok := kpc.mutation.K8sNamespaceId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8spod.FieldK8sNamespaceId,
		})
		_node.K8sNamespaceId = value
	}
	if value, ok := kpc.mutation.K8sObjectId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8spod.FieldK8sObjectId,
		})
		_node.K8sObjectId = value
	}
	return _node, _spec
}

// K8sPodCreateBulk is the builder for creating many K8sPod entities in bulk.
type K8sPodCreateBulk struct {
	config
	builders []*K8sPodCreate
}

// Save creates the K8sPod entities in the database.
func (kpcb *K8sPodCreateBulk) Save(ctx context.Context) ([]*K8sPod, error) {
	specs := make([]*sqlgraph.CreateSpec, len(kpcb.builders))
	nodes := make([]*K8sPod, len(kpcb.builders))
	mutators := make([]Mutator, len(kpcb.builders))
	for i := range kpcb.builders {
		func(i int, root context.Context) {
			builder := kpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*K8sPodMutation)
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
					_, err = mutators[i+1].Mutate(root, kpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, kpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kpcb *K8sPodCreateBulk) SaveX(ctx context.Context) []*K8sPod {
	v, err := kpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kpcb *K8sPodCreateBulk) Exec(ctx context.Context) error {
	_, err := kpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kpcb *K8sPodCreateBulk) ExecX(ctx context.Context) {
	if err := kpcb.Exec(ctx); err != nil {
		panic(err)
	}
}