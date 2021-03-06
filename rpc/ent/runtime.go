// Code generated by entc, DO NOT EDIT.

package ent

import (
	"antd-pro-amis-server/rpc/ent/amisschema"
	"antd-pro-amis-server/rpc/ent/menu"
	"antd-pro-amis-server/rpc/ent/menurole"
	"antd-pro-amis-server/rpc/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	amisschemaFields := schema.AmisSchema{}.Fields()
	_ = amisschemaFields
	// amisschemaDescCreatedAt is the schema descriptor for created_at field.
	amisschemaDescCreatedAt := amisschemaFields[2].Descriptor()
	// amisschema.DefaultCreatedAt holds the default value on creation for the created_at field.
	amisschema.DefaultCreatedAt = amisschemaDescCreatedAt.Default.(func() time.Time)
	// amisschemaDescUpdatedAt is the schema descriptor for updated_at field.
	amisschemaDescUpdatedAt := amisschemaFields[3].Descriptor()
	// amisschema.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	amisschema.DefaultUpdatedAt = amisschemaDescUpdatedAt.Default.(func() time.Time)
	// amisschema.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	amisschema.UpdateDefaultUpdatedAt = amisschemaDescUpdatedAt.UpdateDefault.(func() time.Time)
	menuFields := schema.Menu{}.Fields()
	_ = menuFields
	// menuDescParentID is the schema descriptor for parent_id field.
	menuDescParentID := menuFields[1].Descriptor()
	// menu.DefaultParentID holds the default value on creation for the parent_id field.
	menu.DefaultParentID = menuDescParentID.Default.(int64)
	// menuDescStatus is the schema descriptor for status field.
	menuDescStatus := menuFields[6].Descriptor()
	// menu.DefaultStatus holds the default value on creation for the status field.
	menu.DefaultStatus = menuDescStatus.Default.(bool)
	// menuDescSort is the schema descriptor for sort field.
	menuDescSort := menuFields[8].Descriptor()
	// menu.DefaultSort holds the default value on creation for the sort field.
	menu.DefaultSort = menuDescSort.Default.(int64)
	// menuDescCreatedAt is the schema descriptor for created_at field.
	menuDescCreatedAt := menuFields[9].Descriptor()
	// menu.DefaultCreatedAt holds the default value on creation for the created_at field.
	menu.DefaultCreatedAt = menuDescCreatedAt.Default.(func() time.Time)
	// menuDescUpdatedAt is the schema descriptor for updated_at field.
	menuDescUpdatedAt := menuFields[10].Descriptor()
	// menu.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	menu.DefaultUpdatedAt = menuDescUpdatedAt.Default.(func() time.Time)
	// menu.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	menu.UpdateDefaultUpdatedAt = menuDescUpdatedAt.UpdateDefault.(func() time.Time)
	menuroleFields := schema.MenuRole{}.Fields()
	_ = menuroleFields
	// menuroleDescCreatedAt is the schema descriptor for created_at field.
	menuroleDescCreatedAt := menuroleFields[1].Descriptor()
	// menurole.DefaultCreatedAt holds the default value on creation for the created_at field.
	menurole.DefaultCreatedAt = menuroleDescCreatedAt.Default.(func() time.Time)
	// menuroleDescUpdatedAt is the schema descriptor for updated_at field.
	menuroleDescUpdatedAt := menuroleFields[2].Descriptor()
	// menurole.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	menurole.DefaultUpdatedAt = menuroleDescUpdatedAt.Default.(func() time.Time)
	// menurole.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	menurole.UpdateDefaultUpdatedAt = menuroleDescUpdatedAt.UpdateDefault.(func() time.Time)
}
