package alert

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)
func TestFilterAlerts(t *testing.T) {
	alerts := []Alert{
		{ID: "1", Timestamp: time.Now().Add(-time.Minute * 5), Service: "service1", Component: "component1", Severity: "critical", Metric: "cpu", Value: 90, Threshold: 80, Description: "High CPU usage"},
		{ID: "2", Timestamp: time.Now().Add(-time.Minute * 10), Service: "service2", Component: "component2", Severity: "warning", Metric: "memory", Value: 70, Threshold: 60, Description: "High memory usage"},
		{ID: "3", Timestamp: time.Now().Add(-time.Hour), Service: "service1", Component: "component1", Severity: "info", Metric: "disk", Value: 30, Threshold: 50, Description: "Disk usage normal"},
	}

	filtered := FilterAlerts(alerts, "", "", 0)
	assert.Equal(t, len(alerts), len(filtered))

	filtered = FilterAlerts(alerts, "critical", "", 0)
	assert.Equal(t, 1, len(filtered))
	assert.Equal(t, alerts[0].ID, filtered[0].ID)

	filtered = FilterAlerts(alerts, "", "service1", 0)
	assert.Equal(t, 2, len(filtered))

	filtered = FilterAlerts(alerts, "", "", 11)
	assert.Equal(t, 2, len(filtered))
}