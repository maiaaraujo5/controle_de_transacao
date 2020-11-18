package request

import "github.com/labstack/echo/v4"

type Transaction struct {
	AccountID       int     `json:"account_id"`
	OperationTypeID int     `json:"operation_type_id"`
	Amount          float32 `json:"amount"`
}

func NewTransaction(c echo.Context) (*Transaction, error) {
	transaction := new(Transaction)
	err := c.Bind(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
