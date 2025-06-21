package alert

import "time"

func FilterAlerts(alerts []Alert, severity, service string, withinMinutes int) []Alert {
	var filtered []Alert
	now := time.Now().UTC()

	for _, alert := range alerts {
		if (severity == "" || alert.Severity == severity) &&
			(service == "" || alert.Service == service) &&
			(withinMinutes == 0 || now.Sub(alert.Timestamp).Minutes() <= float64(withinMinutes)) {
			filtered = append(filtered, alert)
		}
	}

	return filtered
}
