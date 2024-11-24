package domain

import "time"

type Device struct {
	ID      ID
	History []StatesHistory
}

func (d *Device) AddStates(states ...State) {
	if len(states) == 0 {
		return
	}
	now := time.Now()
	d.History = append(d.History, StatesHistory{
		States:    states,
		Timestamp: now,
	})
}

func (d *Device) CurrentStates() ([]State, bool) {
	if len(d.History) == 0 {
		return nil, false
	}
	return d.History[len(d.History)-1].States, true
}

type StatesHistory struct {
	States    []State
	Timestamp time.Time
}

type State struct {
	Name  string
	Value any
}
