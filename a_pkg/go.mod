module a_pkg

go 1.24.3

require (
	bazil.org/fuse v0.0.0-20230120002735-62a210ff1fd5
	dubbo.apache.org/dubbo-go/v3 v3.0.2
	github.com/FerretDB/FerretDB v1.3.0
	github.com/argoproj/argo-cd v1.8.7
	//github.com/Mirantis/cri-dockerd latest
	github.com/beego/bee/v2 v2.3.0
	github.com/beego/beego/v2 v2.3.3
	github.com/blend/go-sdk v1.20220411.4-0.20220608163541-be8198318397
	github.com/caarlos0/starcharts v1.6.5-0.20221003175303-951f7eea37bf
	github.com/carlmjohnson/requests v0.24.3
	github.com/cilium/cilium v1.16.4
	github.com/cloudwego/hertz v0.9.3
	github.com/cloudwego/kitex v0.11.3
	github.com/cockroachdb/cockroach v0.0.0-20250408171622-f3e357c1ca10
	github.com/cockroachdb/pebble/v2 v2.0.3
	github.com/codegangsta/cli v1.22.16
	github.com/compose-spec/compose-go/v2 v2.6.2
	github.com/containerd/cgroups/v3 v3.0.5
	github.com/containerd/containerd v1.7.24
	github.com/containerd/containerd/v2 v2.1.0
	github.com/containerd/nerdctl v1.7.5
	github.com/containernetworking/cni v1.3.0
	github.com/containers/buildah v1.39.4
	github.com/containers/fetchit v0.0.2-0.20221010155311-f5ffd62ecd0c
	github.com/containers/image/v5 v5.34.3
	github.com/containers/podman/v4 v4.9.4
	github.com/containers/podman/v5 v5.4.2
	github.com/containers/skopeo v1.12.0
	github.com/coocood/freecache v1.2.4
	github.com/cri-o/cri-o v1.32.5
	github.com/dgraph-io/badger/v3 v3.2103.5
	github.com/distribution/distribution/v3 v3.0.0
	github.com/docker-slim/docker-slim v0.0.0-20221114191106-a4aca0c03e46
	github.com/docker/buildx v0.23.0
	github.com/docker/cli v28.1.1+incompatible
	github.com/docker/compose/v2 v2.36.0
	github.com/docker/docker v28.1.1+incompatible
	github.com/docker/machine v0.16.2
	github.com/docker/swarm v1.2.9
	// github.com/OpenIMSDK/Open-IM-Server/v2 a4fe45d
	// github.com/LockGit/gochat v0.0.0-20221024110244-4284a5a8971e
	github.com/douyu/jupiter v0.9.0
	github.com/dromara/carbon/v2 v2.5.3
	github.com/elastic/go-elasticsearch/v7 v7.17.10
	github.com/emitter-io/emitter v0.0.0-20240202195010-996fa4af0e95 //v3.1
	github.com/flower-corp/lotusdb v1.0.1-0.20220729120541-8d44d1d6c0ab
	github.com/flower-corp/rosedb v1.1.2-0.20220817135557-31bf5acf7397
	github.com/fluxcd/flux2/v2 v2.2.3
	github.com/gin-gonic/gin v1.10.0
	github.com/go-co-op/gocron v1.30.1
	github.com/go-kit/kit v0.13.1-0.20231222231659-844c3d2de01a
	github.com/go-nunu/nunu v1.0.4
	github.com/go-redsync/redsync v1.4.2
	github.com/gofiber/fiber/v2 v2.52.5
	github.com/goharbor/harbor/src v0.0.0-20240403034559-9778176ff1ee
	github.com/gohugoio/hugo v0.145.0
	github.com/golang-jwt/jwt/v4 v4.5.2
	github.com/golang-jwt/jwt/v5 v5.2.2
	github.com/golang-module/dongle v0.2.4
	github.com/goodrain/rainbond v0.0.0-20241203084856-572501181669
	github.com/google/btree v1.1.3 // indirect
	github.com/google/pprof v0.0.0-20250501235452-c0086092b71a
	github.com/gookit/config/v2 v2.1.6
	github.com/gookit/filter v1.1.3
	github.com/gookit/gcli/v3 v3.0.7-0.20220904045221-3d6cf1f66110
	github.com/gookit/rux v1.3.4
	github.com/gookit/validate v1.4.4
	github.com/hashicorp/consul v1.20.5
	github.com/hashicorp/nomad v1.10.0
	github.com/hashicorp/packer v1.10.1
	github.com/hashicorp/raft v1.7.3
	github.com/hashicorp/vault v1.15.6
	github.com/hdt3213/godis v1.2.8
	github.com/influxdata/flux v0.194.3
	github.com/influxdata/influxdb/v2 v2.5.0
	github.com/influxdata/telegraf v1.28.1
	github.com/ipfs/kubo v0.26.0
	github.com/jmoiron/sqlx v1.4.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/k3s-io/k3s v1.33.1-0.20250430173654-63ab8e534cdf
	github.com/karmada-io/karmada v1.13.1
	github.com/kata-containers/kata-containers v0.0.0-20221110134924-56641bc2303c
	github.com/kataras/iris/v12 v12.2.11
	github.com/kataras/muxie v1.1.2
	github.com/kedacore/keda/v2 v2.13.1
	github.com/koderover/zadig v1.18.0
	github.com/kubecube-io/kubecube v1.5.1
	github.com/kubernetes-sigs/cri-tools v0.0.0
	github.com/labstack/echo/v4 v4.12.0
	github.com/leaanthony/slicer v1.6.0
	github.com/lima-vm/lima v0.13.1-0.20221108071032-4a4cc62749c1
	github.com/linkerd/linkerd2 v0.5.1-0.20221110190454-75cbed2469cc
	github.com/lxc/lxd v0.0.0-20221110184730-a10417833a8e
	github.com/markbates/goth v1.74.2
	github.com/minio/md5-simd v1.1.2
	github.com/minio/minio v0.0.0-20250228193308-11507d46da0c
	github.com/moby/buildkit v0.21.1
	github.com/moby/swarmkit/v2 v2.0.0-20250103191802-8c1959736554
	github.com/monaco-io/request v1.0.16
	github.com/nats-io/jwt/v2 v2.7.3
	github.com/nats-io/nats-server/v2 v2.10.27
	github.com/no-src/gofs v0.7.1
	github.com/nsqio/nsq v1.2.2-0.20220923154746-393f7b994f0c
	github.com/oam-dev/kubevela v1.9.9
	github.com/olivere/elastic/v7 v7.0.32
	github.com/opencontainers/runc v1.2.6
	github.com/opencontainers/runtime-tools v0.9.1-0.20241108202711-f7e3563b0271
	github.com/opencontainers/umoci v0.4.8-0.20221029002459-fb2db51251ac
	github.com/opensearch-project/opensearch-go/v2 v2.3.0
	github.com/panjf2000/gnet v1.6.0
	github.com/pingcap/tidb v1.1.0-beta.0.20250116073434-fea86c8e35ad
	github.com/pingcap/tiflow v0.0.0-20221107160750-0ad56ca6594a
	github.com/polarismesh/polaris v1.17.2-0.20230531031309-90b39c8785f4
	//github.com/rancher/k3s v1.33.1-0.20250430173654-63ab8e534cdf
	github.com/rancher/rancher v0.0.0-20240226212747-4119767ddf14
	github.com/rancher/rke v1.5.5
	github.com/rancher/rke2 v1.25.4-0.20221102171456-ddb87a9b6684
	github.com/reugn/go-streams v0.7.0
	github.com/revel/revel v1.1.0
	github.com/rezmoss/axios4go v0.6.2
	github.com/rkt/rkt v1.30.1-0.20200224141603-171c416fac02
	github.com/rook/rook v1.13.7
	github.com/rpcxio/basalt v0.0.0-20200813040018-a7e186f1ae4c
	github.com/rpcxio/did v0.0.0-20200223142455-7b9772571fae
	github.com/rpcxio/rpcx-examples v1.1.7-0.20220606035200-e356ac23accc
	github.com/rpcxio/rpcx-gateway v0.0.0-20220611140137-f4a622048e82
	github.com/rqlite/rqlite/v8 v8.36.18
	github.com/sirupsen/logrus v1.9.3
	github.com/smallnest/rpcx v1.7.11
	github.com/sqlc-dev/sqlc v1.28.0
	github.com/stretchr/testify v1.10.0
	github.com/syncthing/syncthing v1.23.5
	github.com/syndtr/goleveldb v1.0.1-0.20220721030215-126854af5e6d
	github.com/tinode/chat v0.20.1-0.20221011125331-85ecad89d2db
	github.com/unionj-cloud/go-doudou/v2 v2.3.0
	github.com/volatiletech/authboss/v3 v3.2.1
	github.com/volatiletech/sqlboiler/v4 v4.13.0
	github.com/xuri/excelize/v2 v2.8.1
	go-micro.dev/v5 v5.3.0
	go.mongodb.org/mongo-driver v1.17.3
	golang.org/x/oauth2 v0.30.0 // indirect
	golang.org/x/term v0.32.0 // indirect
	golang.org/x/text v0.25.0
	helm.sh/helm/v3 v3.17.3
	k8s.io/minikube v1.28.1-0.20221202001624-44ea7f43205e
	k8s.io/mount-utils v0.32.4
	kubesphere.io/kubesphere v0.0.0-20250303081139-2ed83e77d41c //v0.0.0-20221102022257-ff59b80c2807
	sigs.k8s.io/kubefed v0.10.0
	vitess.io/vitess v0.10.3-0.20250311021806-209d940ad923
)

