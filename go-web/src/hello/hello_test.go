package main

import "testing"

func TestHello(t *testing.T) {

	got  := greeting("Code.education Rocks!");
	want := "<b>Code.education Rocks!</b>";
	
	if got != want {
		t.Errorf("resultado da função incorreto %s, resultado correto: %s", got, want)
	}
}