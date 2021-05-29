package render

import (
	"embed"
	"github.com/asaldana-uoc/tfg-uoc-app/web"
	"html/template"
	"log"
	"net/http"
)

//go:embed html
var templatesHTML embed.FS

const templatesPath = "html/"

// RenderTemplate Funció que s'encarrega de realitzar la renderització de la pàgina HTML a mostrar segons els paràmetres
// que rebi d'entrada (plantilla a utilitzar i conjunt de variables a tenir en compte)
func RenderTemplate(w http.ResponseWriter, tmpl string, data *web.DataHTML) error {
	templateFile := templatesPath + tmpl
	parsedTemplate, err:= template.ParseFS(templatesHTML, templateFile)
	// Si no es troba la plantilla, es retorna un missatge i codi d'error HTTP/503
	if err != nil {
		http.Error(w, "Servei no disponible", http.StatusServiceUnavailable)
		log.Print(err.Error())
		return err
	}

	err2 := parsedTemplate.Execute(w, data)
	// Si no es pot processar la plantilla, es retorna un missatge i el codi d'error HTTP/503
	if err2 != nil {
		http.Error(w, "Petició no vàlida", http.StatusServiceUnavailable)
		log.Print(err.Error())
		return err
	}
	return nil
}