require (
	github.com/Workiva/go-datastructures v1.1.5
	github.com/blevesearch/bleve/v2 v2.4.2
	github.com/bombsimon/logrusr/v4 v4.1.0
	github.com/cilium/ebpf v0.17.3
	github.com/containerd/containerd/api v1.9.0
	github.com/containerd/continuity v0.4.5
	github.com/containerd/go-cni v1.1.12
	github.com/containerd/platforms v1.0.0-rc.1
	github.com/containerd/ttrpc v1.2.7
	github.com/containerd/typeurl/v2 v2.2.3
	github.com/containernetworking/plugins v1.7.1
	github.com/containers/common v0.62.3
	github.com/containers/ocicrypt v1.2.1
	github.com/containers/storage v1.58.0
	github.com/coreos/go-systemd/v22 v22.5.1-0.20231103132048-7d375ecc2b09
	github.com/cpuguy83/go-md2man/v2 v2.0.7
	github.com/cyphar/filepath-securejoin v0.4.1
	github.com/daichi-m/go18ds v1.12.1
	github.com/dchest/captcha v1.0.0
	github.com/dgraph-io/ristretto v0.1.1 // indirect
	github.com/emicklei/go-restful/v3 v3.12.2
	github.com/emirpasic/gods v1.18.1
	github.com/go-jet/jet/v2 v2.13.0
	github.com/godbus/dbus/v5 v5.1.1-0.20241109141217-c266b19b28e9
	github.com/golang-migrate/migrate/v4 v4.18.2
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/hashicorp/consul-awsauth v0.0.0-20250130185352-0a5f57fe920a
	github.com/hashicorp/consul/api v1.32.0
	github.com/hashicorp/go-msgpack/v2 v2.1.3
	github.com/hashicorp/mdns v1.0.5
	github.com/hashicorp/memberlist v0.5.3
	github.com/hashicorp/serf v0.10.2
	github.com/ikeikeikeike/go-sitemap-generator/v2 v2.0.2
	github.com/iris-contrib/middleware/cors v0.0.0-20240502084239-34f27409ce72
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
	github.com/lionsoul2014/ip2region/binding/golang v0.0.0-20240510055607-89e20ab7b6c6
	github.com/mlogclub/simple v1.2.30
	github.com/moby/sys/capability v0.4.0
	github.com/moby/sys/mount v0.3.4
	github.com/moby/sys/mountinfo v0.7.2
	github.com/moby/sys/sequential v0.6.0
	github.com/moby/sys/signal v0.7.1
	github.com/moby/sys/user v0.4.0
	github.com/moby/term v0.5.2
	github.com/mtrmac/gpgme v0.1.2 // indirect
	github.com/onsi/ginkgo/v2 v2.23.4
	github.com/onsi/gomega v1.37.0
	github.com/opencontainers/image-spec v1.1.1
	github.com/opencontainers/runtime-spec v1.2.1
	github.com/rqlite/go-sqlite3 v1.38.0
	github.com/rqlite/sql v0.0.0-20241111133259-a4122fabb196
	github.com/secure-systems-lab/go-securesystemslib v0.9.0
	github.com/smartystreets/assertions v1.13.0 // indirect
	github.com/spf13/afero v1.14.0
	github.com/vbauerster/mpb/v5 v5.4.0 // indirect
	github.com/vektra/mockery/v2 v2.53.3
	golang.org/x/crypto v0.38.0
	golang.org/x/exp v0.0.0-20250408133849-7e4ce0ab07d0
	golang.org/x/net v0.40.0
	golang.org/x/sync v0.14.0
	golang.org/x/sys v0.33.0
	golang.org/x/tools v0.33.0
	google.golang.org/genproto v0.0.0-20250324211829-b45e905df463
	//github.com/harness/gitness/v2 v2.26.0
	//github.com/drone/drone/v2 latest
	modernc.org/cc/v4 v4.26.1
	modernc.org/ccgo/v4 v4.28.0
	modernc.org/ccorpus2 v1.5.2
	modernc.org/fileutil v1.3.1
	modernc.org/gc/v2 v2.6.5
	modernc.org/golex v1.1.0
	modernc.org/libc v1.65.3
	modernc.org/memory v1.10.0
	modernc.org/opt v0.1.4
	modernc.org/sortutil v1.2.1
	moul.io/http2curl v1.0.0 // indirect
// https://github.com/micro/plugins
)

