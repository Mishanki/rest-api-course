package model

type CreateArgs struct {
	A int `json:"A"`
	B int `json:"B"`
	C int `json:"C"`
}

type CalcArgs struct {
	CreateArgs
	Nroots int `json:"Nroots"`
}
