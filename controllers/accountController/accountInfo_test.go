package AccountController_test

import (
	accountcontroller "JHETBackend/controllers/accountController"
	"log"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_GetAInfo(t *testing.T) {
	var c gin.Context
	c.Set("AccountID", uint64(14))
	log.Printf("%v", accountcontroller.GetAccountInfo(&c))
}
