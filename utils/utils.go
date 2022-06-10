package utils

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
)

var (
	allNumRe  = regexp.MustCompile("[^0-9]")
	allLetRe  = regexp.MustCompile("[^a-zA-Z]")
	allSame14 = regexp.MustCompile("0{14}|1{14}|2{14}|3{14}|4{14}|5{14}|6{14}|7{14}|8{14}|9{14}")
	allSame11 = regexp.MustCompile("0{11}|1{11}|2{11}|3{11}|4{11}|5{11}|6{11}|7{11}|8{11}|9{11}")
)

func IsValidCNPJ(cnpj string) bool {
	n := allNumRe.ReplaceAllString(cnpj, "")
	if n == "" {
		return false
	}
	if len(n) != 14 {
		return false
	}
	if allSame14.Match([]byte(n)) {
		return false
	}
	if n[:3] == "000" {
		return false
	}
	size := len(n) - 2
	numbers := n[0:size]
	digits := n[size:]
	var sum int
	pos := size - 7
	for i := size; i >= 1; i-- {
		num, _ := strconv.Atoi(string(numbers[size-i]))
		sum += num * pos
		pos = pos - 1
		if pos < 2 {
			pos = 9
		}
	}
	var result int
	if sum%11 < 2 {
		result = 0
	} else {
		result = 11 - sum%11
	}
	x, _ := strconv.Atoi(string(digits[0]))
	if result != x {
		return false
	}
	size = size + 1
	numbers = n[0:size]
	sum = 0
	pos = size - 7
	for i := size; i >= 1; i-- {
		num, _ := strconv.Atoi(string(numbers[size-i]))
		sum += num * pos
		pos = pos - 1
		if pos < 2 {
			pos = 9
		}
	}
	if sum%11 < 2 {
		result = 0
	} else {
		result = 11 - sum%11
	}
	num, _ := strconv.Atoi(string(digits[1]))
	if result != num {
		return false
	}
	return true
}

func IsValidCPF(cpf string) bool {
	cpf = allNumRe.ReplaceAllString(cpf, "")
	if cpf == "" {
		return false
	}
	if len(cpf) != 11 {
		return false
	}
	if allSame11.Match([]byte(cpf)) {
		return false
	}
	var sum int
	var res int
	for i := 1; i <= 9; i++ {
		num, _ := strconv.Atoi(cpf[i-1 : i])
		sum = sum + num*(11-i)
	}
	res = (sum * 10) % 11
	if (res == 10) || (res == 11) {
		res = 0
	}
	num, _ := strconv.Atoi(cpf[9:10])
	if res != num {
		return false
	}
	sum = 0
	for i := 1; i <= 10; i++ {
		num, _ := strconv.Atoi(cpf[i-1 : i])
		sum = sum + num*(12-i)
	}
	res = (sum * 10) % 11
	if (res == 10) || (res == 11) {
		res = 0
	}
	num, _ = strconv.Atoi(cpf[10:11])
	if res != num {
		return false
	}
	return true
}

func OnlyNumbers(s string) string {
	return allNumRe.ReplaceAllString(s, "")
}

func NormalizeCPFCNPJ(cpfcnpj string) (string, bool) {

	cpfcnpj = OnlyNumbers(cpfcnpj)
	if IsValidCNPJ(cpfcnpj) {
		return cpfcnpj, true
	}

	if len(cpfcnpj) > 11 {
		cpfcnpj = cpfcnpj[3:]
	}
	if IsValidCPF(cpfcnpj) {
		return cpfcnpj, true
	}

	return cpfcnpj, false

}

func RespondWithError(w http.ResponseWriter, code, errorCode int, message string) {
	RespondWithJSON(w, code, map[string]interface{}{"error": message, "code": errorCode})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
