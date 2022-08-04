# hptx-samples
Simply, you can build the environment by running the `make run` command. then, you can test by `curl http://localhost:8003/createSoRollback`, check whether the database data is consistent.
## benchmark
```shell
ab -t 10 -c 5 localhost:8003/createSoCommit
```