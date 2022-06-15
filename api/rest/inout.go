package rest

import (
	"api/domains/cpfcnpj/transport"
	"api/utils"
	"net/http"
)

func CpfCnpjHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// ROTA PARA INSERIR CPF OU CNPJ
	if r.Method == http.MethodPost {
		transport.RecebeCPFCNPJ(w, r)
		return
	}

	// ROTA PARA BUSCAR CPF OU CNPJ
	if r.Method == http.MethodGet {
		transport.BuscaDados(w, r)
		return
	}

	utils.RespondWithError(w, http.StatusBadRequest, 0, "Método não permitido")

}

// ROTA PARA BUSCAR TODOS OS CPF OU CNPJ
func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodGet {
		transport.GetAllValues(w, r)
		return
	}

	utils.RespondWithError(w, http.StatusBadRequest, 0, "Método não permitido")
}
