package middleware

import (
	"bytes"
	"io"
	"time"

	"bit-labs.cn/flex-admin/app/model"
	"bit-labs.cn/flex-admin/app/service"
	"github.com/gin-gonic/gin"
)

func OperationLog(logSvc *service.LogService) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		var body string
		if c.Request.Body != nil {
			b, _ := io.ReadAll(c.Request.Body)
			body = string(b)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(b))
		}
		c.Next()
		cost := time.Since(start).Milliseconds()
		status := c.Writer.Status()
		var user *model.User
		if v, ok := c.Get("user"); ok {
			user = v.(*model.User)
		}
		_ = logSvc.RecordOperation(c, user, status, int(cost), body)
	}
}
