package html

import (
	"bytes"
	"html/template"
	"path/filepath"
)

func Render(path string, tpl string, data any, files ...string) (string, error) {
	for i, file := range files {
		files[i] = filepath.Join(path, file)
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	if err := t.ExecuteTemplate(&b, tpl, data); err != nil {
		return "", err
	}
	return b.String(), nil
}
