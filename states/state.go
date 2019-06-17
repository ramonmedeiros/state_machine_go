package states

import (
    "fmt"
)

type ScooterState struct {
    Name string
    batteryLevel int
}

func (state *ScooterState) Next() (bool, error) {
    return false, fmt.Errorf("no next state implemented")
}

func (state *ScooterState) IsValid() (bool, error)  {
    // valid: return null error
    return true, nil

    return false, fmt.Errorf("")
}
