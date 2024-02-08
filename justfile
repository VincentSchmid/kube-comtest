username := schmivin
version := "0.1.0"

producer-build:
    cd producer && \
    sudo docker build . -t {{username}}/comtest-producer:{{version}}

consumer-build:
    cd consumer && \
    sudo docker build . -t {{username}}/comtest-consumer:{{version}}

build-all: producer-build && consumer-build

push-images:
    docker push {{username}}/comtest-producer:{{version}}
    docker push {{username}}/comtest-consumer:{{version}}

kube-setup:
    -kubectl create namespace comtest

producer-kube:
    kubectl apply -f kubernetes/producer-deployment.yaml
    kubectl apply -f kubernetes/producer-service.yaml

consumer-kube:
    kubectl apply -f kubernetes/consumer-deployment.yaml

kube-destroy:
    -kubectl delete -f kubernetes/producer-deployment.yaml
    -kubectl delete -f kubernetes/producer-service.yaml
    -kubectl delete -f kubernetes/consumer-deployment.yaml
    -kubectl delete namespace comtest

kube-deploy: kube-setup producer-kube && consumer-kube
    kubectl wait --for=condition=available --timeout=10s deployment/producer -n comtest
