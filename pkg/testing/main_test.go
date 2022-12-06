package testing

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func newTestServer() {

}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
