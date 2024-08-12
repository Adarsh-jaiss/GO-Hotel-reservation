#

1. setup the namespace : `kubectl config set-context --current --namespace=default`
2. Deploy the container and service :

```bash
#pass the .env file here
kubectl create configmap my-env-config --from-env-file=.env
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml

kubectl get deployemmt
kubectl get pods
kubectl get services

kubectl port-forward svc/book-my-hotel-service 8080:80
```

3. Deleting the deployemts

```bash

kubectl get deployments -n default 
kubectl delete deployment <deployemetn-name>

kubectl get services
kubectl delete services <service-name>



