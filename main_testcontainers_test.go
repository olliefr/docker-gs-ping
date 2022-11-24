package main

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// The tests in this suite check if the business requirements
// have been satisfied. Each test checks a single requirement.

// Every test in this suite is isolated from all other tests,
// starting a different container for each test.

// Notice that, because Testcontainers for Go uses a random por
// for each container, we can have as many tests decoupled from
// each other using their own container.

// To verify that a container has started, the tests wait for the
// container to be ready, checking that the application is listening
// on the 8080/tcp port. This is possible thanks to the wait strategies
// provided by Testcontainers for Go.

// Requirement 1: The application response must contain a heart ("<3")
func TestRespondsWithLoveTestcontainers(t *testing.T) {
	req := testcontainers.ContainerRequest{
		Image:        "docker.io/olliefr/docker-gs-ping:latest",
		Env:          map[string]string{},
		ExposedPorts: []string{"8080/tcp"},
		WaitingFor:   wait.ForHTTP("/").WithPort("8080/tcp"), // wait for port to be ready
	}

	ctx := context.Background()
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err, "could not start container")

	endpoint, err := container.PortEndpoint(ctx, "8080/tcp", "http")
	require.NoError(t, err, "port not available")

	t.Cleanup(func() {
		require.NoError(t, container.Terminate(ctx), "failed to remove container")
	})

	var resp *http.Response

	resp, err = http.Get(endpoint)
	require.NoError(t, err, "HTTP error")
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code")

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "failed to read HTTP body")

	// Finally, test the business requirement!
	require.Contains(t, string(body), "<3", "does not respond with love?")
}

// Requirement 2: The application must include a health-check end-point at /ping,
// responding with JSON document { "Status": "OK" } if everything is well.
func TestHealthCheckTestcontainers(t *testing.T) {
	req := testcontainers.ContainerRequest{
		Image:        "docker.io/olliefr/docker-gs-ping:latest",
		Env:          map[string]string{},
		ExposedPorts: []string{"8080/tcp"},
		WaitingFor:   wait.ForHTTP("/").WithPort("8080/tcp"), // wait for port to be ready
	}

	ctx := context.Background()
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err, "could not start container")

	endpoint, err := container.PortEndpoint(ctx, "8080/tcp", "http")
	require.NoError(t, err, "port not available")

	t.Cleanup(func() {
		require.NoError(t, container.Terminate(ctx), "failed to remove container")
	})

	var resp *http.Response
	resp, err = http.Get(endpoint + "/ping")
	require.NoError(t, err, "HTTP error")
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code")

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "failed to read HTTP body")

	// Finally, test the business requirement!
	require.JSONEq(t, `{"Status":"OK"}`, string(body), "does not respond with valid JSON?")
}
