package helper

import (
	"context"
	"os"
	"strconv"
	"time"
)

// If t == 0 then value get from env variable CONTEXT_TIMEOUT
// else using t as value
func ContextTimeout(t int) (c context.Context, cf context.CancelFunc) {

	if t == 0 {
		t, _ = strconv.Atoi(os.Getenv("CONTEXT_TIMEOUT"))
	}

	c, cf = context.WithTimeout(context.Background(), time.Duration(t)*time.Second)
	return
}
