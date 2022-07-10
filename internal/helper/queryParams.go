package helper

import (
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
)

type QueryParamsInterface interface {
	Get(q url.Values, meta map[string]interface{}) (bson.M, map[string]interface{})
}

type queryParams struct {
}

func QueryParams() QueryParamsInterface {
	return &queryParams{}
}

// Get func
func (queryParams) Get(q url.Values, meta map[string]interface{}) (bson.M, map[string]interface{}) {
	var filter = bson.M{}
	for k := range q {
		switch k {
		case "limit", "start", "sort", "sortBy", "skip":
			continue
		}
		v := q.Get(k)
		filter[k] = v
		meta[k] = v
	}

	return filter, meta
}
