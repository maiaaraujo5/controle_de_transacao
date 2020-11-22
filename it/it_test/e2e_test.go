package it_test

import (
	"bou.ke/monkey"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/fx/module/runner"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/provider/postgre/model"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
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
		Database: "pismo-teste",
		Addr:     "localhost:5432",
		PoolSize: 5,
	}
	s.dbConn = pg.Connect(options)

	err := os.Setenv("environment", "test")
	s.NoError(err)

	time.Sleep(1 * time.Second) //Tempo para esperar a imagem docker do banco subir

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

func (s *e2eTestSuite) Test_EndToEnd_Try_To_Create_Account_With_Document_Number_That_Already_Exist_in_Another_Account() {
	_, err := s.dbConn.Model(&model.Account{
		ID:             1,
		DocumentNumber: "123",
	}).Insert()

	s.NoError(err)

	requestStr := `{"document_number": "123"}`
	req, err := http.NewRequest(echo.POST, "http://localhost:8080/accounts", strings.NewReader(requestStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusConflict, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.Equal(`{"status_code":409,"message":"the document number is already register in another account"}`, strings.Trim(string(byteBody), "\n"))
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

func (s *e2eTestSuite) TestEndToEnd_Try_To_Create_Account_With_An_Malformed_Json_Body() {
	requestStr := `{"document_number:" "123"}`
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

func (s *e2eTestSuite) TestEndToEnd_Try_To_Create_Account_With_An_Incorrect_Content_Type() {
	requestStr := `{"document_number": "123"}`
	req, err := http.NewRequest(echo.POST, "http://localhost:8080/accounts", strings.NewReader(requestStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
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

func (s *e2eTestSuite) TestEndToEnd_Try_Recover_Account_that_does_not_exists_in_database() {

	req, err := http.NewRequest(echo.GET, "http://localhost:8080/accounts/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.xfts1LNO-o8YTY9SmuoyakqTBtuCOTYNF7sDrkH_9-g")
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusNotFound, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.Equal(`{"status_code":404,"message":"the account was not found"}`, strings.Trim(string(byteBody), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}

func (s *e2eTestSuite) TestEndToEnd_Create_Transaction_With_Cash_Purchase() {

	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2020, time.November, 21, 1, 2, 3, 4, time.UTC)
	})

	_, err := s.dbConn.Model(&model.Account{
		ID:             1,
		DocumentNumber: "123",
	}).Insert()
	s.NoError(err)
	requestStr := `{"account_id":1,"operation_type_id":1,"amount":126.85}`

	req, err := http.NewRequest(echo.POST, "http://localhost:8080/transactions", strings.NewReader(requestStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.xfts1LNO-o8YTY9SmuoyakqTBtuCOTYNF7sDrkH_9-g")
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusCreated, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.Equal(`{"id":1,"account_id":1,"operation_type_id":1,"amount":-126.85,"event_date":"2020-11-21T01:02:03.000000004Z"}`, strings.Trim(string(byteBody), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}

func (s *e2eTestSuite) TestEndToEnd_Create_Transaction_With_Installment_Purchase() {

	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2020, time.November, 21, 1, 2, 3, 4, time.UTC)
	})

	_, err := s.dbConn.Model(&model.Account{
		ID:             1,
		DocumentNumber: "123",
	}).Insert()
	s.NoError(err)
	requestStr := `{"account_id":1,"operation_type_id":2,"amount":126.85}`

	req, err := http.NewRequest(echo.POST, "http://localhost:8080/transactions", strings.NewReader(requestStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.xfts1LNO-o8YTY9SmuoyakqTBtuCOTYNF7sDrkH_9-g")
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusCreated, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.Equal(`{"id":1,"account_id":1,"operation_type_id":2,"amount":-126.85,"event_date":"2020-11-21T01:02:03.000000004Z"}`, strings.Trim(string(byteBody), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}

func (s *e2eTestSuite) TestEndToEnd_Create_Transaction_With_Withdraw() {

	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2020, time.November, 21, 1, 2, 3, 4, time.UTC)
	})

	_, err := s.dbConn.Model(&model.Account{
		ID:             1,
		DocumentNumber: "123",
	}).Insert()
	s.NoError(err)
	requestStr := `{"account_id":1,"operation_type_id":3,"amount":126.85}`

	req, err := http.NewRequest(echo.POST, "http://localhost:8080/transactions", strings.NewReader(requestStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.xfts1LNO-o8YTY9SmuoyakqTBtuCOTYNF7sDrkH_9-g")
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusCreated, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.Equal(`{"id":1,"account_id":1,"operation_type_id":3,"amount":-126.85,"event_date":"2020-11-21T01:02:03.000000004Z"}`, strings.Trim(string(byteBody), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}

func (s *e2eTestSuite) TestEndToEnd_Create_Transaction_With_Payment() {

	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2020, time.November, 21, 1, 2, 3, 4, time.UTC)
	})

	_, err := s.dbConn.Model(&model.Account{
		ID:             1,
		DocumentNumber: "123",
	}).Insert()
	s.NoError(err)
	requestStr := `{"account_id":1,"operation_type_id":4,"amount":126.85}`

	req, err := http.NewRequest(echo.POST, "http://localhost:8080/transactions", strings.NewReader(requestStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.xfts1LNO-o8YTY9SmuoyakqTBtuCOTYNF7sDrkH_9-g")
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusCreated, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.Equal(`{"id":1,"account_id":1,"operation_type_id":4,"amount":126.85,"event_date":"2020-11-21T01:02:03.000000004Z"}`, strings.Trim(string(byteBody), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}

func (s *e2eTestSuite) TestEndToEnd_Create_Transaction_Without_A_Token_Authorization() {

	requestStr := `{"account_id":1,"operation_type_id":4,"amount":126.85}`

	req, err := http.NewRequest(echo.POST, "http://localhost:8080/transactions", strings.NewReader(requestStr))
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

func (s *e2eTestSuite) TestEndToEnd_Create_Transaction_With_A_Json_Malformed() {

	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2020, time.November, 21, 1, 2, 3, 4, time.UTC)
	})

	_, err := s.dbConn.Model(&model.Account{
		ID:             1,
		DocumentNumber: "123",
	}).Insert()
	s.NoError(err)
	requestStr := `{"account_id:"1,"operation_type_id":4,"amount":126.85}`

	req, err := http.NewRequest(echo.POST, "http://localhost:8080/transactions", strings.NewReader(requestStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.xfts1LNO-o8YTY9SmuoyakqTBtuCOTYNF7sDrkH_9-g")
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

func (s *e2eTestSuite) TestEndToEnd_Try_To_Create_A_Transaction_With_One_Invalid_Account() {
	requestStr := `{"account_id":1,"operation_type_id":1,"amount":126.85}`

	req, err := http.NewRequest(echo.POST, "http://localhost:8080/transactions", strings.NewReader(requestStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.xfts1LNO-o8YTY9SmuoyakqTBtuCOTYNF7sDrkH_9-g")
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusNotFound, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.Equal(`{"status_code":404,"message":"the account was not found"}`, strings.Trim(string(byteBody), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}

func (s *e2eTestSuite) TestEndToEnd_Try_To_Create_Transaction_With_Invalid_Operation_Type_Id() {

	requestStr := `{"account_id":1,"operation_type_id":200,"amount":126.85}`

	req, err := http.NewRequest(echo.POST, "http://localhost:8080/transactions", strings.NewReader(requestStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.xfts1LNO-o8YTY9SmuoyakqTBtuCOTYNF7sDrkH_9-g")
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusBadRequest, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.Equal(`{"status_code":400,"message":"the operation type is invalid"}`, strings.Trim(string(byteBody), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}

func (s *e2eTestSuite) TestEndToEnd_Try_To_Create_Transaction_Without_Account_Id() {

	requestStr := `{"operation_type_id":200,"amount":126.85}`

	req, err := http.NewRequest(echo.POST, "http://localhost:8080/transactions", strings.NewReader(requestStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.xfts1LNO-o8YTY9SmuoyakqTBtuCOTYNF7sDrkH_9-g")
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

func (s *e2eTestSuite) TestEndToEnd_Try_To_Create_Transaction_Without_Operation_Type_Id() {

	requestStr := `{"account_id":1,"amount":126.85}`

	req, err := http.NewRequest(echo.POST, "http://localhost:8080/transactions", strings.NewReader(requestStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.xfts1LNO-o8YTY9SmuoyakqTBtuCOTYNF7sDrkH_9-g")
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

func (s *e2eTestSuite) TestEndToEnd_Try_To_Create_Transaction_Without_Amount() {

	requestStr := `{"account_id":1,"operation_type_id":200}`

	req, err := http.NewRequest(echo.POST, "http://localhost:8080/transactions", strings.NewReader(requestStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.xfts1LNO-o8YTY9SmuoyakqTBtuCOTYNF7sDrkH_9-g")
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

func (s *e2eTestSuite) TestEndToEnd_Recover_Transaction_With_Operation_Type_Id_Cash_Purchase() {

	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2020, time.November, 21, 1, 2, 3, 4, time.UTC)
	})

	_, err := s.dbConn.Model(&model.Transaction{
		ID:              1,
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          -100,
		EventDate:       time.Now(),
	}).Insert()

	req, err := http.NewRequest(echo.GET, "http://localhost:8080/transactions/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.xfts1LNO-o8YTY9SmuoyakqTBtuCOTYNF7sDrkH_9-g")
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.Equal(`{"id":1,"account_id":1,"operation_type_id":1,"amount":-100,"event_date":"2020-11-21T01:02:03Z"}`, strings.Trim(string(byteBody), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}

func (s *e2eTestSuite) TestEndToEnd_Recover_Transaction_With_Operation_Type_Id_Installment_Purchase() {

	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2020, time.November, 21, 1, 2, 3, 4, time.UTC)
	})

	_, err := s.dbConn.Model(&model.Transaction{
		ID:              1,
		AccountID:       1,
		OperationTypeID: 2,
		Amount:          -100,
		EventDate:       time.Now(),
	}).Insert()

	req, err := http.NewRequest(echo.GET, "http://localhost:8080/transactions/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.xfts1LNO-o8YTY9SmuoyakqTBtuCOTYNF7sDrkH_9-g")
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.Equal(`{"id":1,"account_id":1,"operation_type_id":2,"amount":-100,"event_date":"2020-11-21T01:02:03Z"}`, strings.Trim(string(byteBody), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}

func (s *e2eTestSuite) TestEndToEnd_Recover_Transaction_With_Operation_Type_Id_Withdraw() {

	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2020, time.November, 21, 1, 2, 3, 4, time.UTC)
	})

	_, err := s.dbConn.Model(&model.Transaction{
		ID:              1,
		AccountID:       1,
		OperationTypeID: 3,
		Amount:          -100,
		EventDate:       time.Now(),
	}).Insert()

	req, err := http.NewRequest(echo.GET, "http://localhost:8080/transactions/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.xfts1LNO-o8YTY9SmuoyakqTBtuCOTYNF7sDrkH_9-g")
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.Equal(`{"id":1,"account_id":1,"operation_type_id":3,"amount":-100,"event_date":"2020-11-21T01:02:03Z"}`, strings.Trim(string(byteBody), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}

func (s *e2eTestSuite) TestEndToEnd_Recover_Transaction_With_Operation_Type_Id_Payment() {

	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2020, time.November, 21, 1, 2, 3, 4, time.UTC)
	})

	_, err := s.dbConn.Model(&model.Transaction{
		ID:              1,
		AccountID:       1,
		OperationTypeID: 4,
		Amount:          100,
		EventDate:       time.Now(),
	}).Insert()

	req, err := http.NewRequest(echo.GET, "http://localhost:8080/transactions/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.xfts1LNO-o8YTY9SmuoyakqTBtuCOTYNF7sDrkH_9-g")
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.Equal(`{"id":1,"account_id":1,"operation_type_id":4,"amount":100,"event_date":"2020-11-21T01:02:03Z"}`, strings.Trim(string(byteBody), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}

func (s *e2eTestSuite) TestEndToEnd_Try_To_Recover_A_Transaction_That_Does_Not_Exists_In_Database() {

	req, err := http.NewRequest(echo.GET, "http://localhost:8080/transactions/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.xfts1LNO-o8YTY9SmuoyakqTBtuCOTYNF7sDrkH_9-g")
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusNotFound, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.Equal(`{"status_code":404,"message":"the transaction not exists"}`, strings.Trim(string(byteBody), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}
