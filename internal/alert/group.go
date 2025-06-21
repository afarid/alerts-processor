package alert

func GroupByComponent(alerts []Alert) map[string][]Alert {
	grouped := make(map[string][]Alert)
	for _, alert := range alerts {
		grouped[alert.Component] = append(grouped[alert.Component], alert)
	}
	return grouped
}
