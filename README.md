# k8s-reliability-operator

Golang Kubernetes operator for automated reliability enforcement.
Monitors distributed workloads, detects stalled jobs, and triggers
fault-recovery actions automatically.

## What this solves
- Detects stalled long-running distributed jobs before they breach SLOs
- Auto-restarts failed workloads up to a configurable limit
- Escalates to on-call when auto-recovery is exhausted
- Runs controlled Chaos Engineering experiments safely
- Tracks DORA metrics across all managed workloads

## Architecture
