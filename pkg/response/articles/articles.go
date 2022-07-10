package articles

import (
	"article-cache/pkg/model/articles"
)

type (
	Response struct {
		Data []articles.ModelDB `json:"data"`
	}
)
