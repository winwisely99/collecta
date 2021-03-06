// Code generated by entc, DO NOT EDIT.

package account

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/minskylab/collecta/ent/predicate"
	"github.com/minskylab/collecta/uuid"
)

// ID filters vertices based on their identifier.
func ID(id uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Sub applies equality check predicate on the "sub" field. It's identical to SubEQ.
func Sub(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSub), v))
	})
}

// RemoteID applies equality check predicate on the "remoteID" field. It's identical to RemoteIDEQ.
func RemoteID(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemoteID), v))
	})
}

// Secret applies equality check predicate on the "secret" field. It's identical to SecretEQ.
func Secret(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSecret), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Account {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Account(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Account {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Account(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// SubEQ applies the EQ predicate on the "sub" field.
func SubEQ(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSub), v))
	})
}

// SubNEQ applies the NEQ predicate on the "sub" field.
func SubNEQ(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSub), v))
	})
}

// SubIn applies the In predicate on the "sub" field.
func SubIn(vs ...string) predicate.Account {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Account(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSub), v...))
	})
}

// SubNotIn applies the NotIn predicate on the "sub" field.
func SubNotIn(vs ...string) predicate.Account {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Account(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSub), v...))
	})
}

// SubGT applies the GT predicate on the "sub" field.
func SubGT(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSub), v))
	})
}

// SubGTE applies the GTE predicate on the "sub" field.
func SubGTE(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSub), v))
	})
}

// SubLT applies the LT predicate on the "sub" field.
func SubLT(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSub), v))
	})
}

// SubLTE applies the LTE predicate on the "sub" field.
func SubLTE(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSub), v))
	})
}

// SubContains applies the Contains predicate on the "sub" field.
func SubContains(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSub), v))
	})
}

// SubHasPrefix applies the HasPrefix predicate on the "sub" field.
func SubHasPrefix(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSub), v))
	})
}

// SubHasSuffix applies the HasSuffix predicate on the "sub" field.
func SubHasSuffix(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSub), v))
	})
}

// SubEqualFold applies the EqualFold predicate on the "sub" field.
func SubEqualFold(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSub), v))
	})
}

// SubContainsFold applies the ContainsFold predicate on the "sub" field.
func SubContainsFold(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSub), v))
	})
}

// RemoteIDEQ applies the EQ predicate on the "remoteID" field.
func RemoteIDEQ(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemoteID), v))
	})
}

// RemoteIDNEQ applies the NEQ predicate on the "remoteID" field.
func RemoteIDNEQ(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRemoteID), v))
	})
}

// RemoteIDIn applies the In predicate on the "remoteID" field.
func RemoteIDIn(vs ...string) predicate.Account {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Account(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRemoteID), v...))
	})
}

// RemoteIDNotIn applies the NotIn predicate on the "remoteID" field.
func RemoteIDNotIn(vs ...string) predicate.Account {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Account(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRemoteID), v...))
	})
}

// RemoteIDGT applies the GT predicate on the "remoteID" field.
func RemoteIDGT(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRemoteID), v))
	})
}

// RemoteIDGTE applies the GTE predicate on the "remoteID" field.
func RemoteIDGTE(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRemoteID), v))
	})
}

// RemoteIDLT applies the LT predicate on the "remoteID" field.
func RemoteIDLT(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRemoteID), v))
	})
}

// RemoteIDLTE applies the LTE predicate on the "remoteID" field.
func RemoteIDLTE(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRemoteID), v))
	})
}

// RemoteIDContains applies the Contains predicate on the "remoteID" field.
func RemoteIDContains(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRemoteID), v))
	})
}

// RemoteIDHasPrefix applies the HasPrefix predicate on the "remoteID" field.
func RemoteIDHasPrefix(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRemoteID), v))
	})
}

// RemoteIDHasSuffix applies the HasSuffix predicate on the "remoteID" field.
func RemoteIDHasSuffix(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRemoteID), v))
	})
}

// RemoteIDEqualFold applies the EqualFold predicate on the "remoteID" field.
func RemoteIDEqualFold(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRemoteID), v))
	})
}

// RemoteIDContainsFold applies the ContainsFold predicate on the "remoteID" field.
func RemoteIDContainsFold(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRemoteID), v))
	})
}

// SecretEQ applies the EQ predicate on the "secret" field.
func SecretEQ(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSecret), v))
	})
}

// SecretNEQ applies the NEQ predicate on the "secret" field.
func SecretNEQ(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSecret), v))
	})
}

// SecretIn applies the In predicate on the "secret" field.
func SecretIn(vs ...string) predicate.Account {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Account(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSecret), v...))
	})
}

// SecretNotIn applies the NotIn predicate on the "secret" field.
func SecretNotIn(vs ...string) predicate.Account {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Account(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSecret), v...))
	})
}

// SecretGT applies the GT predicate on the "secret" field.
func SecretGT(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSecret), v))
	})
}

// SecretGTE applies the GTE predicate on the "secret" field.
func SecretGTE(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSecret), v))
	})
}

// SecretLT applies the LT predicate on the "secret" field.
func SecretLT(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSecret), v))
	})
}

// SecretLTE applies the LTE predicate on the "secret" field.
func SecretLTE(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSecret), v))
	})
}

// SecretContains applies the Contains predicate on the "secret" field.
func SecretContains(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSecret), v))
	})
}

// SecretHasPrefix applies the HasPrefix predicate on the "secret" field.
func SecretHasPrefix(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSecret), v))
	})
}

// SecretHasSuffix applies the HasSuffix predicate on the "secret" field.
func SecretHasSuffix(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSecret), v))
	})
}

// SecretIsNil applies the IsNil predicate on the "secret" field.
func SecretIsNil() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSecret)))
	})
}

// SecretNotNil applies the NotNil predicate on the "secret" field.
func SecretNotNil() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSecret)))
	})
}

// SecretEqualFold applies the EqualFold predicate on the "secret" field.
func SecretEqualFold(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSecret), v))
	})
}

// SecretContainsFold applies the ContainsFold predicate on the "secret" field.
func SecretContainsFold(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSecret), v))
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Person) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
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
func Not(p predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		p(s.Not())
	})
}
