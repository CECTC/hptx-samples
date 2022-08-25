# hptx-samples

Simply, you can build the environment by running the `make run` command. then, you can test by `curl http://localhost:8003/createSoRollback`, check whether the database data is consistent.

## benchmark
```shell
ab -t 10 -c 5 localhost:8003/createSoCommit
```

## Specific

We implemented gin middleware to start global transactions, commit and rollback global transactions, we can use `hptxGin.GlobalTransaction(60000)` instead of `tm.Implement(svc.ProxySvc)`.

https://github.com/CECTC/hptx-samples/blob/feba75dfc46dd5ec4220f018b1dc91e2595c7543/http_gin_middleware/aggregation_svc/main.go#L19
