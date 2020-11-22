# **API Controle de Transação**
Essa api tem como objetivo controlar as contas e suas transações

## **Como Rodar a Aplicação**
A aplicação depende do postgres para realizar a persistência dos dados e partindo deste princípio temos dois modos para rodar a nossa aplicação: 

#### make docker-run
Se você já tem uma instância do postgre rodando em sua máquina ou na nuvem pode facilmente apontar para sua instância, rodar o script de criação das tabelas e subir a aplicação com o comando ***make docker-run***

#### make docker-run-with-providers-dependencies
Se você não tem uma instância do postgre ou simplesmente quer rodar a aplicação do jeito mais rápido possível. Basta executar o comando ***make docker-run-with-providers-dependencies*** que irá buildar e subir uma imagem do postgre com todas as tabelas necessárias já criadas e no final a aplicação rodará pronta para receber os requests

##Requests
### **Authorization**
