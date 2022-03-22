// Code generated by entc, DO NOT EDIT.

package template

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/notification/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// LangID applies equality check predicate on the "lang_id" field. It's identical to LangIDEQ.
func LangID(v uuid.UUID) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLangID), v))
	})
}

// UsedFor applies equality check predicate on the "used_for" field. It's identical to UsedForEQ.
func UsedFor(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsedFor), v))
	})
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v uint32) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v uint32) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// DeleteAt applies equality check predicate on the "delete_at" field. It's identical to DeleteAtEQ.
func DeleteAt(v uint32) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContent), v))
	})
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Template {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Template(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldContent), v...))
	})
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Template {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Template(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldContent), v...))
	})
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContent), v))
	})
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContent), v))
	})
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContent), v))
	})
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContent), v))
	})
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContent), v))
	})
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContent), v))
	})
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContent), v))
	})
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContent), v))
	})
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContent), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.Template {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Template(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.Template {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Template(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// LangIDEQ applies the EQ predicate on the "lang_id" field.
func LangIDEQ(v uuid.UUID) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLangID), v))
	})
}

// LangIDNEQ applies the NEQ predicate on the "lang_id" field.
func LangIDNEQ(v uuid.UUID) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLangID), v))
	})
}

// LangIDIn applies the In predicate on the "lang_id" field.
func LangIDIn(vs ...uuid.UUID) predicate.Template {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Template(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLangID), v...))
	})
}

// LangIDNotIn applies the NotIn predicate on the "lang_id" field.
func LangIDNotIn(vs ...uuid.UUID) predicate.Template {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Template(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLangID), v...))
	})
}

// LangIDGT applies the GT predicate on the "lang_id" field.
func LangIDGT(v uuid.UUID) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLangID), v))
	})
}

// LangIDGTE applies the GTE predicate on the "lang_id" field.
func LangIDGTE(v uuid.UUID) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLangID), v))
	})
}

// LangIDLT applies the LT predicate on the "lang_id" field.
func LangIDLT(v uuid.UUID) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLangID), v))
	})
}

// LangIDLTE applies the LTE predicate on the "lang_id" field.
func LangIDLTE(v uuid.UUID) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLangID), v))
	})
}

// UsedForEQ applies the EQ predicate on the "used_for" field.
func UsedForEQ(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsedFor), v))
	})
}

// UsedForNEQ applies the NEQ predicate on the "used_for" field.
func UsedForNEQ(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUsedFor), v))
	})
}

// UsedForIn applies the In predicate on the "used_for" field.
func UsedForIn(vs ...string) predicate.Template {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Template(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUsedFor), v...))
	})
}

// UsedForNotIn applies the NotIn predicate on the "used_for" field.
func UsedForNotIn(vs ...string) predicate.Template {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Template(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUsedFor), v...))
	})
}

// UsedForGT applies the GT predicate on the "used_for" field.
func UsedForGT(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUsedFor), v))
	})
}

// UsedForGTE applies the GTE predicate on the "used_for" field.
func UsedForGTE(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUsedFor), v))
	})
}

// UsedForLT applies the LT predicate on the "used_for" field.
func UsedForLT(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUsedFor), v))
	})
}

// UsedForLTE applies the LTE predicate on the "used_for" field.
func UsedForLTE(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUsedFor), v))
	})
}

// UsedForContains applies the Contains predicate on the "used_for" field.
func UsedForContains(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUsedFor), v))
	})
}

// UsedForHasPrefix applies the HasPrefix predicate on the "used_for" field.
func UsedForHasPrefix(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUsedFor), v))
	})
}

// UsedForHasSuffix applies the HasSuffix predicate on the "used_for" field.
func UsedForHasSuffix(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUsedFor), v))
	})
}

// UsedForEqualFold applies the EqualFold predicate on the "used_for" field.
func UsedForEqualFold(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUsedFor), v))
	})
}

// UsedForContainsFold applies the ContainsFold predicate on the "used_for" field.
func UsedForContainsFold(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUsedFor), v))
	})
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v uint32) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v uint32) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...uint32) predicate.Template {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Template(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateAt), v...))
	})
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...uint32) predicate.Template {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Template(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateAt), v...))
	})
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v uint32) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateAt), v))
	})
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v uint32) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateAt), v))
	})
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v uint32) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateAt), v))
	})
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v uint32) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateAt), v))
	})
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v uint32) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v uint32) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...uint32) predicate.Template {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Template(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...uint32) predicate.Template {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Template(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v uint32) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v uint32) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v uint32) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v uint32) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateAt), v))
	})
}

// DeleteAtEQ applies the EQ predicate on the "delete_at" field.
func DeleteAtEQ(v uint32) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtNEQ applies the NEQ predicate on the "delete_at" field.
func DeleteAtNEQ(v uint32) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtIn applies the In predicate on the "delete_at" field.
func DeleteAtIn(vs ...uint32) predicate.Template {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Template(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtNotIn applies the NotIn predicate on the "delete_at" field.
func DeleteAtNotIn(vs ...uint32) predicate.Template {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Template(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtGT applies the GT predicate on the "delete_at" field.
func DeleteAtGT(v uint32) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtGTE applies the GTE predicate on the "delete_at" field.
func DeleteAtGTE(v uint32) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLT applies the LT predicate on the "delete_at" field.
func DeleteAtLT(v uint32) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLTE applies the LTE predicate on the "delete_at" field.
func DeleteAtLTE(v uint32) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleteAt), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Template {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Template(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Template {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Template(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Template) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Template) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
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
func Not(p predicate.Template) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		p(s.Not())
	})
}
