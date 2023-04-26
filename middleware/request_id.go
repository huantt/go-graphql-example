package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func AddRequestId(c *gin.Context) {
	requestId := time.Now().UnixMilli()
	c.Set(RequestIdContextKey, strconv.FormatInt(requestId, 10))
}

const RequestIdContextKey = "request_id"
