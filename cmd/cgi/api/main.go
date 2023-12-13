package main

import (
	"fmt"
	"os"

	"xie.sh.cn/panabit-ttyd/v2/pkg/html"
	"xie.sh.cn/panabit-ttyd/v2/pkg/ttyd"
)

const (
	CgiPrefix             = "CGI_"
	RequestMethodKey      = "REQUEST_METHOD"
	RequestMethodFallback = "GET"

	DefaultTemplatePath     = "./template"
	HttpTemplate            = "http.tpl"
	HtmlTemplate            = "html.tpl"
	PartialRedirectTemplate = "redirect.tpl"
)

func main() {
	// TODO: use daemon
	if _, err := ttyd.Start(); err != nil {
		os.Exit(1)
	}
	if err := render(); err != nil {
		os.Exit(1)
	}
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

func render() error {
	d := struct {
		Title    string
		Location string
	}{
		Title:    "Redirecting...",
		Location: "http://192.168.0.200:7681", // TODO: getenv
	}
	s, err := html.Render(
		DefaultTemplatePath,
		HttpTemplate,
		d,
		HttpTemplate, HtmlTemplate, PartialRedirectTemplate,
	)
	if err != nil {
		return err
	}
	fmt.Println(s)
	return nil
}
