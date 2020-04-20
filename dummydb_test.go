package main

import "testing"

func TestGetEventByID(t *testing.T) {
	res, err := getEventByID("1")
	if err != nil {
		t.Errorf("Failing during execution %s", err)
	}
	if res.Title != "Festa da Uva de Vinhedo" {
		t.Error("Method is not working")
	} else {
		t.Logf("resultado é: %s", res.Title)
	}
}
