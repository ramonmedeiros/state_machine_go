package states_test

import (
    "fmt"
    "testing"
    "ramonmedeiros/state_machine_go/states"
)


func TestBasic(t *testing.T) {
    state := states.ScooterState{}
    ret, err := state.Next()

   fmt.Println(ret)
   if err == nil {
        t.Fatalf("State expected to not have next function %v", ret)
    }

}

