package main

import "testing"

func TestAddHtml(t *testing.T) {
	s := AddHtml("a")
	if s != "<html><body>a</body></html>" {
		t.Errorf("AddHtml(%s) == %s, want %s", "a", s, "<html><body>a</body></html>")
	}
}
