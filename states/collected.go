package states

import (
	"fmt"
	"github.com/ramonmedeiros/state_machine_go/users"
	"reflect"
	"time"
)

type ScooterCollected struct {
	ScooterState
}

func (state *ScooterCollected) Next() (interface{}, error) {
	if (reflect.TypeOf(state.User) != reflect.TypeOf(&users.Hunter{})) {
		return false, fmt.Errorf("Hunter user is expected")
	}

	dropped := ScooterDropped{}
	dropped.Name = state.Name
	dropped.User = state.User
	dropped.BatteryLevel = state.BatteryLevel
	dropped.LastStateChange = time.Now()

	return dropped, nil
}
