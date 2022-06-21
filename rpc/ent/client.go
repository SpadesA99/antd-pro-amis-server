// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"antd-pro-amis-server/rpc/ent/migrate"

	"antd-pro-amis-server/rpc/ent/amisschema"
	"antd-pro-amis-server/rpc/ent/menu"
	"antd-pro-amis-server/rpc/ent/menurole"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// AmisSchema is the client for interacting with the AmisSchema builders.
	AmisSchema *AmisSchemaClient
	// Menu is the client for interacting with the Menu builders.
	Menu *MenuClient
	// MenuRole is the client for interacting with the MenuRole builders.
	MenuRole *MenuRoleClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.AmisSchema = NewAmisSchemaClient(c.config)
	c.Menu = NewMenuClient(c.config)
	c.MenuRole = NewMenuRoleClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		AmisSchema: NewAmisSchemaClient(cfg),
		Menu:       NewMenuClient(cfg),
		MenuRole:   NewMenuRoleClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		AmisSchema: NewAmisSchemaClient(cfg),
		Menu:       NewMenuClient(cfg),
		MenuRole:   NewMenuRoleClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		AmisSchema.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.AmisSchema.Use(hooks...)
	c.Menu.Use(hooks...)
	c.MenuRole.Use(hooks...)
}

// AmisSchemaClient is a client for the AmisSchema schema.
type AmisSchemaClient struct {
	config
}

