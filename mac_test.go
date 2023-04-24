package bluetooth

import (
	"testing"
)

func TestParseMAC(t *testing.T) {
	_, e := ParseMAC("11:22:33:AA:BB:CC")
	if e != nil {
		t.Errorf("expected nil but got %v", e)
	}

	_, e = ParseMAC("11:22:33:AA:BB:C")
	if e != errInvalidMAC {
		t.Errorf("expected ErrInvalidMAC but got %v", e)
	}

	_, e = ParseMAC("")
	if e != errInvalidMAC {
		t.Errorf("expected ErrInvalidMAC but got %v", e)
	}
}
