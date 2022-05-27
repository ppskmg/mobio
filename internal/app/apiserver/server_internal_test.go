package apiserver

import (
	"testing"
)

type testClient struct {
	testServer  *server
}

// testServer tc, teardownPostgres, teardownRedis, teardownRedisBlackList, teardownRedisRefreshBlackList
func testServer(t *testing.T) (tc *testClient) {

	tc = &testClient{
		testServer:  newServer(),
	}

	return tc

}
