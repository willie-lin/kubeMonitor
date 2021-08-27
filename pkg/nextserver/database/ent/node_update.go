// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/container"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/node"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/predicate"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/proces"
)

// NodeUpdate is the builder for updating Node entities.
type NodeUpdate struct {
	config
	hooks    []Hook
	mutation *NodeMutation
}

// Where appends a list predicates to the NodeUpdate builder.
func (nu *NodeUpdate) Where(ps ...predicate.Node) *NodeUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetUpdatedAt sets the "updated_at" field.
func (nu *NodeUpdate) SetUpdatedAt(t time.Time) *NodeUpdate {
	nu.mutation.SetUpdatedAt(t)
	return nu
}

// SetDeletedAt sets the "deleted_at" field.
func (nu *NodeUpdate) SetDeletedAt(t time.Time) *NodeUpdate {
	nu.mutation.SetDeletedAt(t)
	return nu
}

// SetHost sets the "host" field.
func (nu *NodeUpdate) SetHost(s string) *NodeUpdate {
	nu.mutation.SetHost(s)
	return nu
}

// SetIpv4 sets the "ipv4" field.
func (nu *NodeUpdate) SetIpv4(s string) *NodeUpdate {
	nu.mutation.SetIpv4(s)
	return nu
}

// SetIpv6 sets the "ipv6" field.
func (nu *NodeUpdate) SetIpv6(s string) *NodeUpdate {
	nu.mutation.SetIpv6(s)
	return nu
}

// SetPublicIpv4 sets the "public_ipv4" field.
func (nu *NodeUpdate) SetPublicIpv4(s string) *NodeUpdate {
	nu.mutation.SetPublicIpv4(s)
	return nu
}

// SetPublicIpv6 sets the "public_ipv6" field.
func (nu *NodeUpdate) SetPublicIpv6(s string) *NodeUpdate {
	nu.mutation.SetPublicIpv6(s)
	return nu
}

// SetOs sets the "os" field.
func (nu *NodeUpdate) SetOs(s string) *NodeUpdate {
	nu.mutation.SetOs(s)
	return nu
}

// SetPlatform sets the "platform" field.
func (nu *NodeUpdate) SetPlatform(s string) *NodeUpdate {
	nu.mutation.SetPlatform(s)
	return nu
}

// SetPlatformFamily sets the "platformFamily" field.
func (nu *NodeUpdate) SetPlatformFamily(s string) *NodeUpdate {
	nu.mutation.SetPlatformFamily(s)
	return nu
}

// SetInfo sets the "info" field.
func (nu *NodeUpdate) SetInfo(s []string) *NodeUpdate {
	nu.mutation.SetInfo(s)
	return nu
}

// SetUUID sets the "uuid" field.
func (nu *NodeUpdate) SetUUID(s string) *NodeUpdate {
	nu.mutation.SetUUID(s)
	return nu
}

// SetDescription sets the "description" field.
func (nu *NodeUpdate) SetDescription(s string) *NodeUpdate {
	nu.mutation.SetDescription(s)
	return nu
}

// SetDisabled sets the "disabled" field.
func (nu *NodeUpdate) SetDisabled(b bool) *NodeUpdate {
	nu.mutation.SetDisabled(b)
	return nu
}

// SetAgentId sets the "agentId" field.
func (nu *NodeUpdate) SetAgentId(u uint) *NodeUpdate {
	nu.mutation.ResetAgentId()
	nu.mutation.SetAgentId(u)
	return nu
}

// AddAgentId adds u to the "agentId" field.
func (nu *NodeUpdate) AddAgentId(u uint) *NodeUpdate {
	nu.mutation.AddAgentId(u)
	return nu
}

// SetClusterId sets the "clusterId" field.
func (nu *NodeUpdate) SetClusterId(u uint) *NodeUpdate {
	nu.mutation.ResetClusterId()
	nu.mutation.SetClusterId(u)
	return nu
}

