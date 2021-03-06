// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/k8sstatefulset"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/predicate"
)

// K8sStatefulSetUpdate is the builder for updating K8sStatefulSet entities.
type K8sStatefulSetUpdate struct {
	config
	hooks    []Hook
	mutation *K8sStatefulSetMutation
}

// Where appends a list predicates to the K8sStatefulSetUpdate builder.
func (kssu *K8sStatefulSetUpdate) Where(ps ...predicate.K8sStatefulSet) *K8sStatefulSetUpdate {
	kssu.mutation.Where(ps...)
	return kssu
}

// SetUpdatedAt sets the "updated_at" field.
func (kssu *K8sStatefulSetUpdate) SetUpdatedAt(t time.Time) *K8sStatefulSetUpdate {
	kssu.mutation.SetUpdatedAt(t)
	return kssu
}

// SetDeletedAt sets the "deleted_at" field.
func (kssu *K8sStatefulSetUpdate) SetDeletedAt(t time.Time) *K8sStatefulSetUpdate {
	kssu.mutation.SetDeletedAt(t)
	return kssu
}

// SetName sets the "name" field.
func (kssu *K8sStatefulSetUpdate) SetName(s string) *K8sStatefulSetUpdate {
	kssu.mutation.SetName(s)
	return kssu
}

// SetK8sClusterId sets the "k8sClusterId" field.
func (kssu *K8sStatefulSetUpdate) SetK8sClusterId(u uint) *K8sStatefulSetUpdate {
	kssu.mutation.ResetK8sClusterId()
	kssu.mutation.SetK8sClusterId(u)
	return kssu
}

// AddK8sClusterId adds u to the "k8sClusterId" field.
func (kssu *K8sStatefulSetUpdate) AddK8sClusterId(u uint) *K8sStatefulSetUpdate {
	kssu.mutation.AddK8sClusterId(u)
	return kssu
}

// SetK8sNamespaceId sets the "k8sNamespaceId" field.
func (kssu *K8sStatefulSetUpdate) SetK8sNamespaceId(u uint) *K8sStatefulSetUpdate {
	kssu.mutation.ResetK8sNamespaceId()
	kssu.mutation.SetK8sNamespaceId(u)
	return kssu
}

// AddK8sNamespaceId adds u to the "k8sNamespaceId" field.
func (kssu *K8sStatefulSetUpdate) AddK8sNamespaceId(u uint) *K8sStatefulSetUpdate {
	kssu.mutation.AddK8sNamespaceId(u)
	return kssu
}

// SetK8sObjectId sets the "k8sObjectId" field.
func (kssu *K8sStatefulSetUpdate) SetK8sObjectId(u uint) *K8sStatefulSetUpdate {
	kssu.mutation.ResetK8sObjectId()
	kssu.mutation.SetK8sObjectId(u)
	return kssu
}

// AddK8sObjectId adds u to the "k8sObjectId" field.
func (kssu *K8sStatefulSetUpdate) AddK8sObjectId(u uint) *K8sStatefulSetUpdate {
	kssu.mutation.AddK8sObjectId(u)
	return kssu
}

