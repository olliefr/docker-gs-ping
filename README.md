# docker-gs-ping

A simple Go server example for [Docker's Go Language Guide](https://docs.docker.com/language/golang/).

Notable features:

* A multi-stage Docker build using the official Go image to build a static binary for the application and then create a very lean deployment image with the smallest possible attack surface, running the application as a non-root user.
* (Kind of) behavior-driven testing of application's business requirements with proper isolation between tests using `ory/dockertest`.

Planned:

* CI pipeline using GitHub Actions.
* CD pipeline using GitHub Actions.
* Building Go modules and Docker images with `goreleaser`

## License

[Apache-2.0 License](LICENSE)
