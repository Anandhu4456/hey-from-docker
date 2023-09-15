package main

import (
	"net/http"
	"text/template"
)

var tmpl *template.Template
func main(){
	tmpl =template.Must(template.ParseGlob("template/*"))
	http.HandleFunc("/",sample)
	http.ListenAndServe(":8080",nil)
}
func sample(w http.ResponseWriter, req *http.Request){
	tmpl.ExecuteTemplate(w,"first.html",nil)
}