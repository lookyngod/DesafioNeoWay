# Teste Técnico


# Desafio

Desenvolver uma aplicação de validação de CPF/CNPJ que deve conter uma interface (UI) para
gerenciamento de CPF/CNPJ (CRUD) com a possibilidade filtros, ordenação e marcação de alguns em uma
blocklist.

## Requisitos

- Docker-compose 1.29 ou compátivel com versão 3.8 do docker-compose.yml
- SGDB (Sistema Gerenciados de Banco de Dados) compátivel para Postgres
indico o Beekeeper, fácil de usar
- Extensão LiveServer no VScode

PS: Para logs direto do BackEnd, recomendo usar Postman com os endpoints inseridos no Router.go


## Como executar o projeto?

- Antes de mais nada abra seu terminal e clone este repositorio em sua maquina

- Instale o LiveServer na extensão para rodar o front end e inicie o GO Server no Rodapé do VSCode

- Em seguida suba os containers rodando o comando abaixo na pasta raiz do projeto:

    ```docker-compose up ```
- Entre na pasta '/api' 

- Em seguida ainda na pasta api rode o seguinte comando no terminal:
    ```go run router.go```
- *Prontinho, sua aplicação está no ar!*
## Como testar a aplicação?
- Por padrão sua aplicação front irá rodar na porta 5500 ```http://localhost:5500/frontend```
- A aplicação contém um prefixo rota '/api' para teste via Postman ```http://localhost:5000/api```
- *Tudo pronto, vamos testar!*
