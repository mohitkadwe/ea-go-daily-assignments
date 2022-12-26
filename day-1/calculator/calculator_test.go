package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	cal := calculator{1, 2}
	exp := 3
	res := cal.add()
	if res != exp {
		t.Errorf("%d was expect but got %d .\n", exp, res)
	}
}

func TestSubtract(t *testing.T) {
	cal := calculator{2, 1}
	exp := 1
	res := cal.sub()
	if res != exp {
		t.Errorf("%d was expect but got %d .\n", exp, res)
	}
}

func TestMultiply(t *testing.T) {
	cal := calculator{2, 2}
	exp := 4
	res := cal.mul()
	if res != exp {
		t.Errorf("%d was expect but got %d.\n", exp, res)
	}
}

func TestDivide(t *testing.T) {
	cal := calculator{2, 2}
	exp := 1
	res := cal.div()
	if res != exp {
		t.Errorf("%d was expect but got %d.\n", exp, res)
	}
}

func TestTrigonometricSin(t *testing.T) {
	cal := calculator{1, 0}
	exp := 0.8414709848078965
	res := cal.trigonometricSin()
	if res != exp {
		t.Errorf("%f was expect but got %f.\n", exp, res)
	}
}

func TestTrigonometricCos(t *testing.T) {
	cal := calculator{1, 0}
	exp := 0.5403023058681398
	res := cal.trigonometricCos()
	if res != exp {
		t.Errorf("%f was expect but got %f.\n", exp, res)
	}
}

func TestTrigonometricTan(t *testing.T) {
	cal := calculator{1, 0}
	exp := 1.557407724654902
	res := cal.trigonometricTan()
	if res != exp {
		t.Errorf("%f was expect but got %f.\n", exp, res)
	}
}

func TestSquareRoot(t *testing.T) {
	cal := calculator{4, 0}
	exp := 2.0
	res := cal.squareRoot()
	if res != exp {
		t.Errorf("%f was expect but got %f.\n", exp, res)
	}
}
