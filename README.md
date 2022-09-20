# Sample Application on Okteto

## Requirements

- [Install kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)
- Setup Okteto
    - Signup Okteto
    - [Install Okteto CLI](https://www.okteto.com/docs/getting-started/#installing-okteto-cli)
    - [Install Okteto CLI](https://www.okteto.com/docs/getting-started/#installing-okteto-cli)
    - [Setup kubeconfig](https://www.okteto.com/docs/cloud/credentials/)

## Usage

```sh
$ kubectl apply -f deploy.yaml
deployment.apps/sample-deploy created

# wait for depoying all pods
$ kubectl get pods -w
NAME                             READY   STATUS              RESTARTS   AGE
sample-deploy-7f99785569-4f4jl   0/1     ContainerCreating   0          3s
sample-deploy-7f99785569-7smz4   0/1     ContainerCreating   0          3s
sample-deploy-7f99785569-96467   0/1     ContainerCreating   0          3s
sample-deploy-7f99785569-7smz4   1/1     Running             0          5s
sample-deploy-7f99785569-4f4jl   1/1     Running             0          5s
sample-deploy-7f99785569-96467   1/1     Running             0          6s

$ kubectl apply service.yaml
NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
sample-svc   ClusterIP   10.152.25.196   <none>        443/TCP   2s

# get assigned hostname 
$ kubectl get ingress
NAME                CLASS   HOSTS                                         ADDRESS        PORTS     AGE
okteto-sample-svc   nginx   sample-svc-oriishitakahiro.cloud.okteto.net   35.225.69.73   80, 443   3m

# request ip address of reached pod
$ curl 'https://sample-svc-oriishitakahiro.cloud.okteto.net/ip'
```