require (
	github.com/hashicorp/cap v0.9.0
	github.com/hashicorp/cli v1.1.7
	github.com/hashicorp/consul/sdk v0.16.2
	github.com/hashicorp/go-bexpr v0.1.14
	github.com/hashicorp/go-connlimit v0.3.1
	github.com/hashicorp/go-discover v1.0.0
	github.com/hashicorp/go-envparse v0.1.0
	github.com/hashicorp/go-getter v1.7.8
	github.com/hashicorp/go-memdb v1.3.5
	github.com/hashicorp/go-plugin v1.6.3
	github.com/hashicorp/go-secure-stdlib/listenerutil v0.1.10
	github.com/hashicorp/go-set/v3 v3.0.0
	github.com/hashicorp/hcl v1.0.1-vault-7
	github.com/hashicorp/hcl/v2 v2.23.0
	github.com/hashicorp/net-rpc-msgpackrpc/v2 v2.0.1
	github.com/hashicorp/nomad/api v0.0.0-20250425125325-3e688cf9282f
	github.com/hashicorp/raft-boltdb/v2 v2.3.1
	github.com/hashicorp/vault/api v1.16.0
	github.com/k3s-io/kine v0.13.14
	go.etcd.io/bbolt v1.4.0
	go.etcd.io/etcd/v3 v3.6.0-alpha.0
)

