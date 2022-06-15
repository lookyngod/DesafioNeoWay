package cpfcnpj

import (
	"api/config"
	"api/models"
	"api/utils"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	ErroCPFCNPJDuplicado = "pq: duplicate key value violates unique constraint \"cpfcnpj_pkey\""
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

func InsereCPFCNPJ(db *sql.DB, cpfcnpj string) error {
	sqlStatement := fmt.Sprintf("INSERT INTO cpfcnpj(cpfcnpj, iscpf) VALUES('%s', '%t');", cpfcnpj, utils.IsValidCPF(cpfcnpj))
	_, err := db.Exec(sqlStatement)
	if err != nil {
		if err.Error() == ErroCPFCNPJDuplicado {
			return fmt.Errorf("CPF/CNPJ já cadastrado no Postgres")
		}
		return fmt.Errorf("falha na execução do insert de CPF/CNPJ no postgres: %v", err)
	}
	return nil
}

func BuscaCPFCNPJ(db *sql.DB, cpfcnpj string) (models.CPFCNPJ, error) {
	var row models.CPFCNPJ
	err := db.QueryRow("SELECT cpfcnpj, iscpf FROM cpfcnpj WHERE cpfcnpj=$1", cpfcnpj).Scan(&row.CPFCNPJ, &row.IsCPF)
	if err != nil {
		return row, fmt.Errorf("falha na execução da busca de cpf no postgres: %v", err)
	}
	return row, nil
}

func BuscaTodosCPFCNPJ(db *sql.DB) ([]models.CPFCNPJ, error) {
	var registros []models.CPFCNPJ
	rows, err := db.Query("SELECT * FROM cpfcnpj")
	if err != nil {
		return nil, fmt.Errorf("falha na execução da busca de todos os cpf no postgres: %v", err)
	}
	for rows.Next() {
		var row models.CPFCNPJ
		err = rows.Scan(&row.CPFCNPJ, &row.IsCPF)
		if err != nil {
			return nil, fmt.Errorf("falha na execução da busca de todos os cpf no postgres: %v", err)
		}
		registros = append(registros, row)
	}
	return registros, nil
}
