// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/k8sdeployment"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/predicate"
)

// K8sDeploymentUpdate is the builder for updating K8sDeployment entities.
type K8sDeploymentUpdate struct {
	config
	hooks    []Hook
	mutation *K8sDeploymentMutation
}

// Where appends a list predicates to the K8sDeploymentUpdate builder.
func (kdu *K8sDeploymentUpdate) Where(ps ...predicate.K8sDeployment) *K8sDeploymentUpdate {
	kdu.mutation.Where(ps...)
	return kdu
}

// SetUpdatedAt sets the "updated_at" field.
func (kdu *K8sDeploymentUpdate) SetUpdatedAt(t time.Time) *K8sDeploymentUpdate {
	kdu.mutation.SetUpdatedAt(t)
	return kdu
}

// SetDeletedAt sets the "deleted_at" field.
func (kdu *K8sDeploymentUpdate) SetDeletedAt(t time.Time) *K8sDeploymentUpdate {
	kdu.mutation.SetDeletedAt(t)
	return kdu
}

// SetName sets the "name" field.
func (kdu *K8sDeploymentUpdate) SetName(s string) *K8sDeploymentUpdate {
	kdu.mutation.SetName(s)
	return kdu
}

// SetK8sClusterId sets the "k8sClusterId" field.
func (kdu *K8sDeploymentUpdate) SetK8sClusterId(u uint) *K8sDeploymentUpdate {
	kdu.mutation.ResetK8sClusterId()
	kdu.mutation.SetK8sClusterId(u)
	return kdu
}

// AddK8sClusterId adds u to the "k8sClusterId" field.
func (kdu *K8sDeploymentUpdate) AddK8sClusterId(u uint) *K8sDeploymentUpdate {
	kdu.mutation.AddK8sClusterId(u)
	return kdu
}

// SetK8sNamespaceId sets the "k8sNamespaceId" field.
func (kdu *K8sDeploymentUpdate) SetK8sNamespaceId(u uint) *K8sDeploymentUpdate {
	kdu.mutation.ResetK8sNamespaceId()
	kdu.mutation.SetK8sNamespaceId(u)
	return kdu
}

// AddK8sNamespaceId adds u to the "k8sNamespaceId" field.
func (kdu *K8sDeploymentUpdate) AddK8sNamespaceId(u uint) *K8sDeploymentUpdate {
	kdu.mutation.AddK8sNamespaceId(u)
	return kdu
}

// SetK8sObjectId sets the "k8sObjectId" field.
func (kdu *K8sDeploymentUpdate) SetK8sObjectId(u uint) *K8sDeploymentUpdate {
	kdu.mutation.ResetK8sObjectId()
	kdu.mutation.SetK8sObjectId(u)
	return kdu
}

// AddK8sObjectId adds u to the "k8sObjectId" field.
func (kdu *K8sDeploymentUpdate) AddK8sObjectId(u uint) *K8sDeploymentUpdate {
	kdu.mutation.AddK8sObjectId(u)
	return kdu
}

