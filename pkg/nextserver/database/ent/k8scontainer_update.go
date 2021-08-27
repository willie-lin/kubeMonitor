// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/k8scontainer"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/predicate"
)

// K8sContainerUpdate is the builder for updating K8sContainer entities.
type K8sContainerUpdate struct {
	config
	hooks    []Hook
	mutation *K8sContainerMutation
}

// Where appends a list predicates to the K8sContainerUpdate builder.
func (kcu *K8sContainerUpdate) Where(ps ...predicate.K8sContainer) *K8sContainerUpdate {
	kcu.mutation.Where(ps...)
	return kcu
}

// SetUpdatedAt sets the "updated_at" field.
func (kcu *K8sContainerUpdate) SetUpdatedAt(t time.Time) *K8sContainerUpdate {
	kcu.mutation.SetUpdatedAt(t)
	return kcu
}

// SetDeletedAt sets the "deleted_at" field.
func (kcu *K8sContainerUpdate) SetDeletedAt(t time.Time) *K8sContainerUpdate {
	kcu.mutation.SetDeletedAt(t)
	return kcu
}

// SetName sets the "name" field.
func (kcu *K8sContainerUpdate) SetName(s string) *K8sContainerUpdate {
	kcu.mutation.SetName(s)
	return kcu
}

// SetImage sets the "image" field.
func (kcu *K8sContainerUpdate) SetImage(s string) *K8sContainerUpdate {
	kcu.mutation.SetImage(s)
	return kcu
}

// SetContainerType sets the "containerType" field.
func (kcu *K8sContainerUpdate) SetContainerType(s string) *K8sContainerUpdate {
	kcu.mutation.SetContainerType(s)
	return kcu
}

// SetContainerId sets the "containerId" field.
func (kcu *K8sContainerUpdate) SetContainerId(s string) *K8sContainerUpdate {
	kcu.mutation.SetContainerId(s)
	return kcu
}

// SetK8sClusterId sets the "k8sClusterId" field.
func (kcu *K8sContainerUpdate) SetK8sClusterId(u uint) *K8sContainerUpdate {
	kcu.mutation.ResetK8sClusterId()
	kcu.mutation.SetK8sClusterId(u)
	return kcu
}

// AddK8sClusterId adds u to the "k8sClusterId" field.
func (kcu *K8sContainerUpdate) AddK8sClusterId(u uint) *K8sContainerUpdate {
	kcu.mutation.AddK8sClusterId(u)
	return kcu
}

// SetK8sNamespaceId sets the "k8sNamespaceId" field.
func (kcu *K8sContainerUpdate) SetK8sNamespaceId(u uint) *K8sContainerUpdate {
	kcu.mutation.ResetK8sNamespaceId()
	kcu.mutation.SetK8sNamespaceId(u)
	return kcu
}

// AddK8sNamespaceId adds u to the "k8sNamespaceId" field.
func (kcu *K8sContainerUpdate) AddK8sNamespaceId(u uint) *K8sContainerUpdate {
	kcu.mutation.AddK8sNamespaceId(u)
	return kcu
}

// SetK8sPodId sets the "K8sPodId" field.
func (kcu *K8sContainerUpdate) SetK8sPodId(u uint) *K8sContainerUpdate {
	kcu.mutation.ResetK8sPodId()
	kcu.mutation.SetK8sPodId(u)
	return kcu
}

// AddK8sPodId adds u to the "K8sPodId" field.
func (kcu *K8sContainerUpdate) AddK8sPodId(u uint) *K8sContainerUpdate {
	kcu.mutation.AddK8sPodId(u)
	return kcu
}

