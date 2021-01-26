# CNAB Operator

🚨 This repository has been archived. Our operator efforts are being continued in https://github.com/getporter/operator 🚨

## Running locally

### Prerequisites

- [`operator-sdk`][operator_sdk_tool]
- [`dep`][dep_tool] version v0.5.0+.
- [`git`][git_tool]
- [`go`][go_tool] version v1.10+.
- [`docker`][docker_tool] version 17.03+.
- [`kubectl`][kubectl_tool] version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### Building and running

For now, you can run a local OCI registry, push a CNAB bundle using a tool that implements the `coras` package ([`coras`][coras] or this branch of [`duffle`][duffle_branch]).


You must first create the [Custom Resource Definition for a CNAB Bundle][crd] (this is a placeholder that only contain the name and URL you can find the bundle in an OCI registry), then create an [instance of the Bundle CRD][cr] with the remote bundle you pushed. At this point, the controller will pull the bundle periodically and, for now, display it.

```
$ operator-sdk up local
INFO[0000] Running the operator locally.
INFO[0000] Using namespace cnab-operator.
{"level":"info","ts":1556804757.218317,"logger":"cmd","msg":"Go Version: go1.12.1"}
{"level":"info","ts":1556804757.218915,"logger":"cmd","msg":"Go OS/Arch: darwin/amd64"}
{"level":"info","ts":1556804757.218924,"logger":"cmd","msg":"Version of operator-sdk: v0.7.0+git"}
{"level":"info","ts":1556804757.224808,"logger":"leader","msg":"Trying to become the leader."}
{"level":"info","ts":1556804757.22487,"logger":"leader","msg":"Skipping leader election; not running in a cluster."}
{"level":"info","ts":1556804759.023943,"logger":"cmd","msg":"Registering Components."}
{"level":"info","ts":1556804759.02464,"logger":"kubebuilder.controller","msg":"Starting EventSource","controller":"bundle-controller","source":"kind source: /, Kind="}
{"level":"info","ts":1556804759.025194,"logger":"kubebuilder.controller","msg":"Starting EventSource","controller":"bundle-controller","source":"kind source: /, Kind="}
{"level":"info","ts":1556804760.586217,"logger":"cmd","msg":"failed to initialize service object for metrics: OPERATOR_NAME must be set"}
{"level":"info","ts":1556804760.586272,"logger":"cmd","msg":"Starting the Cmd."}
{"level":"info","ts":1556804760.69733,"logger":"kubebuilder.controller","msg":"Starting Controller","controller":"bundle-controller"}
{"level":"info","ts":1556804760.8028731,"logger":"kubebuilder.controller","msg":"Starting workers","controller":"bundle-controller","worker count":1}
{"level":"info","ts":1556804760.803205,"logger":"controller_bundle","msg":"Reconciling Bundle","Request.Namespace":"cnab-operator","Request.Name":"example-bundle"}
{"level":"info","ts":1556804760.803304,"logger":"controller_bundle","msg":"Reconciliating for bundle: testing-bundle from localhost:5000/cnab-operator-tetst:v1"}
WARN[0003] reference for unknown type: application/vnd.cnab.bundle.thin.v1-wd+json  digest="sha256:7882e4c8af49e3f13575bc7dbc41f442fb90f3bd86a4f3327649a937912f4d8a" mediatype=application/vnd.cnab.bundle.thin.v1-wd+json size=1071
WARN[0003] unknown type: application/vnd.oci.image.config.v1+json
WARN[0003] encountered unknown type application/vnd.cnab.bundle.thin.v1-wd+json; children may not be fetched
{"level":"info","ts":1556804760.8466072,"logger":"controller_bundle","msg":"Pulled bundle from OCI registry: {\"name\":\"helloworld-testdata\",\"version\":\"0.1.2\",\"description\":\"An example 'thin' helloworld Cloud-Native Application Bundle\",\"maintainers\":[{\"name\":\"Matt Butcher\",\"email\":\"matt.butcher@microsoft.com\",\"url\":\"https://example.com\"}],\"invocationImages\":[{\"imageType\":\"\",\"image\":\"localhost:5000/cnab-operator-tetst:cnabregistry.azurecr.io-cnab-meeting-docker.io-library-alpine-latest\",\"originalImage\":\"cnabregistry.azurecr.io/cnab-meeting:docker.io-library-alpine-latest\",\"digest\":\"sha256:5c40b3c27b9f13c873fefb2139765c56ce97fd50230f1f2d5c91e55dec171907\"}],\"images\":{\"my-microservice\":{\"imageType\":\"\",\"image\":\"localhost:5000/cnab-operator-tetst:cnabregistry.azurecr.io-cnab-meeting-docker.io-library-ubuntu-latest\",\"originalImage\":\"cnabregistry.azurecr.io/cnab-meeting:docker.io-library-ubuntu-latest\",\"digest\":\"sha256:ab6cb8de3ad7bb33e2534677f865008535427390b117d7939193f8d1a6613e34\",\"description\":\"my microservice\"}},\"parameters\":{\"backend_port\":{\"type\":\"integer\",\"destination\":{\"env\":\"BACKEND_PORT\"}}},\"credentials\":{\"hostkey\":{\"path\":\"/etc/hostkey.txt\",\"env\":\"HOST_KEY\"}}}"}
```

## Running in a Kubernetes cluster

WIP it. WIP it good. 😎

[operator_sdk_tool]: https://github.com/operator-framework/operator-sdk
[dep_tool]:https://golang.github.io/dep/docs/installation.html
[git_tool]:https://git-scm.com/downloads
[go_tool]:https://golang.org/dl/
[docker_tool]:https://docs.docker.com/install/
[kubectl_tool]:https://kubernetes.io/docs/tasks/tools/install-kubectl/

[coras]: https://github.com/radu-matei/coras
[duffle_branch]: https://github.com/radu-matei/duffle/tree/push-coras
[crd]: deploy/crds/cnab_v1alpha1_bundle_crd.yaml
[cr]: deploy/crds/cnab_v1alpha1_bundle_cr.yaml
