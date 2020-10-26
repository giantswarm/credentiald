package service

import (
	"sync"

	"github.com/giantswarm/k8sclient/v4/pkg/k8srestconfig"
	"github.com/giantswarm/microendpoint/service/version"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/giantswarm/credentiald/v2/flag"
	"github.com/giantswarm/credentiald/v2/service/collector"
	"github.com/giantswarm/credentiald/v2/service/creator"
	"github.com/giantswarm/credentiald/v2/service/lister"
	"github.com/giantswarm/credentiald/v2/service/searcher"
)

type Config struct {
	Flag   *flag.Flag
	Logger micrologger.Logger
	Viper  *viper.Viper

	Description string
	GitCommit   string
	ProjectName string
	Source      string
	Version     string
}

type Service struct {
	Creator  *creator.Service
	Lister   *lister.Service
	Searcher *searcher.Service
	Version  *version.Service

	bootOnce sync.Once
}

// New creates a new service with given configuration.
func New(config Config) (*Service, error) {
	if config.Flag == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Flag must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.Viper == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Viper must not be empty", config)
	}

	if config.ProjectName == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.ProjectName must not be empty", config)
	}

	var err error

	var restConfig *rest.Config
	{
		c := k8srestconfig.Config{
			Logger: config.Logger,

			Address:    config.Viper.GetString(config.Flag.Service.Kubernetes.Address),
			InCluster:  config.Viper.GetBool(config.Flag.Service.Kubernetes.InCluster),
			KubeConfig: config.Viper.GetString(config.Flag.Service.Kubernetes.KubeConfig),
			TLS: k8srestconfig.ConfigTLS{
				CAFile:  config.Viper.GetString(config.Flag.Service.Kubernetes.TLS.CAFile),
				CrtFile: config.Viper.GetString(config.Flag.Service.Kubernetes.TLS.CrtFile),
				KeyFile: config.Viper.GetString(config.Flag.Service.Kubernetes.TLS.KeyFile),
			},
		}

		restConfig, err = k8srestconfig.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	k8sClient, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	{
		var credentialdCollector *collector.Collector

		c := collector.Config{
			K8sClient: k8sClient,
			Logger:    config.Logger,
		}
		credentialdCollector, err = collector.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		prometheus.MustRegister(credentialdCollector)
	}

	var creatorService *creator.Service
	{
		c := creator.Config{
			K8sClient: k8sClient,
			Logger:    config.Logger,

			IDLength:         config.Viper.GetInt(config.Flag.Service.Secrets.IDLength),
			NameFormat:       config.Viper.GetString(config.Flag.Service.Secrets.NameFormat),
			SecretsNamespace: config.Viper.GetString(config.Flag.Service.Secrets.Namespace),
		}

		creatorService, err = creator.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var listerService *lister.Service
	{
		c := lister.Config{
			K8sClient: k8sClient,
			Logger:    config.Logger,
		}

		listerService, err = lister.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var searcherService *searcher.Service
	{
		c := searcher.Config{
			K8sClient: k8sClient,
			Logger:    config.Logger,
		}

		searcherService, err = searcher.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var versionService *version.Service
	{
		c := version.Config{
			Description: config.Description,
			GitCommit:   config.GitCommit,
			Name:        config.ProjectName,
			Source:      config.Source,
			Version:     config.Version,
		}

		versionService, err = version.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	s := &Service{
		Creator:  creatorService,
		Lister:   listerService,
		Searcher: searcherService,
		Version:  versionService,

		bootOnce: sync.Once{},
	}

	return s, nil
}

func (s *Service) Boot() {
	s.bootOnce.Do(func() {})
}
