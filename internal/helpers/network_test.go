package helpers

import (
	"net"
	"testing"
)

// TestGetLocalIP Funció de test unitari sobre la funció GetLocalIP
func TestGetLocalIP(t *testing.T) {
	ipAddress, err := GetLocalIP()

	// Si el format d'adreça IP obtingut no és correcte, es llança un error
	if net.ParseIP(ipAddress) == nil {
		t.Errorf("L'adreça IP %s obtinguda no és vàlida", ipAddress)
	}

	// Si a l'hora d'obtenir l'adreça IP trobem un error inesperat, es llança un error
	if err != nil {
		t.Error("error al obtenir l'adreça IP d'aquest sistema", err)
	}
}
