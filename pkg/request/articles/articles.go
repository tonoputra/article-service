package articles

import "article-cache/pkg/model/articles"

type (
	Create struct {
		Data articles.ModelDB `bson:"data" json:"data"`
	}
)
