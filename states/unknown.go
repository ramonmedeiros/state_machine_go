package states

import "fmt"

type ScooterUnknown struct {
	ScooterState
}

func (state *ScooterUnknown) Next() (bool, error) {
	return false, fmt.Errorf("no next state implemented")
}

func (state *ScooterUnknown) IsValid() (bool, error) {
	return false, fmt.Errorf("Unknown state")
}
