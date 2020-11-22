package it_test

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/fx/module/runner"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/provider/postgre/model"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type e2eTestSuite struct {
	suite.Suite
	dbConn *pg.DB
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, new(e2eTestSuite))
}

func (s *e2eTestSuite) SetupSuite() {
	options := &pg.Options{
		User:     "postgres",
		Password: "docker",
		Database: "pismo",
		Addr:     "localhost:5432",
		PoolSize: 5,
	}
	s.dbConn = pg.Connect(options)

	go runner.RunApplication()
}
func (s *e2eTestSuite) SetupTest() {
	models := []interface{}{
		(*model.Account)(nil),
		(*model.Transaction)(nil),
	}

	for _, m := range models {
		err := s.dbConn.Model(m).DropTable(&orm.DropTableOptions{
			Cascade: true,
		})
		err = s.dbConn.Model(m).CreateTable(&orm.CreateTableOptions{
			Temp: false,
		})

		s.NoError(err)
	}
}

func (s *e2eTestSuite) Test_EndToEnd_Create_Account() {
	requestStr := `{"document_number": "123"}`
	req, err := http.NewRequest(echo.POST, "http://localhost:8080/accounts", strings.NewReader(requestStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusCreated, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.Equal(`{"account_id":1,"document_number":"123"}`, strings.Trim(string(byteBody), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}

func (s *e2eTestSuite) TestEndToEnd_Try_To_Create_Account_Without_Document_Number() {
	requestStr := `{}`
	req, err := http.NewRequest(echo.POST, "http://localhost:8080/accounts", strings.NewReader(requestStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusBadRequest, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.Equal(`{"status_code":400,"message":"bad request"}`, strings.Trim(string(byteBody), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}

func (s *e2eTestSuite) TestEndToEnd_Recover_Account() {

	_, err := s.dbConn.Model(&model.Account{
		ID:             1,
		DocumentNumber: "123",
	}).Insert()

	s.NoError(err)

	req, err := http.NewRequest(echo.GET, "http://localhost:8080/accounts/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.xfts1LNO-o8YTY9SmuoyakqTBtuCOTYNF7sDrkH_9-g")
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.Equal(`{"account_id":1,"document_number":"123"}`, strings.Trim(string(byteBody), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}
func (s *e2eTestSuite) TestEndToEnd_Try_To_Recover_Account_Without_Token_Authorization() {

	_, err := s.dbConn.Model(&model.Account{
		ID:             1,
		DocumentNumber: "123",
	}).Insert()

	s.NoError(err)

	req, err := http.NewRequest(echo.GET, "http://localhost:8080/accounts/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusBadRequest, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.Equal(`{"message":"missing or malformed jwt"}`, strings.Trim(string(byteBody), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}