// NewAmisSchemaClient returns a client for the AmisSchema from the given config.
func NewAmisSchemaClient(c config) *AmisSchemaClient {
	return &AmisSchemaClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `amisschema.Hooks(f(g(h())))`.
func (c *AmisSchemaClient) Use(hooks ...Hook) {
	c.hooks.AmisSchema = append(c.hooks.AmisSchema, hooks...)
}

// Create returns a create builder for AmisSchema.
func (c *AmisSchemaClient) Create() *AmisSchemaCreate {
	mutation := newAmisSchemaMutation(c.config, OpCreate)
	return &AmisSchemaCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AmisSchema entities.
func (c *AmisSchemaClient) CreateBulk(builders ...*AmisSchemaCreate) *AmisSchemaCreateBulk {
	return &AmisSchemaCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AmisSchema.
func (c *AmisSchemaClient) Update() *AmisSchemaUpdate {
	mutation := newAmisSchemaMutation(c.config, OpUpdate)
	return &AmisSchemaUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AmisSchemaClient) UpdateOne(as *AmisSchema) *AmisSchemaUpdateOne {
	mutation := newAmisSchemaMutation(c.config, OpUpdateOne, withAmisSchema(as))
	return &AmisSchemaUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AmisSchemaClient) UpdateOneID(id int) *AmisSchemaUpdateOne {
	mutation := newAmisSchemaMutation(c.config, OpUpdateOne, withAmisSchemaID(id))
	return &AmisSchemaUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AmisSchema.
func (c *AmisSchemaClient) Delete() *AmisSchemaDelete {
	mutation := newAmisSchemaMutation(c.config, OpDelete)
	return &AmisSchemaDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AmisSchemaClient) DeleteOne(as *AmisSchema) *AmisSchemaDeleteOne {
	return c.DeleteOneID(as.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AmisSchemaClient) DeleteOneID(id int) *AmisSchemaDeleteOne {
	builder := c.Delete().Where(amisschema.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AmisSchemaDeleteOne{builder}
}

// Query returns a query builder for AmisSchema.
func (c *AmisSchemaClient) Query() *AmisSchemaQuery {
	return &AmisSchemaQuery{
		config: c.config,
	}
}

// Get returns a AmisSchema entity by its id.
func (c *AmisSchemaClient) Get(ctx context.Context, id int) (*AmisSchema, error) {
	return c.Query().Where(amisschema.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AmisSchemaClient) GetX(ctx context.Context, id int) *AmisSchema {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AmisSchemaClient) Hooks() []Hook {
	return c.hooks.AmisSchema
}

// MenuClient is a client for the Menu schema.
type MenuClient struct {
	config
}

// NewMenuClient returns a client for the Menu from the given config.
func NewMenuClient(c config) *MenuClient {
	return &MenuClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `menu.Hooks(f(g(h())))`.
func (c *MenuClient) Use(hooks ...Hook) {
	c.hooks.Menu = append(c.hooks.Menu, hooks...)
}

// Create returns a create builder for Menu.
func (c *MenuClient) Create() *MenuCreate {
	mutation := newMenuMutation(c.config, OpCreate)
	return &MenuCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Menu entities.
func (c *MenuClient) CreateBulk(builders ...*MenuCreate) *MenuCreateBulk {
	return &MenuCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Menu.
func (c *MenuClient) Update() *MenuUpdate {
	mutation := newMenuMutation(c.config, OpUpdate)
	return &MenuUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MenuClient) UpdateOne(m *Menu) *MenuUpdateOne {
	mutation := newMenuMutation(c.config, OpUpdateOne, withMenu(m))
	return &MenuUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MenuClient) UpdateOneID(id int) *MenuUpdateOne {
	mutation := newMenuMutation(c.config, OpUpdateOne, withMenuID(id))
	return &MenuUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Menu.
func (c *MenuClient) Delete() *MenuDelete {
	mutation := newMenuMutation(c.config, OpDelete)
	return &MenuDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MenuClient) DeleteOne(m *Menu) *MenuDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MenuClient) DeleteOneID(id int) *MenuDeleteOne {
	builder := c.Delete().Where(menu.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MenuDeleteOne{builder}
}

// Query returns a query builder for Menu.
func (c *MenuClient) Query() *MenuQuery {
	return &MenuQuery{
		config: c.config,
	}
}

// Get returns a Menu entity by its id.
func (c *MenuClient) Get(ctx context.Context, id int) (*Menu, error) {
	return c.Query().Where(menu.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MenuClient) GetX(ctx context.Context, id int) *Menu {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MenuClient) Hooks() []Hook {
	return c.hooks.Menu
}

// MenuRoleClient is a client for the MenuRole schema.
type MenuRoleClient struct {
	config
}

// NewMenuRoleClient returns a client for the MenuRole from the given config.
func NewMenuRoleClient(c config) *MenuRoleClient {
	return &MenuRoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `menurole.Hooks(f(g(h())))`.
func (c *MenuRoleClient) Use(hooks ...Hook) {
	c.hooks.MenuRole = append(c.hooks.MenuRole, hooks...)
}

// Create returns a create builder for MenuRole.
func (c *MenuRoleClient) Create() *MenuRoleCreate {
	mutation := newMenuRoleMutation(c.config, OpCreate)
	return &MenuRoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MenuRole entities.
func (c *MenuRoleClient) CreateBulk(builders ...*MenuRoleCreate) *MenuRoleCreateBulk {
	return &MenuRoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MenuRole.
func (c *MenuRoleClient) Update() *MenuRoleUpdate {
	mutation := newMenuRoleMutation(c.config, OpUpdate)
	return &MenuRoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MenuRoleClient) UpdateOne(mr *MenuRole) *MenuRoleUpdateOne {
	mutation := newMenuRoleMutation(c.config, OpUpdateOne, withMenuRole(mr))
	return &MenuRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MenuRoleClient) UpdateOneID(id int) *MenuRoleUpdateOne {
	mutation := newMenuRoleMutation(c.config, OpUpdateOne, withMenuRoleID(id))
	return &MenuRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MenuRole.
func (c *MenuRoleClient) Delete() *MenuRoleDelete {
	mutation := newMenuRoleMutation(c.config, OpDelete)
	return &MenuRoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MenuRoleClient) DeleteOne(mr *MenuRole) *MenuRoleDeleteOne {
	return c.DeleteOneID(mr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MenuRoleClient) DeleteOneID(id int) *MenuRoleDeleteOne {
	builder := c.Delete().Where(menurole.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MenuRoleDeleteOne{builder}
}

// Query returns a query builder for MenuRole.
func (c *MenuRoleClient) Query() *MenuRoleQuery {
	return &MenuRoleQuery{
		config: c.config,
	}
}

// Get returns a MenuRole entity by its id.
func (c *MenuRoleClient) Get(ctx context.Context, id int) (*MenuRole, error) {
	return c.Query().Where(menurole.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MenuRoleClient) GetX(ctx context.Context, id int) *MenuRole {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MenuRoleClient) Hooks() []Hook {
	return c.hooks.MenuRole
}
