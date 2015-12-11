package handler

import (
	"html/template"
	"sync"
)

// TemplateHandler is a basic type for handler using template
type TemplateHandler struct {
	TemplateName string
	once         sync.Once
	Template     *template.Template
}

// TemplateHandlerWData extends TemplateHandler providing
// data, passed into template
type TemplateHandlerWData struct {
	TemplateHandler
	Data interface{}
}
