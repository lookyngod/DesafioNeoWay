##  Desafio:
Desenvolver uma aplicação de validação de CPF/CNPJ que deve conter uma interface (UI) para gerenciamento de CPF/CNPJ (CRUD) com a possibilidade filtros, ordenação e marcação de alguns em uma blocklist.

A aplicação API tem como responsabilidade fornecer uma tabela com CPF e CNPJ quando inseridos.

## Stack

- Golang
- Docker
- Docker Compose
- Extensão LiveServer
 
 PS: Para logs direto do BackEnd, recomendo usar Postman com os endpoints inseridos no Router.go


## Instalação

A aplicação necessite de um ambiente com [Golang](https://go.dev/doc/install) 1.17+ para rodar
e do [Docker Compose](https://docs.docker.com/compose/install/)


Instale as dependências e para rodar a aplicação use o passo-a-passo abaixo:

### Passo 1:
Instale o LiveServer na extensão para rodar o front end e inicie o Live Server no Rodapé do VSCode

### Passo 2:
Suba os containers rodando o comando abaixo na pasta raiz do projeto:
```sh
docker-compose up
```


### Passo 3:
Após a inicialização do docker-compose, entre na pasta '/api' e em um novo terminal, rode o comando:
```sh
go run router.go
```

- *Prontinho, sua aplicação está no ar!*

## Como testar a aplicação?
- Por padrão sua aplicação front irá rodar na porta 5500 ```http://localhost:5500/frontend```
- A aplicação contém um prefixo rota '/api' para teste via Postman ```http://localhost:5000/api```


# Métodos
Requisições para a API devem seguir os padrões:
| Método | Descrição |
|---|---|
| `GET` | Retorna informações de um cpf ou cnpj na tabela |
| `POST` | Utilizado para validar e inserir na tabela um cpf/cnpj |
| `DELETE` | Deleta um cpf/cnpj da tabela  |

## Notes

Por se tratar de uma linguagem em que não há uma "regra" de arquitetura, utilizei algumas premissas da comunidade e de uma forma mais familiar devido a minha última experiência.