require (
	github.com/KimMachineGun/automemlimit v0.7.1
	github.com/cncf/xds/go v0.0.0-20250326154945-ae57f3c0d45f
	github.com/digitalocean/godo v1.136.0
	github.com/envoyproxy/go-control-plane/envoy v1.32.4
	github.com/envoyproxy/protoc-gen-validate v1.2.1
	github.com/golang/protobuf v1.5.4
	github.com/gophercloud/gophercloud/v2 v2.6.0
	github.com/hetznercloud/hcloud-go/v2 v2.19.1
	github.com/ionos-cloud/sdk-go/v6 v6.3.2
	github.com/klauspost/compress v1.18.0
	github.com/knadh/koanf/maps v0.1.2 // indirect
	github.com/knadh/koanf/v2 v2.1.2 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor v0.124.1
	github.com/ovh/go-ovh v1.7.0
	github.com/prometheus/alertmanager v0.28.1
	github.com/prometheus/client_golang v1.22.0
	github.com/prometheus/exporter-toolkit v0.14.0
	github.com/prometheus/sigv4 v0.1.2
	go.opentelemetry.io/collector/component v1.30.0
	go.opentelemetry.io/collector/consumer v1.30.0
	go.opentelemetry.io/collector/pdata v1.30.0
	go.opentelemetry.io/collector/processor v1.30.0
	go.opentelemetry.io/collector/semconv v0.124.0
	go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace v0.60.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.60.0
	go.opentelemetry.io/otel v1.35.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.35.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.35.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.35.0
	go.opentelemetry.io/otel/metric v1.35.0
	go.opentelemetry.io/otel/sdk v1.35.0
	go.opentelemetry.io/otel/trace v1.35.0
	google.golang.org/api v0.228.0
	google.golang.org/genproto/googleapis/api v0.0.0-20250414145226-207652e42e2e
	google.golang.org/grpc v1.72.0
	google.golang.org/grpc/examples v0.0.0-20250513095349-b89909b7bd0d
	google.golang.org/protobuf v1.36.6
)

