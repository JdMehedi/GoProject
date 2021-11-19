package handler

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)
type Handler struct{
	templates *template.Template
	db *sqlx.DB
}

func New(db *sqlx.DB) *mux.Router{
	h:= &Handler{
		db: db,
   }
   h.parseTemplate()
   r :=mux.NewRouter() 

	r.HandleFunc("/", h.Home)
	r.HandleFunc("/categories/create", h.createCategory)
	r.HandleFunc("/categories/store", h.storeCategory)
	r.HandleFunc("/categories/{id}/edit", h.editCategory)
	r.HandleFunc("/categories/{id}/update", h.updateCategory)
	r.HandleFunc("/categories/{id}/delete", h.deleteCategory)

	r.NotFoundHandler = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err:= h.templates.ExecuteTemplate(rw,"404.html", nil); err !=nil{
			http.Error(rw, err.Error(),http.StatusInternalServerError)
			return
		}
	})
	return r
}

func (h *Handler) parseTemplate(){
h.templates = template.Must(template.ParseFiles(
	"templates/create-category.html",
	"templates/index-category.html",
	"templates/edit-category.html",
	"templates/404.html",
))
}

