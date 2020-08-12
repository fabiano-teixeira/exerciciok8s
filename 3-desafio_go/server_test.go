package main

import "testing"

func TestGreeting(t *testing.T) {
	obtido := greeting("Teste")
	esperado := "<b>Teste</b>"

	if obtido != esperado {
		t.Errorf("greeting('TesteGreeting') \n obtido: %v \n esperado:  \n%v", obtido, esperado)
	}
}

func TestGreetingCodeEducationRocks(t *testing.T) {
	obtido := greeting("CodeEducation Rocks!")
	esperado := "<b>CodeEducation Rocks!</b>"

	if obtido != esperado {
		t.Errorf("greeting('CodeEducation Rocks!') \n obtido: %v \n esperado:  \n%v", obtido, esperado)
	}
}