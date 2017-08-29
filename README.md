# Golang-Microservice.
```A simple microservice written in Go with sweetness of Docker Containers & Kubernetes Orchestration.```

__A small weekend exercise for personal practice.__

Includes:
 
 #### Creating a simple go microservice.
    
    - Implementing a simple *http server*.
    - JSON *Marshalling & Unmarshalling* with query parameters.
    - Unit tests for test microservice.
    - API tests with **POSTMAN**.
    - Implementing HTTP verbs, _GET | POST | PUT | DELETE_.

 #### Containerizing Go Microservice on __Docker__.

    - building Docker-Images then running/improving/maintaining docker containers.
    - pushing the docker image on docker-hub.
    - with variations of Dockerfile, (generic one & more optimized version)
    - using docker-compose, ( w/ separate componenets: app + nginx)
    - docker commands for _CPU & Mem contraints_ and more docker CLI experiments.
 
 ###% Deploying & Scaling Go microservice on __Kubernetes__.

    - Learning basic architecture of kubernetes. 
    - Using Kubernetes CLI with Kubernetes cluster. (kubectl, minikube)
    - Deploy Go microservice using images on docker hub.
    - Implementing Deployment & Service descriptors.
    - Horizonatally scaling Deployments with replicas and Rolling Updates.
