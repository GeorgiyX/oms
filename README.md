# What is it

This is the repository with the ordering system services. The services are written as homework of the Route 256 course from Ozone Tech. Written on Go 1.19. Using gRPC and Kafka for communication, postgresql as storage, jaeger for tracing, prometheus and grafana for metrics and alerts. Services can be started via docker-compose.



# Directories

* `checkout/` - customer service for managing the shopping cart and placing orders. Linked by gRPC to product_service (external service) and loms.
* `loms/` - internal order management service. Linked to checkout via gRPC and notifications via Kafka.
* `notifications/` - order notification service. Just reads messages from Kafka topic.
* `libs/` - implementations of various tools for services
    * `cache` -  in-memory cache implementation
    * `config` - help parse yaml config to struct
    * `cron` - lib for periodical tasks
    * `db` -  wrapper over db interface with transactions support
    * `errors` - errors handling
    * `grpc` - gRPC interceptors for client and server
    * `httpaux` - wrappers over http client and server
    * `kafka` - kafka producer and consumer
    * `middleware` - logger and ratelimiter middleware
    * `ratelimiter` - ratelimiter lib
    * `validator` - singleton and checker for `github.com/go-playground/validator/v10`
    * `workerpool` - implemplementation of goroutine worker pool



# Run

```
# lint
make precommit
# test
make test
# run
make run-all
```

