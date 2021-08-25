// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/k8sreplicaset"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/predicate"
)

// K8sReplicaSetUpdate is the builder for updating K8sReplicaSet entities.
type K8sReplicaSetUpdate struct {
	config
	hooks    []Hook
	mutation *K8sReplicaSetMutation
}

// Where appends a list predicates to the K8sReplicaSetUpdate builder.
func (krsu *K8sReplicaSetUpdate) Where(ps ...predicate.K8sReplicaSet) *K8sReplicaSetUpdate {
	krsu.mutation.Where(ps...)
	return krsu
}

// SetUpdatedAt sets the "updated_at" field.
func (krsu *K8sReplicaSetUpdate) SetUpdatedAt(t time.Time) *K8sReplicaSetUpdate {
	krsu.mutation.SetUpdatedAt(t)
	return krsu
}

// SetDeletedAt sets the "deleted_at" field.
func (krsu *K8sReplicaSetUpdate) SetDeletedAt(t time.Time) *K8sReplicaSetUpdate {
	krsu.mutation.SetDeletedAt(t)
	return krsu
}

// SetName sets the "name" field.
func (krsu *K8sReplicaSetUpdate) SetName(s string) *K8sReplicaSetUpdate {
	krsu.mutation.SetName(s)
	return krsu
}

// SetK8sClusterId sets the "k8sClusterId" field.
func (krsu *K8sReplicaSetUpdate) SetK8sClusterId(u uint) *K8sReplicaSetUpdate {
	krsu.mutation.ResetK8sClusterId()
	krsu.mutation.SetK8sClusterId(u)
	return krsu
}

// AddK8sClusterId adds u to the "k8sClusterId" field.
func (krsu *K8sReplicaSetUpdate) AddK8sClusterId(u uint) *K8sReplicaSetUpdate {
	krsu.mutation.AddK8sClusterId(u)
	return krsu
}

// SetK8sNamespaceId sets the "k8sNamespaceId" field.
func (krsu *K8sReplicaSetUpdate) SetK8sNamespaceId(u uint) *K8sReplicaSetUpdate {
	krsu.mutation.ResetK8sNamespaceId()
	krsu.mutation.SetK8sNamespaceId(u)
	return krsu
}

// AddK8sNamespaceId adds u to the "k8sNamespaceId" field.
func (krsu *K8sReplicaSetUpdate) AddK8sNamespaceId(u uint) *K8sReplicaSetUpdate {
	krsu.mutation.AddK8sNamespaceId(u)
	return krsu
}

// SetK8sDeploymentId sets the "k8sDeploymentId" field.
func (krsu *K8sReplicaSetUpdate) SetK8sDeploymentId(u uint) *K8sReplicaSetUpdate {
	krsu.mutation.ResetK8sDeploymentId()
	krsu.mutation.SetK8sDeploymentId(u)
	return krsu
}

// AddK8sDeploymentId adds u to the "k8sDeploymentId" field.
func (krsu *K8sReplicaSetUpdate) AddK8sDeploymentId(u uint) *K8sReplicaSetUpdate {
	krsu.mutation.AddK8sDeploymentId(u)
	return krsu
}

// SetK8sObjectId sets the "k8sObjectId" field.
func (krsu *K8sReplicaSetUpdate) SetK8sObjectId(u uint) *K8sReplicaSetUpdate {
	krsu.mutation.ResetK8sObjectId()
	krsu.mutation.SetK8sObjectId(u)
	return krsu
}

// AddK8sObjectId adds u to the "k8sObjectId" field.
func (krsu *K8sReplicaSetUpdate) AddK8sObjectId(u uint) *K8sReplicaSetUpdate {
	krsu.mutation.AddK8sObjectId(u)
	return krsu
}

