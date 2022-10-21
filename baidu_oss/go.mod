module baidu_oss

go 1.15

require (
	dubbo.apache.org/dubbo-go/v3 v3.0.2
	github.com/Mirantis/cri-dockerd v0.0.0
	github.com/cloudwego/hertz v0.3.2
	github.com/cloudwego/kitex v0.4.2
	github.com/cockroachdb/cockroach v0.0.0-20221019051325-bce7b6fee1bb
	github.com/cockroachdb/pebble v0.0.0-20221018213927-bdc28bc8dac2
	github.com/containerd/containerd v1.6.8
	github.com/coocood/freecache v1.2.2
	github.com/dgraph-io/badger/v3 v3.0.0-20221014230948-7a3a87a78191
	github.com/elastic/go-elasticsearch/v7 v7.17.0
	github.com/gookit/config/v2 v2.1.6
	github.com/gookit/filter v1.1.3
	github.com/gookit/gcli/v3 v3.0.7-0.20220904045221-3d6cf1f66110
	github.com/gookit/rux v1.3.4
	github.com/gookit/validate v1.4.4
	github.com/hdt3213/godis v1.2.8
	github.com/julienschmidt/httprouter v1.3.0
	github.com/k3s-io/k3s v1.25.3-0.20220922032136-53c268d8eb90
	github.com/kubernetes-sigs/cri-tools v0.0.0
	github.com/leaanthony/slicer v1.6.0
	github.com/nats-io/jwt/v2 v2.2.1-0.20220113022732-58e87895b296
	github.com/nsqio/nsq v1.2.2-0.20220923154746-393f7b994f0c
	github.com/panjf2000/gnet v1.6.0
	github.com/reugn/go-streams v0.7.0
	github.com/rpcxio/basalt v0.0.0-20200813040018-a7e186f1ae4c
	github.com/rpcxio/did v0.0.0-20200223142455-7b9772571fae
	github.com/rpcxio/rpcx-examples v1.1.7-0.20220606035200-e356ac23accc
	github.com/rpcxio/rpcx-gateway v0.0.0-20220611140137-f4a622048e82
	github.com/smallnest/rpcx v1.7.11
	github.com/syndtr/goleveldb v1.0.1-0.20220721030215-126854af5e6d
	go.mongodb.org/mongo-driver v1.10.0-prerelease.0.20221010212311-cb7150a53b0a
	k8s.io/mount-utils v0.25.2
	github.com/olivere/elastic/v7 v7.0.32
	github.com/hashicorp/raft v1.3.11
	github.com/rqlite/rqlite v0.0.0-20221021022757-ad7f383d8386
)

replace (
	// k8s.io/mount-utils => k8s.io/mount-utils v0.20.4
	github.com/Mirantis/cri-dockerd => github.com/Mirantis/cri-dockerd v0.2.6

	github.com/abourget/teamcity => github.com/cockroachdb/teamcity v0.0.0-20180905144921-8ca25c33eb11
	github.com/kubernetes-sigs/cri-tools => github.com/kubernetes-sigs/cri-tools v1.25.0
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

	vitess.io/vitess => github.com/cockroachdb/vitess v0.0.0-20210218160543-54524729cc82
)
