package rlpString

import (
	"testing"
)

func TestNewRlpString(t *testing.T) {
	var value = "hello world"
	rString := CreateRlpString(value)

	if rString.AsString() != "hello world" {
		t.Errorf("Timesmape incorrect, got: %s, want: %s.", rString.AsString(), "hello world")
	}
}
