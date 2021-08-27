// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/incidentbasicrule"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/predicate"
)

// IncidentBasicRuleUpdate is the builder for updating IncidentBasicRule entities.
type IncidentBasicRuleUpdate struct {
	config
	hooks    []Hook
	mutation *IncidentBasicRuleMutation
}

// Where appends a list predicates to the IncidentBasicRuleUpdate builder.
func (ibru *IncidentBasicRuleUpdate) Where(ps ...predicate.IncidentBasicRule) *IncidentBasicRuleUpdate {
	ibru.mutation.Where(ps...)
	return ibru
}

// SetUpdatedAt sets the "updated_at" field.
func (ibru *IncidentBasicRuleUpdate) SetUpdatedAt(t time.Time) *IncidentBasicRuleUpdate {
	ibru.mutation.SetUpdatedAt(t)
	return ibru
}

// SetDeletedAt sets the "deleted_at" field.
func (ibru *IncidentBasicRuleUpdate) SetDeletedAt(t time.Time) *IncidentBasicRuleUpdate {
	ibru.mutation.SetDeletedAt(t)
	return ibru
}

// SetName sets the "name" field.
func (ibru *IncidentBasicRuleUpdate) SetName(s string) *IncidentBasicRuleUpdate {
	ibru.mutation.SetName(s)
	return ibru
}

// SetDescription sets the "Description" field.
func (ibru *IncidentBasicRuleUpdate) SetDescription(s string) *IncidentBasicRuleUpdate {
	ibru.mutation.SetDescription(s)
	return ibru
}

// SetQuery sets the "query" field.
func (ibru *IncidentBasicRuleUpdate) SetQuery(s string) *IncidentBasicRuleUpdate {
	ibru.mutation.SetQuery(s)
	return ibru
}

