package states

import (
    "fmt"
    "testing"
)


func TestBasic(t *testing.T) {
    state := ScooterState{}
    ret, err := state.Next()

   fmt.Println(ret)
   if err == nil {
        t.Fatalf("State expected to not have next function %v", ret)
    }

}

