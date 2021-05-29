package handlers

import (
	"fmt"
	"github.com/asaldana-uoc/tfg-uoc-app/pkg/render"
	"github.com/asaldana-uoc/tfg-uoc-app/web"
	"log"
	"net"
	"net/http"
	"os"
)

// Home Handler que s'encarrega d'implementar la lògica de la pàgina principal on es mostra una pàgina HTML
// amb la informació del TFG i el hostname i l'adreça IP del sistema. Es limita les possibles URIs a gestionar amb
// la condició inicial que analitza la URL de la petició. Un cop validat, es renderitza el contingut de la pàgina HTML
// passant les variables hostname i ipAddress.
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Printf("Petició rebuda a la URL %s però no es troba disponible", r.URL.Path)
		http.Error(w, "Pàgina no trobada.", http.StatusNotFound)
		return
	}
	log.Printf("Petició rebuda a la URL %s", r.URL.Path)

	stringMap := make(map[string]string)
	hostname, _ := os.Hostname()
	stringMap["hostname"] = hostname
	stringMap["ipAddress"] = getLocalIP()

	render.RenderTemplate(w, "home.html.tmpl", &web.HTMLData{StringMap: stringMap})
}


// Status Handler que s'encarrega d'implementar la lògica per obtenir l'estat del servei per a retornar-lo
//al balancejador o al element que comprovi el health check
func Status(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/estat" {
		log.Printf("Petició rebuda a la URL %s però no es troba disponible", r.URL.Path)
		http.Error(w, "Pàgina no trobada.", http.StatusNotFound)
		return
	}
	//http.(w, "Service Unavailable", http.StatusServiceUnavailable)
	log.Printf("Petició rebuda a la URL %s", r.URL.Path)
	fmt.Fprintf(w, "Health check OK")

}

/* Funció auxiliar que retorna l'adreça IP de l'equip on s'està executant el codi */
func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Print("Error al obtenir les adreces IPs locals", err.Error())
	}

	for _, a := range addrs {
		if ip, ok := a.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				ipAddress := ip.IP.String()
				log.Printf("Indentificada l'adreça IP %s en el host local", ipAddress)
				return ipAddress
			}
		}
	}
	return ""
}
