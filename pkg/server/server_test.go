package server

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// before
	res := m.Run()
	// after

	os.Exit(res)
}
