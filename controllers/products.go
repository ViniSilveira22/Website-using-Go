package controllers

import (
	"Site/models"
	"html/template"
	"net/http"
	"strconv"
)

const situation int = 301

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.SelectProducts()
	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		checkErr(err)
		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		checkErr(err)
		models.InsertProduct(name, description, price, quantity)
	}
	http.Redirect(w, r, "/", situation)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := getId(r)
	models.DeleteProduct(id)
	http.Redirect(w, r, "/", situation)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := getId(r)
	product := models.EditProduct(id)
	temp.ExecuteTemplate(w, "Edit", product)
	http.Redirect(w, r, "/", situation)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, err := strconv.Atoi(r.FormValue("id"))
		checkErr(err)
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		checkErr(err)
		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		checkErr(err)
		models.UpdateProduct(id, quantity, name, description, price)
	}
	http.Redirect(w, r, "/", situation)
}

func getId(r *http.Request) string {
	return r.URL.Query().Get("id")
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
