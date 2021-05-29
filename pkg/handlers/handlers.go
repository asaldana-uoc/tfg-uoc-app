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
