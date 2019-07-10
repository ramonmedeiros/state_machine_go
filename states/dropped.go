package states

import (
    "time"
    "fmt"
    "reflect"
    "github.com/ramonmedeiros/state_machine_go/users"
)

type ScooterDropped struct {
	ScooterState
}

func (state *ScooterDropped) Next() (interface{}, error) {
    if (reflect.TypeOf(state.User) != reflect.TypeOf(&users.Hunter{})) {
        return false, fmt.Errorf("Hunter user is expected")
    }

    dropped := ScooterReady{}
    dropped.Name = state.Name
    dropped.User = nil
    dropped.BatteryLevel = 100
    dropped.LastStateChange = time.Now()

    return dropped, nil
}
