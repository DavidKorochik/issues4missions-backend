package handlers

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func newTestServer(recorder *httptest.ResponseRecorder) *gin.Engine {
	_, r := gin.CreateTestContext(recorder)
	return r
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
