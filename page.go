package aveecare

import (
	"context"
	"net/url"
)

// listEnvelope is the wire shape of every list response.
type listEnvelope[T any] struct {
	Object     string  `json:"object"`
	Data       []T     `json:"data"`
	NextCursor *string `json:"next_cursor"`
	HasMore    bool    `json:"has_more"`
}

// Page is one page of a list response.
//
// Walk every item on every page with the resource's ListAutoPaging method (or
// Page.AutoPaging), or page by hand:
//
//	for page.HasNextPage() {
//		page, err = page.NextPage(ctx)
//		...
//	}
type Page[T any] struct {
	// Object is always "list".
	Object string
	// Data holds the objects on this page.
	Data []T
	// NextCursor is the cursor of the following page; empty on the last page.
	NextCursor string
	// HasMore reports whether another page exists.
	HasMore bool

	fetch func(ctx context.Context, cursor string) (*Page[T], error)
}

// HasNextPage reports whether NextPage can fetch another page.
func (p *Page[T]) HasNextPage() bool { return p.HasMore && p.NextCursor != "" }

// NextPage fetches the following page with the same filters. It returns
// ErrNoMorePages when there is none; check HasNextPage first.
func (p *Page[T]) NextPage(ctx context.Context) (*Page[T], error) {
	if !p.HasNextPage() || p.fetch == nil {
		return nil, ErrNoMorePages
	}
	return p.fetch(ctx, p.NextCursor)
}

// AutoPaging returns an Iterator over this page's items and every following
// page, fetching pages lazily.
func (p *Page[T]) AutoPaging(ctx context.Context) *Iterator[T] {
	return newIterator(ctx, func(context.Context) (*Page[T], error) { return p, nil })
}

// getPage requests one list page and arms it to fetch the next.
func getPage[T any](ctx context.Context, c *core, path string, query url.Values, opts []RequestOption) (*Page[T], error) {
	ro := newRequestOptions(opts)
	var env listEnvelope[T]
	if err := c.do(ctx, requestSpec{method: "GET", path: path, query: query}, ro, &env); err != nil {
		return nil, err
	}
	page := &Page[T]{Object: env.Object, Data: env.Data}
	if page.Data == nil {
		page.Data = []T{}
	}
	if env.NextCursor != nil {
		page.NextCursor = *env.NextCursor
	}
	page.HasMore = env.HasMore && page.NextCursor != ""
	page.fetch = func(ctx context.Context, cursor string) (*Page[T], error) {
		next := make(url.Values, len(query)+1)
		for k, vs := range query {
			next[k] = append([]string(nil), vs...)
		}
		next.Set("cursor", cursor)
		return getPage[T](ctx, c, path, next, opts)
	}
	return page, nil
}

// Iterator walks every item of a list across all its pages. Pages are fetched
// lazily as the iteration reaches them.
//
//	it := client.Patients.ListAutoPaging(ctx, &aveecare.PatientListParams{Limit: aveecare.Int(100)})
//	for it.Next() {
//		patient := it.Item()
//		...
//	}
//	if err := it.Err(); err != nil {
//		...
//	}
//
// With Go 1.23 or newer, Iterator.All also supports range-over-func.
type Iterator[T any] struct {
	ctx   context.Context
	first func(context.Context) (*Page[T], error)
	page  *Page[T]
	idx   int
	cur   T
	err   error
	done  bool
}

func newIterator[T any](ctx context.Context, first func(context.Context) (*Page[T], error)) *Iterator[T] {
	return &Iterator[T]{ctx: ctx, first: first}
}

// Next advances to the next item, fetching another page when needed. It
// returns false at the end of the list or on error; check Err afterwards.
func (it *Iterator[T]) Next() bool {
	if it.err != nil || it.done {
		return false
	}
	for {
		if it.page == nil {
			page, err := it.first(it.ctx)
			if err != nil {
				it.err = err
				return false
			}
			it.page, it.idx = page, 0
		}
		if it.idx < len(it.page.Data) {
			it.cur = it.page.Data[it.idx]
			it.idx++
			return true
		}
		if !it.page.HasNextPage() {
			it.done = true
			return false
		}
		page, err := it.page.NextPage(it.ctx)
		if err != nil {
			it.err = err
			return false
		}
		it.page, it.idx = page, 0
	}
}

// Item returns the current item. It is only valid after Next returned true.
func (it *Iterator[T]) Item() T { return it.cur }

// Err returns the error that stopped the iteration, if any.
func (it *Iterator[T]) Err() error { return it.err }
