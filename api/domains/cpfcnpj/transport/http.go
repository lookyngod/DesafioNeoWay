package transport

import (
	"api/domains/cpfcnpj"
	"api/models"
	"api/utils"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

//HANDLERS CPFCNPJ E POSTGRES

func GetDados(w http.ResponseWriter, r *http.Request) {
	var dados models.CPFCNPJ
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0, "Erro ao ler o corpo da requisição")
		return
	}

	if err := json.Unmarshal(body, &dados); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0, "Falha ao realizar unmarshal do cliente")
		return
	}

	pg, err := cpfcnpj.ConnectPostgres()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, 0, "Erro ao conectar ao banco de dados")
		return
	}

	cpfcnpjmask := utils.RemoveMask(dados.CPFCNPJ)

	v, err := cpfcnpj.GetCPFCNPJ(pg, cpfcnpjmask)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, v)
}

func ReceiveCPFCNPJ(w http.ResponseWriter, r *http.Request) {
	var dados models.CPFCNPJ
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0, "Erro ao ler o corpo da requisição")
		return
	}

	if err := json.Unmarshal(body, &dados); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0, "Falha ao realizar unmarshal do cliente")
		return
	}

	pg, err := cpfcnpj.ConnectPostgres()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, 0, "Erro ao conectar ao banco de dados")
		return
	}

	cpfcnpjmask := utils.RemoveMask(dados.CPFCNPJ)

	if len(cpfcnpjmask) == 11 {

		if !utils.IsValidCPF(cpfcnpjmask) {
			utils.RespondWithError(w, http.StatusBadRequest, 0, "CPF inválido!")
			return
		}

	} else if len(cpfcnpjmask) == 14 {

		if !utils.IsValidCNPJ(cpfcnpjmask) {
			utils.RespondWithError(w, http.StatusBadRequest, 0, "CNPJ inválido!")
			return
		}

	} else if len(cpfcnpjmask) != 14 && len(cpfcnpjmask) != 11 {
		utils.RespondWithError(w, http.StatusBadRequest, 0, "CPF/CNPJ inválido!")
		return
	}

	err = cpfcnpj.InsertCPFCNPJ(pg, cpfcnpjmask)
	if err != nil {
		utils.RespondWithError(w, http.StatusOK, 0, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, "CPF/CNPJ inserido com sucesso!")
}

func GetAllValues(w http.ResponseWriter, r *http.Request) {
	pg, err := cpfcnpj.ConnectPostgres()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, 0, "Erro ao conectar ao banco de dados")
		return
	}

	dados, err := cpfcnpj.GetAllCPFCNPJ(pg)
	if err != nil {
		utils.RespondWithError(w, http.StatusOK, 0, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, dados)
}

func DeleteDados(w http.ResponseWriter, r *http.Request) {
	pg, err := cpfcnpj.ConnectPostgres()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, 0, "Erro ao conectar ao banco de dados")
		return
	}

	err = cpfcnpj.DeleteCPFCNPJ(pg, r.FormValue("cpfcnpj"))
	if err != nil {
		utils.RespondWithError(w, http.StatusOK, 0, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, "CPF/CNPJ deletado com sucesso!")

}

func CreateTable(w http.ResponseWriter, r *http.Request) {
	pg, err := cpfcnpj.ConnectPostgres()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, 0, "Erro ao conectar ao banco de dados")
		return
	}

	err = cpfcnpj.CreateTable(pg)
	if err != nil {
		utils.RespondWithError(w, http.StatusOK, 0, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, "Tabela criada com sucesso!")
}
