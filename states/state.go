package states

import (
    "fmt"
    "time"
)

type ScooterState struct {
    Name string
    User interface{}
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

