package http

import (
	"github.com/labstack/echo/v4"

	"github.com/Yangiboev/go-clean-architecture/internal/comments"
	"github.com/Yangiboev/go-clean-architecture/internal/middleware"
)

// Map comments routes
func MapCommentsRoutes(commGroup *echo.Group, h comments.Handlers, mw *middleware.MiddlewareManager) {
	commGroup.POST("", h.Create(), mw.AuthSessionMiddleware, mw.CSRF)
	commGroup.DELETE("/:comment_id", h.Delete(), mw.AuthSessionMiddleware, mw.CSRF)
	commGroup.PUT("/:comment_id", h.Update(), mw.AuthSessionMiddleware, mw.CSRF)
	commGroup.GET("/:comment_id", h.GetByID())
	commGroup.GET("/byNewsId/:news_id", h.GetAllByNewsID())
}
