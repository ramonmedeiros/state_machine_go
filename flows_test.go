package main_test

import (
       "github.com/ramonmedeiros/state_machine_go/users"
       "github.com/ramonmedeiros/state_machine_go/states"
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

    // create ready state and add user
    ready := states.ScooterReady{}
    user := users.User{}
    ready.BatteryLevel = 100
    ready.LastStateChange = time.Now()
    t.Log("State is", reflect.TypeOf(ready).String())

    // validate
    status, err := ready.IsValid()
    if status == false {
        t.Fatalf("%v", err)
    }

    // ready state with
    ready.User = user
    ret, err := ready.Next()
    if ret == false {
        t.Fatalf("%v", err)
    }
    riding := ret.(states.ScooterRiding)
    t.Log("State is", reflect.TypeOf(riding).String())

    // reset user and get back to ready
    riding.User = nil
    ret2, err := riding.Next()
    if ret2 == false {
        t.Fatalf("%v", err)
    }
    ready2 := ret2.(states.ScooterReady)
    t.Log("State is", reflect.TypeOf(ready2).String())

    // go to riding
    ready2.User = user
    ret3, err := ready2.Next()
    if ret3 == false {
        t.Fatalf("%v", err)
    }
    riding2 := ret3.(states.ScooterRiding)
    t.Log("State is", reflect.TypeOf(riding2).String())

    // go to battery low
    riding2.BatteryLevel = 19
    ret4, err := riding2.Next()
    if ret4 == false {
        t.Fatalf("%v", err)
    }
    batLow := ret4.(states.ScooterBatteryLow)
    t.Log("State is", reflect.TypeOf(batLow).String())

    if (reflect.TypeOf(batLow) != reflect.TypeOf(states.ScooterBatteryLow{})) {
        t.Fatalf("Expected BatteryLow, found %v", reflect.TypeOf(batLow))
    }

}

