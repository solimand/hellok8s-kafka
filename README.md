# helloworld-k8s
Tutorials and exercises with Kubernetes

## Deployment
We describe different exercises

1. ConfigSet with ```hellogo``` container
    - ```docker build -t localhub/localhellogo .\hellogo\```
    - ```kubectl apply -f .\hellogo\auto-configMap.yaml```
    - ```kubectl apply -f .\auto-deployment.yaml```
    - ```kubectl apply -f .\yaml\first-service.yaml```
    - ```kubectl port-forward service/demo 9999:8888```
