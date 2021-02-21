package kubeclient

import (
	"context"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func DeleteServer(ctx context.Context, name string) error {
	tmp := ""
	if err := clientset.NetworkingV1beta1().Ingresses(defaultNamespace).Delete(ctx, name, v1.DeleteOptions{}); err != nil {
		tmp += err.Error()
		tmp += ";"
	}
	if err := clientset.CoreV1().Services(defaultNamespace).Delete(ctx, name, v1.DeleteOptions{}); err != nil {
		tmp += err.Error()
		tmp += ";"
	}
	if err := clientset.AppsV1().Deployments(defaultNamespace).Delete(ctx, name, v1.DeleteOptions{}); err != nil {
		tmp += err.Error()
		tmp += ";"
	}
	if tmp != "" {
		return fmt.Errorf(tmp)
	}
	return nil
}
