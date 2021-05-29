package helpers

import (
	"net"
	"testing"
)

func TestGetLocalIP(t *testing.T) {
	ipAddress, err := GetLocalIP()

	if net.ParseIP(ipAddress) == nil {
		t.Errorf("L'adreça IP %s obtinguda no és vàlida", ipAddress)
	}

	if err != nil {
		t.Error("error al obtenir l'adreça IP d'aquest sistema", err)
	}
}
