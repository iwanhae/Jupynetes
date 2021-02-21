package kubeclient

import (
	"context"
	"path/filepath"

	"github.com/iwanhae/Jupynetes/pkg/config"
	"github.com/rs/zerolog/log"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var clientset *kubernetes.Clientset
var defaultNamespace string

var domainUpper string
var domainPrefix string

//Init Initialize k8s client configurations
func Init(ctx context.Context, c *config.Configs) {
	defaultNamespace = c.Kubernetes.Namespace

	domainUpper = c.Domain.Upper
	domainPrefix = c.Domain.Prefix

	config, err := rest.InClusterConfig()
	if err == nil {
		log.Info().Msg("parsed incluster kubeconfig")
	} else {
		log.Info().Msg("fail to parse incluster kubeconfig")
		config, err = clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
		if err != nil {
			log.Fatal().Err(err).Msg("fail to connect kubernetes cluster:fail to parse kubeconfig")
		}
		log.Info().Msg("parsed kubeconfig from ~/.kube/config")
	}

	// create the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal().Err(err).Msg("fail to make kubernetes clientset")
	}

	var namespacePresent bool
	nss, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal().Err(err).Msg("fail to retreive namespace from cluster")
	}
	for _, ns := range nss.Items {
		log.Debug().Str("namespace", ns.Name).Msg("detect namespace")
		if ns.Name == c.Kubernetes.Namespace {
			namespacePresent = true
		}
	}
	if namespacePresent == false {
		if c.Kubernetes.CreateNamespaceIfNotPresent == false {
			log.Fatal().Str("namespace", c.Kubernetes.Namespace).Msg("requested namespace not present")
		}

		ns, err := clientset.CoreV1().Namespaces().Create(context.TODO(), &v1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: c.Kubernetes.Namespace,
			},
		}, metav1.CreateOptions{})

		if err != nil {
			log.Fatal().Err(err).Str("namespace", c.Kubernetes.Namespace).Msg("fail to create requested namespace")
		}

		log.Info().Interface("namespace_object", ns).Msg("namespace created")
	}
	log.Info().Msgf("using namespace %q", c.Kubernetes.Namespace)
}
