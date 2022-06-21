/*
 * @Author       : SpadesA.yanjuan9998@gmail.com
 * @Date         : 2022-06-15 11:48:27
 * @LastEditors  : SpadesA.yanjuan9998@gmail.com
 * @LastEditTime : 2022-06-20 10:25:19
 * @FilePath     : \antd-pro-amis-server\rpc\ent\schema\amisschema.go
 */
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AmisSchema holds the schema definition for the AmisSchema entity.
type AmisSchema struct {
	ent.Schema
}

// Fields of the AmisSchema.
func (AmisSchema) Fields() []ent.Field {
	return []ent.Field{
		field.String("router"),
		field.Text("schema"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the AmisSchema.
func (AmisSchema) Edges() []ent.Edge {
	return nil
}
