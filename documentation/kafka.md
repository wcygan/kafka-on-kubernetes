# Kafka

## Installation

Run [install-strimzi.sh](../scripts/install-strimzi.sh)

```
./scripts/install-strimzi.sh
```

## Deployment

Run [deploy.sh](../scripts/deploy.sh)

```
./scripts/deploy.sh
```

Check the deployment with [check-kafka.sh](../scripts/check-kafka.sh)

```
./scripts/check-kafka.sh
```

You should wait until you see something like this:

```
NAME                                                   READY   STATUS    RESTARTS   AGE
pod/example-cluster-entity-operator-74bdb84f54-4hdbt   2/2     Running   0          86s
pod/example-cluster-kafka-0                            1/1     Running   0          108s
pod/example-cluster-kafka-1                            1/1     Running   0          108s
pod/example-cluster-kafka-2                            1/1     Running   0          108s
pod/example-cluster-zookeeper-0                        1/1     Running   0          2m10s
pod/example-cluster-zookeeper-1                        1/1     Running   0          2m10s
pod/example-cluster-zookeeper-2                        1/1     Running   0          2m10s
pod/strimzi-cluster-operator-7bb5468c59-t55rq          1/1     Running   0          20m

NAME                                       TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)                                        AGE
service/example-cluster-kafka-bootstrap    ClusterIP   10.111.31.67   <none>        9091/TCP,9092/TCP,9093/TCP                     108s
service/example-cluster-kafka-brokers      ClusterIP   None           <none>        9090/TCP,9091/TCP,8443/TCP,9092/TCP,9093/TCP   108s
service/example-cluster-zookeeper-client   ClusterIP   10.96.68.62    <none>        2181/TCP                                       2m11s
service/example-cluster-zookeeper-nodes    ClusterIP   None           <none>        2181/TCP,2888/TCP,3888/TCP                     2m11s

NAME                                              READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/example-cluster-entity-operator   1/1     1            1           86s
deployment.apps/strimzi-cluster-operator          1/1     1            1           20m

NAME                                                         DESIRED   CURRENT   READY   AGE
replicaset.apps/example-cluster-entity-operator-74bdb84f54   1         1         1       86s
replicaset.apps/strimzi-cluster-operator-7bb5468c59          1         1         1       20m
```

## Undeploy

Run [undeploy.sh](../scripts/undeploy.sh)

```
./scripts/undeploy.sh
```