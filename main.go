package main

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/microkit/command"
	microserver "github.com/giantswarm/microkit/server"
	"github.com/giantswarm/micrologger"

	"github.com/giantswarm/credentiald/flag"
	"github.com/giantswarm/credentiald/server"
	"github.com/giantswarm/credentiald/service"
)

var (
	description = "credentiald manages credentials for cloud environments."
	f           = flag.New()
	gitCommit   = "n/a"
	name        = "credentiald"
	source      = "https://github.com/giantswarm/credentiald"
)

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
			return microerror.Maskf(err, "micrologger.New")
		}
	}

	newServerFactory := func(v *viper.Viper) microserver.Server {
		var newService *service.Service
		{
			c := service.Config{
				Flag:   f,
				Logger: newLogger,
				Viper:  v,

				Description: description,
				GitCommit:   gitCommit,
				ProjectName: name,
				Source:      source,
			}

			newService, err = service.New(c)
			if err != nil {
				panic(fmt.Sprintf("%#v\n", microerror.Maskf(err, "service.New")))
			}

			go newService.Boot()
		}

		var newServer microserver.Server
		{
			c := server.Config{
				Logger:  newLogger,
				Service: newService,
				Viper:   v,

				ProjectName: name,
			}

			newServer, err = server.New(c)
			if err != nil {
				panic(fmt.Sprintf("%#v\n", microerror.Maskf(err, "server.New")))
			}
		}

		return newServer
	}

	var newCommand command.Command
	{
		c := command.Config{
			Logger:        newLogger,
			ServerFactory: newServerFactory,

			Description: description,
			GitCommit:   gitCommit,
			Name:        name,
			Source:      source,
		}

		newCommand, err = command.New(c)
		if err != nil {
			return microerror.Maskf(err, "command.New")
		}
	}

	newCommand.CobraCommand().Execute()

	return nil
}
