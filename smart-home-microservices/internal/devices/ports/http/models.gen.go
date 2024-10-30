// Package http provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package http

const (
	Api_keyScopes = "api_key.Scopes"
)

// Defines values for DeviceType.
const (
	Gates  DeviceType = "gates"
	Heater DeviceType = "heater"
	Light  DeviceType = "light"
)

// Device defines model for Device.
type Device struct {
	Id     string     `json:"id"`
	Name   string     `json:"name"`
	On     bool       `json:"on"`
	Online bool       `json:"online"`
	States *[]string  `json:"states,omitempty"`
	Type   DeviceType `json:"type"`
}

// DeviceType defines model for Device.Type.
type DeviceType string

// Home defines model for Home.
type Home struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// Scenario defines model for Scenario.
type Scenario struct {
	Actions   *[]State `json:"actions,omitempty"`
	Activated *bool    `json:"activated,omitempty"`
	Id        *string  `json:"id,omitempty"`
	Name      *string  `json:"name,omitempty"`
	Trigger   *State   `json:"trigger,omitempty"`
}

// State defines model for State.
type State struct {
	DeviceId string `json:"device_id"`
	Name     string `json:"name"`
	Value    string `json:"value"`
}

// DeviceUpdate defines model for DeviceUpdate.
type DeviceUpdate struct {
	Name *string `json:"name,omitempty"`
	On   *bool   `json:"on,omitempty"`
}

// States defines model for States.
type States = []State

// SetDeviceStatesJSONBody defines parameters for SetDeviceStates.
type SetDeviceStatesJSONBody = []State

// UpdateDeviceJSONBody defines parameters for UpdateDevice.
type UpdateDeviceJSONBody struct {
	Name *string `json:"name,omitempty"`
	On   *bool   `json:"on,omitempty"`
}

// ToggleDeviceJSONBody defines parameters for ToggleDevice.
type ToggleDeviceJSONBody = bool

// ActivateScenarioJSONBody defines parameters for ActivateScenario.
type ActivateScenarioJSONBody = bool

// SetDeviceStatesJSONRequestBody defines body for SetDeviceStates for application/json ContentType.
type SetDeviceStatesJSONRequestBody = SetDeviceStatesJSONBody

// CreateDeviceJSONRequestBody defines body for CreateDevice for application/json ContentType.
type CreateDeviceJSONRequestBody = Device

// UpdateDeviceJSONRequestBody defines body for UpdateDevice for application/json ContentType.
type UpdateDeviceJSONRequestBody UpdateDeviceJSONBody

// ToggleDeviceJSONRequestBody defines body for ToggleDevice for application/json ContentType.
type ToggleDeviceJSONRequestBody = ToggleDeviceJSONBody

// CreateHomeJSONRequestBody defines body for CreateHome for application/json ContentType.
type CreateHomeJSONRequestBody = Home

// CreateScenarioJSONRequestBody defines body for CreateScenario for application/json ContentType.
type CreateScenarioJSONRequestBody = Scenario

// ActivateScenarioJSONRequestBody defines body for ActivateScenario for application/json ContentType.
type ActivateScenarioJSONRequestBody = ActivateScenarioJSONBody