// Mutation returns the IncidentBasicRuleMutation object of the builder.
func (ibru *IncidentBasicRuleUpdate) Mutation() *IncidentBasicRuleMutation {
	return ibru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ibru *IncidentBasicRuleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ibru.defaults()
	if len(ibru.hooks) == 0 {
		affected, err = ibru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IncidentBasicRuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ibru.mutation = mutation
			affected, err = ibru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ibru.hooks) - 1; i >= 0; i-- {
			if ibru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ibru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ibru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ibru *IncidentBasicRuleUpdate) SaveX(ctx context.Context) int {
	affected, err := ibru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ibru *IncidentBasicRuleUpdate) Exec(ctx context.Context) error {
	_, err := ibru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ibru *IncidentBasicRuleUpdate) ExecX(ctx context.Context) {
	if err := ibru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ibru *IncidentBasicRuleUpdate) defaults() {
	if _, ok := ibru.mutation.UpdatedAt(); !ok {
		v := incidentbasicrule.UpdateDefaultUpdatedAt()
		ibru.mutation.SetUpdatedAt(v)
	}
	if _, ok := ibru.mutation.DeletedAt(); !ok {
		v := incidentbasicrule.UpdateDefaultDeletedAt()
		ibru.mutation.SetDeletedAt(v)
	}
}

func (ibru *IncidentBasicRuleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   incidentbasicrule.Table,
			Columns: incidentbasicrule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: incidentbasicrule.FieldID,
			},
		},
	}
	if ps := ibru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ibru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: incidentbasicrule.FieldUpdatedAt,
		})
	}
	if value, ok := ibru.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: incidentbasicrule.FieldDeletedAt,
		})
	}
	if value, ok := ibru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: incidentbasicrule.FieldName,
		})
	}
	if value, ok := ibru.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: incidentbasicrule.FieldDescription,
		})
	}
	if value, ok := ibru.mutation.Query(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: incidentbasicrule.FieldQuery,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ibru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{incidentbasicrule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// IncidentBasicRuleUpdateOne is the builder for updating a single IncidentBasicRule entity.
type IncidentBasicRuleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IncidentBasicRuleMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ibruo *IncidentBasicRuleUpdateOne) SetUpdatedAt(t time.Time) *IncidentBasicRuleUpdateOne {
	ibruo.mutation.SetUpdatedAt(t)
	return ibruo
}

// SetDeletedAt sets the "deleted_at" field.
func (ibruo *IncidentBasicRuleUpdateOne) SetDeletedAt(t time.Time) *IncidentBasicRuleUpdateOne {
	ibruo.mutation.SetDeletedAt(t)
	return ibruo
}

// SetName sets the "name" field.
func (ibruo *IncidentBasicRuleUpdateOne) SetName(s string) *IncidentBasicRuleUpdateOne {
	ibruo.mutation.SetName(s)
	return ibruo
}

// SetDescription sets the "Description" field.
func (ibruo *IncidentBasicRuleUpdateOne) SetDescription(s string) *IncidentBasicRuleUpdateOne {
	ibruo.mutation.SetDescription(s)
	return ibruo
}

// SetQuery sets the "query" field.
func (ibruo *IncidentBasicRuleUpdateOne) SetQuery(s string) *IncidentBasicRuleUpdateOne {
	ibruo.mutation.SetQuery(s)
	return ibruo
}

// Mutation returns the IncidentBasicRuleMutation object of the builder.
func (ibruo *IncidentBasicRuleUpdateOne) Mutation() *IncidentBasicRuleMutation {
	return ibruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ibruo *IncidentBasicRuleUpdateOne) Select(field string, fields ...string) *IncidentBasicRuleUpdateOne {
	ibruo.fields = append([]string{field}, fields...)
	return ibruo
}

// Save executes the query and returns the updated IncidentBasicRule entity.
func (ibruo *IncidentBasicRuleUpdateOne) Save(ctx context.Context) (*IncidentBasicRule, error) {
	var (
		err  error
		node *IncidentBasicRule
	)
	ibruo.defaults()
	if len(ibruo.hooks) == 0 {
		node, err = ibruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IncidentBasicRuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ibruo.mutation = mutation
			node, err = ibruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ibruo.hooks) - 1; i >= 0; i-- {
			if ibruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ibruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ibruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ibruo *IncidentBasicRuleUpdateOne) SaveX(ctx context.Context) *IncidentBasicRule {
	node, err := ibruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ibruo *IncidentBasicRuleUpdateOne) Exec(ctx context.Context) error {
	_, err := ibruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ibruo *IncidentBasicRuleUpdateOne) ExecX(ctx context.Context) {
	if err := ibruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ibruo *IncidentBasicRuleUpdateOne) defaults() {
	if _, ok := ibruo.mutation.UpdatedAt(); !ok {
		v := incidentbasicrule.UpdateDefaultUpdatedAt()
		ibruo.mutation.SetUpdatedAt(v)
	}
	if _, ok := ibruo.mutation.DeletedAt(); !ok {
		v := incidentbasicrule.UpdateDefaultDeletedAt()
		ibruo.mutation.SetDeletedAt(v)
	}
}

func (ibruo *IncidentBasicRuleUpdateOne) sqlSave(ctx context.Context) (_node *IncidentBasicRule, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   incidentbasicrule.Table,
			Columns: incidentbasicrule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: incidentbasicrule.FieldID,
			},
		},
	}
	id, ok := ibruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing IncidentBasicRule.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ibruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, incidentbasicrule.FieldID)
		for _, f := range fields {
			if !incidentbasicrule.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != incidentbasicrule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ibruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ibruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: incidentbasicrule.FieldUpdatedAt,
		})
	}
	if value, ok := ibruo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: incidentbasicrule.FieldDeletedAt,
		})
	}
	if value, ok := ibruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: incidentbasicrule.FieldName,
		})
	}
	if value, ok := ibruo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: incidentbasicrule.FieldDescription,
		})
	}
	if value, ok := ibruo.mutation.Query(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: incidentbasicrule.FieldQuery,
		})
	}
	_node = &IncidentBasicRule{config: ibruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ibruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{incidentbasicrule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}