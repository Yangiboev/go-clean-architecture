package middleware

import (
	"github.com/Yangiboev/go-clean-architecture/config"
	"github.com/Yangiboev/go-clean-architecture/internal/auth"
	"github.com/Yangiboev/go-clean-architecture/internal/session"
	"github.com/Yangiboev/go-clean-architecture/pkg/logger"
)

// Middleware manager
type MiddlewareManager struct {
	sessUC  session.UCSession
	authUC  auth.UseCase
	cfg     *config.Config
	origins []string
	logger  logger.Logger
}

// Middleware manager constructor
func NewMiddlewareManager(sessUC session.UCSession, authUC auth.UseCase, cfg *config.Config, origins []string, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{sessUC: sessUC, authUC: authUC, cfg: cfg, origins: origins, logger: logger}
}
