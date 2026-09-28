package kubernetes

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
)

func ListPods(client *clientset.Clientset) ([]corev1.Pod, error) {
	pods, err := client.CoreV1().Pods("").List(
		context.Background(),
		metav1.ListOptions{},
	)
	if err != nil {
		return nil, err
	}

	return pods.Items, nil
}
