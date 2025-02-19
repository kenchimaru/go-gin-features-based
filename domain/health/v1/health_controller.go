package health

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type HealthController struct{}

func NewHealthController() *HealthController {
    return &HealthController{}
}

func (ctr *HealthController) HealthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "UP"})
}
