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

func TestOrderRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewOrder()
	order.ID = uuid.NewV4().String()
	order.FilePath = "path"
	order.CreatedAt = time.Now()

	repo := repositories.OrderRepositoryDb{Db:db}
	repo.Insert(order)

	v, err := repo.Find(order.ID)

	require.NotEmpty(t, v.ID)
	require.Nil(t, err)
	require.Equal(t, v.ID, Order.ID)
}
