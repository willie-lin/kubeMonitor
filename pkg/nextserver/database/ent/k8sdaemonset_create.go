// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/k8sdaemonset"
)

// K8sDaemonSetCreate is the builder for creating a K8sDaemonSet entity.
type K8sDaemonSetCreate struct {
	config
	mutation *K8sDaemonSetMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (kdsc *K8sDaemonSetCreate) SetCreatedAt(t time.Time) *K8sDaemonSetCreate {
	kdsc.mutation.SetCreatedAt(t)
	return kdsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (kdsc *K8sDaemonSetCreate) SetNillableCreatedAt(t *time.Time) *K8sDaemonSetCreate {
	if t != nil {
		kdsc.SetCreatedAt(*t)
	}
	return kdsc
}

// SetUpdatedAt sets the "updated_at" field.
func (kdsc *K8sDaemonSetCreate) SetUpdatedAt(t time.Time) *K8sDaemonSetCreate {
	kdsc.mutation.SetUpdatedAt(t)
	return kdsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (kdsc *K8sDaemonSetCreate) SetNillableUpdatedAt(t *time.Time) *K8sDaemonSetCreate {
	if t != nil {
		kdsc.SetUpdatedAt(*t)
	}
	return kdsc
}

// SetDeletedAt sets the "deleted_at" field.
func (kdsc *K8sDaemonSetCreate) SetDeletedAt(t time.Time) *K8sDaemonSetCreate {
	kdsc.mutation.SetDeletedAt(t)
	return kdsc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (kdsc *K8sDaemonSetCreate) SetNillableDeletedAt(t *time.Time) *K8sDaemonSetCreate {
	if t != nil {
		kdsc.SetDeletedAt(*t)
	}
	return kdsc
}

// SetName sets the "name" field.
func (kdsc *K8sDaemonSetCreate) SetName(s string) *K8sDaemonSetCreate {
	kdsc.mutation.SetName(s)
	return kdsc
}

// SetK8sClusterId sets the "k8sClusterId" field.
func (kdsc *K8sDaemonSetCreate) SetK8sClusterId(u uint) *K8sDaemonSetCreate {
	kdsc.mutation.SetK8sClusterId(u)
	return kdsc
}

// SetK8sNamespaceId sets the "k8sNamespaceId" field.
func (kdsc *K8sDaemonSetCreate) SetK8sNamespaceId(u uint) *K8sDaemonSetCreate {
	kdsc.mutation.SetK8sNamespaceId(u)
	return kdsc
}

// SetK8sObjectId sets the "k8sObjectId" field.
func (kdsc *K8sDaemonSetCreate) SetK8sObjectId(u uint) *K8sDaemonSetCreate {
	kdsc.mutation.SetK8sObjectId(u)
	return kdsc
}

// SetID sets the "id" field.
func (kdsc *K8sDaemonSetCreate) SetID(u uint) *K8sDaemonSetCreate {
	kdsc.mutation.SetID(u)
	return kdsc
}

// Mutation returns the K8sDaemonSetMutation object of the builder.
func (kdsc *K8sDaemonSetCreate) Mutation() *K8sDaemonSetMutation {
	return kdsc.mutation
}

// Save creates the K8sDaemonSet in the database.
func (kdsc *K8sDaemonSetCreate) Save(ctx context.Context) (*K8sDaemonSet, error) {
	var (
		err  error
		node *K8sDaemonSet
	)
	kdsc.defaults()
	if len(kdsc.hooks) == 0 {
		if err = kdsc.check(); err != nil {
			return nil, err
		}
		node, err = kdsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*K8sDaemonSetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kdsc.check(); err != nil {
				return nil, err
			}
			kdsc.mutation = mutation
			if node, err = kdsc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(kdsc.hooks) - 1; i >= 0; i-- {
			if kdsc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kdsc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kdsc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (kdsc *K8sDaemonSetCreate) SaveX(ctx context.Context) *K8sDaemonSet {
	v, err := kdsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kdsc *K8sDaemonSetCreate) Exec(ctx context.Context) error {
	_, err := kdsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kdsc *K8sDaemonSetCreate) ExecX(ctx context.Context) {
	if err := kdsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kdsc *K8sDaemonSetCreate) defaults() {
	if _, ok := kdsc.mutation.CreatedAt(); !ok {
		v := k8sdaemonset.DefaultCreatedAt()
		kdsc.mutation.SetCreatedAt(v)
	}
	if _, ok := kdsc.mutation.UpdatedAt(); !ok {
		v := k8sdaemonset.DefaultUpdatedAt()
		kdsc.mutation.SetUpdatedAt(v)
	}
	if _, ok := kdsc.mutation.DeletedAt(); !ok {
		v := k8sdaemonset.DefaultDeletedAt()
		kdsc.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kdsc *K8sDaemonSetCreate) check() error {
	if _, ok := kdsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := kdsc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := kdsc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "deleted_at"`)}
	}
	if _, ok := kdsc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := kdsc.mutation.Name(); ok {
		if err := k8sdaemonset.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := kdsc.mutation.K8sClusterId(); !ok {
		return &ValidationError{Name: "k8sClusterId", err: errors.New(`ent: missing required field "k8sClusterId"`)}
	}
	if _, ok := kdsc.mutation.K8sNamespaceId(); !ok {
		return &ValidationError{Name: "k8sNamespaceId", err: errors.New(`ent: missing required field "k8sNamespaceId"`)}
	}
	if _, ok := kdsc.mutation.K8sObjectId(); !ok {
		return &ValidationError{Name: "k8sObjectId", err: errors.New(`ent: missing required field "k8sObjectId"`)}
	}
	return nil
}

func (kdsc *K8sDaemonSetCreate) sqlSave(ctx context.Context) (*K8sDaemonSet, error) {
	_node, _spec := kdsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kdsc.driver, _spec); err != nil {
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

func (kdsc *K8sDaemonSetCreate) createSpec() (*K8sDaemonSet, *sqlgraph.CreateSpec) {
	var (
		_node = &K8sDaemonSet{config: kdsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: k8sdaemonset.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: k8sdaemonset.FieldID,
			},
		}
	)
	if id, ok := kdsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := kdsc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8sdaemonset.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := kdsc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8sdaemonset.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := kdsc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8sdaemonset.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := kdsc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: k8sdaemonset.FieldName,
		})
		_node.Name = value
	}
	if value, ok := kdsc.mutation.K8sClusterId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sdaemonset.FieldK8sClusterId,
		})
		_node.K8sClusterId = value
	}
	if value, ok := kdsc.mutation.K8sNamespaceId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sdaemonset.FieldK8sNamespaceId,
		})
		_node.K8sNamespaceId = value
	}
	if value, ok := kdsc.mutation.K8sObjectId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sdaemonset.FieldK8sObjectId,
		})
		_node.K8sObjectId = value
	}
	return _node, _spec
}

// K8sDaemonSetCreateBulk is the builder for creating many K8sDaemonSet entities in bulk.
type K8sDaemonSetCreateBulk struct {
	config
	builders []*K8sDaemonSetCreate
}

// Save creates the K8sDaemonSet entities in the database.
func (kdscb *K8sDaemonSetCreateBulk) Save(ctx context.Context) ([]*K8sDaemonSet, error) {
	specs := make([]*sqlgraph.CreateSpec, len(kdscb.builders))
	nodes := make([]*K8sDaemonSet, len(kdscb.builders))
	mutators := make([]Mutator, len(kdscb.builders))
	for i := range kdscb.builders {
		func(i int, root context.Context) {
			builder := kdscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*K8sDaemonSetMutation)
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
					_, err = mutators[i+1].Mutate(root, kdscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kdscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, kdscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kdscb *K8sDaemonSetCreateBulk) SaveX(ctx context.Context) []*K8sDaemonSet {
	v, err := kdscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kdscb *K8sDaemonSetCreateBulk) Exec(ctx context.Context) error {
	_, err := kdscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kdscb *K8sDaemonSetCreateBulk) ExecX(ctx context.Context) {
	if err := kdscb.Exec(ctx); err != nil {
		panic(err)
	}
}