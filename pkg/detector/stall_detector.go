package detector

import (
	"fmt"
	"time"
)

// StalledJob represents a job that has exceeded the stall threshold
type StalledJob struct {
	JobName        string
	RunningSeconds float64
	TimeoutSeconds int
	Severity       string
	ProblemPods    []ProblemPod
}

// ProblemPod represents a pod in a problematic state
type ProblemPod struct {
	Name         string
	Reason       string
	RestartCount int
	Message      string
}

// StallDetector monitors long-running Kubernetes jobs for stalls
type StallDetector struct {
	Namespace    string
	StallTimeout time.Duration
}

// NewStallDetector creates a new StallDetector
func NewStallDetector(namespace string, timeout time.Duration) *StallDetector {
	return &StallDetector{
		Namespace:    namespace,
		StallTimeout: timeout,
	}
}

// IsStalled checks if a job has been running longer than the stall timeout
func (d *StallDetector) IsStalled(jobName string, startTime time.Time) bool {
	runningDuration := time.Since(startTime)
	return runningDuration >= d.StallTimeout
}

// AnalyzeJob checks a job and returns stall info if stalled
func (d *StallDetector) AnalyzeJob(jobName string, startTime time.Time, problemPods []ProblemPod) *StalledJob {
	if !d.IsStalled(jobName, startTime) {
		return nil
	}

	runningSeconds := time.Since(startTime).Seconds()
	severity := "WARNING"
	if len(problemPods) > 0 {
		severity = "CRITICAL"
	}

	return &StalledJob{
		JobName:        jobName,
		RunningSeconds: runningSeconds,
		TimeoutSeconds: int(d.StallTimeout.Seconds()),
		Severity:       severity,
		ProblemPods:    problemPods,
	}
}

// Report prints a stall report for a job
func (s *StalledJob) Report() {
	fmt.Printf("\n%s\n", "=======================================================")
	fmt.Printf("STALLED JOB DETECTED\n")
	fmt.Printf("%s\n", "=======================================================")
	fmt.Printf("  Job Name:        %s\n", s.JobName)
	fmt.Printf("  Running:         %.0f seconds\n", s.RunningSeconds)
	fmt.Printf("  Timeout:         %d seconds\n", s.TimeoutSeconds)
	fmt.Printf("  Severity:        %s\n", s.Severity)
	fmt.Printf("  Problem Pods:    %d\n", len(s.ProblemPods))
	for _, pod := range s.ProblemPods {
		fmt.Printf("    - %s: %s (restarts: %d)\n", pod.Name, pod.Reason, pod.RestartCount)
	}
	fmt.Printf("%s\n\n", "=======================================================")
}
