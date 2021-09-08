package middleware

import (
	"bytes"
	"devops/log"
	"devops/model"
	"devops/myerr"
	"devops/res"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-uuid"
	"github.com/willf/pad"
)

func Logging() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now().UTC()
		path := context.Request.URL.Path

		reqUrl := context.Request.URL.Path
		if strings.Contains(reqUrl, "image") {
			//context.Next()
			return
		}

		var bodies []byte
		if context.Request.Body != nil {
			bodies, _ = ioutil.ReadAll(context.Request.Body)
		}
		context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodies))
		method := context.Request.Method
		ip := context.ClientIP()

		blw := &model.BodyLoggerWriter{
			ResponseWriter: context.Writer,
			Body:           bytes.NewBufferString(""),
		}
		context.Writer = blw

		context.Next()

		end := time.Now().UTC()
		sub := end.Sub(start)
		code := -100
		message := ""
		var response res.Response
		if err := json.Unmarshal(blw.Body.Bytes(), &response); err != nil {
			code = myerr.InternalServerError.Code
			message = err.Error()
			//log.Error(err)
		} else {
			code = response.Code
			message = response.Message
		}
		log.Info(fmt.Sprintf("%-13s | %-12s | %-6s %-20s | {code: %d, message: %s}", sub, ip, pad.Right(method, 3, ""), path, code, message))
	}
}

func ProcessTraceID() gin.HandlerFunc {
	return func(context *gin.Context) {
		traceId := context.Request.Header.Get("X-Trace-Id")
		if traceId == "" {
			u4id, _ := uuid.GenerateUUID()
			traceId = u4id
		}
		context.Set("X-Trace-Id", traceId)
		context.Writer.Header().Set("X-Trace-Id", traceId)
		context.Next()
	}
}
