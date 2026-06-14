package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func LoggerMiddleware() gin.HandlerFunc {
	logPath := "logs/http.log"

	if err := os.MkdirAll(filepath.Dir(logPath), os.ModePerm); err != nil {
		panic(err)
	}

	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		panic(err)
	}

	logger := zerolog.New(logFile).With().Timestamp().Logger()
	return func(ctx *gin.Context) {
		start := time.Now()

		contentTye := ctx.GetHeader("Content-Type")
		//
		requestBody := make(map[string]any)

		var formfiles []map[string]any

		if strings.HasPrefix(contentTye, "multipart/form-data") {
			log.Println("multipart/form-data")

			if err := ctx.Request.ParseMultipartForm(32 << 20); err == nil && ctx.Request.MultipartForm != nil {
				for key, vals := range ctx.Request.MultipartForm.Value {
					if len(vals) == 1 {
						requestBody[key] = vals[0]
					} else {
						requestBody[key] = vals
					}
				}
				//file
				for field, files := range ctx.Request.MultipartForm.File {
					for _, f := range files {
						formfiles = append(formfiles, map[string]any{
							"field":        field,
							"filename":     f.Filename,
							"size":         formatfileSize(f.Size),
							"content_type": f.Header.Get("COntent-type"),
						})
					}
				}

				if len(formfiles) > 0 {
					requestBody["form_files"] = formfiles
				}
			}
		} else {
			bodyBytes, err := io.ReadAll(ctx.Request.Body)
			if err != nil {
				logger.Error().Err(err).Msg("FAILED to read request body")
			}

			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			if strings.HasPrefix(contentTye, "application/json") {
				_ = json.Unmarshal(bodyBytes, &requestBody)
			} else {
				//
				values, _ := url.ParseQuery(string(bodyBytes))
				log.Println(values)
				for key, vals := range values {
					if len(vals) == 1 {
						requestBody[key] = vals[0]
					} else {
						requestBody[key] = vals
					}
				}
			}
		}

		ctx.Next()

		duration := time.Since(start)

		logEvent := logger.Info()

		statusCode := ctx.Writer.Status()

		if statusCode >= 500 {
			logEvent = logger.Error()
		} else if statusCode >= 400 {
			logEvent = logger.Warn()
		}

		logEvent.Str("method", ctx.Request.Method).
			Str("path", ctx.Request.URL.Path).
			Str("query", ctx.Request.URL.RawQuery).
			Str("ip", ctx.ClientIP()).
			Str("user_agent", ctx.Request.UserAgent()).
			Str("referer", ctx.Request.Referer()).
			Str("protocol", ctx.Request.Proto).
			Str("host", ctx.Request.Host).
			Str("remote_addr", ctx.Request.RemoteAddr).
			Str("request_uri", ctx.Request.RequestURI).
			Int64("content_length", ctx.Request.ContentLength).
			Interface("header", ctx.Request.Header).
			Int("status_code", statusCode).
			Interface("body", requestBody).
			Int64("status_code", duration.Microseconds()).Msg("HTTP Request Log")
	}
}

func formatfileSize(size int64) string {
	switch {
	case size >= 1<<20:
		return fmt.Sprintf("%.2f MB", float64(size)/(1<<20))
	case size <= 1<<20:
		return fmt.Sprintf("%.2f KB", float64(size)/(1<<10))
	default:
		return fmt.Sprintf("%d B", size)
	}
}
