# kafka-test
testing out kafka with go

## zookeeper setup
bin/zookeeper-server-start.sh config/zookeeper.properties


# Simple block diagram of kafka cluster
![alt text](https://www.sohamkamani.com/e0b11d2ab5b62a78cb7fd26d3a9d279b/basic_arch.svg "Block_Diagram")

Each block(specified below) in the diagram can be a different system/service on the network

-Producer-- can be different service
-Consumer-- can be different service
-Brokers-- can be different service

## Zookeeper
Zookeper is used by kafka to maintain the state between the nodes in the cluster

## Kafka brokers
Brokers store and emit data. They are the pipes in the data pipeline

## Producers
Producers insert data into the kafka cluster

## Consumers
Consumers read data sent from brokers

## config/zookeeper.properties
Check port zookeeper runs on. Default :2181

## config/server.properties-- 3 unique properties for each broker
1. broker.id = 0 (broker ID)
2. listeners = PLAINTEXT://localhost:9093 (PORT Number)
3. log.dirs = /tmp/kafka-logs

## Creating a topic
Topics are groups of paritions. At least one, or multiple partitions can exist on a topic

bin/kafka-topics.sh --create --topic [topic-name] --zookeeper localhost:2181 --partitions [X] --replication-factor [N]

**If N brokers are down, consumers will no longer be able to receive messages**

## Creating a producer
bin/kafka-console-producer.sh --broker-list [list of PORTS ex- localhost:9093,localhost:9094,localhost:9095] --topic [topic-name]

## Creating a consumer
**Bootstrap server must be any one of the brokers in the cluster

bin/kafka-console-consumer.sh --bootstrap-server [localhost:PORT] --topic [topic-name] --from-beginning [optional] --group [group-name]

EX:
bin/kafka-console-consumer.sh --bootstrap-server localhost:9097 --topic my-kafka-topic --from-beginning 

