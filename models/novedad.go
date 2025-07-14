package models

type Novedad struct {
	Id                int    `json:"Id"`
	NumeroSolicitud   string `json:"NumeroSolicitud"`
	ContratoId        int    `json:"ContratoId"`
	NumeroCdpId       int    `json:"NumeroCdpId"`
	Motivo            string `json:"Motivo"`
	Aclaracion        string `json:"Aclaracion"`
	Observacion       string `json:"Observacion"`
	Vigencia          int    `json:"Vigencia"`
	VigenciaCdp       int    `json:"VigenciaCdp"`
	FechaCreacion     string `json:"FechaCreacion"`
	FechaModificacion string `json:"FechaModificacion"`
	Activo            bool   `json:"Activo"`
	TipoNovedad       int    `json:"TipoNovedad"`
	OficioSupervisor  string `json:"OficioSupervisor"`
	OficioOrdenador   string `json:"OficioOrdenador"`
	Estado            string `json:"Estado"`
	EnlaceDocumento   string `json:"EnlaceDocumento"`
}
