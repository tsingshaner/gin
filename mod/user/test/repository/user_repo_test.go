package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tsingshaner/gin/e2e/container"
	"github.com/tsingshaner/gin/mod/user/dto"
	"github.com/tsingshaner/gin/mod/user/internal/entity"
	"github.com/tsingshaner/gin/shared"
)

func TestUserRepo(t *testing.T) {
	t.Parallel()

	server, cleanup := container.NewAppWithCleanup()
	t.Cleanup(cleanup)

	userRepo := server.Providers().Repo.User

	t.Run("Save", func(t *testing.T) {
		t.Parallel()

		user, err := userRepo.Save(entity.New(entity.WithUser(&dto.User{
			Username: "Save",
			Password: "Save",
			Nickname: "Save",
		})))

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Greater(t, user.ID, shared.ID(0))
		assert.Nil(t, user.Role)
		assert.Equal(t, "Save", user.Username)
		assert.Equal(t, "Save", user.Nickname)
		assert.Equal(t, "Save", user.Password)

		id := user.ID
		user, err = userRepo.Save(entity.New(entity.WithUser(&dto.User{
			ID:       id,
			Username: "SaveAsUpdate",
			Password: "SaveAsUpdate",
			Nickname: "SaveAsUpdate",
		})))

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, id, user.ID)
		assert.Nil(t, user.Role)
		assert.Equal(t, "SaveAsUpdate", user.Username)
		assert.Equal(t, "SaveAsUpdate", user.Password)
		assert.Equal(t, "SaveAsUpdate", user.Nickname)
	})

	t.Run("QueryByID", func(t *testing.T) {
		t.Parallel()

		u, err := userRepo.Save(entity.New(entity.WithUser(&dto.User{
			Username: "QueryByID",
			Password: "QueryByID",
		})))
		assert.NoError(t, err)

		user, err := userRepo.GetUserByID(u.ID)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, u.ID, user.ID)
		assert.Equal(t, u.Username, user.Username)
		assert.Equal(t, u.Password, user.Password)
	})

	t.Run("CreateWithASoftDeletedUsername", func(t *testing.T) {
		t.Parallel()

		user, err := userRepo.Save(entity.New(entity.WithUser(&dto.User{
			Username: "CreateWithASoftDeletedUsername",
			Password: "Create",
		})))

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Greater(t, user.ID, shared.ID(0))

		assert.NoError(t, userRepo.Remove(user.ID))

		user.ID = 0
		user, err = userRepo.Save(user)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Greater(t, user.ID, shared.ID(0))
		assert.Equal(t, "CreateWithASoftDeletedUsername", user.Username)
	})
}
