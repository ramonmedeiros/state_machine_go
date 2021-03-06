package main_test

import (
	"github.com/ramonmedeiros/state_machine_go/states"
	"github.com/ramonmedeiros/state_machine_go/users"
	"reflect"
	"testing"
	"time"
)

func mockTimeTo21() {
	// mock time to 21:30
	states.Now = func() time.Time {
		return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 21, 00, 1, 0, time.Now().Location())
	}
}

func TestNormalUserUntilBatteryLow(t *testing.T) {
	// mock time to avoid problems with bounty time
	mockTimeTo21()

	// vars to receive content
	var ret interface{}
	var status bool
	var err error
	var ready states.ScooterReady
	var riding states.ScooterRiding

	// create ready state and add user
	ready = states.ScooterReady{}
	user := users.User{}
	ready.BatteryLevel = 100
	ready.LastStateChange = time.Now()
	t.Log("State is", reflect.TypeOf(ready).String())

	// validate ready
	status, err = ready.IsValid()
	if status == false {
		t.Fatalf("%v", err)
	}

	// ready state with
	ready.User = user
	ret, err = ready.Next()
	if ret == false {
		t.Fatalf("%v", err)
	}
	riding = ret.(states.ScooterRiding)
	t.Log("State is", reflect.TypeOf(riding).String())

	// validate riding
	status, err = riding.IsValid()
	if status == false {
		t.Fatalf("%v", err)
	}

	// reset user and get back to ready
	riding.User = nil
	ret, err = riding.Next()
	if ret == false {
		t.Fatalf("%v", err)
	}
	ready = ret.(states.ScooterReady)
	t.Log("State is", reflect.TypeOf(ready).String())

	// validate ready2
	status, err = ready.IsValid()
	if status == false {
		t.Fatalf("%v", err)
	}

	// go to riding
	ready.User = user
	ret, err = ready.Next()
	if ret == false {
		t.Fatalf("%v", err)
	}
	riding = ret.(states.ScooterRiding)
	t.Log("State is", reflect.TypeOf(riding).String())

	// validate riding2
	status, err = riding.IsValid()
	if status == false {
		t.Fatalf("%v", err)
	}

	// go to battery low
	riding.BatteryLevel = 19
	ret, err = riding.Next()
	if ret == false {
		t.Fatalf("%v", err)
	}
	batLow := ret.(states.ScooterBatteryLow)
	batLow.User = nil
	t.Log("State is", reflect.TypeOf(batLow).String())

	// validate batLow
	status, err = batLow.IsValid()
	if status == false {
		t.Fatalf("%v", err)
	}

	if (reflect.TypeOf(batLow) != reflect.TypeOf(states.ScooterBatteryLow{})) {
		t.Fatalf("Expected BatteryLow, found %v", reflect.TypeOf(batLow))
	}

}

func TestBatteryLowToCollected(t *testing.T) {
	// vars to receive content
	var ret interface{}
	var status bool
	var err error

	// create ready state and add user
	batLow := states.ScooterBatteryLow{}
	batLow.BatteryLevel = 19
	batLow.LastStateChange = time.Now()
	t.Log("State is", reflect.TypeOf(batLow).String())

	// validate batlow
	status, err = batLow.IsValid()
	if status == false {
		t.Fatalf("%v", err)
	}

	// next state of BatteryLow
	ret, err = batLow.Next()
	if ret == false {
		t.Fatalf("%v", err)
	}
	bounty := ret.(states.ScooterBounty)
	t.Log("State is", reflect.TypeOf(bounty).String())

	// set Hunter user and validate Bounty
	bounty.User = users.Hunter{}
	status, err = bounty.IsValid()
	if status == false {
		t.Fatalf("%v", err)
	}

	// next state of bounty
	ret, err = bounty.Next()
	if ret == false {
		t.Fatalf("%v", err)
	}
	collected := ret.(states.ScooterCollected)
	t.Log("State is", reflect.TypeOf(collected).String())

	// validate Collected
	status, err = collected.IsValid()
	if status == false {
		t.Fatalf("%v", err)
	}

	// next state of Collected
	ret, err = collected.Next()
	if ret == false {
		t.Fatalf("%v", err)
	}
	dropped := ret.(states.ScooterDropped)
	t.Log("State is", reflect.TypeOf(dropped).String())

	// validate dropped
	status, err = dropped.IsValid()
	if status == false {
		t.Fatalf("%v", err)
	}

	// next state of dropped
	ret, err = dropped.Next()
	if ret == false {
		t.Fatalf("%v", err)
	}
	ready := ret.(states.ScooterReady)
	t.Log("State is", reflect.TypeOf(ready).String())
}
