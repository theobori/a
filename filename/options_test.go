package filename

import (
	"testing"

	"github.com/theobori/a/internal"
)

func TestNewOptions(t *testing.T) {
	_, err := NewOptions("__", "_", internal.NewEmptyHashset[string](), false)
	if err != nil {
		t.Fatal(err)
	}

	// err should not be nil here
	_, err = NewOptions("_", "__", internal.NewEmptyHashset[string](), false)
	if err == nil {
		t.Fatal(err)
	}
}
