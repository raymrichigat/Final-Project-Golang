package homecontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}
