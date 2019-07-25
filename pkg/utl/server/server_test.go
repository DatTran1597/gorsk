package server_test

import (
	"testing"

	"github.com/DatTran1597/gorsk/pkg/utl/server"
)

// Improve tests
func TestNew(t *testing.T) {
	e := server.New()
	if e == nil {
		t.Errorf("Server should not be nil")
	}
}