// Mutation returns the K8sReplicaSetMutation object of the builder.
func (krsu *K8sReplicaSetUpdate) Mutation() *K8sReplicaSetMutation {
	return krsu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (krsu *K8sReplicaSetUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	krsu.defaults()
	if len(krsu.hooks) == 0 {
		if err = krsu.check(); err != nil {
			return 0, err
		}
		affected, err = krsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*K8sReplicaSetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = krsu.check(); err != nil {
				return 0, err
			}
			krsu.mutation = mutation
			affected, err = krsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(krsu.hooks) - 1; i >= 0; i-- {
			if krsu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = krsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, krsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (krsu *K8sReplicaSetUpdate) SaveX(ctx context.Context) int {
	affected, err := krsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (krsu *K8sReplicaSetUpdate) Exec(ctx context.Context) error {
	_, err := krsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (krsu *K8sReplicaSetUpdate) ExecX(ctx context.Context) {
	if err := krsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (krsu *K8sReplicaSetUpdate) defaults() {
	if _, ok := krsu.mutation.UpdatedAt(); !ok {
		v := k8sreplicaset.UpdateDefaultUpdatedAt()
		krsu.mutation.SetUpdatedAt(v)
	}
	if _, ok := krsu.mutation.DeletedAt(); !ok {
		v := k8sreplicaset.UpdateDefaultDeletedAt()
		krsu.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (krsu *K8sReplicaSetUpdate) check() error {
	if v, ok := krsu.mutation.Name(); ok {
		if err := k8sreplicaset.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (krsu *K8sReplicaSetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   k8sreplicaset.Table,
			Columns: k8sreplicaset.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: k8sreplicaset.FieldID,
			},
		},
	}
	if ps := krsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := krsu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8sreplicaset.FieldUpdatedAt,
		})
	}
	if value, ok := krsu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8sreplicaset.FieldDeletedAt,
		})
	}
	if value, ok := krsu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: k8sreplicaset.FieldName,
		})
	}
	if value, ok := krsu.mutation.K8sClusterId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sreplicaset.FieldK8sClusterId,
		})
	}
	if value, ok := krsu.mutation.AddedK8sClusterId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sreplicaset.FieldK8sClusterId,
		})
	}
	if value, ok := krsu.mutation.K8sNamespaceId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sreplicaset.FieldK8sNamespaceId,
		})
	}
	if value, ok := krsu.mutation.AddedK8sNamespaceId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sreplicaset.FieldK8sNamespaceId,
		})
	}
	if value, ok := krsu.mutation.K8sDeploymentId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sreplicaset.FieldK8sDeploymentId,
		})
	}
	if value, ok := krsu.mutation.AddedK8sDeploymentId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sreplicaset.FieldK8sDeploymentId,
		})
	}
	if value, ok := krsu.mutation.K8sObjectId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sreplicaset.FieldK8sObjectId,
		})
	}
	if value, ok := krsu.mutation.AddedK8sObjectId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sreplicaset.FieldK8sObjectId,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, krsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{k8sreplicaset.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// K8sReplicaSetUpdateOne is the builder for updating a single K8sReplicaSet entity.
type K8sReplicaSetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *K8sReplicaSetMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (krsuo *K8sReplicaSetUpdateOne) SetUpdatedAt(t time.Time) *K8sReplicaSetUpdateOne {
	krsuo.mutation.SetUpdatedAt(t)
	return krsuo
}

// SetDeletedAt sets the "deleted_at" field.
func (krsuo *K8sReplicaSetUpdateOne) SetDeletedAt(t time.Time) *K8sReplicaSetUpdateOne {
	krsuo.mutation.SetDeletedAt(t)
	return krsuo
}

// SetName sets the "name" field.
func (krsuo *K8sReplicaSetUpdateOne) SetName(s string) *K8sReplicaSetUpdateOne {
	krsuo.mutation.SetName(s)
	return krsuo
}

// SetK8sClusterId sets the "k8sClusterId" field.
func (krsuo *K8sReplicaSetUpdateOne) SetK8sClusterId(u uint) *K8sReplicaSetUpdateOne {
	krsuo.mutation.ResetK8sClusterId()
	krsuo.mutation.SetK8sClusterId(u)
	return krsuo
}

// AddK8sClusterId adds u to the "k8sClusterId" field.
func (krsuo *K8sReplicaSetUpdateOne) AddK8sClusterId(u uint) *K8sReplicaSetUpdateOne {
	krsuo.mutation.AddK8sClusterId(u)
	return krsuo
}

// SetK8sNamespaceId sets the "k8sNamespaceId" field.
func (krsuo *K8sReplicaSetUpdateOne) SetK8sNamespaceId(u uint) *K8sReplicaSetUpdateOne {
	krsuo.mutation.ResetK8sNamespaceId()
	krsuo.mutation.SetK8sNamespaceId(u)
	return krsuo
}

// AddK8sNamespaceId adds u to the "k8sNamespaceId" field.
func (krsuo *K8sReplicaSetUpdateOne) AddK8sNamespaceId(u uint) *K8sReplicaSetUpdateOne {
	krsuo.mutation.AddK8sNamespaceId(u)
	return krsuo
}

// SetK8sDeploymentId sets the "k8sDeploymentId" field.
func (krsuo *K8sReplicaSetUpdateOne) SetK8sDeploymentId(u uint) *K8sReplicaSetUpdateOne {
	krsuo.mutation.ResetK8sDeploymentId()
	krsuo.mutation.SetK8sDeploymentId(u)
	return krsuo
}

// AddK8sDeploymentId adds u to the "k8sDeploymentId" field.
func (krsuo *K8sReplicaSetUpdateOne) AddK8sDeploymentId(u uint) *K8sReplicaSetUpdateOne {
	krsuo.mutation.AddK8sDeploymentId(u)
	return krsuo
}

// SetK8sObjectId sets the "k8sObjectId" field.
func (krsuo *K8sReplicaSetUpdateOne) SetK8sObjectId(u uint) *K8sReplicaSetUpdateOne {
	krsuo.mutation.ResetK8sObjectId()
	krsuo.mutation.SetK8sObjectId(u)
	return krsuo
}

// AddK8sObjectId adds u to the "k8sObjectId" field.
func (krsuo *K8sReplicaSetUpdateOne) AddK8sObjectId(u uint) *K8sReplicaSetUpdateOne {
	krsuo.mutation.AddK8sObjectId(u)
	return krsuo
}

// Mutation returns the K8sReplicaSetMutation object of the builder.
func (krsuo *K8sReplicaSetUpdateOne) Mutation() *K8sReplicaSetMutation {
	return krsuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (krsuo *K8sReplicaSetUpdateOne) Select(field string, fields ...string) *K8sReplicaSetUpdateOne {
	krsuo.fields = append([]string{field}, fields...)
	return krsuo
}

// Save executes the query and returns the updated K8sReplicaSet entity.
func (krsuo *K8sReplicaSetUpdateOne) Save(ctx context.Context) (*K8sReplicaSet, error) {
	var (
		err  error
		node *K8sReplicaSet
	)
	krsuo.defaults()
	if len(krsuo.hooks) == 0 {
		if err = krsuo.check(); err != nil {
			return nil, err
		}
		node, err = krsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*K8sReplicaSetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = krsuo.check(); err != nil {
				return nil, err
			}
			krsuo.mutation = mutation
			node, err = krsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(krsuo.hooks) - 1; i >= 0; i-- {
			if krsuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = krsuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, krsuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (krsuo *K8sReplicaSetUpdateOne) SaveX(ctx context.Context) *K8sReplicaSet {
	node, err := krsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (krsuo *K8sReplicaSetUpdateOne) Exec(ctx context.Context) error {
	_, err := krsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (krsuo *K8sReplicaSetUpdateOne) ExecX(ctx context.Context) {
	if err := krsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (krsuo *K8sReplicaSetUpdateOne) defaults() {
	if _, ok := krsuo.mutation.UpdatedAt(); !ok {
		v := k8sreplicaset.UpdateDefaultUpdatedAt()
		krsuo.mutation.SetUpdatedAt(v)
	}
	if _, ok := krsuo.mutation.DeletedAt(); !ok {
		v := k8sreplicaset.UpdateDefaultDeletedAt()
		krsuo.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (krsuo *K8sReplicaSetUpdateOne) check() error {
	if v, ok := krsuo.mutation.Name(); ok {
		if err := k8sreplicaset.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (krsuo *K8sReplicaSetUpdateOne) sqlSave(ctx context.Context) (_node *K8sReplicaSet, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   k8sreplicaset.Table,
			Columns: k8sreplicaset.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: k8sreplicaset.FieldID,
			},
		},
	}
	id, ok := krsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing K8sReplicaSet.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := krsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, k8sreplicaset.FieldID)
		for _, f := range fields {
			if !k8sreplicaset.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != k8sreplicaset.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := krsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := krsuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8sreplicaset.FieldUpdatedAt,
		})
	}
	if value, ok := krsuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8sreplicaset.FieldDeletedAt,
		})
	}
	if value, ok := krsuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: k8sreplicaset.FieldName,
		})
	}
	if value, ok := krsuo.mutation.K8sClusterId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sreplicaset.FieldK8sClusterId,
		})
	}
	if value, ok := krsuo.mutation.AddedK8sClusterId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sreplicaset.FieldK8sClusterId,
		})
	}
	if value, ok := krsuo.mutation.K8sNamespaceId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sreplicaset.FieldK8sNamespaceId,
		})
	}
	if value, ok := krsuo.mutation.AddedK8sNamespaceId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sreplicaset.FieldK8sNamespaceId,
		})
	}
	if value, ok := krsuo.mutation.K8sDeploymentId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sreplicaset.FieldK8sDeploymentId,
		})
	}
	if value, ok := krsuo.mutation.AddedK8sDeploymentId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sreplicaset.FieldK8sDeploymentId,
		})
	}
	if value, ok := krsuo.mutation.K8sObjectId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sreplicaset.FieldK8sObjectId,
		})
	}
	if value, ok := krsuo.mutation.AddedK8sObjectId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8sreplicaset.FieldK8sObjectId,
		})
	}
	_node = &K8sReplicaSet{config: krsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, krsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{k8sreplicaset.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
