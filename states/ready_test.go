package states_test

import (
    "time"
    "testing"
    "reflect"
    "github.com/ramonmedeiros/state_machine_go/states"
)

var tnow = time.Now()

func TestUnknownState(t *testing.T) {
    twodaysago := tnow.Add(time.Hour * -48)
    state := states.ScooterReady{}
    state.Name = "test-name"
    state.User = nil
    state.BatteryLevel = 100
    state.LastStateChange = twodaysago

    newstate, _ := state.Next()

    if (reflect.TypeOf(newstate) != reflect.TypeOf(states.ScooterUnknown{})) {
        t.Fatalf("Expected unknow, found %v", reflect.TypeOf(newstate))
    }
}
