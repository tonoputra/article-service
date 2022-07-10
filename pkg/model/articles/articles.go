package articles

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	ModelDB struct {
		ID          *primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
		ArticleId   int                 `bson:"article_id" json:"article_id"`
		AuthorName  string              `bson:"author_name" json:"author_name"`
		Title       string              `bson:"title" json:"title"`
		ArticleBody string              `bson:"article_body" json:"article_body"`
		CreatedAt   time.Time           `bson:"created_at" json:"created_at"`
	}
)
