package main

import (
	"os"
	"path/filepath"
	"text/template"
)

const (
	DefaultTemplatePath  = "./template"
	HttpTemplate         = "http.tpl"
	HtmlTemplate         = "html.tpl"
	PartialIndexTemplate = "index.tpl"
)

func main() {
	loc := "http://192.168.0.200:7681"
	b := struct {
		Title          string
		WindowLocation string
	}{
		Title:          "ttyd Launcher",
		WindowLocation: loc, // TODO
	}
	if err := render(DefaultTemplatePath, b); err != nil {
		os.Exit(1)
	}
}

func render(path string, data any) error {
	tpl, err := template.ParseFiles(
		filepath.Join(path, HttpTemplate),
		filepath.Join(path, HtmlTemplate),
		filepath.Join(path, PartialIndexTemplate),
	)
	if err != nil {
		return err
	}
	if err := tpl.ExecuteTemplate(os.Stdout, HttpTemplate, data); err != nil {
		return err
	}
	return nil
}
