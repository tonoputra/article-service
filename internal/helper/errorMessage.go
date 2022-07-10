package helper

import "fmt"

const (
	ERR_MSG_TEMPLATE         = "%s: %s. %v"
	ERR_MSG_MONGO_FIND_ALL   = "Error find all"
	ERR_MSG_MONGO_COUNT_ALL  = "Error count all"
	ERR_MSG_MONGO_COUNT_DOCS = "Error count documents"
	ERR_MSG_MONGO_FIND_ONE   = "Error find one"
	ERR_MSG_MONGO_INSERT     = "Error insert"
	ERR_MSG_MONGO_UPDATE     = "Error update"
	ERR_MSG_MONGO_DELETE     = "Error delete"
	ERR_MSG_MONGO_DECODE     = "Error decode mongo result"
	ERR_MSG_MONGO_AGGREGATE  = "Error aggregate data"
	ERR_REDIS_GET_DATA       = "Error get redis data"
	ERR_REDIS_SET_DATA       = "Error set redis data"
	ERR_REDIS_UNMARSHAL      = "Error unmarshal redis response"
	REDIS_VAL_EMPTY          = "Redis value is empty"
	REDIS_NIL                = "Redis key doesn't exist :"

	// general message
	GET_REDIS_DATA = "Get data from redis."
)

type (
	ErrorMessageInterface interface {
		Msg() string
	}
	errorMessage struct {
		mn string // model name
	}
)

func ErrorMessage(s string) ErrorMessageInterface {
	return &errorMessage{
		mn: s,
	}
}

// MongoErrorDecode func
func (em *errorMessage) Msg() string {
	return fmt.Sprintf("%s.%s.", em, ERR_MSG_MONGO_DECODE)
}
