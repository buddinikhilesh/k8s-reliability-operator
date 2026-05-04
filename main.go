package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// Config holds operator configuration
type Config struct {
	Namespace      string
	StallTimeout   time.Duration
	MaxRestarts    int
	SLOTarget      float64
	ChaosEnabled   bool
	MetricsPort    int
	ResyncInterval time.Duration
}

func main() {
	cfg := Config{}

	flag.StringVar(&cfg.Namespace, "namespace", "default", "Kubernetes namespace to watch")
	flag.DurationVar(&cfg.StallTimeout, "stall-timeout", 5*time.Minute, "Duration before a job is considered stalled")
	flag.IntVar(&cfg.MaxRestarts, "max-restarts", 3, "Maximum auto-restarts before escalating to on-call")
	flag.Float64Var(&cfg.SLOTarget, "slo-target", 99.9, "SLO target percentage for job completion")
	flag.BoolVar(&cfg.ChaosEnabled, "chaos", false, "Enable chaos engineering mode")
	flag.IntVar(&cfg.MetricsPort, "metrics-port", 8080, "Port to expose Prometheus metrics")
	flag.DurationVar(&cfg.ResyncInterval, "resync-interval", 30*time.Second, "Reconcile loop interval")
	flag.Parse()

	fmt.Printf("Starting k8s-reliability-operator\n")
	fmt.Printf("  Namespace:      %s\n", cfg.Namespace)
	fmt.Printf("  Stall Timeout:  %s\n", cfg.StallTimeout)
	fmt.Printf("  SLO Target:     %.2f%%\n", cfg.SLOTarget)
	fmt.Printf("  Chaos Enabled:  %v\n", cfg.ChaosEnabled)
	fmt.Printf("  Metrics Port:   %d\n", cfg.MetricsPort)

	// Start controller loop
	controller := NewWorkloadController(cfg)
	if err := controller.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Controller failed: %v\n", err)
		os.Exit(1)
	}
}
