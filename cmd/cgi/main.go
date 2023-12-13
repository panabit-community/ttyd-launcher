package main

import (
	"fmt"
	"os"

	"xie.sh.cn/panabit-ttyd/v2/pkg/html"
)

const (
	DefaultTemplatePath  = "./template"
	HttpTemplate         = "http.tpl"
	HtmlTemplate         = "html.tpl"
	PartialIndexTemplate = "index.tpl"
)

func main() {
	if err := render(); err != nil {
		os.Exit(1)
	}
}

func render() error {
	d := struct {
		Title string
	}{
		Title: "ttyd Launcher",
	}
	s, err := html.Render(
		DefaultTemplatePath,
		HttpTemplate,
		d,
		HttpTemplate, HtmlTemplate, PartialIndexTemplate,
	)
	if err != nil {
		return err
	}
	fmt.Println(s)
	return nil
}
