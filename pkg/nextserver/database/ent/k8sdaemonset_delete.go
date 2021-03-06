// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/k8sdaemonset"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/predicate"
)

// K8sDaemonSetDelete is the builder for deleting a K8sDaemonSet entity.
type K8sDaemonSetDelete struct {
	config
	hooks    []Hook
	mutation *K8sDaemonSetMutation
}

// Where appends a list predicates to the K8sDaemonSetDelete builder.
func (kdsd *K8sDaemonSetDelete) Where(ps ...predicate.K8sDaemonSet) *K8sDaemonSetDelete {
	kdsd.mutation.Where(ps...)
	return kdsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (kdsd *K8sDaemonSetDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(kdsd.hooks) == 0 {
		affected, err = kdsd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*K8sDaemonSetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			kdsd.mutation = mutation
			affected, err = kdsd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(kdsd.hooks) - 1; i >= 0; i-- {
			if kdsd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kdsd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kdsd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (kdsd *K8sDaemonSetDelete) ExecX(ctx context.Context) int {
	n, err := kdsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (kdsd *K8sDaemonSetDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: k8sdaemonset.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: k8sdaemonset.FieldID,
			},
		},
	}
	if ps := kdsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, kdsd.driver, _spec)
}

// K8sDaemonSetDeleteOne is the builder for deleting a single K8sDaemonSet entity.
type K8sDaemonSetDeleteOne struct {
	kdsd *K8sDaemonSetDelete
}

// Exec executes the deletion query.
func (kdsdo *K8sDaemonSetDeleteOne) Exec(ctx context.Context) error {
	n, err := kdsdo.kdsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{k8sdaemonset.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (kdsdo *K8sDaemonSetDeleteOne) ExecX(ctx context.Context) {
	kdsdo.kdsd.ExecX(ctx)
}
