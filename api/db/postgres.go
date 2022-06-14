package db

import (
	"api/config"
	"api/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	ErroCPFDuplicado  = "pq: duplicate key value violates unique constraint \"cpf_pkey\""
	ErroCNPJDuplicado = "pq: duplicate key value violates unique constraint \"cnpj_pkey\""
	ErroCNPJNotFound  = "sql: no rows in result set"
)

func ConectarPostgres() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.StringConexaoPostgres)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func InsereCNPJPostgres(db *sql.DB, cnpj string) error {
	sqlStatement := fmt.Sprintf("INSERT INTO cnpj(cnpj) VALUES('%s');", cnpj)
	res, err := db.Exec(sqlStatement)
	if err != nil {
		if err.Error() == ErroCNPJDuplicado {
			return fmt.Errorf("CNPJ já cadastrado no Postgres")

		}
		return fmt.Errorf("falha na execução do Insert de CNPJ no Postgres: %v", err)
	}

	log.Printf("CNPJ inserido no postgres: %s", res)

	return nil

}

func InsereCPFPostgres(db *sql.DB, cpf string) error {
	sqlStatement := fmt.Sprintf("INSERT INTO cpf(cpf) VALUES('%s');", cpf)
	res, err := db.Exec(sqlStatement)
	if err != nil {
		if err.Error() == ErroCPFDuplicado {
			return fmt.Errorf("CPF já cadastrado no Postgres")

		}
		return fmt.Errorf("falha na execução do insert de CPF no postgres: %v", err)
	}

	log.Printf("CPF inserido no postgres: %s", res)

	return nil

}

func BuscaCpfPostgres(db *sql.DB, cpf string) (models.CPFCNPJ, error) {
	var row models.CPFCNPJ
	err := db.QueryRow("SELECT cpf, iscpf from cpf where cpf=$1", cpf).Scan(&row.CPFCNPJ, &row.IsCPF)
	if err != nil {
		return row, fmt.Errorf("falha na execução da busca de cpf no postgres: %v", err)
	}

	log.Printf("Busca de CPF no postgres: %v", row)

	return row, nil
}

func BuscaCnpjPostgres(db *sql.DB, cnpj string) (models.CPFCNPJ, error) {
	var row models.CPFCNPJ
	err := db.QueryRow("SELECT cnpj, iscpf from cnpj where cnpj=$1", cnpj).Scan(&row.CPFCNPJ, &row.IsCPF)
	if err != nil {
		if err.Error() == ErroCNPJNotFound {
			return row, fmt.Errorf("CNPJ não existente no Postgres")

		}
		return row, fmt.Errorf("falha na execução da busca de CNPJ no postgres: %v", err)
	}

	log.Printf("Busca de CNPJ no postgres: %v", row)

	return row, nil
}
