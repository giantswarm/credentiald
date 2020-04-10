package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/viper"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/microkit/command"
	microserver "github.com/giantswarm/microkit/server"
	"github.com/giantswarm/micrologger"

	"github.com/giantswarm/credentiald/flag"
	"github.com/giantswarm/credentiald/pkg/project"
	"github.com/giantswarm/credentiald/server"
	"github.com/giantswarm/credentiald/service"
)

var (
	f *flag.Flag = flag.New()
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	err := mainWithError()
	if err != nil {
		panic(fmt.Sprintf("%#v\n", err))
	}
}

func mainWithError() (err error) {
	var newLogger micrologger.Logger
	{
		newLogger, err = micrologger.New(micrologger.Config{})
		if err != nil {
			return microerror.Mask(err)
		}
	}

	newServerFactory := func(v *viper.Viper) microserver.Server {
		var newService *service.Service
		{
			c := service.Config{
				Flag:   f,
				Logger: newLogger,
				Viper:  v,

				Description: project.Description(),
				GitCommit:   project.GitSHA(),
				ProjectName: project.Name(),
				Source:      project.Source(),
				Version:     project.Version(),
			}

			newService, err = service.New(c)
			if err != nil {
				panic(fmt.Sprintf("%#v\n", microerror.Mask(err)))
			}

			go newService.Boot()
		}

		var newServer microserver.Server
		{
			c := server.Config{
				Logger:  newLogger,
				Service: newService,
				Viper:   v,

				ProjectName: project.Name(),
			}

			newServer, err = server.New(c)
			if err != nil {
				panic(fmt.Sprintf("%#v\n", microerror.Mask(err)))
			}
		}

		return newServer
	}

	var newCommand command.Command
	{
		c := command.Config{
			Logger:        newLogger,
			ServerFactory: newServerFactory,

			Description: project.Description(),
			GitCommit:   project.GitSHA(),
			Name:        project.Name(),
			Source:      project.Source(),
			Version:     project.Version(),
		}

		newCommand, err = command.New(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	daemonCommand := newCommand.DaemonCommand().CobraCommand()

	daemonCommand.PersistentFlags().String(f.Service.Kubernetes.Address, "", "Address used to connect to Kubernetes.")
	daemonCommand.PersistentFlags().Bool(f.Service.Kubernetes.InCluster, true, "Whether to use the in-cluster config to authenticate with Kubernetes.")
	daemonCommand.PersistentFlags().String(f.Service.Kubernetes.KubeConfig, "", "KubeConfig used to connect to Kubernetes. When empty other settings are used.")
	daemonCommand.PersistentFlags().String(f.Service.Kubernetes.TLS.CAFile, "", "Certificate authority file path to use to authenticate with Kubernetes.")
	daemonCommand.PersistentFlags().String(f.Service.Kubernetes.TLS.CrtFile, "", "Certificate file path to use to authenticate with Kubernetes.")
	daemonCommand.PersistentFlags().String(f.Service.Kubernetes.TLS.KeyFile, "", "Key file path to use to authenticate with Kubernetes.")

	daemonCommand.PersistentFlags().String(f.Service.Secrets.Namespace, "giantswarm", "Namespace to store secrets in.")
	daemonCommand.PersistentFlags().String(f.Service.Secrets.NameFormat, "credential-%s", "Format to name secrets with.")
	daemonCommand.PersistentFlags().Int(f.Service.Secrets.IDLength, 6, "Length of credential IDs.")

	newCommand.CobraCommand().Execute()

	return nil
}
