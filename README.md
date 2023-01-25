# operator-scaling
project for creating a crd that managed a deploymrnt that project is too simple and there is alot pf hard coded things - we can change on it 
refernce for creating the operator 
https://betterprogramming.pub/build-a-kubernetes-operator-in-10-minutes-11eec1492d30
steps that need to do to run it : 
first cd to new-pro
1. run make manifests
2. this is to run the controller
2.1  make install - can see the crd with that command  kubectl get crds
2.2 make run - its need to be up on the background 
on other terminal 
3. kubectl apply -f config/samples  --validate=false
to run it in our test need to add the crd to the tnf_config.yaml.