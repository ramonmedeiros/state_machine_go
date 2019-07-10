package states

import (
    "time"
    "fmt"
    "reflect"
    "github.com/ramonmedeiros/state_machine_go/users"
)

type ScooterBounty struct {
	ScooterState
}

func (state *ScooterBounty) Next() (interface{}, error) {
    if (reflect.TypeOf(state.User) != reflect.TypeOf(&users.Hunter{})) {
        return false, fmt.Errorf("Hunter user is expected")
    }

    collected := ScooterCollected{}
    collected.Name = state.Name
    collected.User = state.User
    collected.BatteryLevel = state.BatteryLevel
    collected.LastStateChange = time.Now()

    return collected, nil
}
