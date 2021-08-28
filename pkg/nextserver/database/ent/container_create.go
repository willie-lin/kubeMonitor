// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/container"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/node"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/proces"
)

// ContainerCreate is the builder for creating a Container entity.
type ContainerCreate struct {
	config
	mutation *ContainerMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *ContainerCreate) SetCreatedAt(t time.Time) *ContainerCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *ContainerCreate) SetNillableCreatedAt(t *time.Time) *ContainerCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *ContainerCreate) SetUpdatedAt(t time.Time) *ContainerCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *ContainerCreate) SetNillableUpdatedAt(t *time.Time) *ContainerCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *ContainerCreate) SetDeletedAt(t time.Time) *ContainerCreate {
	cc.mutation.SetDeletedAt(t)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *ContainerCreate) SetNillableDeletedAt(t *time.Time) *ContainerCreate {
	if t != nil {
		cc.SetDeletedAt(*t)
	}
	return cc
}

// SetType sets the "type" field.
func (cc *ContainerCreate) SetType(s string) *ContainerCreate {
	cc.mutation.SetType(s)
	return cc
}

// SetContainerId sets the "containerId" field.
func (cc *ContainerCreate) SetContainerId(s string) *ContainerCreate {
	cc.mutation.SetContainerId(s)
	return cc
}

// SetName sets the "name" field.
func (cc *ContainerCreate) SetName(s string) *ContainerCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetImage sets the "image" field.
func (cc *ContainerCreate) SetImage(s string) *ContainerCreate {
	cc.mutation.SetImage(s)
	return cc
}

// SetInfo sets the "info" field.
func (cc *ContainerCreate) SetInfo(s []string) *ContainerCreate {
	cc.mutation.SetInfo(s)
	return cc
}

// SetClusterId sets the "clusterId" field.
func (cc *ContainerCreate) SetClusterId(u uint) *ContainerCreate {
	cc.mutation.SetClusterId(u)
	return cc
}

// SetNodeId sets the "nodeId" field.
func (cc *ContainerCreate) SetNodeId(u uint) *ContainerCreate {
	cc.mutation.SetNodeId(u)
	return cc
}

// SetID sets the "id" field.
func (cc *ContainerCreate) SetID(u uint) *ContainerCreate {
	cc.mutation.SetID(u)
	return cc
}

// AddProcesIDs adds the "process" edge to the Proces entity by IDs.
func (cc *ContainerCreate) AddProcesIDs(ids ...uint) *ContainerCreate {
	cc.mutation.AddProcesIDs(ids...)
	return cc
}

// AddProcess adds the "process" edges to the Proces entity.
func (cc *ContainerCreate) AddProcess(p ...*Proces) *ContainerCreate {
	ids := make([]uint, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddProcesIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the Node entity by ID.
func (cc *ContainerCreate) SetOwnerID(id uint) *ContainerCreate {
	cc.mutation.SetOwnerID(id)
	return cc
}

// SetNillableOwnerID sets the "owner" edge to the Node entity by ID if the given value is not nil.
func (cc *ContainerCreate) SetNillableOwnerID(id *uint) *ContainerCreate {
	if id != nil {
		cc = cc.SetOwnerID(*id)
	}
	return cc
}

// SetOwner sets the "owner" edge to the Node entity.
func (cc *ContainerCreate) SetOwner(n *Node) *ContainerCreate {
	return cc.SetOwnerID(n.ID)
}

// Mutation returns the ContainerMutation object of the builder.
func (cc *ContainerCreate) Mutation() *ContainerMutation {
	return cc.mutation
}

// Save creates the Container in the database.
func (cc *ContainerCreate) Save(ctx context.Context) (*Container, error) {
	var (
		err  error
		node *Container
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContainerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ContainerCreate) SaveX(ctx context.Context) *Container {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ContainerCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ContainerCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ContainerCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := container.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := container.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.DeletedAt(); !ok {
		v := container.DefaultDeletedAt()
		cc.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ContainerCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := cc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "deleted_at"`)}
	}
	if _, ok := cc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "type"`)}
	}
	if _, ok := cc.mutation.ContainerId(); !ok {
		return &ValidationError{Name: "containerId", err: errors.New(`ent: missing required field "containerId"`)}
	}
	if v, ok := cc.mutation.ContainerId(); ok {
		if err := container.ContainerIdValidator(v); err != nil {
			return &ValidationError{Name: "containerId", err: fmt.Errorf(`ent: validator failed for field "containerId": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := container.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Image(); !ok {
		return &ValidationError{Name: "image", err: errors.New(`ent: missing required field "image"`)}
	}
	if v, ok := cc.mutation.Image(); ok {
		if err := container.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf(`ent: validator failed for field "image": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Info(); !ok {
		return &ValidationError{Name: "info", err: errors.New(`ent: missing required field "info"`)}
	}
	if _, ok := cc.mutation.ClusterId(); !ok {
		return &ValidationError{Name: "clusterId", err: errors.New(`ent: missing required field "clusterId"`)}
	}
	if _, ok := cc.mutation.NodeId(); !ok {
		return &ValidationError{Name: "nodeId", err: errors.New(`ent: missing required field "nodeId"`)}
	}
	return nil
}

func (cc *ContainerCreate) sqlSave(ctx context.Context) (*Container, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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

func (cc *ContainerCreate) createSpec() (*Container, *sqlgraph.CreateSpec) {
	var (
		_node = &Container{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: container.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: container.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: container.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: container.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: container.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := cc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: container.FieldType,
		})
		_node.Type = value
	}
	if value, ok := cc.mutation.ContainerId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: container.FieldContainerId,
		})
		_node.ContainerId = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: container.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.Image(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: container.FieldImage,
		})
		_node.Image = value
	}
	if value, ok := cc.mutation.Info(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: container.FieldInfo,
		})
		_node.Info = value
	}
	if value, ok := cc.mutation.ClusterId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: container.FieldClusterId,
		})
		_node.ClusterId = value
	}
	if value, ok := cc.mutation.NodeId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: container.FieldNodeId,
		})
		_node.NodeId = value
	}
	if nodes := cc.mutation.ProcessIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   container.ProcessTable,
			Columns: []string{container.ProcessColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: proces.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   container.OwnerTable,
			Columns: []string{container.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: node.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.node_containers = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ContainerCreateBulk is the builder for creating many Container entities in bulk.
type ContainerCreateBulk struct {
	config
	builders []*ContainerCreate
}

// Save creates the Container entities in the database.
func (ccb *ContainerCreateBulk) Save(ctx context.Context) ([]*Container, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Container, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ContainerMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ContainerCreateBulk) SaveX(ctx context.Context) []*Container {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ContainerCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ContainerCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
