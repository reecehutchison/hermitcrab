package http

import "testing"

func TestSayHi(t *testing.T) {
	if SayHi() != "Hi" {
		t.Errorf("got %s, want Hi", SayHi())
	}
}