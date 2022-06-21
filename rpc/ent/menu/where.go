// Code generated by entc, DO NOT EDIT.

package menu

import (
	"antd-pro-amis-server/rpc/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// MenuName applies equality check predicate on the "menu_name" field. It's identical to MenuNameEQ.
func MenuName(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMenuName), v))
	})
}

// ParentID applies equality check predicate on the "parent_id" field. It's identical to ParentIDEQ.
func ParentID(v int64) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldParentID), v))
	})
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPath), v))
	})
}

// Icon applies equality check predicate on the "icon" field. It's identical to IconEQ.
func Icon(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIcon), v))
	})
}

// Layout applies equality check predicate on the "layout" field. It's identical to LayoutEQ.
func Layout(v bool) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLayout), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v bool) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// HideInMenu applies equality check predicate on the "hide_in_menu" field. It's identical to HideInMenuEQ.
func HideInMenu(v bool) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHideInMenu), v))
	})
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v int64) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSort), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// MenuNameEQ applies the EQ predicate on the "menu_name" field.
func MenuNameEQ(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMenuName), v))
	})
}

// MenuNameNEQ applies the NEQ predicate on the "menu_name" field.
func MenuNameNEQ(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMenuName), v))
	})
}

// MenuNameIn applies the In predicate on the "menu_name" field.
func MenuNameIn(vs ...string) predicate.Menu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Menu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMenuName), v...))
	})
}

// MenuNameNotIn applies the NotIn predicate on the "menu_name" field.
func MenuNameNotIn(vs ...string) predicate.Menu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Menu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMenuName), v...))
	})
}

// MenuNameGT applies the GT predicate on the "menu_name" field.
func MenuNameGT(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMenuName), v))
	})
}

// MenuNameGTE applies the GTE predicate on the "menu_name" field.
func MenuNameGTE(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMenuName), v))
	})
}

// MenuNameLT applies the LT predicate on the "menu_name" field.
func MenuNameLT(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMenuName), v))
	})
}

// MenuNameLTE applies the LTE predicate on the "menu_name" field.
func MenuNameLTE(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMenuName), v))
	})
}

// MenuNameContains applies the Contains predicate on the "menu_name" field.
func MenuNameContains(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMenuName), v))
	})
}

// MenuNameHasPrefix applies the HasPrefix predicate on the "menu_name" field.
func MenuNameHasPrefix(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMenuName), v))
	})
}

// MenuNameHasSuffix applies the HasSuffix predicate on the "menu_name" field.
func MenuNameHasSuffix(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMenuName), v))
	})
}

// MenuNameEqualFold applies the EqualFold predicate on the "menu_name" field.
func MenuNameEqualFold(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMenuName), v))
	})
}

// MenuNameContainsFold applies the ContainsFold predicate on the "menu_name" field.
func MenuNameContainsFold(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMenuName), v))
	})
}

// ParentIDEQ applies the EQ predicate on the "parent_id" field.
func ParentIDEQ(v int64) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldParentID), v))
	})
}

// ParentIDNEQ applies the NEQ predicate on the "parent_id" field.
func ParentIDNEQ(v int64) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldParentID), v))
	})
}

// ParentIDIn applies the In predicate on the "parent_id" field.
func ParentIDIn(vs ...int64) predicate.Menu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Menu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldParentID), v...))
	})
}

// ParentIDNotIn applies the NotIn predicate on the "parent_id" field.
func ParentIDNotIn(vs ...int64) predicate.Menu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Menu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldParentID), v...))
	})
}

// ParentIDGT applies the GT predicate on the "parent_id" field.
func ParentIDGT(v int64) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldParentID), v))
	})
}

// ParentIDGTE applies the GTE predicate on the "parent_id" field.
func ParentIDGTE(v int64) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldParentID), v))
	})
}

// ParentIDLT applies the LT predicate on the "parent_id" field.
func ParentIDLT(v int64) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldParentID), v))
	})
}

// ParentIDLTE applies the LTE predicate on the "parent_id" field.
func ParentIDLTE(v int64) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldParentID), v))
	})
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPath), v))
	})
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPath), v))
	})
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.Menu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Menu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPath), v...))
	})
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.Menu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Menu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPath), v...))
	})
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPath), v))
	})
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPath), v))
	})
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPath), v))
	})
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPath), v))
	})
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPath), v))
	})
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPath), v))
	})
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPath), v))
	})
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPath), v))
	})
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPath), v))
	})
}

// IconEQ applies the EQ predicate on the "icon" field.
func IconEQ(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIcon), v))
	})
}

// IconNEQ applies the NEQ predicate on the "icon" field.
func IconNEQ(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIcon), v))
	})
}

// IconIn applies the In predicate on the "icon" field.
func IconIn(vs ...string) predicate.Menu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Menu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIcon), v...))
	})
}

// IconNotIn applies the NotIn predicate on the "icon" field.
func IconNotIn(vs ...string) predicate.Menu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Menu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIcon), v...))
	})
}

// IconGT applies the GT predicate on the "icon" field.
func IconGT(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIcon), v))
	})
}

// IconGTE applies the GTE predicate on the "icon" field.
func IconGTE(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIcon), v))
	})
}

// IconLT applies the LT predicate on the "icon" field.
func IconLT(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIcon), v))
	})
}

// IconLTE applies the LTE predicate on the "icon" field.
func IconLTE(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIcon), v))
	})
}

// IconContains applies the Contains predicate on the "icon" field.
func IconContains(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIcon), v))
	})
}

// IconHasPrefix applies the HasPrefix predicate on the "icon" field.
func IconHasPrefix(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIcon), v))
	})
}

// IconHasSuffix applies the HasSuffix predicate on the "icon" field.
func IconHasSuffix(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIcon), v))
	})
}

// IconEqualFold applies the EqualFold predicate on the "icon" field.
func IconEqualFold(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIcon), v))
	})
}

// IconContainsFold applies the ContainsFold predicate on the "icon" field.
func IconContainsFold(v string) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIcon), v))
	})
}

// LayoutEQ applies the EQ predicate on the "layout" field.
func LayoutEQ(v bool) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLayout), v))
	})
}

// LayoutNEQ applies the NEQ predicate on the "layout" field.
func LayoutNEQ(v bool) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLayout), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v bool) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v bool) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// HideInMenuEQ applies the EQ predicate on the "hide_in_menu" field.
func HideInMenuEQ(v bool) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHideInMenu), v))
	})
}

// HideInMenuNEQ applies the NEQ predicate on the "hide_in_menu" field.
func HideInMenuNEQ(v bool) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHideInMenu), v))
	})
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v int64) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSort), v))
	})
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v int64) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSort), v))
	})
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...int64) predicate.Menu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Menu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSort), v...))
	})
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...int64) predicate.Menu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Menu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSort), v...))
	})
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v int64) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSort), v))
	})
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v int64) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSort), v))
	})
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v int64) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSort), v))
	})
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v int64) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSort), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Menu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Menu(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.Menu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Menu(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Menu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Menu(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.Menu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Menu(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Menu) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Menu) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
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
func Not(p predicate.Menu) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		p(s.Not())
	})
}