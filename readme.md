# *API Controle de Transação*
Essa api tem como objetivo controlar as contas e suas transações

## *Requisitos necessários*
Os requisitos necessários para rodar a aplicação são: 

* [Docker](https://docs.docker.com/get-docker/)
* [Docker-compose](https://docs.docker.com/compose/install/)

**Clique nos links acima para instalar as dependências caso ainda não tenha em seu computador**

## *Como Rodar a Aplicação*

A aplicação depende do postgres para realizar a persistência dos dados e partindo deste princípio temos dois modos para rodar a nossa aplicação: 

#### make docker-run
Se você já tem uma instância do postgre rodando em sua máquina ou na nuvem pode facilmente apontar para sua instância, rodar o script de criação das tabelas e subir a aplicação com o comando ***make docker-run***

#### make docker-run-with-providers-dependencies
Se você não tem uma instância do postgre ou simplesmente quer rodar a aplicação do jeito mais rápido possível. Basta executar o comando ***make docker-run-with-providers-dependencies*** que irá buildar e subir uma imagem do postgre com todas as tabelas necessárias já criadas e no final a aplicação rodará pronta para receber os requests

## *Requests*

#### Postman
Caso você utilize postman para realizar requisições em suas apis basta importar essa [collection](https://www.getpostman.com/collections/62090174474357926179)

#### Swagger
Para visualizar as especificações da api acesse este [swagger](https://app.swaggerhub.com/apis-docs/maia.araujo51/controle-de-transacao/1.0#/Transactions/RecoverTransaction)

#### Authorization
A maioria dos endpoints presentes na aplicação necessitam de autententicação. Somente o endpoint de criação de conta que não é necessário enviar o header authorization.

| Header | Type | Description | Exemplo
| :--- | :--- | :--- | :---
| `Authorization` | `string` |Header de Authorization | Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.xfts1LNO-o8YTY9SmuoyakqTBtuCOTYNF7sDrkH_9-g

 