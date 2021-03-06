package main

import (
	"context"
	"fmt"
	"net/http"
)

const requestIDKey int = 0

func WithRequestID(next http.Handler) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		// 从 header 中提取 request-id
		reqID := req.Header.Get("X-Request-ID")
		// 创建 valueCtx。使用自定义的类型，不容易冲突
		ctx := context.WithValue(req.Context(), requestIDKey, reqID)
		// 创建新的请求
		req = req.WithContext(ctx)
		// 调用 HTTP 处理函数
		next.ServeHTTP(rw, req)
	}
}

// 获取 request-id
func GetRequestID(ctx context.Context) string {
	requestIDKey, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		return ""
	} else {
		return requestIDKey
	}
}

func Handle(rw http.ResponseWriter, req *http.Request) {
	// 拿到 reqId，后面可以记录日志等等
	reqID := GetRequestID(req.Context())
	fmt.Println("reqID", reqID)
}

//func main() {
//	handler := WithRequestID(http.HandlerFunc(Handle))
//	http.HandleFunc("/test", handler)
//	_ = http.ListenAndServe(":8080", nil)
//}

func main() {
	handler := WithRequestID(http.HandlerFunc(Handle))
	// 新建多路复用器
	mux := http.NewServeMux()
	mux.HandleFunc("/test", handler)
	// 创建路由
	_ = http.ListenAndServe(":8080", mux)
}
