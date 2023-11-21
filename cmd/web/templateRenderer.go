package main

import (
	"io"
	"io/fs"
	"log"
	"path/filepath"
	"text/template"

	"github.com/labstack/echo/v4"
)

var funcMap = template.FuncMap{
	"numList":   numList,
	"dateStamp": getDate,
}

type TemplateRenderer struct {
	templates *template.Template
}

func NewEmbedTemplateRenderer(fsys fs.FS) *TemplateRenderer {
	templates, err := parseEmbeddedTemplates(fsys, ".")
	if err != nil {
		log.Fatal(err)
	}
	return &TemplateRenderer{
		templates: templates,
	}
}

func (t *TemplateRenderer) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func parseEmbeddedTemplates(fsys fs.FS, path string) (*template.Template, error) {
	var tmpl *template.Template
	err := fs.WalkDir(fsys, path, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			b, err := fs.ReadFile(fsys, path)
			if err != nil {
				return err
			}
			name := filepath.ToSlash(path)
			if tmpl == nil {
				tmpl = template.New(name).Funcs(funcMap)
				_, err = tmpl.Parse(string(b))
			} else {
				_, err = tmpl.New(name).Funcs(funcMap).Parse(string(b))
			}
			if err != nil {
				return err
			}
		}
		return nil
	})
	return tmpl, err
}