require (
	github.com/BurntSushi/toml v1.5.0
	github.com/Kong/kong v0.0.0-20250527125949-77aa98b7a876
	github.com/argoproj/argo-cd/v3 v3.0.4
	github.com/cheggaaa/pb/v3 v3.1.7
	github.com/concourse/concourse v4.2.3+incompatible
	github.com/go-sql-driver/mysql v1.9.1
	github.com/gocolly/colly/v2 v2.2.0
	github.com/google/go-containerregistry v0.20.5
	github.com/google/ko v0.18.0
	github.com/linuxkit/linuxkit v1.6.0
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/sigstore/cosign/v2 v2.5.0
	github.com/slimtoolkit/slim v0.0.0-20250520142417-a35f6b52ce53
	github.com/traefik/traefik/v3 v3.4.1
	//github.com/prometheus/prometheus v2.5.0+incompatible
	github.com/urfave/cli/v3 v3.3.3
	gopkg.in/urfave/cli.v1 v1.20.0
	gvisor.dev/gvisor v0.0.0-20250510022024-e4c059533a2a
)
require (
	github.com/gobuffalo/buffalo v1.1.2
	github.com/gobuffalo/packr/v2 v2.8.3
	github.com/gobuffalo/helpers v0.6.10
	github.com/gobuffalo/httptest v1.5.2
	github.com/gobuffalo/logger v1.0.7
	github.com/gobuffalo/meta v0.3.3
	github.com/gobuffalo/nulls v0.4.2
	github.com/gobuffalo/plush/v5 v5.0.5
	github.com/gobuffalo/pop/v6 v6.1.1
	github.com/gobuffalo/refresh v1.13.3
	github.com/gobuffalo/tags/v3 v3.1.4
)

require (
	github.com/kubeedge/kubeedge        latest
    github.com/fatedier/frp             latest
    github.com/harness/gitness/v3       latest
    github.com/dagger/dagger            latest
    github.com/abiosoft/colima          latest
    github.com/gogf/gf/v2               latest
    github.com/gogf/examples            latest
    github.com/milvus-io/milvus/v2      latest
    github.com/mudler/LocalAI/v2        latest
    github.com/canonical/lxd/v6         latest
    github.com/wanzo-mini/mini-rpc      latest
    github.com/wanzo-mini/mini-balancer latest
)