// AddClusterId adds u to the "clusterId" field.
func (nu *NodeUpdate) AddClusterId(u uint) *NodeUpdate {
	nu.mutation.AddClusterId(u)
	return nu
}

// AddContainerIDs adds the "containers" edge to the Container entity by IDs.
func (nu *NodeUpdate) AddContainerIDs(ids ...uint) *NodeUpdate {
	nu.mutation.AddContainerIDs(ids...)
	return nu
}

// AddContainers adds the "containers" edges to the Container entity.
func (nu *NodeUpdate) AddContainers(c ...*Container) *NodeUpdate {
	ids := make([]uint, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return nu.AddContainerIDs(ids...)
}

// AddProcesIDs adds the "process" edge to the Proces entity by IDs.
func (nu *NodeUpdate) AddProcesIDs(ids ...uint) *NodeUpdate {
	nu.mutation.AddProcesIDs(ids...)
	return nu
}

// AddProcess adds the "process" edges to the Proces entity.
func (nu *NodeUpdate) AddProcess(p ...*Proces) *NodeUpdate {
	ids := make([]uint, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return nu.AddProcesIDs(ids...)
}

// Mutation returns the NodeMutation object of the builder.
func (nu *NodeUpdate) Mutation() *NodeMutation {
	return nu.mutation
}

// ClearContainers clears all "containers" edges to the Container entity.
func (nu *NodeUpdate) ClearContainers() *NodeUpdate {
	nu.mutation.ClearContainers()
	return nu
}

// RemoveContainerIDs removes the "containers" edge to Container entities by IDs.
func (nu *NodeUpdate) RemoveContainerIDs(ids ...uint) *NodeUpdate {
	nu.mutation.RemoveContainerIDs(ids...)
	return nu
}

// RemoveContainers removes "containers" edges to Container entities.
func (nu *NodeUpdate) RemoveContainers(c ...*Container) *NodeUpdate {
	ids := make([]uint, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return nu.RemoveContainerIDs(ids...)
}

// ClearProcess clears all "process" edges to the Proces entity.
func (nu *NodeUpdate) ClearProcess() *NodeUpdate {
	nu.mutation.ClearProcess()
	return nu
}

// RemoveProcesIDs removes the "process" edge to Proces entities by IDs.
func (nu *NodeUpdate) RemoveProcesIDs(ids ...uint) *NodeUpdate {
	nu.mutation.RemoveProcesIDs(ids...)
	return nu
}

// RemoveProcess removes "process" edges to Proces entities.
func (nu *NodeUpdate) RemoveProcess(p ...*Proces) *NodeUpdate {
	ids := make([]uint, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return nu.RemoveProcesIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NodeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	nu.defaults()
	if len(nu.hooks) == 0 {
		if err = nu.check(); err != nil {
			return 0, err
		}
		affected, err = nu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nu.check(); err != nil {
				return 0, err
			}
			nu.mutation = mutation
			affected, err = nu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nu.hooks) - 1; i >= 0; i-- {
			if nu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NodeUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NodeUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NodeUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nu *NodeUpdate) defaults() {
	if _, ok := nu.mutation.UpdatedAt(); !ok {
		v := node.UpdateDefaultUpdatedAt()
		nu.mutation.SetUpdatedAt(v)
	}
	if _, ok := nu.mutation.DeletedAt(); !ok {
		v := node.UpdateDefaultDeletedAt()
		nu.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nu *NodeUpdate) check() error {
	if v, ok := nu.mutation.Host(); ok {
		if err := node.HostValidator(v); err != nil {
			return &ValidationError{Name: "host", err: fmt.Errorf("ent: validator failed for field \"host\": %w", err)}
		}
	}
	if v, ok := nu.mutation.Ipv4(); ok {
		if err := node.Ipv4Validator(v); err != nil {
			return &ValidationError{Name: "ipv4", err: fmt.Errorf("ent: validator failed for field \"ipv4\": %w", err)}
		}
	}
	if v, ok := nu.mutation.Ipv6(); ok {
		if err := node.Ipv6Validator(v); err != nil {
			return &ValidationError{Name: "ipv6", err: fmt.Errorf("ent: validator failed for field \"ipv6\": %w", err)}
		}
	}
	if v, ok := nu.mutation.PublicIpv4(); ok {
		if err := node.PublicIpv4Validator(v); err != nil {
			return &ValidationError{Name: "public_ipv4", err: fmt.Errorf("ent: validator failed for field \"public_ipv4\": %w", err)}
		}
	}
	if v, ok := nu.mutation.PublicIpv6(); ok {
		if err := node.PublicIpv6Validator(v); err != nil {
			return &ValidationError{Name: "public_ipv6", err: fmt.Errorf("ent: validator failed for field \"public_ipv6\": %w", err)}
		}
	}
	if v, ok := nu.mutation.Os(); ok {
		if err := node.OsValidator(v); err != nil {
			return &ValidationError{Name: "os", err: fmt.Errorf("ent: validator failed for field \"os\": %w", err)}
		}
	}
	if v, ok := nu.mutation.Platform(); ok {
		if err := node.PlatformValidator(v); err != nil {
			return &ValidationError{Name: "platform", err: fmt.Errorf("ent: validator failed for field \"platform\": %w", err)}
		}
	}
	if v, ok := nu.mutation.PlatformFamily(); ok {
		if err := node.PlatformFamilyValidator(v); err != nil {
			return &ValidationError{Name: "platformFamily", err: fmt.Errorf("ent: validator failed for field \"platformFamily\": %w", err)}
		}
	}
	return nil
}

func (nu *NodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   node.Table,
			Columns: node.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: node.FieldID,
			},
		},
	}
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: node.FieldUpdatedAt,
		})
	}
	if value, ok := nu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: node.FieldDeletedAt,
		})
	}
	if value, ok := nu.mutation.Host(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldHost,
		})
	}
	if value, ok := nu.mutation.Ipv4(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldIpv4,
		})
	}
	if value, ok := nu.mutation.Ipv6(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldIpv6,
		})
	}
	if value, ok := nu.mutation.PublicIpv4(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldPublicIpv4,
		})
	}
	if value, ok := nu.mutation.PublicIpv6(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldPublicIpv6,
		})
	}
	if value, ok := nu.mutation.Os(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldOs,
		})
	}
	if value, ok := nu.mutation.Platform(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldPlatform,
		})
	}
	if value, ok := nu.mutation.PlatformFamily(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldPlatformFamily,
		})
	}
	if value, ok := nu.mutation.Info(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: node.FieldInfo,
		})
	}
	if value, ok := nu.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldUUID,
		})
	}
	if value, ok := nu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldDescription,
		})
	}
	if value, ok := nu.mutation.Disabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: node.FieldDisabled,
		})
	}
	if value, ok := nu.mutation.AgentId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: node.FieldAgentId,
		})
	}
	if value, ok := nu.mutation.AddedAgentId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: node.FieldAgentId,
		})
	}
	if value, ok := nu.mutation.ClusterId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: node.FieldClusterId,
		})
	}
	if value, ok := nu.mutation.AddedClusterId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: node.FieldClusterId,
		})
	}
	if nu.mutation.ContainersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.RemovedContainersIDs(); len(nodes) > 0 && !nu.mutation.ContainersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.ContainersIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nu.mutation.ProcessCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.RemovedProcessIDs(); len(nodes) > 0 && !nu.mutation.ProcessCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.ProcessIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{node.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// NodeUpdateOne is the builder for updating a single Node entity.
type NodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NodeMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (nuo *NodeUpdateOne) SetUpdatedAt(t time.Time) *NodeUpdateOne {
	nuo.mutation.SetUpdatedAt(t)
	return nuo
}

// SetDeletedAt sets the "deleted_at" field.
func (nuo *NodeUpdateOne) SetDeletedAt(t time.Time) *NodeUpdateOne {
	nuo.mutation.SetDeletedAt(t)
	return nuo
}

// SetHost sets the "host" field.
func (nuo *NodeUpdateOne) SetHost(s string) *NodeUpdateOne {
	nuo.mutation.SetHost(s)
	return nuo
}

// SetIpv4 sets the "ipv4" field.
func (nuo *NodeUpdateOne) SetIpv4(s string) *NodeUpdateOne {
	nuo.mutation.SetIpv4(s)
	return nuo
}

// SetIpv6 sets the "ipv6" field.
func (nuo *NodeUpdateOne) SetIpv6(s string) *NodeUpdateOne {
	nuo.mutation.SetIpv6(s)
	return nuo
}

// SetPublicIpv4 sets the "public_ipv4" field.
func (nuo *NodeUpdateOne) SetPublicIpv4(s string) *NodeUpdateOne {
	nuo.mutation.SetPublicIpv4(s)
	return nuo
}

// SetPublicIpv6 sets the "public_ipv6" field.
func (nuo *NodeUpdateOne) SetPublicIpv6(s string) *NodeUpdateOne {
	nuo.mutation.SetPublicIpv6(s)
	return nuo
}

// SetOs sets the "os" field.
func (nuo *NodeUpdateOne) SetOs(s string) *NodeUpdateOne {
	nuo.mutation.SetOs(s)
	return nuo
}

// SetPlatform sets the "platform" field.
func (nuo *NodeUpdateOne) SetPlatform(s string) *NodeUpdateOne {
	nuo.mutation.SetPlatform(s)
	return nuo
}

// SetPlatformFamily sets the "platformFamily" field.
func (nuo *NodeUpdateOne) SetPlatformFamily(s string) *NodeUpdateOne {
	nuo.mutation.SetPlatformFamily(s)
	return nuo
}

// SetInfo sets the "info" field.
func (nuo *NodeUpdateOne) SetInfo(s []string) *NodeUpdateOne {
	nuo.mutation.SetInfo(s)
	return nuo
}

// SetUUID sets the "uuid" field.
func (nuo *NodeUpdateOne) SetUUID(s string) *NodeUpdateOne {
	nuo.mutation.SetUUID(s)
	return nuo
}

// SetDescription sets the "description" field.
func (nuo *NodeUpdateOne) SetDescription(s string) *NodeUpdateOne {
	nuo.mutation.SetDescription(s)
	return nuo
}

// SetDisabled sets the "disabled" field.
func (nuo *NodeUpdateOne) SetDisabled(b bool) *NodeUpdateOne {
	nuo.mutation.SetDisabled(b)
	return nuo
}

// SetAgentId sets the "agentId" field.
func (nuo *NodeUpdateOne) SetAgentId(u uint) *NodeUpdateOne {
	nuo.mutation.ResetAgentId()
	nuo.mutation.SetAgentId(u)
	return nuo
}

// AddAgentId adds u to the "agentId" field.
func (nuo *NodeUpdateOne) AddAgentId(u uint) *NodeUpdateOne {
	nuo.mutation.AddAgentId(u)
	return nuo
}

// SetClusterId sets the "clusterId" field.
func (nuo *NodeUpdateOne) SetClusterId(u uint) *NodeUpdateOne {
	nuo.mutation.ResetClusterId()
	nuo.mutation.SetClusterId(u)
	return nuo
}

// AddClusterId adds u to the "clusterId" field.
func (nuo *NodeUpdateOne) AddClusterId(u uint) *NodeUpdateOne {
	nuo.mutation.AddClusterId(u)
	return nuo
}

// AddContainerIDs adds the "containers" edge to the Container entity by IDs.
func (nuo *NodeUpdateOne) AddContainerIDs(ids ...uint) *NodeUpdateOne {
	nuo.mutation.AddContainerIDs(ids...)
	return nuo
}

// AddContainers adds the "containers" edges to the Container entity.
func (nuo *NodeUpdateOne) AddContainers(c ...*Container) *NodeUpdateOne {
	ids := make([]uint, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return nuo.AddContainerIDs(ids...)
}

// AddProcesIDs adds the "process" edge to the Proces entity by IDs.
func (nuo *NodeUpdateOne) AddProcesIDs(ids ...uint) *NodeUpdateOne {
	nuo.mutation.AddProcesIDs(ids...)
	return nuo
}

// AddProcess adds the "process" edges to the Proces entity.
func (nuo *NodeUpdateOne) AddProcess(p ...*Proces) *NodeUpdateOne {
	ids := make([]uint, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return nuo.AddProcesIDs(ids...)
}

// Mutation returns the NodeMutation object of the builder.
func (nuo *NodeUpdateOne) Mutation() *NodeMutation {
	return nuo.mutation
}

// ClearContainers clears all "containers" edges to the Container entity.
func (nuo *NodeUpdateOne) ClearContainers() *NodeUpdateOne {
	nuo.mutation.ClearContainers()
	return nuo
}

// RemoveContainerIDs removes the "containers" edge to Container entities by IDs.
func (nuo *NodeUpdateOne) RemoveContainerIDs(ids ...uint) *NodeUpdateOne {
	nuo.mutation.RemoveContainerIDs(ids...)
	return nuo
}

// RemoveContainers removes "containers" edges to Container entities.
func (nuo *NodeUpdateOne) RemoveContainers(c ...*Container) *NodeUpdateOne {
	ids := make([]uint, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return nuo.RemoveContainerIDs(ids...)
}

// ClearProcess clears all "process" edges to the Proces entity.
func (nuo *NodeUpdateOne) ClearProcess() *NodeUpdateOne {
	nuo.mutation.ClearProcess()
	return nuo
}

// RemoveProcesIDs removes the "process" edge to Proces entities by IDs.
func (nuo *NodeUpdateOne) RemoveProcesIDs(ids ...uint) *NodeUpdateOne {
	nuo.mutation.RemoveProcesIDs(ids...)
	return nuo
}

// RemoveProcess removes "process" edges to Proces entities.
func (nuo *NodeUpdateOne) RemoveProcess(p ...*Proces) *NodeUpdateOne {
	ids := make([]uint, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return nuo.RemoveProcesIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NodeUpdateOne) Select(field string, fields ...string) *NodeUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Node entity.
func (nuo *NodeUpdateOne) Save(ctx context.Context) (*Node, error) {
	var (
		err  error
		node *Node
	)
	nuo.defaults()
	if len(nuo.hooks) == 0 {
		if err = nuo.check(); err != nil {
			return nil, err
		}
		node, err = nuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nuo.check(); err != nil {
				return nil, err
			}
			nuo.mutation = mutation
			node, err = nuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nuo.hooks) - 1; i >= 0; i-- {
			if nuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NodeUpdateOne) SaveX(ctx context.Context) *Node {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NodeUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NodeUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nuo *NodeUpdateOne) defaults() {
	if _, ok := nuo.mutation.UpdatedAt(); !ok {
		v := node.UpdateDefaultUpdatedAt()
		nuo.mutation.SetUpdatedAt(v)
	}
	if _, ok := nuo.mutation.DeletedAt(); !ok {
		v := node.UpdateDefaultDeletedAt()
		nuo.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nuo *NodeUpdateOne) check() error {
	if v, ok := nuo.mutation.Host(); ok {
		if err := node.HostValidator(v); err != nil {
			return &ValidationError{Name: "host", err: fmt.Errorf("ent: validator failed for field \"host\": %w", err)}
		}
	}
	if v, ok := nuo.mutation.Ipv4(); ok {
		if err := node.Ipv4Validator(v); err != nil {
			return &ValidationError{Name: "ipv4", err: fmt.Errorf("ent: validator failed for field \"ipv4\": %w", err)}
		}
	}
	if v, ok := nuo.mutation.Ipv6(); ok {
		if err := node.Ipv6Validator(v); err != nil {
			return &ValidationError{Name: "ipv6", err: fmt.Errorf("ent: validator failed for field \"ipv6\": %w", err)}
		}
	}
	if v, ok := nuo.mutation.PublicIpv4(); ok {
		if err := node.PublicIpv4Validator(v); err != nil {
			return &ValidationError{Name: "public_ipv4", err: fmt.Errorf("ent: validator failed for field \"public_ipv4\": %w", err)}
		}
	}
	if v, ok := nuo.mutation.PublicIpv6(); ok {
		if err := node.PublicIpv6Validator(v); err != nil {
			return &ValidationError{Name: "public_ipv6", err: fmt.Errorf("ent: validator failed for field \"public_ipv6\": %w", err)}
		}
	}
	if v, ok := nuo.mutation.Os(); ok {
		if err := node.OsValidator(v); err != nil {
			return &ValidationError{Name: "os", err: fmt.Errorf("ent: validator failed for field \"os\": %w", err)}
		}
	}
	if v, ok := nuo.mutation.Platform(); ok {
		if err := node.PlatformValidator(v); err != nil {
			return &ValidationError{Name: "platform", err: fmt.Errorf("ent: validator failed for field \"platform\": %w", err)}
		}
	}
	if v, ok := nuo.mutation.PlatformFamily(); ok {
		if err := node.PlatformFamilyValidator(v); err != nil {
			return &ValidationError{Name: "platformFamily", err: fmt.Errorf("ent: validator failed for field \"platformFamily\": %w", err)}
		}
	}
	return nil
}

func (nuo *NodeUpdateOne) sqlSave(ctx context.Context) (_node *Node, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   node.Table,
			Columns: node.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: node.FieldID,
			},
		},
	}
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Node.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, node.FieldID)
		for _, f := range fields {
			if !node.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != node.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: node.FieldUpdatedAt,
		})
	}
	if value, ok := nuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: node.FieldDeletedAt,
		})
	}
	if value, ok := nuo.mutation.Host(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldHost,
		})
	}
	if value, ok := nuo.mutation.Ipv4(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldIpv4,
		})
	}
	if value, ok := nuo.mutation.Ipv6(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldIpv6,
		})
	}
	if value, ok := nuo.mutation.PublicIpv4(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldPublicIpv4,
		})
	}
	if value, ok := nuo.mutation.PublicIpv6(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldPublicIpv6,
		})
	}
	if value, ok := nuo.mutation.Os(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldOs,
		})
	}
	if value, ok := nuo.mutation.Platform(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldPlatform,
		})
	}
	if value, ok := nuo.mutation.PlatformFamily(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldPlatformFamily,
		})
	}
	if value, ok := nuo.mutation.Info(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: node.FieldInfo,
		})
	}
	if value, ok := nuo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldUUID,
		})
	}
	if value, ok := nuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldDescription,
		})
	}
	if value, ok := nuo.mutation.Disabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: node.FieldDisabled,
		})
	}
	if value, ok := nuo.mutation.AgentId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: node.FieldAgentId,
		})
	}
	if value, ok := nuo.mutation.AddedAgentId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: node.FieldAgentId,
		})
	}
	if value, ok := nuo.mutation.ClusterId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: node.FieldClusterId,
		})
	}
	if value, ok := nuo.mutation.AddedClusterId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: node.FieldClusterId,
		})
	}
	if nuo.mutation.ContainersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.RemovedContainersIDs(); len(nodes) > 0 && !nuo.mutation.ContainersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.ContainersIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nuo.mutation.ProcessCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.RemovedProcessIDs(); len(nodes) > 0 && !nuo.mutation.ProcessCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.ProcessIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Node{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{node.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}