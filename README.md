# docker-gs-ping

A simple Go server/microservice example for [Docker's Go Language Guide](https://docs.docker.com/language/golang/).

Notable features:

* Includes a [multi-stage `Dockerfile`](https://github.com/olliefr/docker-gs-ping/blob/main/Dockerfile.multistage), which actually is a good example of how to build Go binaries _for production releases_.
* Has functional tests for application's business requirements with proper isolation between tests using [`ory/dockertest`](https://github.com/ory/dockertest).
* Has a CI pipeline using GitHub Actions to run functional tests in independent containers.
* Has a CD pipeline using GitHub Actions to publish to Docker Hub.

Planned:

* Building Go modules and Docker images with `goreleaser`

## Running the tests

1. Using dockertest:
```shell
go test ./... -count=1 -tags=dockertest -v    
=== RUN   TestRespondsWithLove
    main_test.go:50: container not ready, waiting...
--- PASS: TestRespondsWithLove (0.90s)
=== RUN   TestHealthCheck
    main_test.go:86: container not ready, waiting...
--- PASS: TestHealthCheck (0.86s)
PASS
ok  	github.com/olliefr/docker-gs-ping	1.947s
```

2. Using Testcontainers for Go:
```shell
go test ./... -count=1 -tags=testcontainers -v
=== RUN   TestRespondsWithLoveTestcontainers
2023/01/25 13:16:16 github.com/testcontainers/testcontainers-go - Connected to docker: 
  Server Version: 20.10.21
  API Version: 1.41
  Operating System: Docker Desktop
  Total Memory: 7851 MB
2023/01/25 13:16:16 Starting container id: b074ee900ace image: docker.io/testcontainers/ryuk:0.3.4
2023/01/25 13:16:16 Waiting for container id b074ee900ace image: docker.io/testcontainers/ryuk:0.3.4
2023/01/25 13:16:16 Container is ready id: b074ee900ace image: docker.io/testcontainers/ryuk:0.3.4
2023/01/25 13:16:16 Starting container id: 894bb575f712 image: docker.io/olliefr/docker-gs-ping:latest
2023/01/25 13:16:16 Waiting for container id 894bb575f712 image: docker.io/olliefr/docker-gs-ping:latest
2023/01/25 13:16:17 Container is ready id: 894bb575f712 image: docker.io/olliefr/docker-gs-ping:latest
--- PASS: TestRespondsWithLoveTestcontainers (1.03s)
=== RUN   TestHealthCheckTestcontainers
2023/01/25 13:16:17 Starting container id: 2da2fa2876c6 image: docker.io/olliefr/docker-gs-ping:latest
2023/01/25 13:16:17 Waiting for container id 2da2fa2876c6 image: docker.io/olliefr/docker-gs-ping:latest
2023/01/25 13:16:17 Container is ready id: 2da2fa2876c6 image: docker.io/olliefr/docker-gs-ping:latest
--- PASS: TestHealthCheckTestcontainers (0.42s)
PASS
ok  	github.com/olliefr/docker-gs-ping	1.753s
```

## Want _moar_?!

There is a more advanced example in [olliefr/docker-gs-ping-roach](https://github.com/olliefr/docker-gs-ping-roach) using [CockroachDB](https://github.com/cockroachdb/cockroach).

## Contributing

This was written for an _introduction_ section of the Docker tutorial and as such it favours brevity and pedagogical clarity over robustness. 

Thus, feedback is welcome, but please no nits or pedantry. Ain't nobody got time for that 🙃

## License

[Apache-2.0 License](LICENSE)
