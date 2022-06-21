/*
 * @Author       : SpadesA.yanjuan9998@gmail.com
 * @Date         : 2022-06-15 17:48:01
 * @LastEditors  : SpadesA.yanjuan9998@gmail.com
 * @LastEditTime : 2022-06-21 13:36:28
 * @FilePath     : \antd-pro-amis-server\rpc\ent\schema\menu.go
 */
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Menu holds the schema definition for the Menu entity.
type Menu struct {
	ent.Schema
}

// Fields of the Menu.
func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.String("menu_name"),
		field.Int64("parent_id").Default(0),
		field.String("path"),
		field.String("icon"),
		field.Strings("roles"),
		field.Bool("layout"),
		field.Bool("status").Default(true),
		field.Bool("hide_in_menu"),
		field.Int64("sort").Default(0),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Menu.
func (Menu) Edges() []ent.Edge {
	return nil
}
