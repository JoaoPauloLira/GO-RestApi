package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-rest-api/models"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personalidade := range models.Personalidades {
		if strconv.Itoa(personalidade.ID) == id {
			json.NewEncoder(w).Encode(personalidade)
		}
	}

}
