package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

//Names of environment variables to read
const (
	EnvDeploy                                = "deploy"
	EnvKubernetesNamespace                   = "kubernetes_Namespace"
	EnvKubernetesCreateNamespaceIfNotPresent = "kubernetes_CreateNamespaceIfNotPresent"
	EnvDatabaseURI                           = "database_uri"
	EnvDatabaseSalt                          = "database_salt"
	EnvSecretKey                             = "secretkey"
	EnvQuotaInstance                         = "quota_instance"
	EnvQuotaCpu                              = "quota_cpu"
	EnvQuotaMemory                           = "quota_memory"
	EnvQuotaNvidiaGpu                        = "quota_nvidia_gpu"
	EnvQuotaStorage                          = "quota_storage"
	EnvDomainUpper                           = "domain_upper"
	EnvDomainPrefix                          = "domain_prefix"
)

//Possible parameters
const (
	EnvDeployDebug = "debug"
	EnvDeployStage = "stage"
	EnvDeployProd  = "prod"
)

//Configs saves configuration needed
type Configs struct {
	Deploy    string // One of debug, stage or prod
	SecretKey string // used for JWT signing

	Kubernetes struct {
		Namespace                   string
		CreateNamespaceIfNotPresent bool
	}
	Database struct {
		URI  string
		Salt string
	}
	Quota struct {
		Instance  int
		Cpu       int
		Memory    int
		NvidiaGpu int
		Storage   int
	}
	Domain struct {
		Upper  string // csuos.ml
		Prefix string // if prefix is "jupy", ingress will be generated likes jupy-servername.csuos.ml
	}
}

//GetConfigs read environment variables and return configuration.
func GetConfigs() *Configs {
	config := Configs{}

	v := viper.New()

	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		v.AutomaticEnv()
	} else {
		v.SetConfigType("dotenv")
		v.SetConfigFile(".env")
		err := v.ReadInConfig()
		if err != nil {
			log.Fatalf("fail to read dotenv file:%s", err.Error())
		}
	}

	v.SetDefault(EnvDeploy, EnvDeployDebug)
	v.SetDefault(EnvSecretKey, "e190ufqe2")
	v.SetDefault(EnvKubernetesNamespace, "jupy")
	v.SetDefault(EnvKubernetesCreateNamespaceIfNotPresent, false)
	v.SetDefault(EnvDatabaseSalt, "ab448a918")

	v.SetDefault(EnvQuotaInstance, -1)
	v.SetDefault(EnvQuotaCpu, -1)
	v.SetDefault(EnvQuotaMemory, -1)
	v.SetDefault(EnvQuotaNvidiaGpu, -1)
	v.SetDefault(EnvQuotaStorage, -1)

	v.SetDefault(EnvDomainUpper, "")
	v.SetDefault(EnvDomainPrefix, "")

	//
	config.Deploy = v.GetString(EnvDeploy)
	config.SecretKey = v.GetString(EnvSecretKey)
	config.Kubernetes.Namespace = v.GetString(EnvKubernetesNamespace)
	config.Kubernetes.CreateNamespaceIfNotPresent = v.GetBool(EnvKubernetesCreateNamespaceIfNotPresent)

	config.Database.URI = v.GetString(EnvDatabaseURI)
	config.Database.Salt = v.GetString(EnvDatabaseSalt)

	config.Quota.Instance = v.GetInt(EnvQuotaInstance)
	config.Quota.Cpu = v.GetInt(EnvQuotaCpu)
	config.Quota.Memory = v.GetInt(EnvQuotaMemory)
	config.Quota.NvidiaGpu = v.GetInt(EnvQuotaNvidiaGpu)
	config.Quota.Storage = v.GetInt(EnvQuotaStorage)

	config.Domain.Upper = v.GetString(EnvDomainUpper)
	config.Domain.Prefix = v.GetString(EnvDomainPrefix)

	return &config
}
