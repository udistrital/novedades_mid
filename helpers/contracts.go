package helpers

func AjustarValorContratoParaAnulacion(initial map[string]interface{}, rows []map[string]interface{}, tipoNovedad int) {
	if tipoNovedad != 1 && tipoNovedad != 2 {
		return
	}
	initialId := GetRowId(initial)
	orig := ParseFloat(Pick(initial, "ValorContrato", "valor_contrato"))
	sumExtras := 0.0
	for _, r := range rows {
		if GetRowId(r) == initialId {
			continue
		}
		sumExtras += ParseFloat(Pick(r, "ValorContrato", "valor_contrato"))
	}
	SetValorContrato(initial, orig+sumExtras)
}