// Mutation returns the K8sContainerMutation object of the builder.
func (kcu *K8sContainerUpdate) Mutation() *K8sContainerMutation {
	return kcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (kcu *K8sContainerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	kcu.defaults()
	if len(kcu.hooks) == 0 {
		if err = kcu.check(); err != nil {
			return 0, err
		}
		affected, err = kcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*K8sContainerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kcu.check(); err != nil {
				return 0, err
			}
			kcu.mutation = mutation
			affected, err = kcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(kcu.hooks) - 1; i >= 0; i-- {
			if kcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (kcu *K8sContainerUpdate) SaveX(ctx context.Context) int {
	affected, err := kcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (kcu *K8sContainerUpdate) Exec(ctx context.Context) error {
	_, err := kcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kcu *K8sContainerUpdate) ExecX(ctx context.Context) {
	if err := kcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kcu *K8sContainerUpdate) defaults() {
	if _, ok := kcu.mutation.UpdatedAt(); !ok {
		v := k8scontainer.UpdateDefaultUpdatedAt()
		kcu.mutation.SetUpdatedAt(v)
	}
	if _, ok := kcu.mutation.DeletedAt(); !ok {
		v := k8scontainer.UpdateDefaultDeletedAt()
		kcu.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kcu *K8sContainerUpdate) check() error {
	if v, ok := kcu.mutation.Name(); ok {
		if err := k8scontainer.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := kcu.mutation.Image(); ok {
		if err := k8scontainer.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf("ent: validator failed for field \"image\": %w", err)}
		}
	}
	if v, ok := kcu.mutation.ContainerType(); ok {
		if err := k8scontainer.ContainerTypeValidator(v); err != nil {
			return &ValidationError{Name: "containerType", err: fmt.Errorf("ent: validator failed for field \"containerType\": %w", err)}
		}
	}
	if v, ok := kcu.mutation.ContainerId(); ok {
		if err := k8scontainer.ContainerIdValidator(v); err != nil {
			return &ValidationError{Name: "containerId", err: fmt.Errorf("ent: validator failed for field \"containerId\": %w", err)}
		}
	}
	return nil
}

func (kcu *K8sContainerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   k8scontainer.Table,
			Columns: k8scontainer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: k8scontainer.FieldID,
			},
		},
	}
	if ps := kcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kcu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8scontainer.FieldUpdatedAt,
		})
	}
	if value, ok := kcu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8scontainer.FieldDeletedAt,
		})
	}
	if value, ok := kcu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: k8scontainer.FieldName,
		})
	}
	if value, ok := kcu.mutation.Image(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: k8scontainer.FieldImage,
		})
	}
	if value, ok := kcu.mutation.ContainerType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: k8scontainer.FieldContainerType,
		})
	}
	if value, ok := kcu.mutation.ContainerId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: k8scontainer.FieldContainerId,
		})
	}
	if value, ok := kcu.mutation.K8sClusterId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8scontainer.FieldK8sClusterId,
		})
	}
	if value, ok := kcu.mutation.AddedK8sClusterId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8scontainer.FieldK8sClusterId,
		})
	}
	if value, ok := kcu.mutation.K8sNamespaceId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8scontainer.FieldK8sNamespaceId,
		})
	}
	if value, ok := kcu.mutation.AddedK8sNamespaceId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8scontainer.FieldK8sNamespaceId,
		})
	}
	if value, ok := kcu.mutation.K8sPodId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8scontainer.FieldK8sPodId,
		})
	}
	if value, ok := kcu.mutation.AddedK8sPodId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8scontainer.FieldK8sPodId,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, kcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{k8scontainer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// K8sContainerUpdateOne is the builder for updating a single K8sContainer entity.
type K8sContainerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *K8sContainerMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (kcuo *K8sContainerUpdateOne) SetUpdatedAt(t time.Time) *K8sContainerUpdateOne {
	kcuo.mutation.SetUpdatedAt(t)
	return kcuo
}

// SetDeletedAt sets the "deleted_at" field.
func (kcuo *K8sContainerUpdateOne) SetDeletedAt(t time.Time) *K8sContainerUpdateOne {
	kcuo.mutation.SetDeletedAt(t)
	return kcuo
}

// SetName sets the "name" field.
func (kcuo *K8sContainerUpdateOne) SetName(s string) *K8sContainerUpdateOne {
	kcuo.mutation.SetName(s)
	return kcuo
}

// SetImage sets the "image" field.
func (kcuo *K8sContainerUpdateOne) SetImage(s string) *K8sContainerUpdateOne {
	kcuo.mutation.SetImage(s)
	return kcuo
}

// SetContainerType sets the "containerType" field.
func (kcuo *K8sContainerUpdateOne) SetContainerType(s string) *K8sContainerUpdateOne {
	kcuo.mutation.SetContainerType(s)
	return kcuo
}

// SetContainerId sets the "containerId" field.
func (kcuo *K8sContainerUpdateOne) SetContainerId(s string) *K8sContainerUpdateOne {
	kcuo.mutation.SetContainerId(s)
	return kcuo
}

// SetK8sClusterId sets the "k8sClusterId" field.
func (kcuo *K8sContainerUpdateOne) SetK8sClusterId(u uint) *K8sContainerUpdateOne {
	kcuo.mutation.ResetK8sClusterId()
	kcuo.mutation.SetK8sClusterId(u)
	return kcuo
}

// AddK8sClusterId adds u to the "k8sClusterId" field.
func (kcuo *K8sContainerUpdateOne) AddK8sClusterId(u uint) *K8sContainerUpdateOne {
	kcuo.mutation.AddK8sClusterId(u)
	return kcuo
}

// SetK8sNamespaceId sets the "k8sNamespaceId" field.
func (kcuo *K8sContainerUpdateOne) SetK8sNamespaceId(u uint) *K8sContainerUpdateOne {
	kcuo.mutation.ResetK8sNamespaceId()
	kcuo.mutation.SetK8sNamespaceId(u)
	return kcuo
}

// AddK8sNamespaceId adds u to the "k8sNamespaceId" field.
func (kcuo *K8sContainerUpdateOne) AddK8sNamespaceId(u uint) *K8sContainerUpdateOne {
	kcuo.mutation.AddK8sNamespaceId(u)
	return kcuo
}

// SetK8sPodId sets the "K8sPodId" field.
func (kcuo *K8sContainerUpdateOne) SetK8sPodId(u uint) *K8sContainerUpdateOne {
	kcuo.mutation.ResetK8sPodId()
	kcuo.mutation.SetK8sPodId(u)
	return kcuo
}

// AddK8sPodId adds u to the "K8sPodId" field.
func (kcuo *K8sContainerUpdateOne) AddK8sPodId(u uint) *K8sContainerUpdateOne {
	kcuo.mutation.AddK8sPodId(u)
	return kcuo
}

// Mutation returns the K8sContainerMutation object of the builder.
func (kcuo *K8sContainerUpdateOne) Mutation() *K8sContainerMutation {
	return kcuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (kcuo *K8sContainerUpdateOne) Select(field string, fields ...string) *K8sContainerUpdateOne {
	kcuo.fields = append([]string{field}, fields...)
	return kcuo
}

// Save executes the query and returns the updated K8sContainer entity.
func (kcuo *K8sContainerUpdateOne) Save(ctx context.Context) (*K8sContainer, error) {
	var (
		err  error
		node *K8sContainer
	)
	kcuo.defaults()
	if len(kcuo.hooks) == 0 {
		if err = kcuo.check(); err != nil {
			return nil, err
		}
		node, err = kcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*K8sContainerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kcuo.check(); err != nil {
				return nil, err
			}
			kcuo.mutation = mutation
			node, err = kcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(kcuo.hooks) - 1; i >= 0; i-- {
			if kcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (kcuo *K8sContainerUpdateOne) SaveX(ctx context.Context) *K8sContainer {
	node, err := kcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (kcuo *K8sContainerUpdateOne) Exec(ctx context.Context) error {
	_, err := kcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kcuo *K8sContainerUpdateOne) ExecX(ctx context.Context) {
	if err := kcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kcuo *K8sContainerUpdateOne) defaults() {
	if _, ok := kcuo.mutation.UpdatedAt(); !ok {
		v := k8scontainer.UpdateDefaultUpdatedAt()
		kcuo.mutation.SetUpdatedAt(v)
	}
	if _, ok := kcuo.mutation.DeletedAt(); !ok {
		v := k8scontainer.UpdateDefaultDeletedAt()
		kcuo.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kcuo *K8sContainerUpdateOne) check() error {
	if v, ok := kcuo.mutation.Name(); ok {
		if err := k8scontainer.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := kcuo.mutation.Image(); ok {
		if err := k8scontainer.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf("ent: validator failed for field \"image\": %w", err)}
		}
	}
	if v, ok := kcuo.mutation.ContainerType(); ok {
		if err := k8scontainer.ContainerTypeValidator(v); err != nil {
			return &ValidationError{Name: "containerType", err: fmt.Errorf("ent: validator failed for field \"containerType\": %w", err)}
		}
	}
	if v, ok := kcuo.mutation.ContainerId(); ok {
		if err := k8scontainer.ContainerIdValidator(v); err != nil {
			return &ValidationError{Name: "containerId", err: fmt.Errorf("ent: validator failed for field \"containerId\": %w", err)}
		}
	}
	return nil
}

func (kcuo *K8sContainerUpdateOne) sqlSave(ctx context.Context) (_node *K8sContainer, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   k8scontainer.Table,
			Columns: k8scontainer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: k8scontainer.FieldID,
			},
		},
	}
	id, ok := kcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing K8sContainer.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := kcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, k8scontainer.FieldID)
		for _, f := range fields {
			if !k8scontainer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != k8scontainer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := kcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kcuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8scontainer.FieldUpdatedAt,
		})
	}
	if value, ok := kcuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: k8scontainer.FieldDeletedAt,
		})
	}
	if value, ok := kcuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: k8scontainer.FieldName,
		})
	}
	if value, ok := kcuo.mutation.Image(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: k8scontainer.FieldImage,
		})
	}
	if value, ok := kcuo.mutation.ContainerType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: k8scontainer.FieldContainerType,
		})
	}
	if value, ok := kcuo.mutation.ContainerId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: k8scontainer.FieldContainerId,
		})
	}
	if value, ok := kcuo.mutation.K8sClusterId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8scontainer.FieldK8sClusterId,
		})
	}
	if value, ok := kcuo.mutation.AddedK8sClusterId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8scontainer.FieldK8sClusterId,
		})
	}
	if value, ok := kcuo.mutation.K8sNamespaceId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8scontainer.FieldK8sNamespaceId,
		})
	}
	if value, ok := kcuo.mutation.AddedK8sNamespaceId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8scontainer.FieldK8sNamespaceId,
		})
	}
	if value, ok := kcuo.mutation.K8sPodId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8scontainer.FieldK8sPodId,
		})
	}
	if value, ok := kcuo.mutation.AddedK8sPodId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: k8scontainer.FieldK8sPodId,
		})
	}
	_node = &K8sContainer{config: kcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, kcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{k8scontainer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}