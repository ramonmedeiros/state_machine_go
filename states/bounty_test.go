package states_test

import (
    "testing"
    "time"
    "reflect"
    "github.com/ramonmedeiros/state_machine_go/states"
    "github.com/ramonmedeiros/state_machine_go/users"
)

func TestBountyNoUser(t* testing.T) {
    state := states.ScooterBounty{}
    state.Name = "test-name"
    state.User = nil
    state.BatteryLevel = 19
    state.LastStateChange = time.Now()

    newstate, _ := state.Next()

    if (newstate != false) {
        t.Fatalf("Expected failure, got %v", newstate)
    }
}

func TestBountyNormalUser(t* testing.T) {
    user := users.User{}
    state := states.ScooterBounty{}
    state.Name = "test-name"
    state.User = &user
    state.BatteryLevel = 19
    state.LastStateChange = time.Now()

    newstate, _ := state.Next()

    if (newstate != false) {
        t.Fatalf("Expected failure, got %v", newstate)
    }
}

func TestBountyToCollected(t* testing.T) {
    user := users.Hunter{}
    state := states.ScooterBounty{}
    state.Name = "test-name"
    state.User = &user
    state.BatteryLevel = 19
    state.LastStateChange = time.Now()

    newstate, _ := state.Next()

    if (reflect.TypeOf(newstate) != reflect.TypeOf(states.ScooterCollected{})) {
        t.Fatalf("Expected Collected, got %v", newstate)
    }
}
