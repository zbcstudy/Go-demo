module a_pkg

go 1.15

require (
	dubbo.apache.org/dubbo-go/v3 v3.0.2
	github.com/FerretDB/FerretDB v1.3.0
	github.com/argoproj/argo-cd v1.8.7
	//github.com/Mirantis/cri-dockerd latest
	github.com/blend/go-sdk v1.20220411.4-0.20220608163541-be8198318397
	github.com/caarlos0/starcharts v1.6.5-0.20221003175303-951f7eea37bf
	github.com/cilium/cilium v1.15.3
	github.com/cloudwego/hertz v0.3.2
	github.com/cloudwego/kitex v0.4.2
	github.com/cockroachdb/cockroach v0.0.0-20221019051325-bce7b6fee1bb
	github.com/cockroachdb/pebble v0.0.0-20221018213927-bdc28bc8dac2
	github.com/containerd/containerd v1.7.13
	github.com/containerd/nerdctl v1.7.5
	github.com/containernetworking/cni v1.1.3-0.20221031153233-76aaefb2cba2
	github.com/containers/buildah v1.35.1
	github.com/containers/fetchit v0.0.2-0.20221010155311-f5ffd62ecd0c
	github.com/containers/image/v5 v5.30.0
	github.com/containers/podman/v4 v4.9.4
	github.com/containers/podman/v5 v5.0.0
	github.com/containers/skopeo v1.12.0
	github.com/coocood/freecache v1.2.4
	github.com/cri-o/cri-o v1.27.4
	github.com/dgraph-io/badger/v3 v3.2103.5
	github.com/dgraph-io/ristretto v0.1.1 // indirect
	github.com/docker-slim/docker-slim v0.0.0-20221114191106-a4aca0c03e46
	github.com/docker/buildx v0.9.2-0.20221110133911-3690cb12e6f9
	github.com/docker/compose/v2 v2.12.3-0.20221108144646-754376916c09
	github.com/docker/machine v0.16.2
	// github.com/OpenIMSDK/Open-IM-Server/v2 a4fe45d
	// github.com/LockGit/gochat v0.0.0-20221024110244-4284a5a8971e
	github.com/douyu/jupiter v0.9.0
	github.com/elastic/go-elasticsearch/v7 v7.17.1
	github.com/emitter-io/emitter v0.0.0-20240202195010-996fa4af0e95 //v3.1
	github.com/flower-corp/lotusdb v1.0.1-0.20220729120541-8d44d1d6c0ab
	github.com/flower-corp/rosedb v1.1.2-0.20220817135557-31bf5acf7397
	github.com/fluxcd/flux2/v2 v2.2.3
	github.com/go-co-op/gocron v1.28.3
	github.com/go-kit/kit v0.13.1-0.20231222231659-844c3d2de01a
	github.com/go-nunu/nunu v1.0.4
	github.com/go-redsync/redsync v1.4.2
	github.com/gobuffalo/buffalo v1.0.1
	github.com/goharbor/harbor/src v0.0.0-20240403034559-9778176ff1ee
	github.com/golang-module/dongle v0.2.4
	github.com/goodrain/rainbond v0.0.0-20221102023605-b0029b2ca126
	github.com/gookit/config/v2 v2.1.6
	github.com/gookit/filter v1.1.3
	github.com/gookit/gcli/v3 v3.0.7-0.20220904045221-3d6cf1f66110
	github.com/gookit/rux v1.3.4
	github.com/gookit/validate v1.4.4
	github.com/hashicorp/consul v1.16.6
	github.com/hashicorp/nomad v1.7.5
	github.com/hashicorp/packer v1.10.1
	github.com/hashicorp/raft v1.5.0
	github.com/hashicorp/vault v1.15.6
	github.com/hdt3213/godis v1.2.8
	github.com/influxdata/flux v0.194.3
	github.com/influxdata/influxdb/v2 v2.5.0
	github.com/influxdata/telegraf v1.28.1
	github.com/ipfs/kubo v0.26.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/kata-containers/kata-containers v0.0.0-20221110134924-56641bc2303c
	github.com/kataras/iris/v12 v12.2.0-beta6.0.20221101212547-4c188d86fdc1
	github.com/kataras/muxie v1.1.2
	github.com/kedacore/keda/v2 v2.13.1
	github.com/koderover/zadig v1.18.0
	github.com/kubecube-io/kubecube v1.5.1
	github.com/kubernetes-sigs/cri-tools v0.0.0
	github.com/labstack/echo/v4 v4.11.4
	github.com/leaanthony/slicer v1.6.0
	github.com/lima-vm/lima v0.13.1-0.20221108071032-4a4cc62749c1
	github.com/linkerd/linkerd2 v0.5.1-0.20221110190454-75cbed2469cc
	github.com/lxc/lxd v0.0.0-20221110184730-a10417833a8e
	github.com/markbates/goth v1.74.2
	github.com/minio/minio v0.0.0-20221026063237-8dd3c41b2a7d
	github.com/moby/buildkit v0.12.5
	github.com/moby/swarmkit/v2 v2.0.0-20221108223242-f34bfc0e146e
	github.com/mtrmac/gpgme v0.1.2 // indirect
	github.com/nats-io/jwt/v2 v2.5.3
	github.com/nats-io/nats-server/v2 v2.10.11
	github.com/no-src/gofs v0.7.1
	github.com/nsqio/nsq v1.2.2-0.20220923154746-393f7b994f0c
	github.com/oam-dev/kubevela v1.9.9
	github.com/olivere/elastic/v7 v7.0.32
	github.com/opencontainers/runtime-tools v0.9.1-0.20230914150019-408c51e934dc
	github.com/opencontainers/umoci v0.4.8-0.20221029002459-fb2db51251ac
	github.com/opensearch-project/opensearch-go/v2 v2.3.0
	github.com/panjf2000/gnet v1.6.0
	github.com/pingcap/tidb v1.1.0-beta.0.20221101122759-863d35290cc4
	github.com/pingcap/tiflow v0.0.0-20221107160750-0ad56ca6594a
	github.com/polarismesh/polaris v1.17.2-0.20230531031309-90b39c8785f4
	github.com/rancher/k3s v1.21.9
	github.com/rancher/rancher v0.0.0-20240226212747-4119767ddf14
	github.com/rancher/rke v1.5.5
	github.com/rancher/rke2 v1.25.4-0.20221102171456-ddb87a9b6684
	github.com/reugn/go-streams v0.7.0
	github.com/rkt/rkt v1.30.1-0.20200224141603-171c416fac02
	github.com/rook/rook v1.13.7
	github.com/rpcxio/basalt v0.0.0-20200813040018-a7e186f1ae4c
	github.com/rpcxio/did v0.0.0-20200223142455-7b9772571fae
	github.com/rpcxio/rpcx-examples v1.1.7-0.20220606035200-e356ac23accc
	github.com/rpcxio/rpcx-gateway v0.0.0-20220611140137-f4a622048e82
	github.com/rqlite/rqlite v0.0.0-20221021022757-ad7f383d8386
	github.com/smallnest/rpcx v1.7.11
	// go-micro.dev/v4 v4.10.2
	github.com/syncthing/syncthing v1.23.5
	github.com/syndtr/goleveldb v1.0.1-0.20220721030215-126854af5e6d
	github.com/tinode/chat v0.20.1-0.20221011125331-85ecad89d2db
	github.com/unionj-cloud/go-doudou/v2 v2.3.0
	github.com/vbauerster/mpb/v5 v5.4.0 // indirect
	github.com/volatiletech/authboss/v3 v3.2.1
	github.com/volatiletech/sqlboiler/v4 v4.13.0
	github.com/xuri/excelize/v2 v2.8.1
	go.mongodb.org/mongo-driver v1.14.0
	helm.sh/helm/v3 v3.14.3
	k8s.io/minikube v1.28.1-0.20221202001624-44ea7f43205e
	k8s.io/mount-utils v0.25.2
	kubesphere.io/kubesphere v0.0.0-20230912034912-a81d22c98875 //v0.0.0-20221102022257-ff59b80c2807
)

