package api_test

import (
	"os"
	"testing"
	"time"

	"github.com/buddhimaaushan/mini_bank/api"
	"github.com/buddhimaaushan/mini_bank/db"
	"github.com/buddhimaaushan/mini_bank/utils"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func newTestServer(t *testing.T, store db.Store) *api.Server {
	config := utils.Config{
		TokenSymmetricKey:   utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := api.NewServer(config, store)
	if err != nil {
		t.Fatal("cannot create server: ", err)
	}
	return server
}
