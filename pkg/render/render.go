package render

import (
	"github.com/asaldana-uoc/tfg-uoc-app/web"
	"html/template"
	"log"
	"net/http"
)

var templatesPath = "./web/html/"

// RenderTemplate Funció que s'encarrega de realitzar la renderització de la pàgina HTML a mostrar segons els paràmetres
// que rebi d'entrada (plantilla a utilitzar i conjunt de variables a tenir en compte)
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