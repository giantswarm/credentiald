module github.com/giantswarm/credentiald/v2

go 1.14

require (
	github.com/giantswarm/k8sclient/v4 v4.1.0
	github.com/giantswarm/microclient v0.2.0
	github.com/giantswarm/microendpoint v0.2.0
	github.com/giantswarm/microerror v0.4.1
	github.com/giantswarm/microkit v1.0.2
	github.com/giantswarm/micrologger v1.1.1
	github.com/giantswarm/operatorkit/v2 v2.0.2
	github.com/go-kit/kit v0.13.0
	github.com/gorilla/mux v1.8.1
	github.com/prometheus/client_golang v1.20.5
	github.com/spf13/viper v1.19.0
	gopkg.in/resty.v1 v1.12.0
	k8s.io/api v0.18.14
	k8s.io/apimachinery v0.18.14
	k8s.io/client-go v0.18.14
)

replace (
	github.com/coreos/etcd => github.com/coreos/etcd v3.3.25+incompatible
	github.com/dgrijalva/jwt-go => github.com/form3tech-oss/jwt-go v3.2.1+incompatible
	github.com/gogo/protobuf v1.3.1 => github.com/gogo/protobuf v1.3.2
	github.com/gorilla/websocket => github.com/gorilla/websocket v1.4.2
	golang.org/x/net => golang.org/x/net v0.18.0
)
