// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/k8spod"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/predicate"
)

// K8sPodUpdate is the builder for updating K8sPod entities.
type K8sPodUpdate struct {
	config
	hooks    []Hook
	mutation *K8sPodMutation
}

// Where appends a list predicates to the K8sPodUpdate builder.
func (kpu *K8sPodUpdate) Where(ps ...predicate.K8sPod) *K8sPodUpdate {
	kpu.mutation.Where(ps...)
	return kpu
}

// SetUpdatedAt sets the "updated_at" field.
func (kpu *K8sPodUpdate) SetUpdatedAt(t time.Time) *K8sPodUpdate {
	kpu.mutation.SetUpdatedAt(t)
	return kpu
}

// SetDeletedAt sets the "deleted_at" field.
func (kpu *K8sPodUpdate) SetDeletedAt(t time.Time) *K8sPodUpdate {
	kpu.mutation.SetDeletedAt(t)
	return kpu
}

// SetName sets the "name" field.
func (kpu *K8sPodUpdate) SetName(s string) *K8sPodUpdate {
	kpu.mutation.SetName(s)
	return kpu
}

// SetQos sets the "qos" field.
func (kpu *K8sPodUpdate) SetQos(s string) *K8sPodUpdate {
	kpu.mutation.SetQos(s)
	return kpu
}

// SetK8sClusterId sets the "k8sClusterId" field.
func (kpu *K8sPodUpdate) SetK8sClusterId(u uint) *K8sPodUpdate {
	kpu.mutation.ResetK8sClusterId()
	kpu.mutation.SetK8sClusterId(u)
	return kpu
}

// AddK8sClusterId adds u to the "k8sClusterId" field.
func (kpu *K8sPodUpdate) AddK8sClusterId(u uint) *K8sPodUpdate {
	kpu.mutation.AddK8sClusterId(u)
	return kpu
}

// SetK8sNamespaceId sets the "k8sNamespaceId" field.
func (kpu *K8sPodUpdate) SetK8sNamespaceId(u uint) *K8sPodUpdate {
	kpu.mutation.ResetK8sNamespaceId()
	kpu.mutation.SetK8sNamespaceId(u)
	return kpu
}

// AddK8sNamespaceId adds u to the "k8sNamespaceId" field.
func (kpu *K8sPodUpdate) AddK8sNamespaceId(u uint) *K8sPodUpdate {
	kpu.mutation.AddK8sNamespaceId(u)
	return kpu
}

// SetK8sObjectId sets the "k8sObjectId" field.
func (kpu *K8sPodUpdate) SetK8sObjectId(u uint) *K8sPodUpdate {
	kpu.mutation.ResetK8sObjectId()
	kpu.mutation.SetK8sObjectId(u)
	return kpu
}

// AddK8sObjectId adds u to the "k8sObjectId" field.
func (kpu *K8sPodUpdate) AddK8sObjectId(u uint) *K8sPodUpdate {
	kpu.mutation.AddK8sObjectId(u)
	return kpu
}

