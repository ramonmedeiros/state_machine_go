package states_test

import (
    "testing"
    "time"
    "reflect"
    "github.com/ramonmedeiros/state_machine_go/states"
    "github.com/ramonmedeiros/state_machine_go/users"
)

func TestDroppedNoUser(t* testing.T) {
    state := states.ScooterDropped{}
    state.Name = "test-name"
    state.User = nil
    state.BatteryLevel = 19
    state.LastStateChange = time.Now()

    newstate, _ := state.Next()

    if (newstate != false) {
        t.Fatalf("Expected failure, got %v", newstate)
    }
}

func TestDroppedNormalUser(t* testing.T) {
    user := users.User{}
    state := states.ScooterDropped{}
    state.Name = "test-name"
    state.User = &user
    state.BatteryLevel = 19
    state.LastStateChange = time.Now()

    newstate, _ := state.Next()

    if (newstate != false) {
        t.Fatalf("Expected failure, got %v", newstate)
    }
}

func TestDroppedToReady(t* testing.T) {
    user := users.Hunter{}
    state := states.ScooterDropped{}
    state.Name = "test-name"
    state.User = &user
    state.BatteryLevel = 19
    state.LastStateChange = time.Now()

    newstate, _ := state.Next()

    // convert interface to ScooterReady
    readyState, _ := newstate.(states.ScooterReady)

    if (reflect.TypeOf(readyState) != reflect.TypeOf(states.ScooterReady{})) {
        t.Fatalf("Expected Dropped, got %v", readyState)
    }

    if (readyState.BatteryLevel != 100) {
        t.Fatalf("Expected BaterryFull, got %v", readyState.BatteryLevel)
    }
}