// Mutation returns the K8sDeploymentMutation object of the builder.
func (kdu *K8sDeploymentUpdate) Mutation() *K8sDeploymentMutation {
	return kdu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (kdu *K8sDeploymentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	kdu.defaults()
	if len(kdu.hooks) == 0 {
		if err = kdu.check(); err != nil {
			return 0, err
		}
		affected, err = kdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*K8sDeploymentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kdu.check(); err != nil {
				return 0, err
			}
			kdu.mutation = mutation
			affected, err = kdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(kdu.hooks) - 1; i >= 0; i-- {
			if kdu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (kdu *K8sDeploymentUpdate) SaveX(ctx context.Context) int {
	affected, err := kdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (kdu *K8sDeploymentUpdate) Exec(ctx context.Context) error {
	_, err := kdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kdu *K8sDeploymentUpdate) ExecX(ctx context.Context) {
	if err := kdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kdu *K8sDeploymentUpdate) defaults() {
	if _, ok := kdu.mutation.UpdatedAt(); !ok {
		v := k8sdeployment.UpdateDefaultUpdatedAt()
		kdu.mutation.SetUpdatedAt(v)
	}
	if _, ok := kdu.mutation.DeletedAt(); !ok {
		v := k8sdeployment.UpdateDefaultDeletedAt()
		kdu.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kdu *K8sDeploymentUpdate) check() error {
	if v, ok := kdu.mutation.Name(); ok {
		if err := k8sdeployment.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (kdu *K8sDeploymentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   k8sdeployment.Table,
			Columns: k8sdeployment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: k8sdeployment.FieldID,
			},
		},
	}
	if ps := kdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kdu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8sdeployment.FieldUpdatedAt,
		})
	}
	if value, ok := kdu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8sdeployment.FieldDeletedAt,
		})
	}
	if value, ok := kdu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: k8sdeployment.FieldName,
		})
	}
	if value, ok := kdu.mutation.K8sClusterId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sdeployment.FieldK8sClusterId,
		})
	}
	if value, ok := kdu.mutation.AddedK8sClusterId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sdeployment.FieldK8sClusterId,
		})
	}
	if value, ok := kdu.mutation.K8sNamespaceId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sdeployment.FieldK8sNamespaceId,
		})
	}
	if value, ok := kdu.mutation.AddedK8sNamespaceId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sdeployment.FieldK8sNamespaceId,
		})
	}
	if value, ok := kdu.mutation.K8sObjectId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sdeployment.FieldK8sObjectId,
		})
	}
	if value, ok := kdu.mutation.AddedK8sObjectId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sdeployment.FieldK8sObjectId,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, kdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{k8sdeployment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// K8sDeploymentUpdateOne is the builder for updating a single K8sDeployment entity.
type K8sDeploymentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *K8sDeploymentMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (kduo *K8sDeploymentUpdateOne) SetUpdatedAt(t time.Time) *K8sDeploymentUpdateOne {
	kduo.mutation.SetUpdatedAt(t)
	return kduo
}

// SetDeletedAt sets the "deleted_at" field.
func (kduo *K8sDeploymentUpdateOne) SetDeletedAt(t time.Time) *K8sDeploymentUpdateOne {
	kduo.mutation.SetDeletedAt(t)
	return kduo
}

// SetName sets the "name" field.
func (kduo *K8sDeploymentUpdateOne) SetName(s string) *K8sDeploymentUpdateOne {
	kduo.mutation.SetName(s)
	return kduo
}

// SetK8sClusterId sets the "k8sClusterId" field.
func (kduo *K8sDeploymentUpdateOne) SetK8sClusterId(u uint) *K8sDeploymentUpdateOne {
	kduo.mutation.ResetK8sClusterId()
	kduo.mutation.SetK8sClusterId(u)
	return kduo
}

// AddK8sClusterId adds u to the "k8sClusterId" field.
func (kduo *K8sDeploymentUpdateOne) AddK8sClusterId(u uint) *K8sDeploymentUpdateOne {
	kduo.mutation.AddK8sClusterId(u)
	return kduo
}

// SetK8sNamespaceId sets the "k8sNamespaceId" field.
func (kduo *K8sDeploymentUpdateOne) SetK8sNamespaceId(u uint) *K8sDeploymentUpdateOne {
	kduo.mutation.ResetK8sNamespaceId()
	kduo.mutation.SetK8sNamespaceId(u)
	return kduo
}

// AddK8sNamespaceId adds u to the "k8sNamespaceId" field.
func (kduo *K8sDeploymentUpdateOne) AddK8sNamespaceId(u uint) *K8sDeploymentUpdateOne {
	kduo.mutation.AddK8sNamespaceId(u)
	return kduo
}

// SetK8sObjectId sets the "k8sObjectId" field.
func (kduo *K8sDeploymentUpdateOne) SetK8sObjectId(u uint) *K8sDeploymentUpdateOne {
	kduo.mutation.ResetK8sObjectId()
	kduo.mutation.SetK8sObjectId(u)
	return kduo
}

// AddK8sObjectId adds u to the "k8sObjectId" field.
func (kduo *K8sDeploymentUpdateOne) AddK8sObjectId(u uint) *K8sDeploymentUpdateOne {
	kduo.mutation.AddK8sObjectId(u)
	return kduo
}

// Mutation returns the K8sDeploymentMutation object of the builder.
func (kduo *K8sDeploymentUpdateOne) Mutation() *K8sDeploymentMutation {
	return kduo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (kduo *K8sDeploymentUpdateOne) Select(field string, fields ...string) *K8sDeploymentUpdateOne {
	kduo.fields = append([]string{field}, fields...)
	return kduo
}

// Save executes the query and returns the updated K8sDeployment entity.
func (kduo *K8sDeploymentUpdateOne) Save(ctx context.Context) (*K8sDeployment, error) {
	var (
		err  error
		node *K8sDeployment
	)
	kduo.defaults()
	if len(kduo.hooks) == 0 {
		if err = kduo.check(); err != nil {
			return nil, err
		}
		node, err = kduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*K8sDeploymentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kduo.check(); err != nil {
				return nil, err
			}
			kduo.mutation = mutation
			node, err = kduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(kduo.hooks) - 1; i >= 0; i-- {
			if kduo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (kduo *K8sDeploymentUpdateOne) SaveX(ctx context.Context) *K8sDeployment {
	node, err := kduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (kduo *K8sDeploymentUpdateOne) Exec(ctx context.Context) error {
	_, err := kduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kduo *K8sDeploymentUpdateOne) ExecX(ctx context.Context) {
	if err := kduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kduo *K8sDeploymentUpdateOne) defaults() {
	if _, ok := kduo.mutation.UpdatedAt(); !ok {
		v := k8sdeployment.UpdateDefaultUpdatedAt()
		kduo.mutation.SetUpdatedAt(v)
	}
	if _, ok := kduo.mutation.DeletedAt(); !ok {
		v := k8sdeployment.UpdateDefaultDeletedAt()
		kduo.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kduo *K8sDeploymentUpdateOne) check() error {
	if v, ok := kduo.mutation.Name(); ok {
		if err := k8sdeployment.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (kduo *K8sDeploymentUpdateOne) sqlSave(ctx context.Context) (_node *K8sDeployment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   k8sdeployment.Table,
			Columns: k8sdeployment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: k8sdeployment.FieldID,
			},
		},
	}
	id, ok := kduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing K8sDeployment.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := kduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, k8sdeployment.FieldID)
		for _, f := range fields {
			if !k8sdeployment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != k8sdeployment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := kduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kduo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8sdeployment.FieldUpdatedAt,
		})
	}
	if value, ok := kduo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8sdeployment.FieldDeletedAt,
		})
	}
	if value, ok := kduo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: k8sdeployment.FieldName,
		})
	}
	if value, ok := kduo.mutation.K8sClusterId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sdeployment.FieldK8sClusterId,
		})
	}
	if value, ok := kduo.mutation.AddedK8sClusterId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sdeployment.FieldK8sClusterId,
		})
	}
	if value, ok := kduo.mutation.K8sNamespaceId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sdeployment.FieldK8sNamespaceId,
		})
	}
	if value, ok := kduo.mutation.AddedK8sNamespaceId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sdeployment.FieldK8sNamespaceId,
		})
	}
	if value, ok := kduo.mutation.K8sObjectId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sdeployment.FieldK8sObjectId,
		})
	}
	if value, ok := kduo.mutation.AddedK8sObjectId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sdeployment.FieldK8sObjectId,
		})
	}
	_node = &K8sDeployment{config: kduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, kduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{k8sdeployment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
