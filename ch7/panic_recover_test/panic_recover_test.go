package panic_recover_test

import (
	"errors"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log("recovered from", err)
		}
	}()
	t.Log("Start")
	panic(errors.New("Something wrong!"))
	// os.Exit(-1)
}
