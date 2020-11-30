package logger

import (
	"time"

	"github.com/gin-gonic/gin"
)

var (
	// XReqIDHeader1 request id header
	XReqIDHeader1 = "X-Reqid"
	// XReqIDHeader2 request id header 2
	XReqIDHeader2 = "X-RequestID"

	// LogCtxKey logger save in gin context
	LogCtxKey = "log_ctx_key"

	// XLogKey log req handle time
	XLogKey = "X-Log"
)

// GinLoggerMiddleware gin web framework logger middleware
func GinLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.GetHeader(XReqIDHeader1)
		if len(reqID) == 0 {
			reqID = c.GetHeader(XReqIDHeader2)
		}

		if len(reqID) == 0 {
			reqID = GenReqID()
		}

		c.Header(XReqIDHeader1, reqID)

		log := newLogger(reqID)
		l := New(reqID)
		c.Set(LogCtxKey, l)
		logReq(log, c)
		now := time.Now()
		c.Next()
		c.Writer.Header().Add(XLogKey, l.XLog())
		logResponse(log, c, now)
	}
}

func logReq(log Logger, c *gin.Context) {
	field := map[string]interface{}{
		"method":    c.Request.Method,
		"path":      c.Request.URL.Path,
		"client_ip": c.ClientIP(),
		"type":      "REQ",
	}

	log.WithField(field).Info("[Started]")
}

func logResponse(log Logger, c *gin.Context, startTime time.Time) {
	field := map[string]interface{}{
		"method":    c.Request.Method,
		"path":      c.Request.URL.Path,
		"status":    c.Writer.Status(),
		"latency":   time.Since(startTime).String(), // 耗时
		"client_ip": c.ClientIP(),
		"type":      "REQ",
	}

	log.WithField(field).Info("[Completed]")
}

// GinRequestLogger if contex not exist, create a new logger
func GinRequestLogger(c *gin.Context) Logger {
	val, ok := c.Get(LogCtxKey)
	if ok {
		return val.(Logger)
	}

	return New()
}
