package articles

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

	db "article-cache/internal/db/mongo"
	"article-cache/internal/helper"
	"article-cache/pkg/model/articles"
	"article-cache/pkg/response"
)

type Controller interface {
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
}

type controller struct {
	m Model
}

func NewController() Controller {
	return &controller{
		m: NewModel(),
	}
}

// FindAll godoc
// @Summary      Find All Articles
// @Description  To get all article list
// @Tags         Articles
// @Accept       json
// @Produce      json
// @Param        start   query     string  true   "used for page"
// @Param        limit   query     string  true   "used for perPage"
// @Param        sort    query     string  false  "asc | desc"
// @Param        sortBy  query     string  false  "used for sorting by field key or title"
// @Param        key     query     string  false  "used for perPage"
// @Success      200     {object}  response.Body{values=[]articles.ModelDB,total=integer,subTotal=integer}
// @Failure      400     {object}  response.Body{errors=model.Object}
// @Router       /v1/articles [get]
func (c *controller) FindAll(ctx echo.Context) error {
	filter := bson.M{}
	qParams := ctx.QueryParams()

	// FindOptions query pagination
	opts, metaArticle, err := db.Pagination().FindOptions(ctx)
	if err != nil {
		return response.WriteError(ctx, response.Body{
			HTTPStatusCode: http.StatusBadRequest,
			Message:        http.StatusText(http.StatusBadRequest),
			Errors:         err,
		})
	}

	for k := range qParams {
		switch k {
		case "limit", "start", "sort", "sortBy", "skip":
			continue
		}
		v := qParams.Get(k)
		filter[k] = v
		metaArticle[k] = v
	}

	total, err := c.m.FindAllCount(ctx, &filter, nil)
	if err != nil {
		statusCode, statusMsg := helper.HTTPCode().MapError(err)
		return response.WriteError(ctx, response.Body{
			HTTPStatusCode: statusCode,
			Message:        statusMsg,
			Errors:         err,
		})
	}

	resArticle, err := c.m.FindAll(ctx, &filter, opts)
	if err != nil {
		statusCode, statusMsg := helper.HTTPCode().MapError(err)
		return response.WriteError(ctx, response.Body{
			HTTPStatusCode: statusCode,
			Message:        statusMsg,
			Errors:         err,
		})
	}

	metaArticle["total"] = total

	return response.WriteSuccess(ctx, response.Body{
		Message:  "Success get article list",
		Data:     resArticle,
		Meta:     metaArticle,
		Total:    int(total),
		SubTotal: len(resArticle),
	})
}


// Create godoc
// @Summary      Create an article
// @Description  To create an article
// @Tags         Articles
// @Accept       json
// @Produce      json
// @Param        payload  body      articles.Create  true  "payload"
// @Success      200      {object}  response.Body{values=articles.Response}
// @Failure      400      {object}  response.Body{errors=model.Object}
// @Router       /v1/articles [post]
func (c *controller) Create(ctx echo.Context) error {
	var err error
	reqArticle := new(articles.ModelDB)
	if err = ctx.Bind(reqArticle); err != nil {
		response.WriteErrorBinding(ctx, err)
		return err
	}

	if err = ctx.Validate(reqArticle); err != nil {
		response.WriteErrorValidate(ctx, err)
		return err
	}

	idObj, err := c.m.Create(ctx, reqArticle)
	if err != nil {
		statusCode, statusMsg := helper.HTTPCode().MapError(err)
		return response.WriteError(ctx, response.Body{
			HTTPStatusCode: statusCode,
			Message:        statusMsg,
			Errors:         err,
		})
	}

	reqArticle.ID = &idObj

	return response.WriteSuccess(ctx, response.Body{
		Message: "Success create data",
		Data:    reqArticle,
	})
}
