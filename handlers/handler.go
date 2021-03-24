package handlers

import (
	"encoding/json"
	"github.com/Mishanki/rest-api-course/internal/core"
	"github.com/Mishanki/rest-api-course/model"
	"github.com/Mishanki/rest-api-course/pkg/helpers"
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
		errMsg := core.ErrorResponseStruct{Success: false, Message: "All params cant be equal to zero", InternalStatusCode: core.INVALID_ARG}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)

		return
	}

	file, _ := json.MarshalIndent(ct, "", " ")
	helpers.MyWrite(file)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ct)

}
