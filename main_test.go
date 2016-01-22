package main

import (
	"testing"
)

func TestStartCar(t *testing.T) {
	c := new(Car)
	c.on = false
	c.Start()
	if !c.on {
		t.Errorf("car did not start")
	}
}

func TestStopCar(t *testing.T) {
	c := new(Car)
	c.on = true
	c.Stop()
	if c.on {
		t.Errorf("car did not stop")
	}
}
