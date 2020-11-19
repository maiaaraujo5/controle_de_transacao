package request

import "github.com/labstack/echo/v4"

type Account struct {
	DocumentNumber string `json:"document_number" validate:"required"`
}

func NewAccount(c echo.Context) (*Account, error) {
	account := new(Account)
	err := c.Bind(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}
