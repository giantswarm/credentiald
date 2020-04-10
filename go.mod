module github.com/giantswarm/credentiald

go 1.14

require (
	github.com/beorn7/perks v1.0.1
	github.com/coreos/go-semver v0.2.0
	github.com/davecgh/go-spew v1.1.1
	github.com/fsnotify/fsnotify v1.4.7
	github.com/giantswarm/k8sclient v0.2.0
	github.com/giantswarm/microclient v0.2.0
	github.com/giantswarm/microendpoint v0.2.0
	github.com/giantswarm/microerror v0.2.0
	github.com/giantswarm/microkit v0.2.0
	github.com/giantswarm/micrologger v0.3.1
	github.com/giantswarm/operatorkit v0.2.0
	github.com/giantswarm/versionbundle v0.2.0
	github.com/go-kit/kit v0.6.0
	github.com/go-logfmt/logfmt v0.4.0
	github.com/go-stack/stack v1.8.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.2
	github.com/google/gofuzz v1.0.0
	github.com/googleapis/gnostic v0.3.1
	github.com/gorilla/mux v1.7.3
	github.com/hashicorp/hcl v1.0.0
	github.com/imdario/mergo v0.3.8
	github.com/inconshreveable/mousetrap v1.0.0
	github.com/json-iterator/go v1.1.8
	github.com/juju/errgo v0.0.0-20140925100237-08cceb5d0b53
	github.com/kr/logfmt v0.0.0-20140226030751-b84e30acd515
	github.com/magiconair/properties v1.8.1
	github.com/matttproud/golang_protobuf_extensions v1.0.1
	github.com/mitchellh/mapstructure v1.1.2
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742
	github.com/pelletier/go-toml v1.6.0
	github.com/prometheus/client_golang v0.9.4
	github.com/prometheus/client_model v0.1.0
	github.com/prometheus/common v0.7.0
	github.com/prometheus/procfs v0.0.8
	github.com/spf13/afero v1.2.2
	github.com/spf13/cast v1.3.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/jwalterweatherman v1.1.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.1
	github.com/subosito/gotenv v1.2.0
	golang.org/x/crypto v0.0.0-20191206172530-e9b2fee46413
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553
	golang.org/x/oauth2 v0.0.0-20191202225959-858c2ad4c8b6
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	golang.org/x/sys v0.0.0-20191206220618-eeba5f6aabab
	golang.org/x/text v0.3.2
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0
	google.golang.org/appengine v1.6.5
	gopkg.in/inf.v0 v0.9.1
	gopkg.in/ini.v1 v1.51.0
	gopkg.in/resty.v1 v1.12.0
	gopkg.in/yaml.v2 v2.2.7
	k8s.io/api v0.0.0-20191114100352-16d7abae0d2a
	k8s.io/apimachinery v0.16.5-beta.1
	k8s.io/client-go v0.0.0-20191114101535-6c5935290e33
	k8s.io/klog v1.0.0
	k8s.io/utils v0.0.0-20191114200735-6ca3b61696b6
	sigs.k8s.io/yaml v1.1.0
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20191114100352-16d7abae0d2a
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20191114105449-027877536833
	k8s.io/apimachinery => k8s.io/apimachinery v0.16.5-beta.1
	k8s.io/apiserver => k8s.io/apiserver v0.0.0-20191114103151-9ca1dc586682
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20191114110141-0a35778df828
	k8s.io/client-go => k8s.io/client-go v0.0.0-20191114101535-6c5935290e33
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20191114112024-4bbba8331835
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.0.0-20191114111741-81bb9acf592d
	k8s.io/code-generator => k8s.io/code-generator v0.16.5-beta.1
	k8s.io/component-base => k8s.io/component-base v0.0.0-20191114102325-35a9586014f7
	k8s.io/cri-api => k8s.io/cri-api v0.16.5-beta.1
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.0.0-20191114112310-0da609c4ca2d
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20191114103820-f023614fb9ea
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.0.0-20191114111510-6d1ed697a64b
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.0.0-20191114110717-50a77e50d7d9
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.0.0-20191114111229-2e90afcb56c7
	k8s.io/kubectl => k8s.io/kubectl v0.0.0-20191114113550-6123e1c827f7
	k8s.io/kubelet => k8s.io/kubelet v0.0.0-20191114110954-d67a8e7e2200
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.0.0-20191114112655-db9be3e678bb
	k8s.io/metrics => k8s.io/metrics v0.0.0-20191114105837-a4a2842dc51b
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.0.0-20191114104439-68caf20693ac
)
