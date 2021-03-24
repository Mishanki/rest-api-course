package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Mishanki/rest-api-course/helpers"
	"github.com/Mishanki/rest-api-course/model"
	"net/http"
)

var GetSolve = func(w http.ResponseWriter, r *http.Request) {
	calc := model.CalcArgs{}

	file := helpers.MyReadFile()
	json.Unmarshal(file, &calc)

	json.NewEncoder(w).Encode(calc)
}

var CreateGrab = func(w http.ResponseWriter, r *http.Request) {
	ct := model.CreateArgs{}
	helpers.BodyJsonDecode(w, r, &ct)

	if ct.A == 0 && ct.B == 0 && ct.C == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "All params cant be 0") // TODO json
		return
	}

	file, _ := json.MarshalIndent(ct, "", " ")
	helpers.MyWrite(file)
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json") // TODO не работает
	json.NewEncoder(w).Encode(ct)

}
