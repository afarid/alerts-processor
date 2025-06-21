package main

import (
	"flag"
	"fmt"
	"github.com/afarid/alerts-processor/internal/alert"
	"log"
)

func main() {
	severity := flag.String("severity", "", "Severity filter")
	service := flag.String("service", "", "Service filter")
	within := flag.Int("within", 0, "Time window in minutes")

	flag.Parse()

	data, err := alert.ParseAlerts("data/sample_alerts.json")
	if err != nil {
		log.Fatalf("Failed to parse alerts: %v", err)
	}

	filtered := alert.FilterAlerts(data.Alerts, *severity, *service, *within)
	grouped := alert.GroupByComponent(filtered)

	for component, alerts := range grouped {
		fmt.Printf("Component: %s \n", component)
		for _, a := range alerts {
			score := alert.CalculatePriority(a)
			fmt.Printf(" - ID: %s, Severity: %s, Score: %f\n", a.ID, a.Severity, score)
		}
	}
}
