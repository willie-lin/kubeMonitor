// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/k8snode"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/predicate"
)

// K8sNodeUpdate is the builder for updating K8sNode entities.
type K8sNodeUpdate struct {
	config
	hooks    []Hook
	mutation *K8sNodeMutation
}

// Where appends a list predicates to the K8sNodeUpdate builder.
func (knu *K8sNodeUpdate) Where(ps ...predicate.K8sNode) *K8sNodeUpdate {
	knu.mutation.Where(ps...)
	return knu
}

// SetUpdatedAt sets the "updated_at" field.
func (knu *K8sNodeUpdate) SetUpdatedAt(t time.Time) *K8sNodeUpdate {
	knu.mutation.SetUpdatedAt(t)
	return knu
}

// SetDeletedAt sets the "deleted_at" field.
func (knu *K8sNodeUpdate) SetDeletedAt(t time.Time) *K8sNodeUpdate {
	knu.mutation.SetDeletedAt(t)
	return knu
}

// SetName sets the "name" field.
func (knu *K8sNodeUpdate) SetName(s string) *K8sNodeUpdate {
	knu.mutation.SetName(s)
	return knu
}

// SetK8sClusterId sets the "k8sClusterId" field.
func (knu *K8sNodeUpdate) SetK8sClusterId(u uint) *K8sNodeUpdate {
	knu.mutation.ResetK8sClusterId()
	knu.mutation.SetK8sClusterId(u)
	return knu
}

// AddK8sClusterId adds u to the "k8sClusterId" field.
func (knu *K8sNodeUpdate) AddK8sClusterId(u uint) *K8sNodeUpdate {
	knu.mutation.AddK8sClusterId(u)
	return knu
}

// SetK8sObjectId sets the "k8sObjectId" field.
func (knu *K8sNodeUpdate) SetK8sObjectId(u uint) *K8sNodeUpdate {
	knu.mutation.ResetK8sObjectId()
	knu.mutation.SetK8sObjectId(u)
	return knu
}

// AddK8sObjectId adds u to the "k8sObjectId" field.
func (knu *K8sNodeUpdate) AddK8sObjectId(u uint) *K8sNodeUpdate {
	knu.mutation.AddK8sObjectId(u)
	return knu
}

