package alert

import (
	"encoding/json"
	"os"
	"time"
)

type Alert struct {
	ID          string    `json:"id"`
	Timestamp   time.Time `json:"timestamp"`
	Service     string    `json:"service"`
	Component   string    `json:"component"`
	Severity    string    `json:"severity"`
	Metric      string    `json:"metric"`
	Value       float64   `json:"value"`
	Threshold   float64   `json:"threshold"`
	Description string    `json:"description"`
}

type AlertData struct {
	Alerts []Alert `json:"alerts"`
}

func ParseAlerts(filename string) (*AlertData, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data AlertData
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
