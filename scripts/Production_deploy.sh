#!/bin/bash
##Tagging Webapi to latest and push to docker hub
docker image pull docker.io/dockercgs/goweb-api:$1
docker image tag dockercgs/goweb-api:$1 dockercgs/goweb-api:latest
docker push dockercgs/goweb-api:latest
##Tagging Notification service to latest and push to docker hub
docker image pull docker.io/dockercgs/notification:$1
docker image tag dockercgs/noification:$1 dockercgs/notification:latest
docker push dockercgs/notification:latest
##Tagging livetail service to latest and push to docker hub
docker image pull docker.io/dockercgs/livetail:$1
docker image tag dockercgs/livetail:$1 dockercgs/livetail:latest
docker push dockercgs/livetail:latest
##Tagging filter service to latest and push to docker hub
docker image pull docker.io/dockercgs/filter-service:$1
docker image tag dockercgs/filter-service:$1 dockercgs/filter-service:latest
docker push dockercgs/filter-service:latest

##Deploy the application
helm upgrade webapp /opt/microservices/helm-chart/webapp-prodchart/ -n logfire-local
helm upgrade --install filter-service /opt/microservices/helm-chart/filter-service/ --namespace logfire-local --set redis.host=redis-master  --set flinkCluster.host=flink-chart-flink-session-deployment-rest --set image.repository=dockercgs/filter-service:latest
helm upgrade --install flink-chart /opt/microservices/helm-chart/flink-chart/ --namespace logfire-local --set flink.image.repository=dockercgs/flink --set flink.image.tag=latest
helm repo add flink-operator-repo https://downloads.apache.org/flink/flink-kubernetes-operator-1.4.0/
helm repo update
helm upgrade --install cp-schema-registry /opt/microservices/helm-chart/cp-schema-registry/ --namespace logfire-local --set kafka.external=true --set kafka.address=my-cluster-kafka-brokers:9092
helm upgrade --install flink-kubernetes-operator flink-operator-repo/flink-kubernetes-operator --namespace logfire-local

