package routes

import (
	"time"

	"github.com/adityjoshi/avinyaa/controllers"
	"github.com/adityjoshi/avinyaa/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/register", controllers.Register)
	incomingRoutes.POST("/login", middleware.RateLimiterMiddleware(2, time.Minute), controllers.Login)
	incomingRoutes.POST("/verify-otp", middleware.RateLimiterMiddleware(2, time.Minute), controllers.VerifyOTP)
	incomingRoutes.POST("/bookAppointment", middleware.AuthRequired("Patient", ""), middleware.OtpAuthRequireed, controllers.CreateAppointment)
}
