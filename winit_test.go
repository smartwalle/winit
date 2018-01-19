package winit

import (
	"testing"
	"os"
)

var client *Winit

func TestMain(m *testing.M) {
	client = New("rebecca", "89435277FA3BA272DE795559998E-", false)
	exitCode := m.Run()
	os.Exit(exitCode)
}