package main

import "os"

const (
	CgiPrefix             = "CGI_"
	RequestMethodKey      = "REQUEST_METHOD"
	RequestMethodFallback = "GET"
)

func main() {

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
