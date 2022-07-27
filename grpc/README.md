# hptx-samples

Simply, you can build the environment by running the `make run` command. then, you can test by `curl http://localhost:8003/createSoRollback`, check whether the database data is consistent.

## benchmark
```shell
ab -t 10 -c 5 localhost:8003/createSoCommit
```

## Specific

There are three microservices, aggregation_svc、order_svc、product_svc, they all communicate with grpc. And, there is also a client used to initiate grpc requests.

We implemented grpc interceptor to start global transactions, commit and rollback global transactions, we can use `hptxGrpc.GlobalTransactionInterceptor(...)` instead of `tm.Implement(svc.ProxySvc)`.

https://github.com/CECTC/hptx-samples/blob/8a09f7a28d96be049b03239a09854ac871c18486/grpc/aggregation_svc/main.go#L48
