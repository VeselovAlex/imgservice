package handlers

import "net/http"

// Home is a homepage handler
type Home struct {
	TemplateHandler
}

func (h Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.TemplateName = "home.html"
	h.TemplateHandler.ServeHTTP(w, r)
}
