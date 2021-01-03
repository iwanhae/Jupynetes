package config

import "github.com/spf13/viper"

//Names of environment variables to read
const (
	EnvDeploy                                = "deploy"
	EnvKubernetesNamespace                   = "kubernetes_Namespace"
	EnvKubernetesCreateNamespaceIfNotPresent = "kubernetes_CreateNamespaceIfNotPresent"
	EnvDatabaseURI                           = "database_uri"
	EnvDatabaseSalt                          = "database_salt"
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
	Database struct {
		URI  string
		Salt string
	}
}

//GetConfigs read environment variables and return configuration.
func GetConfigs() *Configs {
	config := Configs{}

	v := viper.New()
	v.AutomaticEnv()
	v.SetDefault(EnvDeploy, EnvDeployDebug)
	v.SetDefault(EnvKubernetesNamespace, "jupy")
	v.SetDefault(EnvKubernetesCreateNamespaceIfNotPresent, false)
	v.SetDefault(EnvDatabaseSalt, "ab448a918")

	config.Deploy = v.GetString(EnvDeploy)
	config.Kubernetes.Namespace = v.GetString(EnvKubernetesNamespace)
	config.Kubernetes.CreateNamespaceIfNotPresent = v.GetBool(EnvKubernetesCreateNamespaceIfNotPresent)
	config.Database.URI = v.GetString(EnvDatabaseURI)
	config.Database.Salt = v.GetString(EnvDatabaseSalt)

	return &config
}
