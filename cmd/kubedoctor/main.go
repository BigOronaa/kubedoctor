package main

import (
	"fmt"
	"os"

	"github.com/BigOronaa/kubedoctor/internal/diagnostics"
	kubeclient "github.com/BigOronaa/kubedoctor/internal/kubernetes"
)

func main() {
	client, err := kubeclient.NewClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create Kubernetes client: %v\n", err)
		os.Exit(1)
	}

	pods, err := kubeclient.ListPods(client)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to list Kubernetes Pods: %v\n", err)
		os.Exit(1)
	}

	for _, pod := range pods {
		findings := diagnostics.DiagnosePod(pod)

		for _, finding := range findings {
			fmt.Printf(
				"[%s] %s/%s: %s\n",
				finding.Severity,
				finding.Resource,
				finding.Name,
				finding.Problem,
			)
			fmt.Printf("  Explanation: %s\n", finding.Explanation)
			fmt.Printf("  Remediation: %s\n", finding.Remediation)
		}
	}
}
