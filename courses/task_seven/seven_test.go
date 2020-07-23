package main

import (
	"testing"
)

func TestSum(t *testing.T) {
    total := Sum(5, 5)
    if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}

func TestMove(t *testing.T) {
	type test struct {
		input Vehicle
		distance float64
        want  interface{}
	}

	autoMark := Auto{"audi", 50.0, 8.5, 75.0}
	var tr Vehicle = &autoMark

	tests := []test{
        {input: tr, distance: 100.0, want: nil},
        {input: tr, distance: 200.0, want: nil},
        {input: tr, distance: 600.0, want: VehicleError("Too small fuel!!!")},
    }

	for _, ex := range tests {
        got := ex.input.Move(ex.distance)
        if (ex.want != got) {
            t.Fatalf("expected: %v, got: %v", ex.want, got)
        }
    }
}

func TestTankUp(t *testing.T) {
	type test struct {
		input Vehicle
		quantity float64
        want  interface{}
	}

	autoMark := Auto{"audi", 50.0, 8.5, 75.0}
	var tr Vehicle = &autoMark

	tests := []test{
        {input: tr, quantity: 5.0, want: nil},
        {input: tr, quantity: 10.0, want: nil},
        {input: tr, quantity: 50.0, want: VehicleError("Too big fuel!!!")},
    }

	for _, ex := range tests {
        got := ex.input.TankUp(ex.quantity)
        if (ex.want != got) {
            t.Fatalf("expected: %v, got: %v", ex.want, got)
        }
    }
}

func TestMain(t *testing.T) {
	main()
}