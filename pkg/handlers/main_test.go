package handlers

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func newTestServer() (*gin.Engine, *httptest.ResponseRecorder) {
	recorder := httptest.NewRecorder()
	_, r := gin.CreateTestContext(recorder)
	return r, recorder
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
