package states

import (
    "fmt"
    "time"
    "github.com/ramonmedeiros/state_machine_go/users"
)

type ScooterState struct {
    Name string
    User *users.User
    BatteryLevel int
    LastStateChange time.Time
}

func (state *ScooterState) Next() (bool, error) {
    return false, fmt.Errorf("no next state implemented")
}

func (state *ScooterState) IsValid() (bool, error)  {
    // valid: return null error
    return true, nil

    return false, fmt.Errorf("")
}

