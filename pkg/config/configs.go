package config

import "github.com/spf13/viper"

//Names of environment variables to read
const (
	EnvDeploy                                = "deploy"
	EnvKubernetesNamespace                   = "kubernetes_Namespace"
	EnvKubernetesCreateNamespaceIfNotPresent = "kubernetes_CreateNamespaceIfNotPresent"
)

//Possible parameters
const (
	EnvDeployDebug = "debug"
	EnvDeployStage = "stage"
	EnvDeployProd  = "prod"
)

//Configs saves configuration needed
type Configs struct {
	Deploy string // One of debug, stage or prod

	Kubernetes struct {
		Namespace                   string
		CreateNamespaceIfNotPresent bool
	}
}

//GetConfigs read environment variables and return configuration.
func GetConfigs() Configs {
	config := Configs{}

	v := viper.New()
	v.AutomaticEnv()
	v.SetDefault(EnvDeploy, EnvDeployDebug)
	v.SetDefault(EnvKubernetesNamespace, "jupy")
	v.SetDefault(EnvKubernetesCreateNamespaceIfNotPresent, false)

	config.Deploy = v.GetString(EnvDeploy)

	config.Kubernetes.Namespace = v.GetString(EnvKubernetesNamespace)
	config.Kubernetes.CreateNamespaceIfNotPresent = v.GetBool(EnvKubernetesCreateNamespaceIfNotPresent)
	return config
}
