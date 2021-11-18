package handler

import (
	"net/http"
)

type IndexCategory struct{
   Category []Category
}

func (h *Handler) Home (rw http.ResponseWriter, r *http.Request) {
	categories := []Category{}

    h.db.Select(&categories, "SELECT * FROM categories")
	lt := IndexCategory{
		Category:categories,
	}
	if err:= h.templates.ExecuteTemplate(rw,"index-category.html", lt); err !=nil{
		http.Error(rw, err.Error(),http.StatusInternalServerError)
		return
	}
}
