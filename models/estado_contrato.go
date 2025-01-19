package models

import (
	"time"
)

type EstadoContrato struct {
	NombreEstado  string
	FechaRegistro time.Time
	Id            int
}

type CambioEstado struct {
	Estado struct {
		Id int `json:"Id"`
	} `json:"Estado"`
	FechaRegistro  string `json:"FechaRegistro"`
	NumeroContrato string `json:"NumeroContrato"`
	Usuario        string `json:"Usuario"`
	Vigencia       int    `json:"Vigencia"`
}
