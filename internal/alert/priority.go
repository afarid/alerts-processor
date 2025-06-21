package alert

func CalculatePriority(alert Alert) float64 {
	var score float64
	switch alert.Severity {
	case "critical":
		score = 10
	case "warning":
		score = 5
	case "info":
		score = 1
	}
	return (alert.Value / alert.Threshold) * score
}
