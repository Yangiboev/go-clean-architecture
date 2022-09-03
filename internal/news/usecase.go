//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package news

import (
	"context"

	"github.com/google/uuid"

	"github.com/Yangiboev/go-clean-architecture/internal/models"
	"github.com/Yangiboev/go-clean-architecture/pkg/utils"
)

// News use case
type UseCase interface {
	Create(ctx context.Context, news *models.News) (*models.News, error)
	Update(ctx context.Context, news *models.News) (*models.News, error)
	GetNewsByID(ctx context.Context, newsID uuid.UUID) (*models.NewsBase, error)
	Delete(ctx context.Context, newsID uuid.UUID) error
	GetNews(ctx context.Context, pq *utils.PaginationQuery) (*models.NewsList, error)
	SearchByTitle(ctx context.Context, title string, query *utils.PaginationQuery) (*models.NewsList, error)
}