// Mutation returns the K8sNodeMutation object of the builder.
func (knu *K8sNodeUpdate) Mutation() *K8sNodeMutation {
	return knu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (knu *K8sNodeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	knu.defaults()
	if len(knu.hooks) == 0 {
		if err = knu.check(); err != nil {
			return 0, err
		}
		affected, err = knu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*K8sNodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = knu.check(); err != nil {
				return 0, err
			}
			knu.mutation = mutation
			affected, err = knu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(knu.hooks) - 1; i >= 0; i-- {
			if knu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = knu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, knu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (knu *K8sNodeUpdate) SaveX(ctx context.Context) int {
	affected, err := knu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (knu *K8sNodeUpdate) Exec(ctx context.Context) error {
	_, err := knu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (knu *K8sNodeUpdate) ExecX(ctx context.Context) {
	if err := knu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (knu *K8sNodeUpdate) defaults() {
	if _, ok := knu.mutation.UpdatedAt(); !ok {
		v := k8snode.UpdateDefaultUpdatedAt()
		knu.mutation.SetUpdatedAt(v)
	}
	if _, ok := knu.mutation.DeletedAt(); !ok {
		v := k8snode.UpdateDefaultDeletedAt()
		knu.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (knu *K8sNodeUpdate) check() error {
	if v, ok := knu.mutation.Name(); ok {
		if err := k8snode.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (knu *K8sNodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   k8snode.Table,
			Columns: k8snode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: k8snode.FieldID,
			},
		},
	}
	if ps := knu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := knu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8snode.FieldUpdatedAt,
		})
	}
	if value, ok := knu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8snode.FieldDeletedAt,
		})
	}
	if value, ok := knu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: k8snode.FieldName,
		})
	}
	if value, ok := knu.mutation.K8sClusterId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8snode.FieldK8sClusterId,
		})
	}
	if value, ok := knu.mutation.AddedK8sClusterId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8snode.FieldK8sClusterId,
		})
	}
	if value, ok := knu.mutation.K8sObjectId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8snode.FieldK8sObjectId,
		})
	}
	if value, ok := knu.mutation.AddedK8sObjectId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8snode.FieldK8sObjectId,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, knu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{k8snode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// K8sNodeUpdateOne is the builder for updating a single K8sNode entity.
type K8sNodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *K8sNodeMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (knuo *K8sNodeUpdateOne) SetUpdatedAt(t time.Time) *K8sNodeUpdateOne {
	knuo.mutation.SetUpdatedAt(t)
	return knuo
}

// SetDeletedAt sets the "deleted_at" field.
func (knuo *K8sNodeUpdateOne) SetDeletedAt(t time.Time) *K8sNodeUpdateOne {
	knuo.mutation.SetDeletedAt(t)
	return knuo
}

// SetName sets the "name" field.
func (knuo *K8sNodeUpdateOne) SetName(s string) *K8sNodeUpdateOne {
	knuo.mutation.SetName(s)
	return knuo
}

// SetK8sClusterId sets the "k8sClusterId" field.
func (knuo *K8sNodeUpdateOne) SetK8sClusterId(u uint) *K8sNodeUpdateOne {
	knuo.mutation.ResetK8sClusterId()
	knuo.mutation.SetK8sClusterId(u)
	return knuo
}

// AddK8sClusterId adds u to the "k8sClusterId" field.
func (knuo *K8sNodeUpdateOne) AddK8sClusterId(u uint) *K8sNodeUpdateOne {
	knuo.mutation.AddK8sClusterId(u)
	return knuo
}

// SetK8sObjectId sets the "k8sObjectId" field.
func (knuo *K8sNodeUpdateOne) SetK8sObjectId(u uint) *K8sNodeUpdateOne {
	knuo.mutation.ResetK8sObjectId()
	knuo.mutation.SetK8sObjectId(u)
	return knuo
}

// AddK8sObjectId adds u to the "k8sObjectId" field.
func (knuo *K8sNodeUpdateOne) AddK8sObjectId(u uint) *K8sNodeUpdateOne {
	knuo.mutation.AddK8sObjectId(u)
	return knuo
}

// Mutation returns the K8sNodeMutation object of the builder.
func (knuo *K8sNodeUpdateOne) Mutation() *K8sNodeMutation {
	return knuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (knuo *K8sNodeUpdateOne) Select(field string, fields ...string) *K8sNodeUpdateOne {
	knuo.fields = append([]string{field}, fields...)
	return knuo
}

// Save executes the query and returns the updated K8sNode entity.
func (knuo *K8sNodeUpdateOne) Save(ctx context.Context) (*K8sNode, error) {
	var (
		err  error
		node *K8sNode
	)
	knuo.defaults()
	if len(knuo.hooks) == 0 {
		if err = knuo.check(); err != nil {
			return nil, err
		}
		node, err = knuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*K8sNodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = knuo.check(); err != nil {
				return nil, err
			}
			knuo.mutation = mutation
			node, err = knuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(knuo.hooks) - 1; i >= 0; i-- {
			if knuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = knuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, knuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (knuo *K8sNodeUpdateOne) SaveX(ctx context.Context) *K8sNode {
	node, err := knuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (knuo *K8sNodeUpdateOne) Exec(ctx context.Context) error {
	_, err := knuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (knuo *K8sNodeUpdateOne) ExecX(ctx context.Context) {
	if err := knuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (knuo *K8sNodeUpdateOne) defaults() {
	if _, ok := knuo.mutation.UpdatedAt(); !ok {
		v := k8snode.UpdateDefaultUpdatedAt()
		knuo.mutation.SetUpdatedAt(v)
	}
	if _, ok := knuo.mutation.DeletedAt(); !ok {
		v := k8snode.UpdateDefaultDeletedAt()
		knuo.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (knuo *K8sNodeUpdateOne) check() error {
	if v, ok := knuo.mutation.Name(); ok {
		if err := k8snode.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (knuo *K8sNodeUpdateOne) sqlSave(ctx context.Context) (_node *K8sNode, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   k8snode.Table,
			Columns: k8snode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: k8snode.FieldID,
			},
		},
	}
	id, ok := knuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing K8sNode.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := knuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, k8snode.FieldID)
		for _, f := range fields {
			if !k8snode.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != k8snode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := knuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := knuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8snode.FieldUpdatedAt,
		})
	}
	if value, ok := knuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8snode.FieldDeletedAt,
		})
	}
	if value, ok := knuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: k8snode.FieldName,
		})
	}
	if value, ok := knuo.mutation.K8sClusterId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8snode.FieldK8sClusterId,
		})
	}
	if value, ok := knuo.mutation.AddedK8sClusterId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8snode.FieldK8sClusterId,
		})
	}
	if value, ok := knuo.mutation.K8sObjectId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8snode.FieldK8sObjectId,
		})
	}
	if value, ok := knuo.mutation.AddedK8sObjectId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8snode.FieldK8sObjectId,
		})
	}
	_node = &K8sNode{config: knuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, knuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{k8snode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
