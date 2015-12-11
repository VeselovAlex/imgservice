package handlers

import "net/http"

// ImageView is a handler for image representation
type ImageView struct {
	TemplateHandlerWData
}

func (h ImageView) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.TemplateName = "image_view.html"
	// Fill h.Data
	h.TemplateHandlerWData.ServeHTTP(w, r)
}
