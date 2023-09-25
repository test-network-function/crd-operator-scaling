![Project frozen](https://img.shields.io/badge/status-frozen-blue.png) ![Project unmaintained](https://img.shields.io/badge/project-unmaintained-red.svg)
Replaced by [cr-scale-operator](https://github.com/test-network-function/cr-scale-operator)

<h1> DEPRECATED </h1>
# new-pro
// TODO(user): Add simple overview of use/purpose

## Description
// TODO(user): An in-depth paragraph about your project and overview of use

## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster
1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

2. Build and push your image to the location specified by `IMG`:
	
```sh
make docker-build docker-push IMG=<some-registry>/new-pro:tag
```
	
3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/new-pro:tag
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller
UnDeploy the controller to the cluster:

```sh
make undeploy
```

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

### How it works
This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/) 
which provides a reconcile function responsible for synchronizing resources untile the desired state is reached on the cluster 

### Test It Out
1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.


# amals steps
# operator-scaling
project for creating a crd that managed a deploymrnt that project is too simple and there is alot pf hard coded things - we can change on it 
refernce for creating the operator 
https://betterprogramming.pub/build-a-kubernetes-operator-in-10-minutes-11eec1492d30
steps that need to do to run it : 

other useful refernce 
https://medium.com/@thescott111/autoscaling-kubernetes-custom-resource-using-the-hpa-957d00bb7993


1. run 
```sh 
make manifests
```
2. this is to run the controller
2.1  
```sh 
make install 
```
- can see the crd with that command  kubectl get crds
2.2 
```sh 
make run
```
 - its need to be up on the background 
on other terminal 
3. 
```sh 
kubectl apply -f config/samples  --validate=false
```
to run it in our test need to add the crd to the tnf_config.yaml.
