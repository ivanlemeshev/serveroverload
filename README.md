![CI](https://github.com/ivanlemeshev/serveroverload/actions/workflows/ci.yml/badge.svg)

# Server overload, rate limiting, load shedding examples

This repository contains examples of server overload scenarios and how to handle 
them. All the examples are not production-ready and are for educational purposes.
I do not guarantee that the examples are correct, there may be bugs.

## Run HTTP benchmarks

To run the benchmarks, you need to install
[bombardier](https://github.com/codesenberg/bombardier), a CLI tool that can be 
used to test the performance of HTTP services.

Here is the link to the documentation:
https://pkg.go.dev/github.com/codesenberg/bombardier

Then you can run the benchmarks to see how the service behaves under different 
loads.

We will use docker to run the service.

If you want to monitor the service resources during the benchmarks, run the 
following command in a terminal:

```bash
$ docker stats
```

Then run the service in a separate terminal window:

```bash
$ docker build \
    --progress=plain \
    --no-cache \
    -t service:latest \
    -f ./cmd/service/Dockerfile .
$ docker run -d --rm \
    --name service \
    --cpus="0.5" \
    --memory="500m" \
    --memory-swap="500m" \
    -p 8080:8080 \
    service:latest
```

The service will use 200MB of memory and 20% of the one CPU core.

You will see the statistics of the service in the first terminal window and can 
monitor its resources during the benchmarks.

```bash
CONTAINER ID   NAME      CPU %     MEM USAGE / LIMIT   MEM %     NET I/O     BLOCK I/O   PIDS
6a1a9e288a0c   service   0.00%     2.977MiB / 200MiB   1.49%     486B / 0B   0B / 0B     5
```

Run the benchmarks:

```bash
$ bombardier -c 50 -l -d 10s -r 50 http://127.0.0.1:8080/
$ bombardier -c 4000 -l -d 10s -r 4000 http://127.0.0.1:8080/
```

```bash
$ bombardier -c 50 -l -d 10s -r 50 http://127.0.0.1:8080/fixed_window_counter
$ bombardier -c 4000 -l -d 10s -r 4000 http://127.0.0.1:8080/fixed_window_counter
```

```bash
$ bombardier -c 50 -l -d 10s -r 50 http://127.0.0.1:8080/token_bucket
$ bombardier -c 4000 -l -d 10s -r 4000 http://127.0.0.1:8080/token_bucket
```

```bash
$ bombardier -c 50 -l -d 10s -r 50 http://127.0.0.1:8080/sliding_window_log
$ bombardier -c 4000 -l -d 10s -r 4000 http://127.0.0.1:8080/sliding_window_log
```

```bash
$ bombardier -c 50 -l -d 10s -r 50 http://127.0.0.1:8080/sliding_window_counter
$ bombardier -c 4000 -l -d 10s -r 4000 http://127.0.0.1:8080/sliding_window_counter
```

```bash
$ bombardier -c 50 -l -d 10s -r 50 http://127.0.0.1:8080/overload_detector
$ bombardier -c 4000 -l -d 10s -r 4000 http://127.0.0.1:8080/overload_detector
```

Stop the service:

```bash
$ docker stop service
```
