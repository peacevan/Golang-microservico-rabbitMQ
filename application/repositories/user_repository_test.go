package repositories_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"encoder/framework/database"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUserRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	user := domain.NewUser()
	User.ID = uuid.NewV4().String()
	User.Name = "Ivan Amado Cardoso"
	User.CreatedAt = time.Now()

	repo := repositories.UserRepositoryDb{Db: db}
	repo.Insert(user)

	user, err := domain.NewUser("output_path", "Pending",user)
	require.Nil(t, err)

	repoUser := repositories.UserRepositoryDb{Db: db}
	repoUser.Insert(user)

	j, err := repoUser.Find(user.ID)
	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.ID, user.ID)
	require.Equal(t, j.UserID, user.ID)
}

func TestUserRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	user := domain.NewUser()
	user.ID = uuid.NewV4().String()
	user.FilePath = "path"
	user.CreatedAt = time.Now()

	repo := repositories.UserRepositoryDb{Db: db}
	repo.Insert(user)

	user, err := domain.NewUser("output_path", "Pending", user)
	require.Nil(t, err)

	repoUser := repositories.UserRepositoryDb{Db: db}
	repoUser.Insert(user)

	user.Status = "Complete"

	repoUser.Update(user)

	j, err := repoUser.Find(user.ID)
	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.Status, user.Status)
}
