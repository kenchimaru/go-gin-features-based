package health

import (
    "github.com/gin-gonic/gin"
)

func HealthRoutes(r *gin.Engine) {
    controller := NewHealthController()
    r.GET("/api/v1/healthcheck", controller.HealthCheck)
}
