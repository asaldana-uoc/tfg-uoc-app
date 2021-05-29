package render

import (
	"github.com/asaldana-uoc/tfg-uoc-app/web"
	"html/template"
	"log"
	"net/http"
)

const templatesPath = "./web/html/"

func RenderTemplate(w http.ResponseWriter, tmpl string, data *web.HTMLData) error {
	templateFile := templatesPath + tmpl
	parsedTemplate, err:= template.ParseFiles(templateFile)
	if err != nil {
		http.Error(w, "Servei no disponible", http.StatusServiceUnavailable)
		log.Print(err.Error())
		return err
	}

	err2 := parsedTemplate.Execute(w, data)
	if err2 != nil {
		http.Error(w, "Petició no vàlida", http.StatusServiceUnavailable)
		log.Print(err.Error())
		return err
	}
	return nil
}