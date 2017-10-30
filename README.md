# nsq-goclient
Go client example for nsq

## Setup
1. Create docker containers
```
docker-compose up -d --build
```

2. The examples use `go-nsq` so get go get it

```
go get -u -v github.com/bitly/go-nsq
```

3. Start the consumer in a separate terminal, which waits for messages to be published

```
go run command/consumer_example.go
```

4. Start the producer in a separate terminal, which will create 1000 messages

```
go run command/producer.go
```
