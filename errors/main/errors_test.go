package main

import "testing"

func TestGetEndpoint(t *testing.T) {
	if err := GetEndpoint(); err != nil {
		t.Error("AAAAAAAA", err)
	}
}
