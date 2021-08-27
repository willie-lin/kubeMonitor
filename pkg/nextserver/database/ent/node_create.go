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

// NodeCreate is the builder for creating a Node entity.
type NodeCreate struct {
	config
	mutation *NodeMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (nc *NodeCreate) SetCreatedAt(t time.Time) *NodeCreate {
	nc.mutation.SetCreatedAt(t)
	return nc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nc *NodeCreate) SetNillableCreatedAt(t *time.Time) *NodeCreate {
	if t != nil {
		nc.SetCreatedAt(*t)
	}
	return nc
}

// SetUpdatedAt sets the "updated_at" field.
func (nc *NodeCreate) SetUpdatedAt(t time.Time) *NodeCreate {
	nc.mutation.SetUpdatedAt(t)
	return nc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (nc *NodeCreate) SetNillableUpdatedAt(t *time.Time) *NodeCreate {
	if t != nil {
		nc.SetUpdatedAt(*t)
	}
	return nc
}

// SetDeletedAt sets the "deleted_at" field.
func (nc *NodeCreate) SetDeletedAt(t time.Time) *NodeCreate {
	nc.mutation.SetDeletedAt(t)
	return nc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (nc *NodeCreate) SetNillableDeletedAt(t *time.Time) *NodeCreate {
	if t != nil {
		nc.SetDeletedAt(*t)
	}
	return nc
}

// SetHost sets the "host" field.
func (nc *NodeCreate) SetHost(s string) *NodeCreate {
	nc.mutation.SetHost(s)
	return nc
}

// SetIpv4 sets the "ipv4" field.
func (nc *NodeCreate) SetIpv4(s string) *NodeCreate {
	nc.mutation.SetIpv4(s)
	return nc
}

// SetIpv6 sets the "ipv6" field.
func (nc *NodeCreate) SetIpv6(s string) *NodeCreate {
	nc.mutation.SetIpv6(s)
	return nc
}

// SetPublicIpv4 sets the "public_ipv4" field.
func (nc *NodeCreate) SetPublicIpv4(s string) *NodeCreate {
	nc.mutation.SetPublicIpv4(s)
	return nc
}

// SetPublicIpv6 sets the "public_ipv6" field.
func (nc *NodeCreate) SetPublicIpv6(s string) *NodeCreate {
	nc.mutation.SetPublicIpv6(s)
	return nc
}

// SetOs sets the "os" field.
func (nc *NodeCreate) SetOs(s string) *NodeCreate {
	nc.mutation.SetOs(s)
	return nc
}

// SetPlatform sets the "platform" field.
func (nc *NodeCreate) SetPlatform(s string) *NodeCreate {
	nc.mutation.SetPlatform(s)
	return nc
}

// SetPlatformFamily sets the "platformFamily" field.
func (nc *NodeCreate) SetPlatformFamily(s string) *NodeCreate {
	nc.mutation.SetPlatformFamily(s)
	return nc
}

// SetInfo sets the "info" field.
func (nc *NodeCreate) SetInfo(s []string) *NodeCreate {
	nc.mutation.SetInfo(s)
	return nc
}

// SetUUID sets the "uuid" field.
func (nc *NodeCreate) SetUUID(s string) *NodeCreate {
	nc.mutation.SetUUID(s)
	return nc
}

// SetDescription sets the "description" field.
func (nc *NodeCreate) SetDescription(s string) *NodeCreate {
	nc.mutation.SetDescription(s)
	return nc
}

// SetDisabled sets the "disabled" field.
func (nc *NodeCreate) SetDisabled(b bool) *NodeCreate {
	nc.mutation.SetDisabled(b)
	return nc
}

// SetAgentId sets the "agentId" field.
func (nc *NodeCreate) SetAgentId(u uint) *NodeCreate {
	nc.mutation.SetAgentId(u)
	return nc
}

// SetClusterId sets the "clusterId" field.
func (nc *NodeCreate) SetClusterId(u uint) *NodeCreate {
	nc.mutation.SetClusterId(u)
	return nc
}

// SetID sets the "id" field.
func (nc *NodeCreate) SetID(u uint) *NodeCreate {
	nc.mutation.SetID(u)
	return nc
}

// AddContainerIDs adds the "containers" edge to the Container entity by IDs.
func (nc *NodeCreate) AddContainerIDs(ids ...uint) *NodeCreate {
	nc.mutation.AddContainerIDs(ids...)
	return nc
}

// AddContainers adds the "containers" edges to the Container entity.
func (nc *NodeCreate) AddContainers(c ...*Container) *NodeCreate {
	ids := make([]uint, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return nc.AddContainerIDs(ids...)
}

// AddProcesIDs adds the "process" edge to the Proces entity by IDs.
func (nc *NodeCreate) AddProcesIDs(ids ...uint) *NodeCreate {
	nc.mutation.AddProcesIDs(ids...)
	return nc
}

// AddProcess adds the "process" edges to the Proces entity.
func (nc *NodeCreate) AddProcess(p ...*Proces) *NodeCreate {
	ids := make([]uint, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return nc.AddProcesIDs(ids...)
}

// Mutation returns the NodeMutation object of the builder.
func (nc *NodeCreate) Mutation() *NodeMutation {
	return nc.mutation
}

// Save creates the Node in the database.
func (nc *NodeCreate) Save(ctx context.Context) (*Node, error) {
	var (
		err  error
		node *Node
	)
	nc.defaults()
	if len(nc.hooks) == 0 {
		if err = nc.check(); err != nil {
			return nil, err
		}
		node, err = nc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nc.check(); err != nil {
				return nil, err
			}
			nc.mutation = mutation
			if node, err = nc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(nc.hooks) - 1; i >= 0; i-- {
			if nc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NodeCreate) SaveX(ctx context.Context) *Node {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NodeCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NodeCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nc *NodeCreate) defaults() {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		v := node.DefaultCreatedAt()
		nc.mutation.SetCreatedAt(v)
	}
	if _, ok := nc.mutation.UpdatedAt(); !ok {
		v := node.DefaultUpdatedAt()
		nc.mutation.SetUpdatedAt(v)
	}
	if _, ok := nc.mutation.DeletedAt(); !ok {
		v := node.DefaultDeletedAt()
		nc.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NodeCreate) check() error {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := nc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := nc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "deleted_at"`)}
	}
	if _, ok := nc.mutation.Host(); !ok {
		return &ValidationError{Name: "host", err: errors.New(`ent: missing required field "host"`)}
	}
	if v, ok := nc.mutation.Host(); ok {
		if err := node.HostValidator(v); err != nil {
			return &ValidationError{Name: "host", err: fmt.Errorf(`ent: validator failed for field "host": %w`, err)}
		}
	}
	if _, ok := nc.mutation.Ipv4(); !ok {
		return &ValidationError{Name: "ipv4", err: errors.New(`ent: missing required field "ipv4"`)}
	}
	if v, ok := nc.mutation.Ipv4(); ok {
		if err := node.Ipv4Validator(v); err != nil {
			return &ValidationError{Name: "ipv4", err: fmt.Errorf(`ent: validator failed for field "ipv4": %w`, err)}
		}
	}
	if _, ok := nc.mutation.Ipv6(); !ok {
		return &ValidationError{Name: "ipv6", err: errors.New(`ent: missing required field "ipv6"`)}
	}
	if v, ok := nc.mutation.Ipv6(); ok {
		if err := node.Ipv6Validator(v); err != nil {
			return &ValidationError{Name: "ipv6", err: fmt.Errorf(`ent: validator failed for field "ipv6": %w`, err)}
		}
	}
	if _, ok := nc.mutation.PublicIpv4(); !ok {
		return &ValidationError{Name: "public_ipv4", err: errors.New(`ent: missing required field "public_ipv4"`)}
	}
	if v, ok := nc.mutation.PublicIpv4(); ok {
		if err := node.PublicIpv4Validator(v); err != nil {
			return &ValidationError{Name: "public_ipv4", err: fmt.Errorf(`ent: validator failed for field "public_ipv4": %w`, err)}
		}
	}
	if _, ok := nc.mutation.PublicIpv6(); !ok {
		return &ValidationError{Name: "public_ipv6", err: errors.New(`ent: missing required field "public_ipv6"`)}
	}
	if v, ok := nc.mutation.PublicIpv6(); ok {
		if err := node.PublicIpv6Validator(v); err != nil {
			return &ValidationError{Name: "public_ipv6", err: fmt.Errorf(`ent: validator failed for field "public_ipv6": %w`, err)}
		}
	}
	if _, ok := nc.mutation.Os(); !ok {
		return &ValidationError{Name: "os", err: errors.New(`ent: missing required field "os"`)}
	}
	if v, ok := nc.mutation.Os(); ok {
		if err := node.OsValidator(v); err != nil {
			return &ValidationError{Name: "os", err: fmt.Errorf(`ent: validator failed for field "os": %w`, err)}
		}
	}
	if _, ok := nc.mutation.Platform(); !ok {
		return &ValidationError{Name: "platform", err: errors.New(`ent: missing required field "platform"`)}
	}
	if v, ok := nc.mutation.Platform(); ok {
		if err := node.PlatformValidator(v); err != nil {
			return &ValidationError{Name: "platform", err: fmt.Errorf(`ent: validator failed for field "platform": %w`, err)}
		}
	}
	if _, ok := nc.mutation.PlatformFamily(); !ok {
		return &ValidationError{Name: "platformFamily", err: errors.New(`ent: missing required field "platformFamily"`)}
	}
	if v, ok := nc.mutation.PlatformFamily(); ok {
		if err := node.PlatformFamilyValidator(v); err != nil {
			return &ValidationError{Name: "platformFamily", err: fmt.Errorf(`ent: validator failed for field "platformFamily": %w`, err)}
		}
	}
	if _, ok := nc.mutation.Info(); !ok {
		return &ValidationError{Name: "info", err: errors.New(`ent: missing required field "info"`)}
	}
	if _, ok := nc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "uuid"`)}
	}
	if _, ok := nc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "description"`)}
	}
	if _, ok := nc.mutation.Disabled(); !ok {
		return &ValidationError{Name: "disabled", err: errors.New(`ent: missing required field "disabled"`)}
	}
	if _, ok := nc.mutation.AgentId(); !ok {
		return &ValidationError{Name: "agentId", err: errors.New(`ent: missing required field "agentId"`)}
	}
	if _, ok := nc.mutation.ClusterId(); !ok {
		return &ValidationError{Name: "clusterId", err: errors.New(`ent: missing required field "clusterId"`)}
	}
	return nil
}

func (nc *NodeCreate) sqlSave(ctx context.Context) (*Node, error) {
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
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

func (nc *NodeCreate) createSpec() (*Node, *sqlgraph.CreateSpec) {
	var (
		_node = &Node{config: nc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: node.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: node.FieldID,
			},
		}
	)
	if id, ok := nc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := nc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: node.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := nc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: node.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := nc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: node.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := nc.mutation.Host(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldHost,
		})
		_node.Host = value
	}
	if value, ok := nc.mutation.Ipv4(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldIpv4,
		})
		_node.Ipv4 = value
	}
	if value, ok := nc.mutation.Ipv6(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldIpv6,
		})
		_node.Ipv6 = value
	}
	if value, ok := nc.mutation.PublicIpv4(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldPublicIpv4,
		})
		_node.PublicIpv4 = value
	}
	if value, ok := nc.mutation.PublicIpv6(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldPublicIpv6,
		})
		_node.PublicIpv6 = value
	}
	if value, ok := nc.mutation.Os(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldOs,
		})
		_node.Os = value
	}
	if value, ok := nc.mutation.Platform(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldPlatform,
		})
		_node.Platform = value
	}
	if value, ok := nc.mutation.PlatformFamily(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldPlatformFamily,
		})
		_node.PlatformFamily = value
	}
	if value, ok := nc.mutation.Info(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: node.FieldInfo,
		})
		_node.Info = value
	}
	if value, ok := nc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := nc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := nc.mutation.Disabled(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: node.FieldDisabled,
		})
		_node.Disabled = value
	}
	if value, ok := nc.mutation.AgentId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: node.FieldAgentId,
		})
		_node.AgentId = value
	}
	if value, ok := nc.mutation.ClusterId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: node.FieldClusterId,
		})
		_node.ClusterId = value
	}
	if nodes := nc.mutation.ContainersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   node.ContainersTable,
			Columns: []string{node.ContainersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: container.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.ProcessIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   node.ProcessTable,
			Columns: []string{node.ProcessColumn},
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
	return _node, _spec
}

// NodeCreateBulk is the builder for creating many Node entities in bulk.
type NodeCreateBulk struct {
	config
	builders []*NodeCreate
}

// Save creates the Node entities in the database.
func (ncb *NodeCreateBulk) Save(ctx context.Context) ([]*Node, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Node, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NodeMutation)
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
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NodeCreateBulk) SaveX(ctx context.Context) []*Node {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NodeCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NodeCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}