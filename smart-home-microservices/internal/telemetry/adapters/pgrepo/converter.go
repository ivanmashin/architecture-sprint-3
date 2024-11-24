package pgrepo

import (
	"encoding/json"
	"slices"
	"time"

	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/telemetry/domain"
	"github.com/jackc/pgx/v5/pgtype"
)

type Converter struct{}

func (Converter) TelemetryDataToDomain(data []TelemetryDatum) *domain.Device {
	if len(data) == 0 {
		return nil
	}

	device := domain.Device{
		ID:      domain.ID(data[0].DeviceID),
		History: []domain.StatesHistory{},
	}

	lastTimestamp := data[0].Timestamp
	states := make([]domain.State, 0)
	for _, row := range data {
		if row.Timestamp != lastTimestamp {
			device.History = append(device.History, domain.StatesHistory{
				States:    slices.Clone(states),
				Timestamp: lastTimestamp.Time,
			})
			states = states[:0]
			lastTimestamp = row.Timestamp
		}

		var value any
		_ = json.Unmarshal(row.StateValue, &value)

		states = append(states, domain.State{
			Name:  row.StateName,
			Value: value,
		})
	}
	device.History = append(device.History, domain.StatesHistory{
		States:    slices.Clone(states),
		Timestamp: lastTimestamp.Time,
	})

	return &device
}

func (Converter) DomainToInsertDeviceStateParams(deviceID domain.ID, states []domain.State) InsertDeviceStateParams {
	devicesIDs := make([]string, 0, len(states))
	timestamps := make([]pgtype.Timestamp, 0, len(states))
	stateNames := make([]string, 0, len(states))
	stateValues := make([][]byte, 0, len(states))

	now := time.Now()
	for _, h := range states {
		valueData, _ := json.Marshal(h.Value)

		devicesIDs = append(devicesIDs, string(deviceID))
		timestamps = append(timestamps, pgtype.Timestamp{Time: now})
		stateNames = append(stateNames, h.Name)
		stateValues = append(stateValues, valueData)
	}

	return InsertDeviceStateParams{
		DevicesIds:  devicesIDs,
		Timestamps:  timestamps,
		StateNames:  stateNames,
		StateValues: stateValues,
	}
}
