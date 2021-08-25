// Code generated by entc, DO NOT EDIT.

package k8smetric

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// Ts applies equality check predicate on the "ts" field. It's identical to TsEQ.
func Ts(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTs), v))
	})
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v float64) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// EndpointId applies equality check predicate on the "endpointId" field. It's identical to EndpointIdEQ.
func EndpointId(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndpointId), v))
	})
}

// TypeId applies equality check predicate on the "typeId" field. It's identical to TypeIdEQ.
func TypeId(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTypeId), v))
	})
}

// NameId applies equality check predicate on the "nameId" field. It's identical to NameIdEQ.
func NameId(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNameId), v))
	})
}

// LabelId applies equality check predicate on the "labelId" field. It's identical to LabelIdEQ.
func LabelId(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLabelId), v))
	})
}

// K8sClusterId applies equality check predicate on the "k8sClusterId" field. It's identical to K8sClusterIdEQ.
func K8sClusterId(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldK8sClusterId), v))
	})
}

// K8sNodeId applies equality check predicate on the "k8sNodeId" field. It's identical to K8sNodeIdEQ.
func K8sNodeId(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldK8sNodeId), v))
	})
}

// K8sNamespaceId applies equality check predicate on the "k8sNamespaceId" field. It's identical to K8sNamespaceIdEQ.
func K8sNamespaceId(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldK8sNamespaceId), v))
	})
}

// K8sContainerId applies equality check predicate on the "k8sContainerId" field. It's identical to K8sContainerIdEQ.
func K8sContainerId(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldK8sContainerId), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// TsEQ applies the EQ predicate on the "ts" field.
func TsEQ(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTs), v))
	})
}

// TsNEQ applies the NEQ predicate on the "ts" field.
func TsNEQ(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTs), v))
	})
}

// TsIn applies the In predicate on the "ts" field.
func TsIn(vs ...time.Time) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTs), v...))
	})
}

// TsNotIn applies the NotIn predicate on the "ts" field.
func TsNotIn(vs ...time.Time) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTs), v...))
	})
}

// TsGT applies the GT predicate on the "ts" field.
func TsGT(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTs), v))
	})
}

// TsGTE applies the GTE predicate on the "ts" field.
func TsGTE(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTs), v))
	})
}

// TsLT applies the LT predicate on the "ts" field.
func TsLT(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTs), v))
	})
}

// TsLTE applies the LTE predicate on the "ts" field.
func TsLTE(v time.Time) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTs), v))
	})
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v float64) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v float64) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldValue), v))
	})
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...float64) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldValue), v...))
	})
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...float64) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldValue), v...))
	})
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v float64) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldValue), v))
	})
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v float64) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldValue), v))
	})
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v float64) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldValue), v))
	})
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v float64) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldValue), v))
	})
}

// EndpointIdEQ applies the EQ predicate on the "endpointId" field.
func EndpointIdEQ(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndpointId), v))
	})
}

// EndpointIdNEQ applies the NEQ predicate on the "endpointId" field.
func EndpointIdNEQ(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndpointId), v))
	})
}

// EndpointIdIn applies the In predicate on the "endpointId" field.
func EndpointIdIn(vs ...uint) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEndpointId), v...))
	})
}

// EndpointIdNotIn applies the NotIn predicate on the "endpointId" field.
func EndpointIdNotIn(vs ...uint) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEndpointId), v...))
	})
}

// EndpointIdGT applies the GT predicate on the "endpointId" field.
func EndpointIdGT(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndpointId), v))
	})
}

// EndpointIdGTE applies the GTE predicate on the "endpointId" field.
func EndpointIdGTE(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndpointId), v))
	})
}

// EndpointIdLT applies the LT predicate on the "endpointId" field.
func EndpointIdLT(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndpointId), v))
	})
}

// EndpointIdLTE applies the LTE predicate on the "endpointId" field.
func EndpointIdLTE(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndpointId), v))
	})
}

// TypeIdEQ applies the EQ predicate on the "typeId" field.
func TypeIdEQ(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTypeId), v))
	})
}

// TypeIdNEQ applies the NEQ predicate on the "typeId" field.
func TypeIdNEQ(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTypeId), v))
	})
}

// TypeIdIn applies the In predicate on the "typeId" field.
func TypeIdIn(vs ...uint) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTypeId), v...))
	})
}

