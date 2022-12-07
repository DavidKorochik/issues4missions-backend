package testing

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DavidKorochik/issues4missions-backend/pkg/controllers"
	"github.com/gin-gonic/gin"
)

type TestServer struct {
	r         *gin.Engine
	userStore controllers.UserStore
}

func newTestServer(recorder *httptest.ResponseRecorder, userStore controllers.UserStore) *TestServer {
	_, r := gin.CreateTestContext(recorder)

	return &TestServer{
		r:         r,
		userStore: userStore,
	}
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
