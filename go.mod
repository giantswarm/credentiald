module github.com/giantswarm/credentiald/v2

go 1.14

require (
	github.com/giantswarm/k8sclient/v4 v4.0.0
	github.com/giantswarm/microclient v0.2.0
	github.com/giantswarm/microendpoint v0.2.0
	github.com/giantswarm/microerror v0.2.1
	github.com/giantswarm/microkit v0.2.2
	github.com/giantswarm/micrologger v0.3.3
	github.com/giantswarm/operatorkit/v2 v2.0.1
	github.com/go-kit/kit v0.10.0
	github.com/gorilla/mux v1.8.0
	github.com/prometheus/client_golang v1.17.0
	github.com/spf13/viper v1.7.1
	gopkg.in/resty.v1 v1.12.0
	k8s.io/api v0.18.9
	k8s.io/apimachinery v0.18.9
	k8s.io/client-go v0.18.9
)

replace (
	github.com/coreos/etcd => github.com/coreos/etcd v3.3.25+incompatible
	github.com/dgrijalva/jwt-go => github.com/form3tech-oss/jwt-go v3.2.1+incompatible
	github.com/gogo/protobuf v1.3.1 => github.com/gogo/protobuf v1.3.2
	github.com/gorilla/websocket => github.com/gorilla/websocket v1.4.2
	golang.org/x/net => golang.org/x/net v0.18.0
)
