package rest

import (
	"api/domains"
	"api/utils"
	"net/http"
)

func CpfCnpjHandler(w http.ResponseWriter, r *http.Request) {
	// ROTA PARA INSERIR CPF OU CNPJ
	if r.Method == http.MethodPost {
		domains.RecebeCPFCNPJ(w, r)
		return
	}

	// ROTA PARA BUSCAR CPF OU CNPJ
	if r.Method == http.MethodGet {
		domains.BuscaDados(w, r)
		return

	}
	utils.RespondWithError(w, http.StatusBadRequest, 0, "Método não permitido")

}

// ROTA PARA TIMESTAMP DE STATUS
