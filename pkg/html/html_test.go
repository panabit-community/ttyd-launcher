package html

import (
	"fmt"
	"testing"
)

const (
	TestTemplatePath    = "../../static/web/template"
	TestHttpTemplate    = "http.tpl"
	TestHtmlTemplate    = "html.tpl"
	TestPartialTemplate = "index.tpl"
)

func TestRender(t *testing.T) {
	type args struct {
		path  string
		tpl   string
		data  any
		files []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Case 1",
			args: args{
				path: TestTemplatePath,
				tpl:  TestHttpTemplate,
				data: struct {
					Title string
				}{
					Title: "ttyd",
				},
				files: []string{
					TestHttpTemplate,
					TestHtmlTemplate,
					TestPartialTemplate,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := Render(tt.args.path, tt.args.tpl, tt.args.data, tt.args.files...)
			fmt.Println(s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
