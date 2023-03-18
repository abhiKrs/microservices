
## Setup Flink Using Operator

```ps
# Deploy cert manager and wait for it to become active
kubectl create -f https://github.com/jetstack/cert-manager/releases/download/v1.8.2/cert-manager.yaml

helm repo add flink-operator-repo https://downloads.apache.org/flink/flink-kubernetes-operator-1.4.0/
helm repo update

helm install flink-kubernetes-operator flink-operator-repo/flink-kubernetes-operator --namespace logfire-local
```

## Setup Filter Cluster

```ps
# Install sbt by following the link
# https://www.scala-sbt.org/download.html

(cd scala && sbt assemblyPackageDependency)
(cd scala && sbt assembly)

docker build -t logfire/flink -f scala/flink.Dockerfile scala/

# For local building
docker build -t logfire/flink -f scala/flink.dev.Dockerfile scala/

kubectl apply -f scala/flink-session-cluster.yaml -n logfire-local
```

## Setup schema registry

```
helm upgrade --install cp-schema-registry kafka/cp-schema-registry --namespace logfire-local --set kafka.external=true --set kafka.address=my-cluster-kafka-brokers:9092
```

## Setup flink service
```
docker build -t logfire/filter-service -f scala/Dockerfile scala/

# For local building
docker build -t logfire/filter-service -f scala/filterService.dev.Dockerfile scala/


helm upgrade --install filter-service helm-chart/filter-service --namespace logfire-local --set redis.host=redis-master  --set flinkCluster.host=flink-flink-chart-flink-session-deployment-rest --set image.repository=dockercgs/filter-service

# SETUP FOR ROHAN ONLY
helm upgrade --install filter-service helm-chart/filter-service --namespace logfire-local --values scala/helm/filter-service/values-rohan.yaml --set image.repository=logfire/filter-service
```

## Setup redis
```
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update

helm upgrade --install redis bitnami/redis --set auth.enabled=false --namespace logfire-local
#redis-master
export REDIS_PASSWORD=$(kubectl get secret --namespace izac redis -o jsonpath="{.data.redis-password}" | base64 -d)
REDISCLI_AUTH="$REDIS_PASSWORD" redis-cli -h 127.0.0.1 -p 6379
```

## Run test cases
```
sbt "project filterJob" test
```

## Debugging
```
kubectl run debugging --image=prohankumar/debugging:1 -n logfire-local

# Consume an avro topic
kcat -b my-cluster-kafka-brokers:9092 -C -t test -s avro -r http://cp-schema-registry:8081

# Send a filter request
grpcurl -plaintext -d '{"sources":[{"sourceID":"test"}]}' filter-service:80 sh.logfire.FlinkService/SubmitFilterRequest
```


## Testing
```
export JMX_PORT=5557
kafka-avro-console-producer \
 --broker-list my-cluster-kafka-brokers:9092 \
 --topic test  \
 --property schema.registry.url=http://localhost:8081 \
 --property value.schema='{"type":"record","name":"record","fields":[{"name":"platform","type":"string"},{"name":"message","type":"string"},{"name":"msgid","type":"string"},{"name":"pid","type":"int"},{"name":"level","type":"string"},{"name":"dt","type":"string"},{"name":"version","type":"int"}]}'

```