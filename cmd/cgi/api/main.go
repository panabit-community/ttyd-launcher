package main

import (
	"os"

	"xie.sh.cn/panabit-ttyd/v2/pkg/ttyd"
)

const (
	CgiPrefix             = "CGI_"
	RequestMethodKey      = "REQUEST_METHOD"
	RequestMethodFallback = "GET"
)

func main() {
	ttyd.Run() // TODO
}

func parseMethod() string {
	m := os.Getenv(RequestMethodKey)
	if m == "" {
		return RequestMethodFallback
	}
	return m
}

func parseQuery(key string) string {
	v := os.Getenv(CgiPrefix + key)
	return v
}
