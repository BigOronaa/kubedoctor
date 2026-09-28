package diagnostics

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

type Finding struct {
	Severity    string
	Resource    string
	Name        string
	Problem     string
	Explanation string
	Remediation string
}

func DiagnosePod(pod corev1.Pod) []Finding {
	var findings []Finding

	if pod.Status.Phase == corev1.PodPending {
		findings = append(findings, Finding{
			Severity:    "WARNING",
			Resource:    "Pod",
			Name:        pod.Name,
			Problem:     "Pod is Pending",
			Explanation: fmt.Sprintf("Pod has not been scheduled successfully. Reason: %s", pod.Status.Reason),
			Remediation: "Check node capacity, resource requests, node selectors, taints, tolerations, and scheduling events.",
		})
	}

	return findings
}
