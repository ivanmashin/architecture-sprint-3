package domain

type DeviceType string

const (
	Light  DeviceType = "light"
	Heater DeviceType = "heater"
	Gates  DeviceType = "gates"
)

type Home struct {
	ID   ID
	Name string
}

type Device struct {
	ID     ID
	Type   DeviceType
	Name   string
	Online bool
	On     bool
	HomeID ID
}

func (d *Device) UpdateName(newName *string) bool {
	if newName == nil {
		return false
	}
	d.Name = *newName
	return true
}

func (d *Device) UpdateOn(on *bool) {
	if on == nil {
		return
	}
	d.On = *on
}

func (d *Device) Toggle(on bool) {
	d.UpdateOn(&on)
}

type State struct {
	Name  string
	Value any
}
