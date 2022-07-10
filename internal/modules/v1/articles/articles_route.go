package articles

import (
	"article-cache/internal/helper"
)

type Route interface {
	Routes() []helper.Route
}

type route struct {
	c Controller
}

func NewRoute() Route {
	return &route{
		c: NewController(),
	}
}

func (r *route) Routes() []helper.Route {
	return []helper.Route{
		{
			Method:      "GET",
			Pattern:     "/articles",
			HandlerFunc: r.c.FindAll,
		},
		{
			Method:      "POST",
			Pattern:     "/articles",
			HandlerFunc: r.c.Create,
		},
	}
}
