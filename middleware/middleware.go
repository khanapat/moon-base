package middleware

import (
	"context"
	"encoding/json"
	"io"
	"math"
	"moon-base/common"
	"moon-base/response"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	requestInfoMsg    string = "Request Information"
	responseInfoMsg   string = "Response Information"
	contentType       string = "Content-Type"
	applicationJSON   string = "application/json"
	accessControl     string = "Access-Control-Allow-Origin"
	contentLength     string = "Content-Length"
	contentLengthByte int    = 512
)

var (
	tempBody  string
	tempCount int
)

type middleware struct {
	ZapLogger *zap.Logger
}

func NewMiddleware(zapLogger *zap.Logger) *middleware {
	return &middleware{
		ZapLogger: zapLogger,
	}
}

func (m *middleware) BasicAuthenication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, authOK := r.BasicAuth()
		if !authOK {
			w.Header().Set(contentType, applicationJSON)
			w.Header().Set(common.BasicAuthenticationHeader, common.BasicAuthenticationPopup)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response.ResponseContextLocale(r.Context()).AuthenBasicWeb)
			return
		}
		if username != viper.GetString("swagger.user") || password != viper.GetString("swagger.password") {
			w.Header().Set(contentType, applicationJSON)
			w.Header().Set(common.BasicAuthenticationHeader, common.BasicAuthenticationPopup)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response.ResponseContextLocale(r.Context()).AuthenBasicWeb)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m *middleware) ContextLogAndLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get(common.XRequestID) == "" {
			r.Header.Set(common.XRequestID, uuid.New().String())
		}
		logger := m.ZapLogger.With(
			zap.String(common.XRequestID, r.Header.Get(common.XRequestID)),
		)
		if r.Method == http.MethodGet {
			logger.Debug(requestInfoMsg,
				zap.String("method", r.Method),
				zap.String("host", r.Host),
				zap.String("path_uri", r.RequestURI),
				zap.String("remote_addr", r.RemoteAddr),
				// zap.Reflect("header", r.Header),
			)
		} else {
			r.Body = &HackReqBody{
				ReadCloser: r.Body,
				method:     r.Method,
				host:       r.Host,
				requestURI: r.RequestURI,
				remoteAddr: r.RemoteAddr,
				header:     r.Header,
				logger:     logger,
			}
		}
		next.ServeHTTP(&HackResBody{w, logger}, r.WithContext(context.WithValue(r.Context(), common.LoggerKey, logger)))
	})
}

func (m *middleware) ContextLocaleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), common.LocaleKey, r.URL.Query().Get(common.LocaleKey))))
	})
}

func (m *middleware) JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(accessControl, "*")
		w.Header().Set(contentType, applicationJSON)
		next.ServeHTTP(w, r)
	})
}

type HackReqBody struct {
	io.ReadCloser
	method     string
	host       string
	requestURI string
	remoteAddr string
	header     http.Header
	logger     *zap.Logger
}

func (h *HackReqBody) Read(body []byte) (int, error) {
	var n int
	var err error
	defer func() {
		if stringToInt(h.header.Get(contentLength)) > contentLengthByte {
			tempBody += string(body[:n])
			if n < (contentLengthByte * int(math.Pow(2.0, float64(tempCount)))) {
				h.logger.Debug(requestInfoMsg,
					zap.String("body", tempBody),
					zap.String("method", h.method),
					zap.String("host", h.host),
					zap.String("path_uri", h.requestURI),
					zap.String("remote_addr", h.remoteAddr),
					// zap.Reflect("header", h.header),
				)
				tempBody = ""
				tempCount = 0
			} else {
				tempCount++
			}
		} else {
			h.logger.Debug(requestInfoMsg,
				zap.String("body", string(body[:n])),
				zap.String("method", h.method),
				zap.String("host", h.host),
				zap.String("path_uri", h.requestURI),
				zap.String("remote_addr", h.remoteAddr),
				// zap.Reflect("header", h.header),
			)
			tempBody = ""
			tempCount = 0
		}
	}()
	n, err = h.ReadCloser.Read(body)
	return n, err
}

type HackResBody struct {
	http.ResponseWriter
	logger *zap.Logger
}

func (h *HackResBody) Write(b []byte) (int, error) {
	defer func() {
		h.logger.Debug(responseInfoMsg,
			zap.String("body", string(b)),
			// zap.Reflect("header", h.Header),
		)
	}()

	return h.ResponseWriter.Write(b)
}

func stringToInt(numberStr string) int {
	numberInt, err := strconv.Atoi(numberStr)
	if err != nil {
		zap.L().Error(err.Error())
	}
	return numberInt
}
