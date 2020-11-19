package ut

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{1, 4, 9}
	for idx, i := range input {
		got := square(i)
		want := expected[idx]
		if got != want {
			t.Errorf("input: %d, expected: %d, but got %d\n", i, want, got)
		}
	}
}

func TestSquareAssert(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{1, 4, 9}
	for idx, i := range input {
		got := square(i)
		want := expected[idx]
		assert.Equal(t, got, want, "input: %d, expected: %d, but got: %d", i, got, want)
	}
}

func TestErrorInCode(t *testing.T) {
	t.Log("Start")
	t.Error("Error")
	t.Log("End")
}

func TestFatalInCode(t *testing.T) {
	t.Log("Start")
	t.Fatal("Fatal")
	t.Log("End")
}
