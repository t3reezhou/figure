package api

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(res http.ResponseWriter) interface{} {
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	data := struct {
		Title string
		User  string
	}{
		Title: "golang html",
		User:  "test",
	}
	err = t.Execute(res, data)

	if err != nil {
		return err
	}
	return nil
}
