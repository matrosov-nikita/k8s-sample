### Run

Create pods using Deployments:

```sh
$ kubectl create -f deployments/frontend-deployment.yaml
$ kubectl create -f deployments/backend-deployment.yaml
```

Create Services for pods:

```sh
$ kubectl create -f deployments/svc-frontend-lb.yaml
$ kubectl create -f deployments/svc-backend-lb.yaml
```

To run locally in minikube:

```sh
$ minikube start
$ minikube service frontend-lb
```


