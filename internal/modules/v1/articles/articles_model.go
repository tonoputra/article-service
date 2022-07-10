package articles

import (
	db "article-cache/internal/db/mongo"
	"article-cache/internal/helper"
	"article-cache/pkg/model/articles"
	ctype "article-cache/pkg/types/articles"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model interface {
	FindAll(ctx echo.Context, filter *bson.M, opts *options.FindOptions) (res []articles.ModelDB, err error)
	FindAllCount(ctx echo.Context, filter *bson.M, opts *options.CountOptions) (count int64, err error)
	Create(ctx echo.Context, req *articles.ModelDB) (id primitive.ObjectID, err error)
	GetRedis(ctx echo.Context, key string) (res []articles.ModelDB, err error)
	SetRedis(ctx echo.Context, articleList []articles.ModelDB, key string) (err error)
}

type model struct {
}

func NewModel() Model {
	return &model{}
}

// FindAll ...
func (m *model) FindAll(ctx echo.Context, filter *bson.M, opts *options.FindOptions) (res []articles.ModelDB, err error) {
	cto, cancel := helper.ContextTimeout(0)
	defer cancel()

	csr, err := db.Mongo.Conn.Collection(helper.COLL_ARTICLES).Find(cto, filter, opts)
	if err != nil {
		ctx.Logger().Errorf(helper.ERR_MSG_TEMPLATE, ctype.ARTICLES_MODEL_FIND_ALL, helper.ERR_MSG_MONGO_FIND_ALL, err)
		return
	}

	if err = csr.All(cto, &res); err != nil {
		ctx.Logger().Errorf(helper.ERR_MSG_TEMPLATE, ctype.ARTICLES_MODEL_FIND_ALL, helper.ERR_MSG_MONGO_DECODE, err)
		return
	}
	return
}

// FindAllCount ...
func (m *model) FindAllCount(ctx echo.Context, filter *bson.M, opts *options.CountOptions) (count int64, err error) {
	cto, cancel := helper.ContextTimeout(0)
	defer cancel()

	count, err = db.Mongo.Conn.Collection(helper.COLL_ARTICLES).CountDocuments(cto, filter, opts)
	if err != nil {
		ctx.Logger().Errorf(helper.ERR_MSG_TEMPLATE, ctype.ARTICLES_MODEL_COUNT_ALL, helper.ERR_MSG_MONGO_COUNT_ALL, err)
		return
	}

	return
}

// Create ...
func (m *model) Create(ctx echo.Context, req *articles.ModelDB) (id primitive.ObjectID, err error) {
	// cto, cancel := http.Context().ContextTimeout(0)
	// defer cancel()

	csr, err := db.Mongo.Conn.Collection(helper.COLL_ARTICLES).InsertOne(ctx.Request().Context(), req)
	if err != nil {
		ctx.Logger().Errorf(helper.ERR_MSG_TEMPLATE, ctype.ARTICLES_MODEL_CREATE, helper.ERR_MSG_MONGO_INSERT, err)
		return
	}

	id = csr.InsertedID.(primitive.ObjectID)

	return
}

func (m *model) GetRedis(ctx echo.Context, key string) (res []articles.ModelDB, err error) {
	redisKey := fmt.Sprintf("article-service_%s", key)

	result, err := helper.Redis.Get(context.Background(), redisKey)
	if err != nil {
		ctx.Logger().Errorf(helper.ERR_MSG_TEMPLATE, ctype.ARTICLES_MODEL_REDIS_FIND_ALL, helper.ERR_REDIS_GET_DATA, err)
		return
	}

	if result != "" {
		if err = json.Unmarshal([]byte(result), &res); err != nil {
			ctx.Logger().Errorf(helper.ERR_MSG_TEMPLATE, ctype.ARTICLES_MODEL_REDIS_FIND_ALL, helper.ERR_REDIS_UNMARSHAL, err)
			return
		}

		fmt.Println(helper.GET_REDIS_DATA)
		return
	}

	return
}

func (m *model) SetRedis(ctx echo.Context, articleList []articles.ModelDB, key string) (err error) {
	redisKey := fmt.Sprintf("article-service_%s", key)
	values, _ := json.Marshal(articleList)
	err = helper.Redis.Set(context.Background(), redisKey, values, time.Duration(5*time.Minute))
	if err != nil {
		ctx.Logger().Errorf(helper.ERR_MSG_TEMPLATE, ctype.ARTICLES_MODEL_REDIS_FIND_ALL, helper.ERR_REDIS_SET_DATA, err)
		return
	}

	return
}
