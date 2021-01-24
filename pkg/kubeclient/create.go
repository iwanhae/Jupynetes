package kubeclient

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/iwanhae/Jupynetes/pkg/common"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1beta1 "k8s.io/api/networking/v1beta1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/apimachinery/pkg/util/yaml"
)

//DeployServer deploy server based on template
func DeployServer(ctx context.Context, name string, template *common.Template) error {

	err := deployDeploymnet(ctx, name, template)
	if err != nil {
		return err
	}

	err = deployService(ctx, name, template)
	if err != nil {
		return err
	}

	err = deplyIngress(ctx, name, template)
	if err != nil {
		return err
	}

	return nil
}

func deplyIngress(ctx context.Context, name string, template *common.Template) error {
	clientset.NetworkingV1beta1().Ingresses(defaultNamespace).Delete(ctx, name, metav1.DeleteOptions{})
	_, err := clientset.NetworkingV1beta1().Ingresses(defaultNamespace).Create(ctx,
		&networkingv1beta1.Ingress{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Ingress",
				APIVersion: "networking.k8s.io/v1beta1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: defaultNamespace,
			},
			Spec: networkingv1beta1.IngressSpec{
				Rules: []networkingv1beta1.IngressRule{
					{
						Host: fmt.Sprintf("%s-jupynetes.iwanhae.kr", name),
						IngressRuleValue: networkingv1beta1.IngressRuleValue{
							HTTP: &networkingv1beta1.HTTPIngressRuleValue{
								Paths: []networkingv1beta1.HTTPIngressPath{
									{
										Path: "/",
										Backend: networkingv1beta1.IngressBackend{
											ServiceName: name,
											ServicePort: intstr.FromInt(80),
										},
									},
								},
							},
						},
					},
				},
			},
		}, metav1.CreateOptions{})
	return err
}

func deployService(ctx context.Context, name string, template *common.Template) error {
	clientset.CoreV1().Services(defaultNamespace).Delete(ctx, name, metav1.DeleteOptions{})
	_, err := clientset.CoreV1().Services(defaultNamespace).Create(ctx, &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: defaultNamespace,
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				LabelAppName: name,
			},
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Protocol:   corev1.ProtocolTCP,
					Port:       80,
					TargetPort: intstr.FromInt(8080),
				},
			},
		},
	}, metav1.CreateOptions{})
	return err
}

func deployDeploymnet(ctx context.Context, name string, template *common.Template) error {
	variables := make(map[string]string)
	for _, v := range template.Variables {
		variables[v.Name] = v.Value
	}

	deploymentBody := template.Body

	for key, value := range variables {
		deploymentBody = strings.ReplaceAll(
			deploymentBody,
			fmt.Sprintf("{{%s}}", key),
			value,
		)
	}

	deployment := &v1.Deployment{}

	err := yaml.NewYAMLOrJSONDecoder(
		bytes.NewReader([]byte(deploymentBody)),
		1000).Decode(deployment)
	if err != nil {
		return fmt.Errorf("fail to parse deployment:%s", err.Error())
	}

	deployment.Name = name
	deployment.Namespace = defaultNamespace
	deployment.Spec.Template.Labels[LabelAppName] = name

	cpu, err := strconv.Atoi(variables[VariableFlavorCPU])
	if err != nil || cpu == 0 {
		return fmt.Errorf("fail to parse cpu:%s", err)
	}
	mem, err := strconv.Atoi(variables[VariableFlavorMem])
	if err != nil || mem == 0 {
		return fmt.Errorf("fail to parse mem:%s", err)
	}

	limits := deployment.Spec.Template.Spec.Containers[0].Resources.Limits.DeepCopy()

	if limits == nil {
		limits = make(corev1.ResourceList)
	}

	limits[corev1.ResourceCPU] = *resource.NewMilliQuantity(int64(cpu), resource.DecimalSI)
	limits[corev1.ResourceMemory] = *resource.NewScaledQuantity(int64(mem), resource.Mega)
	deployment.Spec.Template.Spec.Containers[0].Resources.Limits = limits

	_, err = clientset.AppsV1().Deployments(defaultNamespace).Get(ctx, name, metav1.GetOptions{})
	if err == nil {
		return fmt.Errorf("deployment %s already exsists", name)
	}
	_, err = clientset.AppsV1().Deployments(defaultNamespace).Create(ctx, deployment, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("fail to deploy deployments:%s", err.Error())
	}
	return nil
}