// Mutation returns the K8sStatefulSetMutation object of the builder.
func (kssu *K8sStatefulSetUpdate) Mutation() *K8sStatefulSetMutation {
	return kssu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (kssu *K8sStatefulSetUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	kssu.defaults()
	if len(kssu.hooks) == 0 {
		if err = kssu.check(); err != nil {
			return 0, err
		}
		affected, err = kssu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*K8sStatefulSetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kssu.check(); err != nil {
				return 0, err
			}
			kssu.mutation = mutation
			affected, err = kssu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(kssu.hooks) - 1; i >= 0; i-- {
			if kssu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kssu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kssu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (kssu *K8sStatefulSetUpdate) SaveX(ctx context.Context) int {
	affected, err := kssu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (kssu *K8sStatefulSetUpdate) Exec(ctx context.Context) error {
	_, err := kssu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kssu *K8sStatefulSetUpdate) ExecX(ctx context.Context) {
	if err := kssu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kssu *K8sStatefulSetUpdate) defaults() {
	if _, ok := kssu.mutation.UpdatedAt(); !ok {
		v := k8sstatefulset.UpdateDefaultUpdatedAt()
		kssu.mutation.SetUpdatedAt(v)
	}
	if _, ok := kssu.mutation.DeletedAt(); !ok {
		v := k8sstatefulset.UpdateDefaultDeletedAt()
		kssu.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kssu *K8sStatefulSetUpdate) check() error {
	if v, ok := kssu.mutation.Name(); ok {
		if err := k8sstatefulset.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (kssu *K8sStatefulSetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   k8sstatefulset.Table,
			Columns: k8sstatefulset.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: k8sstatefulset.FieldID,
			},
		},
	}
	if ps := kssu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kssu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8sstatefulset.FieldUpdatedAt,
		})
	}
	if value, ok := kssu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8sstatefulset.FieldDeletedAt,
		})
	}
	if value, ok := kssu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: k8sstatefulset.FieldName,
		})
	}
	if value, ok := kssu.mutation.K8sClusterId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sstatefulset.FieldK8sClusterId,
		})
	}
	if value, ok := kssu.mutation.AddedK8sClusterId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sstatefulset.FieldK8sClusterId,
		})
	}
	if value, ok := kssu.mutation.K8sNamespaceId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sstatefulset.FieldK8sNamespaceId,
		})
	}
	if value, ok := kssu.mutation.AddedK8sNamespaceId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sstatefulset.FieldK8sNamespaceId,
		})
	}
	if value, ok := kssu.mutation.K8sObjectId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sstatefulset.FieldK8sObjectId,
		})
	}
	if value, ok := kssu.mutation.AddedK8sObjectId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sstatefulset.FieldK8sObjectId,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, kssu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{k8sstatefulset.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// K8sStatefulSetUpdateOne is the builder for updating a single K8sStatefulSet entity.
type K8sStatefulSetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *K8sStatefulSetMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (kssuo *K8sStatefulSetUpdateOne) SetUpdatedAt(t time.Time) *K8sStatefulSetUpdateOne {
	kssuo.mutation.SetUpdatedAt(t)
	return kssuo
}

// SetDeletedAt sets the "deleted_at" field.
func (kssuo *K8sStatefulSetUpdateOne) SetDeletedAt(t time.Time) *K8sStatefulSetUpdateOne {
	kssuo.mutation.SetDeletedAt(t)
	return kssuo
}

// SetName sets the "name" field.
func (kssuo *K8sStatefulSetUpdateOne) SetName(s string) *K8sStatefulSetUpdateOne {
	kssuo.mutation.SetName(s)
	return kssuo
}

// SetK8sClusterId sets the "k8sClusterId" field.
func (kssuo *K8sStatefulSetUpdateOne) SetK8sClusterId(u uint) *K8sStatefulSetUpdateOne {
	kssuo.mutation.ResetK8sClusterId()
	kssuo.mutation.SetK8sClusterId(u)
	return kssuo
}

// AddK8sClusterId adds u to the "k8sClusterId" field.
func (kssuo *K8sStatefulSetUpdateOne) AddK8sClusterId(u uint) *K8sStatefulSetUpdateOne {
	kssuo.mutation.AddK8sClusterId(u)
	return kssuo
}

// SetK8sNamespaceId sets the "k8sNamespaceId" field.
func (kssuo *K8sStatefulSetUpdateOne) SetK8sNamespaceId(u uint) *K8sStatefulSetUpdateOne {
	kssuo.mutation.ResetK8sNamespaceId()
	kssuo.mutation.SetK8sNamespaceId(u)
	return kssuo
}

// AddK8sNamespaceId adds u to the "k8sNamespaceId" field.
func (kssuo *K8sStatefulSetUpdateOne) AddK8sNamespaceId(u uint) *K8sStatefulSetUpdateOne {
	kssuo.mutation.AddK8sNamespaceId(u)
	return kssuo
}

// SetK8sObjectId sets the "k8sObjectId" field.
func (kssuo *K8sStatefulSetUpdateOne) SetK8sObjectId(u uint) *K8sStatefulSetUpdateOne {
	kssuo.mutation.ResetK8sObjectId()
	kssuo.mutation.SetK8sObjectId(u)
	return kssuo
}

// AddK8sObjectId adds u to the "k8sObjectId" field.
func (kssuo *K8sStatefulSetUpdateOne) AddK8sObjectId(u uint) *K8sStatefulSetUpdateOne {
	kssuo.mutation.AddK8sObjectId(u)
	return kssuo
}

// Mutation returns the K8sStatefulSetMutation object of the builder.
func (kssuo *K8sStatefulSetUpdateOne) Mutation() *K8sStatefulSetMutation {
	return kssuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (kssuo *K8sStatefulSetUpdateOne) Select(field string, fields ...string) *K8sStatefulSetUpdateOne {
	kssuo.fields = append([]string{field}, fields...)
	return kssuo
}

// Save executes the query and returns the updated K8sStatefulSet entity.
func (kssuo *K8sStatefulSetUpdateOne) Save(ctx context.Context) (*K8sStatefulSet, error) {
	var (
		err  error
		node *K8sStatefulSet
	)
	kssuo.defaults()
	if len(kssuo.hooks) == 0 {
		if err = kssuo.check(); err != nil {
			return nil, err
		}
		node, err = kssuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*K8sStatefulSetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kssuo.check(); err != nil {
				return nil, err
			}
			kssuo.mutation = mutation
			node, err = kssuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(kssuo.hooks) - 1; i >= 0; i-- {
			if kssuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kssuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kssuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (kssuo *K8sStatefulSetUpdateOne) SaveX(ctx context.Context) *K8sStatefulSet {
	node, err := kssuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (kssuo *K8sStatefulSetUpdateOne) Exec(ctx context.Context) error {
	_, err := kssuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kssuo *K8sStatefulSetUpdateOne) ExecX(ctx context.Context) {
	if err := kssuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kssuo *K8sStatefulSetUpdateOne) defaults() {
	if _, ok := kssuo.mutation.UpdatedAt(); !ok {
		v := k8sstatefulset.UpdateDefaultUpdatedAt()
		kssuo.mutation.SetUpdatedAt(v)
	}
	if _, ok := kssuo.mutation.DeletedAt(); !ok {
		v := k8sstatefulset.UpdateDefaultDeletedAt()
		kssuo.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kssuo *K8sStatefulSetUpdateOne) check() error {
	if v, ok := kssuo.mutation.Name(); ok {
		if err := k8sstatefulset.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (kssuo *K8sStatefulSetUpdateOne) sqlSave(ctx context.Context) (_node *K8sStatefulSet, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   k8sstatefulset.Table,
			Columns: k8sstatefulset.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: k8sstatefulset.FieldID,
			},
		},
	}
	id, ok := kssuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing K8sStatefulSet.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := kssuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, k8sstatefulset.FieldID)
		for _, f := range fields {
			if !k8sstatefulset.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != k8sstatefulset.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := kssuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kssuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8sstatefulset.FieldUpdatedAt,
		})
	}
	if value, ok := kssuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8sstatefulset.FieldDeletedAt,
		})
	}
	if value, ok := kssuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: k8sstatefulset.FieldName,
		})
	}
	if value, ok := kssuo.mutation.K8sClusterId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sstatefulset.FieldK8sClusterId,
		})
	}
	if value, ok := kssuo.mutation.AddedK8sClusterId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sstatefulset.FieldK8sClusterId,
		})
	}
	if value, ok := kssuo.mutation.K8sNamespaceId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sstatefulset.FieldK8sNamespaceId,
		})
	}
	if value, ok := kssuo.mutation.AddedK8sNamespaceId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sstatefulset.FieldK8sNamespaceId,
		})
	}
	if value, ok := kssuo.mutation.K8sObjectId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sstatefulset.FieldK8sObjectId,
		})
	}
	if value, ok := kssuo.mutation.AddedK8sObjectId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sstatefulset.FieldK8sObjectId,
		})
	}
	_node = &K8sStatefulSet{config: kssuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, kssuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{k8sstatefulset.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
