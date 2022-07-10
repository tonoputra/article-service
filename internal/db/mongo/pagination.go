package db

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var PAGINATION_LIMIT = int64(50)
var PAGINATION_SKIP = int64(0)

type PaginationInterface interface {
	FindOptions(ctx echo.Context) (*options.FindOptions, map[string]interface{}, error)

	limit(l string) (*int64, error)
	skip(l string) (*int64, error)
	sort(sort string, sortBy string) map[string]int
}

type pagination struct {
}

func Pagination() PaginationInterface {
	return &pagination{}
}

func (pagination) limit(v string) (*int64, error) {
	if v != "" {
		limit, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		l := int64(limit)
		return &l, nil
	} else {
		return &PAGINATION_LIMIT, nil
	}
}

func (pagination) skip(v string) (*int64, error) {
	if v != "" {
		skip, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		s := int64(skip)
		return &s, nil
	} else {
		return &PAGINATION_SKIP, nil
	}
}

func (pagination) sort(sort string, sortBy string) map[string]int {
	s := -1 // default is desc
	if sort == "asc" {
		s = 1
	}
	return map[string]int{
		sortBy: s,
	}
}

// FindOptions ...
func (p pagination) FindOptions(ctx echo.Context) (*options.FindOptions, map[string]interface{}, error) {
	var opts options.FindOptions
	var err error
	meta := map[string]interface{}{}
	qParams := ctx.QueryParams()

	// LIMIT
	opts.Limit, err = p.limit(qParams.Get("limit"))
	meta["limit"] = opts.Limit
	if err != nil {
		ctx.Logger().Errorf("value query param limit should be number only. %v", err)
		return nil, meta, err
	}

	// SKIP
	opts.Skip, err = p.skip(qParams.Get("start"))
	meta["skip"] = opts.Skip
	if err != nil {
		ctx.Logger().Errorf("value query param skip should be number only. %v", err)
		return nil, meta, err
	}

	// SORT & SORTBY
	if qParams.Get("sort") != "" && qParams.Get("sortBy") != "" {
		opts.Sort = p.sort(qParams.Get("sort"), qParams.Get("sortBy"))
		meta["sort"] = qParams.Get("sort")
		meta["sortBy"] = qParams.Get("sortBy")
	}

	return &opts, meta, nil
}
