# *API Controle de Transação*
Essa api tem como objetivo controlar as contas e suas transações

## *Requisitos necessários*
Os requisitos necessários para rodar a aplicação são: 

* [Docker](https://docs.docker.com/get-docker/)
* [Docker-compose](https://docs.docker.com/compose/install/)

**Clique nos links acima para instalar as dependências caso ainda não tenha no seu computador**

## *Como Rodar a Aplicação*

A aplicação depende do postgres para realizar a persistência dos dados e partindo deste princípio temos dois modos para rodar a nossa aplicação: 

#### make docker-run
Se você já tem uma instância do postgre rodando em sua máquina ou na nuvem pode facilmente apontar para sua instância, rodar o script de criação das tabelas e subir a aplicação com o comando ***make docker-run***

* Altere o arquivo de configuração em **configs/development.yaml**
* Execute o arquivo de scripts sql na sua instância do postgre. O arquivo sql se encontra no pacote **build** e se chama **init.sql**

#### make docker-run-with-providers-dependencies
Se você não tem uma instância do postgre ou simplesmente quer rodar a aplicação do jeito mais rápido possível. Basta executar o comando ***make docker-run-with-providers-dependencies*** que irá buildar e subir uma imagem do postgre com todas as tabelas necessárias já criadas e no final a aplicação rodará pronta para receber os requests

## *Requests*

#### Postman
Caso você utilize postman para realizar requisições em suas apis basta importar essa [collection](https://www.getpostman.com/collections/62090174474357926179)

#### Swagger
Para visualizar as especificações da api acesse este [swagger](https://app.swaggerhub.com/apis-docs/maia.araujo51/controle-de-transacao/1.0#/)

#### Authorization
A maioria dos endpoints presentes na aplicação necessitam de autenticação. Somente o endpoint de criação de conta que não é necessário enviar o header authorization.

| Header | Type | Description | Exemplo
| :--- | :--- | :--- | :---
| `Authorization` | `string` |Header de Authorization | Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.xfts1LNO-o8YTY9SmuoyakqTBtuCOTYNF7sDrkH_9-g

* Caso tenha problemas com essa chave de autorização, é possível gerar uma nova com a assinatura **key-segura** neste: [site](https://jwt.io/)

#### Exemplos de requests Curl
* Criar uma conta
```
curl --location --request POST 'localhost:8080/accounts' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI2ODgzNGYyNy1jNGJhLTQwZmUtODNhYy1mZTUwZDFjYjE5OWQiLCJuYW1lIjoiSm9obiBEb2UiLCJpYXQiOjE1MTYyMzkwMjJ9.r8mJQ-98UskHyMSJ9EGakNI_BEGEka_ZjjZt1bN3h9M' \
--header 'Content-Type: application/json' \
--data-raw '{
	"document_number": "12345689895"
}'
```

* Recuperar uma conta

```
curl --location --request GET 'localhost:8080/accounts/1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.xfts1LNO-o8YTY9SmuoyakqTBtuCOTYNF7sDrkH_9-g' \
--data-raw ''
```

* Criar uma transação

```
curl --location --request POST 'localhost:8080/transactions' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.xfts1LNO-o8YTY9SmuoyakqTBtuCOTYNF7sDrkH_9-g' \
--header 'Content-Type: application/json' \
--data-raw '{
	"account_id": 1,
	"operation_type_id": 4,
	"amount": 126.857
}'
```

* Recuperar uma transação

```
curl --location --request GET 'localhost:8080/transactions/1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.xfts1LNO-o8YTY9SmuoyakqTBtuCOTYNF7sDrkH_9-g' \
--data-raw ''
```

## *Comandos makefile*

####  make build
* Este comando é responsável por realizar o build da aplicação.
####  make test
* Este comando é responsável por rodar todos os testes unitários da aplicação.
####  make integration_tests
* Este comando é responsável por rodar todos os testes de integração presente na aplicação.
####  make run
* Este comando é responsável por levantar a aplicação deixando ela pronta para receber requests.
####  make docker-build
* Este comando é responsável por realizar o build da aplicação e logo em seguida o build da imagem docker.
####  make docker-run
* Este comando realiza o docker-build e logo em seguida levanta a imagem docker com modo network host.
####  make docker-compose-run-dependencies
* Este comando é responsável por subir todos os services presente no docker-compose. Ele sobe o postgres já pronto para ser utilizado pela aplicação.
####  make docker-run-with-providers-dependencies
* Este comando é responsável por subir todos os services presente no docker-compose e também levantar a aplicação em docker deixando ela pronta para receber requests.

## *Bibliotecas utilizadas*
* [go-pg](https://github.com/go-pg/pg)
* [echo](https://github.com/labstack/echo)
* [fx](https://github.com/uber-go/fx)
* [validator](https://github.com/go-playground/validator)