// Mutation returns the K8sPodMutation object of the builder.
func (kpu *K8sPodUpdate) Mutation() *K8sPodMutation {
	return kpu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (kpu *K8sPodUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	kpu.defaults()
	if len(kpu.hooks) == 0 {
		if err = kpu.check(); err != nil {
			return 0, err
		}
		affected, err = kpu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*K8sPodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kpu.check(); err != nil {
				return 0, err
			}
			kpu.mutation = mutation
			affected, err = kpu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(kpu.hooks) - 1; i >= 0; i-- {
			if kpu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kpu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kpu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (kpu *K8sPodUpdate) SaveX(ctx context.Context) int {
	affected, err := kpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (kpu *K8sPodUpdate) Exec(ctx context.Context) error {
	_, err := kpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kpu *K8sPodUpdate) ExecX(ctx context.Context) {
	if err := kpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kpu *K8sPodUpdate) defaults() {
	if _, ok := kpu.mutation.UpdatedAt(); !ok {
		v := k8spod.UpdateDefaultUpdatedAt()
		kpu.mutation.SetUpdatedAt(v)
	}
	if _, ok := kpu.mutation.DeletedAt(); !ok {
		v := k8spod.UpdateDefaultDeletedAt()
		kpu.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kpu *K8sPodUpdate) check() error {
	if v, ok := kpu.mutation.Name(); ok {
		if err := k8spod.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := kpu.mutation.Qos(); ok {
		if err := k8spod.QosValidator(v); err != nil {
			return &ValidationError{Name: "qos", err: fmt.Errorf("ent: validator failed for field \"qos\": %w", err)}
		}
	}
	return nil
}

func (kpu *K8sPodUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   k8spod.Table,
			Columns: k8spod.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: k8spod.FieldID,
			},
		},
	}
	if ps := kpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kpu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8spod.FieldUpdatedAt,
		})
	}
	if value, ok := kpu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8spod.FieldDeletedAt,
		})
	}
	if value, ok := kpu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: k8spod.FieldName,
		})
	}
	if value, ok := kpu.mutation.Qos(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: k8spod.FieldQos,
		})
	}
	if value, ok := kpu.mutation.K8sClusterId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8spod.FieldK8sClusterId,
		})
	}
	if value, ok := kpu.mutation.AddedK8sClusterId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8spod.FieldK8sClusterId,
		})
	}
	if value, ok := kpu.mutation.K8sNamespaceId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8spod.FieldK8sNamespaceId,
		})
	}
	if value, ok := kpu.mutation.AddedK8sNamespaceId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8spod.FieldK8sNamespaceId,
		})
	}
	if value, ok := kpu.mutation.K8sObjectId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8spod.FieldK8sObjectId,
		})
	}
	if value, ok := kpu.mutation.AddedK8sObjectId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8spod.FieldK8sObjectId,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, kpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{k8spod.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// K8sPodUpdateOne is the builder for updating a single K8sPod entity.
type K8sPodUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *K8sPodMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (kpuo *K8sPodUpdateOne) SetUpdatedAt(t time.Time) *K8sPodUpdateOne {
	kpuo.mutation.SetUpdatedAt(t)
	return kpuo
}

// SetDeletedAt sets the "deleted_at" field.
func (kpuo *K8sPodUpdateOne) SetDeletedAt(t time.Time) *K8sPodUpdateOne {
	kpuo.mutation.SetDeletedAt(t)
	return kpuo
}

// SetName sets the "name" field.
func (kpuo *K8sPodUpdateOne) SetName(s string) *K8sPodUpdateOne {
	kpuo.mutation.SetName(s)
	return kpuo
}

// SetQos sets the "qos" field.
func (kpuo *K8sPodUpdateOne) SetQos(s string) *K8sPodUpdateOne {
	kpuo.mutation.SetQos(s)
	return kpuo
}

// SetK8sClusterId sets the "k8sClusterId" field.
func (kpuo *K8sPodUpdateOne) SetK8sClusterId(u uint) *K8sPodUpdateOne {
	kpuo.mutation.ResetK8sClusterId()
	kpuo.mutation.SetK8sClusterId(u)
	return kpuo
}

// AddK8sClusterId adds u to the "k8sClusterId" field.
func (kpuo *K8sPodUpdateOne) AddK8sClusterId(u uint) *K8sPodUpdateOne {
	kpuo.mutation.AddK8sClusterId(u)
	return kpuo
}

// SetK8sNamespaceId sets the "k8sNamespaceId" field.
func (kpuo *K8sPodUpdateOne) SetK8sNamespaceId(u uint) *K8sPodUpdateOne {
	kpuo.mutation.ResetK8sNamespaceId()
	kpuo.mutation.SetK8sNamespaceId(u)
	return kpuo
}

// AddK8sNamespaceId adds u to the "k8sNamespaceId" field.
func (kpuo *K8sPodUpdateOne) AddK8sNamespaceId(u uint) *K8sPodUpdateOne {
	kpuo.mutation.AddK8sNamespaceId(u)
	return kpuo
}

// SetK8sObjectId sets the "k8sObjectId" field.
func (kpuo *K8sPodUpdateOne) SetK8sObjectId(u uint) *K8sPodUpdateOne {
	kpuo.mutation.ResetK8sObjectId()
	kpuo.mutation.SetK8sObjectId(u)
	return kpuo
}

// AddK8sObjectId adds u to the "k8sObjectId" field.
func (kpuo *K8sPodUpdateOne) AddK8sObjectId(u uint) *K8sPodUpdateOne {
	kpuo.mutation.AddK8sObjectId(u)
	return kpuo
}

// Mutation returns the K8sPodMutation object of the builder.
func (kpuo *K8sPodUpdateOne) Mutation() *K8sPodMutation {
	return kpuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (kpuo *K8sPodUpdateOne) Select(field string, fields ...string) *K8sPodUpdateOne {
	kpuo.fields = append([]string{field}, fields...)
	return kpuo
}

// Save executes the query and returns the updated K8sPod entity.
func (kpuo *K8sPodUpdateOne) Save(ctx context.Context) (*K8sPod, error) {
	var (
		err  error
		node *K8sPod
	)
	kpuo.defaults()
	if len(kpuo.hooks) == 0 {
		if err = kpuo.check(); err != nil {
			return nil, err
		}
		node, err = kpuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*K8sPodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kpuo.check(); err != nil {
				return nil, err
			}
			kpuo.mutation = mutation
			node, err = kpuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(kpuo.hooks) - 1; i >= 0; i-- {
			if kpuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kpuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kpuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (kpuo *K8sPodUpdateOne) SaveX(ctx context.Context) *K8sPod {
	node, err := kpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (kpuo *K8sPodUpdateOne) Exec(ctx context.Context) error {
	_, err := kpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kpuo *K8sPodUpdateOne) ExecX(ctx context.Context) {
	if err := kpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kpuo *K8sPodUpdateOne) defaults() {
	if _, ok := kpuo.mutation.UpdatedAt(); !ok {
		v := k8spod.UpdateDefaultUpdatedAt()
		kpuo.mutation.SetUpdatedAt(v)
	}
	if _, ok := kpuo.mutation.DeletedAt(); !ok {
		v := k8spod.UpdateDefaultDeletedAt()
		kpuo.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kpuo *K8sPodUpdateOne) check() error {
	if v, ok := kpuo.mutation.Name(); ok {
		if err := k8spod.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := kpuo.mutation.Qos(); ok {
		if err := k8spod.QosValidator(v); err != nil {
			return &ValidationError{Name: "qos", err: fmt.Errorf("ent: validator failed for field \"qos\": %w", err)}
		}
	}
	return nil
}

func (kpuo *K8sPodUpdateOne) sqlSave(ctx context.Context) (_node *K8sPod, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   k8spod.Table,
			Columns: k8spod.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: k8spod.FieldID,
			},
		},
	}
	id, ok := kpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing K8sPod.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := kpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, k8spod.FieldID)
		for _, f := range fields {
			if !k8spod.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != k8spod.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := kpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kpuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8spod.FieldUpdatedAt,
		})
	}
	if value, ok := kpuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8spod.FieldDeletedAt,
		})
	}
	if value, ok := kpuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: k8spod.FieldName,
		})
	}
	if value, ok := kpuo.mutation.Qos(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: k8spod.FieldQos,
		})
	}
	if value, ok := kpuo.mutation.K8sClusterId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8spod.FieldK8sClusterId,
		})
	}
	if value, ok := kpuo.mutation.AddedK8sClusterId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8spod.FieldK8sClusterId,
		})
	}
	if value, ok := kpuo.mutation.K8sNamespaceId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8spod.FieldK8sNamespaceId,
		})
	}
	if value, ok := kpuo.mutation.AddedK8sNamespaceId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8spod.FieldK8sNamespaceId,
		})
	}
	if value, ok := kpuo.mutation.K8sObjectId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8spod.FieldK8sObjectId,
		})
	}
	if value, ok := kpuo.mutation.AddedK8sObjectId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8spod.FieldK8sObjectId,
		})
	}
	_node = &K8sPod{config: kpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, kpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{k8spod.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
