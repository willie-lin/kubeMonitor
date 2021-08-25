// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/metricname"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/metrictype"
)

// MetricTypeCreate is the builder for creating a MetricType entity.
type MetricTypeCreate struct {
	config
	mutation *MetricTypeMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mtc *MetricTypeCreate) SetCreatedAt(t time.Time) *MetricTypeCreate {
	mtc.mutation.SetCreatedAt(t)
	return mtc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mtc *MetricTypeCreate) SetNillableCreatedAt(t *time.Time) *MetricTypeCreate {
	if t != nil {
		mtc.SetCreatedAt(*t)
	}
	return mtc
}

// SetUpdatedAt sets the "updated_at" field.
func (mtc *MetricTypeCreate) SetUpdatedAt(t time.Time) *MetricTypeCreate {
	mtc.mutation.SetUpdatedAt(t)
	return mtc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mtc *MetricTypeCreate) SetNillableUpdatedAt(t *time.Time) *MetricTypeCreate {
	if t != nil {
		mtc.SetUpdatedAt(*t)
	}
	return mtc
}

// SetDeletedAt sets the "deleted_at" field.
func (mtc *MetricTypeCreate) SetDeletedAt(t time.Time) *MetricTypeCreate {
	mtc.mutation.SetDeletedAt(t)
	return mtc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mtc *MetricTypeCreate) SetNillableDeletedAt(t *time.Time) *MetricTypeCreate {
	if t != nil {
		mtc.SetDeletedAt(*t)
	}
	return mtc
}

// SetName sets the "name" field.
func (mtc *MetricTypeCreate) SetName(s string) *MetricTypeCreate {
	mtc.mutation.SetName(s)
	return mtc
}

// SetID sets the "id" field.
func (mtc *MetricTypeCreate) SetID(u uint) *MetricTypeCreate {
	mtc.mutation.SetID(u)
	return mtc
}

// AddMetricNameIDs adds the "metricNames" edge to the MetricName entity by IDs.
func (mtc *MetricTypeCreate) AddMetricNameIDs(ids ...uint) *MetricTypeCreate {
	mtc.mutation.AddMetricNameIDs(ids...)
	return mtc
}

// AddMetricNames adds the "metricNames" edges to the MetricName entity.
func (mtc *MetricTypeCreate) AddMetricNames(m ...*MetricName) *MetricTypeCreate {
	ids := make([]uint, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mtc.AddMetricNameIDs(ids...)
}

// Mutation returns the MetricTypeMutation object of the builder.
func (mtc *MetricTypeCreate) Mutation() *MetricTypeMutation {
	return mtc.mutation
}

// Save creates the MetricType in the database.
func (mtc *MetricTypeCreate) Save(ctx context.Context) (*MetricType, error) {
	var (
		err  error
		node *MetricType
	)
	mtc.defaults()
	if len(mtc.hooks) == 0 {
		if err = mtc.check(); err != nil {
			return nil, err
		}
		node, err = mtc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MetricTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mtc.check(); err != nil {
				return nil, err
			}
			mtc.mutation = mutation
			if node, err = mtc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mtc.hooks) - 1; i >= 0; i-- {
			if mtc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mtc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mtc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mtc *MetricTypeCreate) SaveX(ctx context.Context) *MetricType {
	v, err := mtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mtc *MetricTypeCreate) Exec(ctx context.Context) error {
	_, err := mtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mtc *MetricTypeCreate) ExecX(ctx context.Context) {
	if err := mtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mtc *MetricTypeCreate) defaults() {
	if _, ok := mtc.mutation.CreatedAt(); !ok {
		v := metrictype.DefaultCreatedAt()
		mtc.mutation.SetCreatedAt(v)
	}
	if _, ok := mtc.mutation.UpdatedAt(); !ok {
		v := metrictype.DefaultUpdatedAt()
		mtc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mtc.mutation.DeletedAt(); !ok {
		v := metrictype.DefaultDeletedAt()
		mtc.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mtc *MetricTypeCreate) check() error {
	if _, ok := mtc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := mtc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := mtc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "deleted_at"`)}
	}
	if _, ok := mtc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	return nil
}

func (mtc *MetricTypeCreate) sqlSave(ctx context.Context) (*MetricType, error) {
	_node, _spec := mtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mtc.driver, _spec); err != nil {
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

func (mtc *MetricTypeCreate) createSpec() (*MetricType, *sqlgraph.CreateSpec) {
	var (
		_node = &MetricType{config: mtc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: metrictype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: metrictype.FieldID,
			},
		}
	)
	if id, ok := mtc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mtc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: metrictype.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := mtc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: metrictype.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := mtc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: metrictype.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := mtc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: metrictype.FieldName,
		})
		_node.Name = value
	}
	if nodes := mtc.mutation.MetricNamesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metrictype.MetricNamesTable,
			Columns: []string{metrictype.MetricNamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: metricname.FieldID,
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

// MetricTypeCreateBulk is the builder for creating many MetricType entities in bulk.
type MetricTypeCreateBulk struct {
	config
	builders []*MetricTypeCreate
}

// Save creates the MetricType entities in the database.
func (mtcb *MetricTypeCreateBulk) Save(ctx context.Context) ([]*MetricType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mtcb.builders))
	nodes := make([]*MetricType, len(mtcb.builders))
	mutators := make([]Mutator, len(mtcb.builders))
	for i := range mtcb.builders {
		func(i int, root context.Context) {
			builder := mtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MetricTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, mtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mtcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mtcb *MetricTypeCreateBulk) SaveX(ctx context.Context) []*MetricType {
	v, err := mtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mtcb *MetricTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := mtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mtcb *MetricTypeCreateBulk) ExecX(ctx context.Context) {
	if err := mtcb.Exec(ctx); err != nil {
		panic(err)
	}
}