replace github.com/axiomhq/hyperloglog/000 => github.com/axiomhq/hyperloglog v0.2.5

replace vitess.io/vitess => vitess.io/vitess v0.10.3-0.20250311021806-209d940ad923

replace (
	//k8s.io/mount-utils => k8s.io/mount-utils v0.20.4
	github.com/Mirantis/cri-dockerd => github.com/Mirantis/cri-dockerd v0.2.6

	github.com/abourget/teamcity => github.com/cockroachdb/teamcity v0.0.0-20180905144921-8ca25c33eb11

	github.com/aquarapid/vaultlib => github.com/mch1307/vaultlib v0.6.1

	github.com/armon/go-metrics => github.com/hashicorp/go-metrics v0.5.4
	github.com/blend/go-sdk => github.com/blend/go-sdk v1.20220411.4-0.20220608163541-be8198318397

	github.com/cloudnativelabs/kube-router/v2 => github.com/cloudnativelabs/kube-router/v2 v2.5.0

	github.com/cockroachdb/gogoproto => github.com/gogo/protobuf v1.3.2
	github.com/codegangsta/cli => github.com/urfave/cli v1.22.14

	github.com/docker/docker => github.com/docker/docker v27.3.1+incompatible
	github.com/kubernetes-incubator/external-storage => github.com/kubernetes-incubator/external-storage v5.5.0+incompatible
	github.com/kubernetes-sigs/cri-tools => github.com/kubernetes-sigs/cri-tools v1.25.0

	github.com/mikefarah/yaml => github.com/go-yaml/yaml v0.0.0-20220527083530-f6f7691b1fde

	github.com/myesui/uuid => github.com/google/uuid v1.3.0
	github.com/opensearch-project/opensearch-go/v2 => github.com/opensearch-project/opensearch-go/v2 v2.1.0
	github.com/operator-framework/operator-lib => github.com/operator-framework/operator-lib v0.11.0
	github.com/pingcap/tidb/pkg/parser => github.com/pingcap/tidb/pkg/parser v0.0.0-20250523074228-78e306d37819
	github.com/polarismesh/go-restful-openapi => github.com/emicklei/go-restful-openapi/v2 v2.9.1
	github.com/projectcalico/api => github.com/projectcalico/api v0.0.0-20230602153125-fb7148692637
	github.com/projectcalico/calico => github.com/projectcalico/calico v2.6.12+incompatible

	github.com/rancher/rancher/pkg/apis => github.com/rancher/rancher/pkg/apis v0.0.0-20221103194918-2d114cccb546
	github.com/rancher/rancher/pkg/client => github.com/rancher/rancher/pkg/client v0.0.0-20221103194918-2d114cccb546

	github.com/sigstore/fulcio => github.com/sigstore/fulcio v1.0.1-0.20221110094946-7cf41b0898d4
	github.com/spegel-org/spegel => github.com/spegel-org/spegel v0.2.0

	go.etcd.io/etcd => github.com/etcd-io/etcd v3.3.27+incompatible

	go.qase.io/client => github.com/rancher/qase-go/client v0.0.0-20231114201952-65195ec001fa
	golang.org/x/time => golang.org/x/time v0.0.0-20220411224347-583f2d630306
	google.golang.org/grpc => google.golang.org/grpc v1.71.1
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
	k8s.io/dynamic-resource-allocation => k8s.io/dynamic-resource-allocation v0.33.1
	k8s.io/endpointslice => k8s.io/endpointslice v0.33.1
	k8s.io/externaljwt => k8s.io/externaljwt v0.33.1
	k8s.io/kms => k8s.io/kms v0.33.1
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
	kubesphere.io/utils => kubesphere.io/utils v0.0.0-20250208102917-f56781425fe8
	sigs.k8s.io/apiserver-network-proxy/konnectivity-client => sigs.k8s.io/apiserver-network-proxy/konnectivity-client v0.0.30
	sigs.k8s.io/cri-tools => sigs.k8s.io/cri-tools v1.33.0

)

