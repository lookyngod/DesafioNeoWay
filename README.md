# TesteNeoWay
Teste Técnico

# API_CPFCNPJ

Desenvolver uma aplicação de validação de CPF/CNPJ que deve conter uma interface (UI) para
gerenciamento de CPF/CNPJ (CRUD) com a possibilidade filtros, ordenação e marcação de alguns em uma
blocklist.

## Requisitos

- Docker
- Docker-compose 1.29 ou compátivel com versão 3.8 do docker-compose.yml
- SGDB (Sistema Gerenciados de Banco de Dados) compátivel para Postgres e MYSQL
indico o Beekeeper, fácil de usar
- Postman para testes de endpoints

## Como executar o projeto?

- Antes de mais nada abra seu terminal e clone este repositorio em sua maquina

- Em seguida suba os containers rodando o comando abaixo na pasta raiz do projeto:

    ```docker-compose up ```
- Entre na pasta '/api' 

- Em seguida ainda na pasta site rode o seguinte comando no terminal:
    ```go run router.go```
*Prontinho, sua aplicação está no ar!*
## Como testar a aplicação?
- Por padrão sua aplicação irá rodar na porta 5000 ```http://localhost:5000```
- A aplicação contém um prefixo rota '/api' ```http://localhost:5000/api```
*Tudo pronto, vamos testar!*
- Primeiro efetue o login do usuario ultilizando o methodo POST```http://localhost:5000/api/login``` adicionando um json no body da requisição, exemplo:
```
{
    "nome":"varejao",
    "senha":"654321"
}
```
- Você receberá um token autenticado como resposta, ultilize este token para ter acesso ao  próximo endpoint cuja o methodo também é POST```http://localhost:5000/api/recebe-dados``` é só inserir ele no 'Barer Token' na aba 'Authorization' do Postman
e enviar um body de acordo com a estrutura:
```
{
    "contacts": [
        {
            "name": "Thauan Mendes",
            "cellphone": "5511960327601"
        }
     ]
}
```
**Prontinho, aplicação testada com sucesso!**
