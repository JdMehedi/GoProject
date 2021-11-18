package handler

import (
	"text/template"

	"github.com/jmoiron/sqlx"
)
type Handler struct{
	templates *template.Template
	db *sqlx.DB
}

func New(db *sqlx.DB) *Handler{
	h:= &Handler{
		db: db,
   }
   h.parseTemplate()
   return h
}

func (h *Handler) parseTemplate(){
h.templates = template.Must(template.ParseFiles(
	"templates/create-category.html",
	"templates/index-category.html",
	"templates/edit-category.html",
))
}

