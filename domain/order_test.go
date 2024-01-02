package domain_test

import (
	"encoder/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestValidateIfOderIsEmpty(t *testing.T) {
	order := domain.NewOder()
	err := order.Validate()

	require.Error(t, err)
}

func TestOrderIdIsNotAUuid(t *testing.T) {
	order := domain.NewOrder()
	order.ID     = "abc"
    order.user_id      = "123"
	order.pair         = "BTC/USD"
	order.amount       = "10"
	order.type_order   = "limit"

	order.CreatedAt = time.Now()
	err := order.Validate()
	require.Error(t, err)
}

func TestorderValidation(t *testing.T) {
	order := domain.NewOrder()

	order.ID = uuid.NewV4().String()
	order.user_id      = "123"
	order.pair         = "BTC/USD"
	order.amount       = "10"
	order.type_order   = "limit"
	order.created_at = time.Now()

	err := order.Validate()
	require.Nil(t, err)
}
