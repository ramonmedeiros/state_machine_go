package states_test

import (
    "testing"
    "github.com/ramonmedeiros/state_machine_go/states"
)


func TestBasic(t *testing.T) {
    state := states.ScooterState{}
    ret, err := state.Next()

   if err == nil {
        t.Fatalf("State expected to not have next function %v", ret)
    }

}

