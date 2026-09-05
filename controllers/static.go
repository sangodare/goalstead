package controllers

import "net/http"



func StaticHandler(tpl Template) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, nil)


	})

}




