# goKafka

### Introduction
http requests are captured via fluentd and POST data is sent to a kafka topic.

### How to run
##### Setup
* dc build
* dc compose up -d
* dc exec broker_one bash
* kafka-topics --create --topic http_logs --bootstrap-server broker_one:9092
* 

##### Monitor kafka topic
* dc exec broker_one bash
* kafka-console-consumer --topic http_logs --from-beginning --bootstrap-server broker_one:9092

##### Send http request
curl -X POST -d 'parikshit, India' http://127.0.0.1:9880/source.http
