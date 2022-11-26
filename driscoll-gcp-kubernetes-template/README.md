## Local set-up
### Download and start minikube
sudo pacman -S minikube
minikube start

### Point docker to minikube instance
eval $(minikube docker-env)
docker images // verification only
docker build -t tag-name ./Dockerfile
// Reference: dzone.com/articles/running-local-docker-images-in-kubernetes-1
