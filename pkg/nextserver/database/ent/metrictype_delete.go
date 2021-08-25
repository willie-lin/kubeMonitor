// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/metrictype"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/predicate"
)

// MetricTypeDelete is the builder for deleting a MetricType entity.
type MetricTypeDelete struct {
	config
	hooks    []Hook
	mutation *MetricTypeMutation
}

// Where appends a list predicates to the MetricTypeDelete builder.
func (mtd *MetricTypeDelete) Where(ps ...predicate.MetricType) *MetricTypeDelete {
	mtd.mutation.Where(ps...)
	return mtd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mtd *MetricTypeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mtd.hooks) == 0 {
		affected, err = mtd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MetricTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mtd.mutation = mutation
			affected, err = mtd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mtd.hooks) - 1; i >= 0; i-- {
			if mtd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mtd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mtd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (mtd *MetricTypeDelete) ExecX(ctx context.Context) int {
	n, err := mtd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mtd *MetricTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: metrictype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: metrictype.FieldID,
			},
		},
	}
	if ps := mtd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, mtd.driver, _spec)
}

// MetricTypeDeleteOne is the builder for deleting a single MetricType entity.
type MetricTypeDeleteOne struct {
	mtd *MetricTypeDelete
}

// Exec executes the deletion query.
func (mtdo *MetricTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := mtdo.mtd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{metrictype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mtdo *MetricTypeDeleteOne) ExecX(ctx context.Context) {
	mtdo.mtd.ExecX(ctx)
}
