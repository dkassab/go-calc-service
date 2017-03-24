# go-calc-service
This is a simple example to show how to use gRPC and gRPC gateway.
* If you want to dockerize the services yourself, make sure to run `glide up` to get all the dependencies before running `docker build`
* To run both services, use `docker-compose up`
* To run the services on Kubernetes, you need to create the deployments, services and the ingress based on the files in the Kubernetes folder. If you're running Kubernetes locally, you can curl to the service using: curl -X POST 192.168.205.11/v1/example/add/2/6 -H 'Host: mydomain.com'
* I run Kubernetes locally using KET: https://github.com/apprenda/kismatic/. The ingress node was added using KET as well.