// TypeIdNotIn applies the NotIn predicate on the "typeId" field.
func TypeIdNotIn(vs ...uint) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTypeId), v...))
	})
}

// TypeIdGT applies the GT predicate on the "typeId" field.
func TypeIdGT(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTypeId), v))
	})
}

// TypeIdGTE applies the GTE predicate on the "typeId" field.
func TypeIdGTE(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTypeId), v))
	})
}

// TypeIdLT applies the LT predicate on the "typeId" field.
func TypeIdLT(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTypeId), v))
	})
}

// TypeIdLTE applies the LTE predicate on the "typeId" field.
func TypeIdLTE(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTypeId), v))
	})
}

// NameIdEQ applies the EQ predicate on the "nameId" field.
func NameIdEQ(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNameId), v))
	})
}

// NameIdNEQ applies the NEQ predicate on the "nameId" field.
func NameIdNEQ(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNameId), v))
	})
}

// NameIdIn applies the In predicate on the "nameId" field.
func NameIdIn(vs ...uint) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNameId), v...))
	})
}

// NameIdNotIn applies the NotIn predicate on the "nameId" field.
func NameIdNotIn(vs ...uint) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNameId), v...))
	})
}

// NameIdGT applies the GT predicate on the "nameId" field.
func NameIdGT(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNameId), v))
	})
}

// NameIdGTE applies the GTE predicate on the "nameId" field.
func NameIdGTE(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNameId), v))
	})
}

// NameIdLT applies the LT predicate on the "nameId" field.
func NameIdLT(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNameId), v))
	})
}

// NameIdLTE applies the LTE predicate on the "nameId" field.
func NameIdLTE(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNameId), v))
	})
}

// LabelIdEQ applies the EQ predicate on the "labelId" field.
func LabelIdEQ(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLabelId), v))
	})
}

// LabelIdNEQ applies the NEQ predicate on the "labelId" field.
func LabelIdNEQ(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLabelId), v))
	})
}

// LabelIdIn applies the In predicate on the "labelId" field.
func LabelIdIn(vs ...uint) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLabelId), v...))
	})
}

// LabelIdNotIn applies the NotIn predicate on the "labelId" field.
func LabelIdNotIn(vs ...uint) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLabelId), v...))
	})
}

// LabelIdGT applies the GT predicate on the "labelId" field.
func LabelIdGT(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLabelId), v))
	})
}

// LabelIdGTE applies the GTE predicate on the "labelId" field.
func LabelIdGTE(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLabelId), v))
	})
}

// LabelIdLT applies the LT predicate on the "labelId" field.
func LabelIdLT(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLabelId), v))
	})
}

// LabelIdLTE applies the LTE predicate on the "labelId" field.
func LabelIdLTE(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLabelId), v))
	})
}

// K8sClusterIdEQ applies the EQ predicate on the "k8sClusterId" field.
func K8sClusterIdEQ(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldK8sClusterId), v))
	})
}

// K8sClusterIdNEQ applies the NEQ predicate on the "k8sClusterId" field.
func K8sClusterIdNEQ(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldK8sClusterId), v))
	})
}

// K8sClusterIdIn applies the In predicate on the "k8sClusterId" field.
func K8sClusterIdIn(vs ...uint) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldK8sClusterId), v...))
	})
}

// K8sClusterIdNotIn applies the NotIn predicate on the "k8sClusterId" field.
func K8sClusterIdNotIn(vs ...uint) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldK8sClusterId), v...))
	})
}

// K8sClusterIdGT applies the GT predicate on the "k8sClusterId" field.
func K8sClusterIdGT(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldK8sClusterId), v))
	})
}

// K8sClusterIdGTE applies the GTE predicate on the "k8sClusterId" field.
func K8sClusterIdGTE(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldK8sClusterId), v))
	})
}

// K8sClusterIdLT applies the LT predicate on the "k8sClusterId" field.
func K8sClusterIdLT(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldK8sClusterId), v))
	})
}

// K8sClusterIdLTE applies the LTE predicate on the "k8sClusterId" field.
func K8sClusterIdLTE(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldK8sClusterId), v))
	})
}

// K8sNodeIdEQ applies the EQ predicate on the "k8sNodeId" field.
func K8sNodeIdEQ(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldK8sNodeId), v))
	})
}

