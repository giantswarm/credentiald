package project

var (
	description = "credentiald manages credentials for cloud environments."
	gitSHA      = "n/a"
	name        = "credentiald"
	source      = "https://github.com/giantswarm/credentiald"
	version     = "2.3.4-dev"
)

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
