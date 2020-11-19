package request

import "github.com/labstack/echo/v4"

type Transaction struct {
	AccountID       int     `json:"account_id" validate:"required"`
	OperationTypeID int     `json:"operation_type_id" validate:"required"`
	Amount          float32 `json:"amount" validate:"required"`
}

func NewTransaction(c echo.Context) (*Transaction, error) {
	transaction := new(Transaction)
	err := c.Bind(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
