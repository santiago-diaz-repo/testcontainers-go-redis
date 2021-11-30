package redis_management

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"testing"
	"time"
)

type redisContainer struct {
	testcontainers.Container
	URI string
}

func setupRedis(ctx context.Context) (*redisContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        "redis:6",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("* Ready to accept connections"),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	mappedPort, err := container.MappedPort(ctx, "6379")
	if err != nil {
		return nil, err
	}

	hostIP, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("redis://%s:%s", hostIP, mappedPort.Port())

	return &redisContainer{Container: container, URI: uri}, nil
}

func flushRedis(client redis.Client) error {
	return client.FlushAll().Err()
}

func TestWithRedis(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	ctx := context.Background()
	redisContainer, err := setupRedis(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer redisContainer.Terminate(ctx)

	options, err := redis.ParseURL(redisContainer.URI)
	if err != nil {
		t.Fatal(err)
	}
	client := redis.NewClient(options)
	defer flushRedis(*client)

	want := "test-containers"
	subject := NewRedisManagement(client)
	subject.Store("test",want,1 * time.Second)
	got := subject.Read("test")
	assert.Equal(t,got,want)
}