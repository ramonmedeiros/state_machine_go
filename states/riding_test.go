package states_test


import (
    "testing"
    "time"
    "reflect"
    "github.com/ramonmedeiros/state_machine_go/states"
    "github.com/ramonmedeiros/state_machine_go/users"
)

func TestKeepRiding(t* testing.T) {
    user := users.User{}
    state := states.ScooterReady{}
    state.Name = "test-name"
    state.User = &user
    state.BatteryLevel = 100
    state.LastStateChange = time.Now()

    newstate, _ := state.Next()

    if (reflect.TypeOf(newstate) != reflect.TypeOf(states.ScooterRiding{})) {
        t.Fatalf("Expected Riding, found %v", reflect.TypeOf(newstate))
    }


}
