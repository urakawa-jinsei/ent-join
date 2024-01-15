// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ent-join/ent/content"
	"github.com/ent-join/ent/contentmoviemetadata"
	"github.com/ent-join/ent/predicate"
)

// ContentMovieMetadataQuery is the builder for querying ContentMovieMetadata entities.
type ContentMovieMetadataQuery struct {
	config
	ctx         *QueryContext
	order       []contentmoviemetadata.OrderOption
	inters      []Interceptor
	predicates  []predicate.ContentMovieMetadata
	withContent *ContentQuery
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ContentMovieMetadataQuery builder.
func (cmmq *ContentMovieMetadataQuery) Where(ps ...predicate.ContentMovieMetadata) *ContentMovieMetadataQuery {
	cmmq.predicates = append(cmmq.predicates, ps...)
	return cmmq
}

// Limit the number of records to be returned by this query.
func (cmmq *ContentMovieMetadataQuery) Limit(limit int) *ContentMovieMetadataQuery {
	cmmq.ctx.Limit = &limit
	return cmmq
}

// Offset to start from.
func (cmmq *ContentMovieMetadataQuery) Offset(offset int) *ContentMovieMetadataQuery {
	cmmq.ctx.Offset = &offset
	return cmmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cmmq *ContentMovieMetadataQuery) Unique(unique bool) *ContentMovieMetadataQuery {
	cmmq.ctx.Unique = &unique
	return cmmq
}

// Order specifies how the records should be ordered.
func (cmmq *ContentMovieMetadataQuery) Order(o ...contentmoviemetadata.OrderOption) *ContentMovieMetadataQuery {
	cmmq.order = append(cmmq.order, o...)
	return cmmq
}

// QueryContent chains the current query on the "content" edge.
func (cmmq *ContentMovieMetadataQuery) QueryContent() *ContentQuery {
	query := (&ContentClient{config: cmmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cmmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cmmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(contentmoviemetadata.Table, contentmoviemetadata.FieldID, selector),
			sqlgraph.To(content.Table, content.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, contentmoviemetadata.ContentTable, contentmoviemetadata.ContentColumn),
		)
		fromU = sqlgraph.SetNeighbors(cmmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ContentMovieMetadata entity from the query.
// Returns a *NotFoundError when no ContentMovieMetadata was found.
func (cmmq *ContentMovieMetadataQuery) First(ctx context.Context) (*ContentMovieMetadata, error) {
	nodes, err := cmmq.Limit(1).All(setContextOp(ctx, cmmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{contentmoviemetadata.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cmmq *ContentMovieMetadataQuery) FirstX(ctx context.Context) *ContentMovieMetadata {
	node, err := cmmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ContentMovieMetadata ID from the query.
// Returns a *NotFoundError when no ContentMovieMetadata ID was found.
func (cmmq *ContentMovieMetadataQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cmmq.Limit(1).IDs(setContextOp(ctx, cmmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{contentmoviemetadata.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cmmq *ContentMovieMetadataQuery) FirstIDX(ctx context.Context) string {
	id, err := cmmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ContentMovieMetadata entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ContentMovieMetadata entity is found.
// Returns a *NotFoundError when no ContentMovieMetadata entities are found.
func (cmmq *ContentMovieMetadataQuery) Only(ctx context.Context) (*ContentMovieMetadata, error) {
	nodes, err := cmmq.Limit(2).All(setContextOp(ctx, cmmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{contentmoviemetadata.Label}
	default:
		return nil, &NotSingularError{contentmoviemetadata.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cmmq *ContentMovieMetadataQuery) OnlyX(ctx context.Context) *ContentMovieMetadata {
	node, err := cmmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ContentMovieMetadata ID in the query.
// Returns a *NotSingularError when more than one ContentMovieMetadata ID is found.
// Returns a *NotFoundError when no entities are found.
func (cmmq *ContentMovieMetadataQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cmmq.Limit(2).IDs(setContextOp(ctx, cmmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{contentmoviemetadata.Label}
	default:
		err = &NotSingularError{contentmoviemetadata.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cmmq *ContentMovieMetadataQuery) OnlyIDX(ctx context.Context) string {
	id, err := cmmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ContentMovieMetadataSlice.
func (cmmq *ContentMovieMetadataQuery) All(ctx context.Context) ([]*ContentMovieMetadata, error) {
	ctx = setContextOp(ctx, cmmq.ctx, "All")
	if err := cmmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ContentMovieMetadata, *ContentMovieMetadataQuery]()
	return withInterceptors[[]*ContentMovieMetadata](ctx, cmmq, qr, cmmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cmmq *ContentMovieMetadataQuery) AllX(ctx context.Context) []*ContentMovieMetadata {
	nodes, err := cmmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ContentMovieMetadata IDs.
func (cmmq *ContentMovieMetadataQuery) IDs(ctx context.Context) (ids []string, err error) {
	if cmmq.ctx.Unique == nil && cmmq.path != nil {
		cmmq.Unique(true)
	}
	ctx = setContextOp(ctx, cmmq.ctx, "IDs")
	if err = cmmq.Select(contentmoviemetadata.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cmmq *ContentMovieMetadataQuery) IDsX(ctx context.Context) []string {
	ids, err := cmmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cmmq *ContentMovieMetadataQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cmmq.ctx, "Count")
	if err := cmmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cmmq, querierCount[*ContentMovieMetadataQuery](), cmmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cmmq *ContentMovieMetadataQuery) CountX(ctx context.Context) int {
	count, err := cmmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cmmq *ContentMovieMetadataQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cmmq.ctx, "Exist")
	switch _, err := cmmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cmmq *ContentMovieMetadataQuery) ExistX(ctx context.Context) bool {
	exist, err := cmmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ContentMovieMetadataQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cmmq *ContentMovieMetadataQuery) Clone() *ContentMovieMetadataQuery {
	if cmmq == nil {
		return nil
	}
	return &ContentMovieMetadataQuery{
		config:      cmmq.config,
		ctx:         cmmq.ctx.Clone(),
		order:       append([]contentmoviemetadata.OrderOption{}, cmmq.order...),
		inters:      append([]Interceptor{}, cmmq.inters...),
		predicates:  append([]predicate.ContentMovieMetadata{}, cmmq.predicates...),
		withContent: cmmq.withContent.Clone(),
		// clone intermediate query.
		sql:  cmmq.sql.Clone(),
		path: cmmq.path,
	}
}

// WithContent tells the query-builder to eager-load the nodes that are connected to
// the "content" edge. The optional arguments are used to configure the query builder of the edge.
func (cmmq *ContentMovieMetadataQuery) WithContent(opts ...func(*ContentQuery)) *ContentMovieMetadataQuery {
	query := (&ContentClient{config: cmmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cmmq.withContent = query
	return cmmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Width int `json:"width,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ContentMovieMetadata.Query().
//		GroupBy(contentmoviemetadata.FieldWidth).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cmmq *ContentMovieMetadataQuery) GroupBy(field string, fields ...string) *ContentMovieMetadataGroupBy {
	cmmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ContentMovieMetadataGroupBy{build: cmmq}
	grbuild.flds = &cmmq.ctx.Fields
	grbuild.label = contentmoviemetadata.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Width int `json:"width,omitempty"`
//	}
//
//	client.ContentMovieMetadata.Query().
//		Select(contentmoviemetadata.FieldWidth).
//		Scan(ctx, &v)
func (cmmq *ContentMovieMetadataQuery) Select(fields ...string) *ContentMovieMetadataSelect {
	cmmq.ctx.Fields = append(cmmq.ctx.Fields, fields...)
	sbuild := &ContentMovieMetadataSelect{ContentMovieMetadataQuery: cmmq}
	sbuild.label = contentmoviemetadata.Label
	sbuild.flds, sbuild.scan = &cmmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ContentMovieMetadataSelect configured with the given aggregations.
func (cmmq *ContentMovieMetadataQuery) Aggregate(fns ...AggregateFunc) *ContentMovieMetadataSelect {
	return cmmq.Select().Aggregate(fns...)
}

func (cmmq *ContentMovieMetadataQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cmmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cmmq); err != nil {
				return err
			}
		}
	}
	for _, f := range cmmq.ctx.Fields {
		if !contentmoviemetadata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cmmq.path != nil {
		prev, err := cmmq.path(ctx)
		if err != nil {
			return err
		}
		cmmq.sql = prev
	}
	return nil
}

func (cmmq *ContentMovieMetadataQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ContentMovieMetadata, error) {
	var (
		nodes       = []*ContentMovieMetadata{}
		_spec       = cmmq.querySpec()
		loadedTypes = [1]bool{
			cmmq.withContent != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ContentMovieMetadata).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ContentMovieMetadata{config: cmmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cmmq.modifiers) > 0 {
		_spec.Modifiers = cmmq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cmmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cmmq.withContent; query != nil {
		if err := cmmq.loadContent(ctx, query, nodes, nil,
			func(n *ContentMovieMetadata, e *Content) { n.Edges.Content = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cmmq *ContentMovieMetadataQuery) loadContent(ctx context.Context, query *ContentQuery, nodes []*ContentMovieMetadata, init func(*ContentMovieMetadata), assign func(*ContentMovieMetadata, *Content)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*ContentMovieMetadata)
	for i := range nodes {
		fk := nodes[i].ID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(content.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cmmq *ContentMovieMetadataQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cmmq.querySpec()
	if len(cmmq.modifiers) > 0 {
		_spec.Modifiers = cmmq.modifiers
	}
	_spec.Node.Columns = cmmq.ctx.Fields
	if len(cmmq.ctx.Fields) > 0 {
		_spec.Unique = cmmq.ctx.Unique != nil && *cmmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cmmq.driver, _spec)
}

func (cmmq *ContentMovieMetadataQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(contentmoviemetadata.Table, contentmoviemetadata.Columns, sqlgraph.NewFieldSpec(contentmoviemetadata.FieldID, field.TypeString))
	_spec.From = cmmq.sql
	if unique := cmmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cmmq.path != nil {
		_spec.Unique = true
	}
	if fields := cmmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contentmoviemetadata.FieldID)
		for i := range fields {
			if fields[i] != contentmoviemetadata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cmmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cmmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cmmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cmmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cmmq *ContentMovieMetadataQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cmmq.driver.Dialect())
	t1 := builder.Table(contentmoviemetadata.Table)
	columns := cmmq.ctx.Fields
	if len(columns) == 0 {
		columns = contentmoviemetadata.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cmmq.sql != nil {
		selector = cmmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cmmq.ctx.Unique != nil && *cmmq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range cmmq.modifiers {
		m(selector)
	}
	for _, p := range cmmq.predicates {
		p(selector)
	}
	for _, p := range cmmq.order {
		p(selector)
	}
	if offset := cmmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cmmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cmmq *ContentMovieMetadataQuery) Modify(modifiers ...func(s *sql.Selector)) *ContentMovieMetadataSelect {
	cmmq.modifiers = append(cmmq.modifiers, modifiers...)
	return cmmq.Select()
}

// ContentMovieMetadataGroupBy is the group-by builder for ContentMovieMetadata entities.
type ContentMovieMetadataGroupBy struct {
	selector
	build *ContentMovieMetadataQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cmmgb *ContentMovieMetadataGroupBy) Aggregate(fns ...AggregateFunc) *ContentMovieMetadataGroupBy {
	cmmgb.fns = append(cmmgb.fns, fns...)
	return cmmgb
}

// Scan applies the selector query and scans the result into the given value.
func (cmmgb *ContentMovieMetadataGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cmmgb.build.ctx, "GroupBy")
	if err := cmmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ContentMovieMetadataQuery, *ContentMovieMetadataGroupBy](ctx, cmmgb.build, cmmgb, cmmgb.build.inters, v)
}

func (cmmgb *ContentMovieMetadataGroupBy) sqlScan(ctx context.Context, root *ContentMovieMetadataQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cmmgb.fns))
	for _, fn := range cmmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cmmgb.flds)+len(cmmgb.fns))
		for _, f := range *cmmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cmmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cmmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ContentMovieMetadataSelect is the builder for selecting fields of ContentMovieMetadata entities.
type ContentMovieMetadataSelect struct {
	*ContentMovieMetadataQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cmms *ContentMovieMetadataSelect) Aggregate(fns ...AggregateFunc) *ContentMovieMetadataSelect {
	cmms.fns = append(cmms.fns, fns...)
	return cmms
}

// Scan applies the selector query and scans the result into the given value.
func (cmms *ContentMovieMetadataSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cmms.ctx, "Select")
	if err := cmms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ContentMovieMetadataQuery, *ContentMovieMetadataSelect](ctx, cmms.ContentMovieMetadataQuery, cmms, cmms.inters, v)
}

func (cmms *ContentMovieMetadataSelect) sqlScan(ctx context.Context, root *ContentMovieMetadataQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cmms.fns))
	for _, fn := range cmms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cmms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cmms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cmms *ContentMovieMetadataSelect) Modify(modifiers ...func(s *sql.Selector)) *ContentMovieMetadataSelect {
	cmms.modifiers = append(cmms.modifiers, modifiers...)
	return cmms
}
