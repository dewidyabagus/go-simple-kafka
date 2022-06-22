package order_test

import (
	"errors"
	"learn/kafka/business/order"
	orderMocks "learn/kafka/business/order/mocks"
	"learn/kafka/utils/configuration"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ErrorMarshal order.NewOrder

var date, _ = time.ParseInLocation("2006-01-02 15:04:05", "2022-06-22 20:00:00", time.Local)

var (
	// Mocking
	orderRepository orderMocks.Repository
	events          orderMocks.Events

	// Instant Service
	orderService = order.NewService(&orderRepository, &events, &configuration.Kafka{})

	// Test Data
	transNo  = "000001"
	newOrder = &order.NewOrder{
		TransactionNo: transNo,
		Items: []order.Item{{
			ItemID:    1,
			ItemPrice: 1000,
			Qty:       2,
		}},
		Date: "2022-06-22 15:30:01",
	}
	orderCollection = []order.Order{
		{
			ID:            1,
			TransactionNo: transNo,
			ItemID:        1,
			ItemPrice:     1000,
			Qty:           2,
			Date:          date,
		},
		{
			ID:            2,
			TransactionNo: transNo,
			ItemID:        2,
			ItemPrice:     1000,
			Qty:           3,
			Date:          date,
		},
	}
)

func TestCreateNewOrder(t *testing.T) {
	t.Run("Expect invalid data", func(t *testing.T) {
		err := orderService.CreateNewOrder(&order.NewOrder{})

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "validate new order")
		assert.Contains(t, err.Error(), "TransactionNo")
	})

	t.Run("Expect error checking transaction no", func(t *testing.T) {
		orderRepository.On("CheckExistingTransNo", mock.AnythingOfType("string")).Return(false, errors.New("error validate trans no")).Once()

		err := orderService.CreateNewOrder(newOrder)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "error validate trans no")
	})

	t.Run("Expect error transaction no alredy exists", func(t *testing.T) {
		orderRepository.On("CheckExistingTransNo", mock.AnythingOfType("string")).Return(true, nil).Once()

		err := orderService.CreateNewOrder(newOrder)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "transaction no already exists")
	})

	t.Run("Expect error marshal json", func(t *testing.T) {

	})
	t.Run("Expect error create new order", func(t *testing.T) {
		orderRepository.On("CheckExistingTransNo", mock.AnythingOfType("string")).Return(false, nil).Once()
		orderRepository.On("CreateNewOrder", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("[]order.Order")).Return(errors.New("SQL Error")).Once()

		err := orderService.CreateNewOrder(newOrder)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "SQL Error")
	})

	t.Run("Success create new order", func(t *testing.T) {
		orderRepository.On("CheckExistingTransNo", mock.AnythingOfType("string")).Return(false, nil).Once()
		orderRepository.On("CreateNewOrder", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("[]order.Order")).Return(nil).Once()
		events.On("SendEventAsync", mock.AnythingOfType("string"), mock.AnythingOfType("[]uint8"), mock.AnythingOfType("[]uint8")).Return(nil).Once()

		err := orderService.CreateNewOrder(newOrder)
		assert.Nil(t, err)
	})
}

func TestGetOrderByTransNo(t *testing.T) {
	t.Run("Expect error and record not found", func(t *testing.T) {
		orderRepository.On("GetOrderByTransNo", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return(nil, errors.New("SQL Error")).Once()

		orders0, err0 := orderService.GetOrderByTransNo(transNo)
		assert.NotNil(t, err0)
		assert.Nil(t, orders0)
		assert.Contains(t, err0.Error(), "SQL Error")

		orderRepository.On("GetOrderByTransNo", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return(nil, errors.New("record not found")).Once()
		orders1, err1 := orderService.GetOrderByTransNo(transNo)
		assert.NotNil(t, err1)
		assert.Nil(t, orders1)
		assert.Contains(t, err1.Error(), "record not found")
	})
	t.Run("Success get transaction", func(t *testing.T) {
		orderRepository.On("GetOrderByTransNo", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return(orderCollection, nil).Once()

		orders, err := orderService.GetOrderByTransNo(transNo)
		assert.Nil(t, err)
		assert.NotNil(t, orders)
		assert.Equal(t, orderCollection, orders)
	})
}
