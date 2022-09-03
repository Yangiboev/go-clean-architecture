package repository

import (
	"context"
	"log"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/Yangiboev/go-clean-architecture/internal/models"
	"github.com/Yangiboev/go-clean-architecture/internal/session"
)

func setupRedis() session.SessRepository {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatal(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	sessRepository := NewSessionRepository(client, nil)
	return sessRepository
}

func createTestSession(sID uuid.UUID) *models.Session {
	return &models.Session{
		SessionID: sID.String(),
		UserID:    sID,
	}
}

func TestSessionRepo_CreateSession(t *testing.T) {
	t.Parallel()

	sessRepository := setupRedis()

	t.Run("CreateSession", func(t *testing.T) {
		sessUUID := uuid.New()

		s, err := sessRepository.CreateSession(context.Background(), createTestSession(sessUUID), 10)
		require.NoError(t, err)
		require.NotEqual(t, s, "")
	})
}

func TestSessionRepo_GetSessionByID(t *testing.T) {
	t.Parallel()

	sessRepository := setupRedis()

	t.Run("GetSessionByID", func(t *testing.T) {
		sessUUID := uuid.New()

		createdSess, err := sessRepository.CreateSession(context.Background(), createTestSession(sessUUID), 10)
		require.NoError(t, err)
		require.NotEqual(t, createdSess, "")

		s, err := sessRepository.GetSessionByID(context.Background(), createdSess)
		require.NoError(t, err)
		require.NotEqual(t, s, "")
	})
}

func TestSessionRepo_DeleteByID(t *testing.T) {
	t.Parallel()

	sessRepository := setupRedis()

	t.Run("DeleteByID", func(t *testing.T) {
		sessUUID := uuid.New()

		createdSess, err := sessRepository.CreateSession(context.Background(), createTestSession(sessUUID), 10)
		require.NoError(t, err)
		require.NotEqual(t, createdSess, "")

		err = sessRepository.DeleteByID(context.Background(), createdSess)
		require.NoError(t, err)
	})
}
