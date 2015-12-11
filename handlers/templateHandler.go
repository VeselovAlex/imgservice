package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

// TemplateDir is a directory, where handlers should search templates
const TemplateDir = "templates"

// TemplateHandler is a basic type for handler using template
type TemplateHandler struct {
	// Name of template file
	TemplateName string
	once         sync.Once
	// Template, loaded from file
	Template *template.Template
}

// TemplateHandlerWData extends TemplateHandler providing
// data, passed into template
type TemplateHandlerWData struct {
	TemplateHandler
	// Data passed to template
	Data interface{}
}

// ServeHTTP method to implement http.Handler interface
func (h TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.ReadRequest) {
	h.once.Do(func() {
		templatePath := filepath.Join(TemplateDir, h.TemplateName)
		h.Template = template.Must(template.ParseFiles(templatePath))
	})
	h.Template.Execute(w, nil)
}

// ServeHTTP method to implement http.Handler interface
func (h TemplateHandlerWData) ServeHTTP(w http.ResponseWriter, r *http.ReadRequest) {
	h.once.Do(func() {
		templatePath := filepath.Join(TemplateDir, h.TemplateName)
		h.Template = template.Must(template.ParseFiles(templatePath))
	})
	h.Template.Execute(w, h.Data)
}