// K8sNodeIdNEQ applies the NEQ predicate on the "k8sNodeId" field.
func K8sNodeIdNEQ(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldK8sNodeId), v))
	})
}

// K8sNodeIdIn applies the In predicate on the "k8sNodeId" field.
func K8sNodeIdIn(vs ...uint) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldK8sNodeId), v...))
	})
}

// K8sNodeIdNotIn applies the NotIn predicate on the "k8sNodeId" field.
func K8sNodeIdNotIn(vs ...uint) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldK8sNodeId), v...))
	})
}

// K8sNodeIdGT applies the GT predicate on the "k8sNodeId" field.
func K8sNodeIdGT(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldK8sNodeId), v))
	})
}

// K8sNodeIdGTE applies the GTE predicate on the "k8sNodeId" field.
func K8sNodeIdGTE(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldK8sNodeId), v))
	})
}

// K8sNodeIdLT applies the LT predicate on the "k8sNodeId" field.
func K8sNodeIdLT(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldK8sNodeId), v))
	})
}

// K8sNodeIdLTE applies the LTE predicate on the "k8sNodeId" field.
func K8sNodeIdLTE(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldK8sNodeId), v))
	})
}

// K8sNamespaceIdEQ applies the EQ predicate on the "k8sNamespaceId" field.
func K8sNamespaceIdEQ(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldK8sNamespaceId), v))
	})
}

// K8sNamespaceIdNEQ applies the NEQ predicate on the "k8sNamespaceId" field.
func K8sNamespaceIdNEQ(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldK8sNamespaceId), v))
	})
}

// K8sNamespaceIdIn applies the In predicate on the "k8sNamespaceId" field.
func K8sNamespaceIdIn(vs ...uint) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldK8sNamespaceId), v...))
	})
}

// K8sNamespaceIdNotIn applies the NotIn predicate on the "k8sNamespaceId" field.
func K8sNamespaceIdNotIn(vs ...uint) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldK8sNamespaceId), v...))
	})
}

// K8sNamespaceIdGT applies the GT predicate on the "k8sNamespaceId" field.
func K8sNamespaceIdGT(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldK8sNamespaceId), v))
	})
}

// K8sNamespaceIdGTE applies the GTE predicate on the "k8sNamespaceId" field.
func K8sNamespaceIdGTE(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldK8sNamespaceId), v))
	})
}

// K8sNamespaceIdLT applies the LT predicate on the "k8sNamespaceId" field.
func K8sNamespaceIdLT(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldK8sNamespaceId), v))
	})
}

// K8sNamespaceIdLTE applies the LTE predicate on the "k8sNamespaceId" field.
func K8sNamespaceIdLTE(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldK8sNamespaceId), v))
	})
}

// K8sContainerIdEQ applies the EQ predicate on the "k8sContainerId" field.
func K8sContainerIdEQ(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldK8sContainerId), v))
	})
}

// K8sContainerIdNEQ applies the NEQ predicate on the "k8sContainerId" field.
func K8sContainerIdNEQ(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldK8sContainerId), v))
	})
}

// K8sContainerIdIn applies the In predicate on the "k8sContainerId" field.
func K8sContainerIdIn(vs ...uint) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldK8sContainerId), v...))
	})
}

// K8sContainerIdNotIn applies the NotIn predicate on the "k8sContainerId" field.
func K8sContainerIdNotIn(vs ...uint) predicate.K8sMetric {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sMetric(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldK8sContainerId), v...))
	})
}

// K8sContainerIdGT applies the GT predicate on the "k8sContainerId" field.
func K8sContainerIdGT(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldK8sContainerId), v))
	})
}

// K8sContainerIdGTE applies the GTE predicate on the "k8sContainerId" field.
func K8sContainerIdGTE(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldK8sContainerId), v))
	})
}

// K8sContainerIdLT applies the LT predicate on the "k8sContainerId" field.
func K8sContainerIdLT(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldK8sContainerId), v))
	})
}

// K8sContainerIdLTE applies the LTE predicate on the "k8sContainerId" field.
func K8sContainerIdLTE(v uint) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldK8sContainerId), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.K8sMetric) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.K8sMetric) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.K8sMetric) predicate.K8sMetric {
	return predicate.K8sMetric(func(s *sql.Selector) {
		p(s.Not())
	})
}
