package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

type formData struct{
	Category Category
	Errors map[string]string
}

type Category struct {
	ID int      `db:"id"`
    Name string `db:"name" json:"name"`
}

func (h *Handler) createCategory (rw http.ResponseWriter, r *http.Request) {
	Errors := map[string]string{}
	category := Category{}
	h.loadCreatedCategoryForm(rw,category,Errors)
}

func (h *Handler) storeCategory (rw http.ResponseWriter, r *http.Request) {

		if err:=r.ParseForm(); err != nil{
			http.Error(rw, err.Error(), http.StatusInternalServerError)
	        return
		}

		category:=r.FormValue("Name")

		categories := Category{
			Name: category,
		}
		if category == ""{
   
		   Errors := map[string]string{
			   "Name":"This filed cannot be null",
		   }
		   h.loadCreatedCategoryForm(rw,categories,Errors)
			  return
		  }
   
	   if len(category) <3 {
   
		   Errors := map[string]string{
			   "Name":"This filed must be greater than or equals 3",
		   }
		   h.loadCreatedCategoryForm(rw,categories,Errors)
		   return	
	   } 
   

	const insertCategory =`INSERT INTO categories (name) VALUES ($1)`

	res :=h.db.MustExec(insertCategory,category)

	if ok, err:=res.RowsAffected(); err !=nil || ok==0{
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	
	http.Redirect(rw,r,"/", http.StatusTemporaryRedirect)
}

func (h *Handler) editCategory(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(rw, "invalid ", http.StatusTemporaryRedirect)
		return
	}
	const getCategory = `SELECT * FROM categories WHERE id = $1`
	var category Category
	h.db.Get(&category, getCategory, id )

	if category.ID == 0 {
		http.Error(rw, "invalid URL", http.StatusInternalServerError)
		return
	}

	h.loadUpdateCategoryForm(rw,category,map[string]string{})

}

func (h *Handler) updateCategory(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(rw, "invalid update", http.StatusTemporaryRedirect)
		return
	}

	const getCategory = `SELECT * FROM categories WHERE id = $1`
	var category Category
	h.db.Get(&category, getCategory, id )

	if category.ID == 0 {
		http.Error(rw, "invalid URL", http.StatusInternalServerError)
		return
	}

	if err :=r.ParseForm(); err !=nil{
		http.Error(rw, err.Error(),http.StatusInternalServerError)
		 }

		 categories := r.FormValue("Name") 

		 if categories == ""{
			Errors := map[string]string{
				"Name":"This filed cannot be null",
			}
			h.loadUpdateCategoryForm(rw,category,Errors)
			   return
		   }
	
		if len(categories) <3 {
	
			Errors := map[string]string{
				"Name":"This filed must be greater than or equals 3",
			}
			h.loadUpdateCategoryForm(rw,category,Errors)	
			return
		} 
	
		const completedCategory = `UPDATE categories SET name = $2 WHERE id = $1`
		res:= h.db.MustExec( completedCategory, id, categories)

		if ok, err:= res.RowsAffected(); err != nil || ok == 0 {
			http.Error(rw, err.Error(),http.StatusInternalServerError)
	
			return
		}

	http.Redirect(rw,r, "/", http.StatusTemporaryRedirect)
}

func (h *Handler) deleteCategory(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	if id == "" {
		http.Error(rw, "invalid update", http.StatusTemporaryRedirect)
		return
	}

	const getcategory = `SELECT * FROM categories WHERE id = $1`
	var category Category
	h.db.Get(&category, getcategory, id )

	if category.ID == 0 {
		http.Error(rw, "invalid URL", http.StatusInternalServerError)
		return
	}

	const deleteCategory =`DELETE FROM categories WHERE id =$1`

	res:= h.db.MustExec( deleteCategory, id)

	if ok, err:= res.RowsAffected(); err != nil || ok == 0 {
		http.Error(rw, err.Error(),http.StatusInternalServerError)

		return
	}


	http.Redirect(rw,r, "/", http.StatusTemporaryRedirect)
}


func (h *Handler) loadCreatedCategoryForm(rw http.ResponseWriter, categories Category, errs map[string]string){

	form:=formData{
		Category: categories,
			Errors: errs,
		}

		if err:= h.templates.ExecuteTemplate(rw,"create-category.html", form); err !=nil{
			http.Error(rw, err.Error(),http.StatusInternalServerError)
			return
		}

}

func (h *Handler) loadUpdateCategoryForm(rw http.ResponseWriter, categories Category, errs map[string]string){

	form:=formData{
			Category: categories,
			Errors: errs,
		}

		if err:= h.templates.ExecuteTemplate(rw,"edit-category.html", form); err !=nil{
			http.Error(rw, err.Error(),http.StatusInternalServerError)
			return
		}

}