replace (
	//k8s.io/mount-utils => k8s.io/mount-utils v0.20.4
	github.com/Mirantis/cri-dockerd => github.com/Mirantis/cri-dockerd v0.2.6

	github.com/abourget/teamcity => github.com/cockroachdb/teamcity v0.0.0-20180905144921-8ca25c33eb11
	github.com/blend/go-sdk => github.com/blend/go-sdk v1.20220411.4-0.20220608163541-be8198318397

	github.com/docker/docker => github.com/docker/docker v20.10.21+incompatible
	github.com/kubernetes-incubator/external-storage => github.com/kubernetes-incubator/external-storage v5.5.0+incompatible
	github.com/kubernetes-sigs/cri-tools => github.com/kubernetes-sigs/cri-tools v1.25.0
	github.com/micro/cli => github.com/go-micro/cli v1.1.4

	github.com/mikefarah/yaml => github.com/go-yaml/yaml v0.0.0-20220527083530-f6f7691b1fde

	github.com/myesui/uuid => github.com/google/uuid v1.3.0
	github.com/opensearch-project/opensearch-go/v2 => github.com/opensearch-project/opensearch-go/v2 v2.1.0
	github.com/operator-framework/operator-lib => github.com/operator-framework/operator-lib v0.11.0
	github.com/polarismesh/go-restful-openapi => github.com/emicklei/go-restful-openapi/v2 v2.9.1

	github.com/projectcalico/api => github.com/projectcalico/api v0.0.0-20230602153125-fb7148692637
	github.com/projectcalico/calico => github.com/projectcalico/calico v2.6.12+incompatible

	github.com/rancher/rancher/pkg/apis => github.com/rancher/rancher/pkg/apis v0.0.0-20221103194918-2d114cccb546
	github.com/rancher/rancher/pkg/client => github.com/rancher/rancher/pkg/client v0.0.0-20221103194918-2d114cccb546

	github.com/sigstore/fulcio => github.com/sigstore/fulcio v1.0.1-0.20221110094946-7cf41b0898d4

	go.etcd.io/etcd => github.com/etcd-io/etcd v3.3.27+incompatible

	go.qase.io/client => github.com/rancher/qase-go/client v0.0.0-20231114201952-65195ec001fa
	golang.org/x/time => golang.org/x/time v0.0.0-20220411224347-583f2d630306
	google.golang.org/grpc => google.golang.org/grpc v1.50.1
	k8s.io/api => k8s.io/kubernetes/staging/src/k8s.io/api v0.0.0-20220823173643-a866cbe2e5bb
	k8s.io/apiextensions-apiserver => k8s.io/kubernetes/staging/src/k8s.io/apiextensions-apiserver v0.0.0-20220823173643-a866cbe2e5bb
	k8s.io/apimachinery => k8s.io/kubernetes/staging/src/k8s.io/apimachinery v0.0.0-20220823173643-a866cbe2e5bb
	k8s.io/apiserver => k8s.io/kubernetes/staging/src/k8s.io/apiserver v0.0.0-20220823173643-a866cbe2e5bb
	k8s.io/cli-runtime => k8s.io/kubernetes/staging/src/k8s.io/cli-runtime v0.0.0-20220823173643-a866cbe2e5bb
	k8s.io/client-go => k8s.io/kubernetes/staging/src/k8s.io/client-go v0.0.0-20220823173643-a866cbe2e5bb
	k8s.io/cloud-provider => k8s.io/kubernetes/staging/src/k8s.io/cloud-provider v0.0.0-20220823173643-a866cbe2e5bb
	k8s.io/cluster-bootstrap => k8s.io/kubernetes/staging/src/k8s.io/cluster-bootstrap v0.0.0-20220823173643-a866cbe2e5bb
	k8s.io/code-generator => k8s.io/kubernetes/staging/src/k8s.io/code-generator v0.0.0-20220823173643-a866cbe2e5bb
	k8s.io/component-base => k8s.io/kubernetes/staging/src/k8s.io/component-base v0.0.0-20220823173643-a866cbe2e5bb
	k8s.io/component-helpers => k8s.io/kubernetes/staging/src/k8s.io/component-helpers v0.0.0-20220823173643-a866cbe2e5bb
	k8s.io/controller-manager => k8s.io/kubernetes/staging/src/k8s.io/controller-manager v0.0.0-20220823173643-a866cbe2e5bb
	k8s.io/cri-api => k8s.io/kubernetes/staging/src/k8s.io/cri-api v0.0.0-20220823173643-a866cbe2e5bb
	k8s.io/csi-translation-lib => k8s.io/kubernetes/staging/src/k8s.io/csi-translation-lib v0.0.0-20220823173643-a866cbe2e5bb
	k8s.io/kube-aggregator => k8s.io/kubernetes/staging/src/k8s.io/kube-aggregator v0.0.0-20220823173643-a866cbe2e5bb
	k8s.io/kube-controller-manager => k8s.io/kubernetes/staging/src/k8s.io/kube-controller-manager v0.0.0-20220823173643-a866cbe2e5bb
	k8s.io/kube-proxy => k8s.io/kubernetes/staging/src/k8s.io/kube-proxy v0.0.0-20220823173643-a866cbe2e5bb
	k8s.io/kube-scheduler => k8s.io/kubernetes/staging/src/k8s.io/kube-scheduler v0.0.0-20220823173643-a866cbe2e5bb
	k8s.io/kubectl => k8s.io/kubernetes/staging/src/k8s.io/kubectl v0.0.0-20220823173643-a866cbe2e5bb
	k8s.io/kubelet => k8s.io/kubernetes/staging/src/k8s.io/kubelet v0.0.0-20220823173643-a866cbe2e5bb
	k8s.io/kubernetes => k8s.io/kubernetes v1.25.0
	k8s.io/legacy-cloud-providers => k8s.io/kubernetes/staging/src/k8s.io/legacy-cloud-providers v0.0.0-20220823173643-a866cbe2e5bb
	k8s.io/metrics => k8s.io/kubernetes/staging/src/k8s.io/metrics v0.0.0-20220823173643-a866cbe2e5bb
	k8s.io/mount-utils => k8s.io/kubernetes/staging/src/k8s.io/mount-utils v0.0.0-20220823173643-a866cbe2e5bb
	k8s.io/pod-security-admission => k8s.io/kubernetes/staging/src/k8s.io/pod-security-admission v0.0.0-20220823173643-a866cbe2e5bb
	k8s.io/sample-apiserver => k8s.io/kubernetes/staging/src/k8s.io/sample-apiserver v0.0.0-20220823173643-a866cbe2e5bb
	kubesphere.io/api => kubesphere.io/api v0.0.0-20221031095409-fdfb4c74e5e9
	kubesphere.io/client-go => kubesphere.io/client-go v0.3.1

	sigs.k8s.io/apiserver-network-proxy/konnectivity-client => sigs.k8s.io/apiserver-network-proxy/konnectivity-client v0.0.30
	sigs.k8s.io/kubefed => sigs.k8s.io/kubefed v0.10.0
	vitess.io/vitess => github.com/cockroachdb/vitess v0.0.0-20210218160543-54524729cc82
)
