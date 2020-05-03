package somewhere

import (
	"strings"
	"testing"
)

func TestSomething(t *testing.T) {
	want := "Somewhere."
	if got := Something(); strings.Contains(got, want) {
		t.Errorf("Failed Test = %q, want %q", got, want)
	} else {
		t.Log("Passed Test = , want ", got, want)
	}
}
