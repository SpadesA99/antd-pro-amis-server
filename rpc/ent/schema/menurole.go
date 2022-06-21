/*
 * @Author       : SpadesA.yanjuan9998@gmail.com
 * @Date         : 2022-06-15 15:16:49
 * @LastEditors  : SpadesA.yanjuan9998@gmail.com
 * @LastEditTime : 2022-06-20 10:26:40
 * @FilePath     : \antd-pro-amis-server\rpc\ent\schema\menurole.go
 */
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MenuRole holds the schema definition for the MenuRole entity.
type MenuRole struct {
	ent.Schema
}

// Fields of the MenuRole.
func (MenuRole) Fields() []ent.Field {
	return []ent.Field{
		field.String("role_name"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the MenuRole.
func (MenuRole) Edges() []ent.Edge {
	return nil
}
