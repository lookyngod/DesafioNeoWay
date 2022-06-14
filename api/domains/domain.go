package domains

import (
	"api/db"
	"api/models"
	"api/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func BuscaDados(w http.ResponseWriter, r *http.Request) {
	var dados models.CPFCNPJ
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Erro ao ler o corpo da requisição")
		return
	}

	if err := json.Unmarshal(body, &dados); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0, "Falha ao realizar unmarshal do cliente")
		return
	}

	pg, err := db.ConectarPostgres()
	if err != nil {
		log.Printf("error: %v", err)

	}

	cpfcnpj := utils.RemoveMask(dados.CPFCNPJ)

	if utils.IsValidCPF(cpfcnpj) {
		v, err := db.BuscaCpfPostgres(pg, cpfcnpj)
		if err != nil {
			utils.RespondWithError(w, http.StatusOK, 0, err.Error())
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, v)
		return
	}

	if utils.IsValidCNPJ(cpfcnpj) {
		v, err := db.BuscaCnpjPostgres(pg, cpfcnpj)
		if err != nil {
			utils.RespondWithError(w, http.StatusOK, 0, err.Error())
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, v)
		return

	}
	utils.RespondWithJSON(w, http.StatusBadRequest, "CPF ou CNPJ não encontrado!")

}

func RecebeCPFCNPJ(w http.ResponseWriter, r *http.Request) {
	var dados models.CPFCNPJ
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Erro ao ler o corpo da requisição")
		return
	}

	if err := json.Unmarshal(body, &dados); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0, "Falha ao realizar unmarshal do cliente")
		return
	}

	pg, err := db.ConectarPostgres()
	if err != nil {
		log.Printf("error: %v", err)

	}

	cpfcnpj := utils.RemoveMask(dados.CPFCNPJ)

	if len(cpfcnpj) == 11 {

		if !utils.IsValidCPF(cpfcnpj) {

			utils.RespondWithJSON(w, http.StatusBadRequest, "CPF inválido!")
			return
		}
		err = db.InsereCPFPostgres(pg, cpfcnpj)
		if err != nil {
			utils.RespondWithError(w, http.StatusOK, 0, err.Error())
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, "CPF inserido com sucesso!")
		return

	} else if len(cpfcnpj) == 14 {

		if !utils.IsValidCNPJ(cpfcnpj) {
			utils.RespondWithError(w, http.StatusBadRequest, 0, "CNPJ inválido!")
			return
		}
		dados.IsCPF = true
		err = db.InsereCNPJPostgres(pg, cpfcnpj)
		if err != nil {
			utils.RespondWithError(w, http.StatusOK, 0, err.Error())
			return
		}
		utils.RespondWithJSON(w, http.StatusBadRequest, "CNPJ inserido com sucesso!")
		return

	}

	utils.RespondWithJSON(w, http.StatusBadRequest, "CPF ou CNPJ inválido!")

}
