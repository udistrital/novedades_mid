// helpers/httpjson.go
package helpers

import (
	"fmt"

	"github.com/udistrital/utils_oas/request"
)

func GetDataListFromURL(url string) ([]map[string]interface{}, error) {
	var raw interface{}
	if err := request.GetJson(url, &raw); err != nil {
		return nil, err
	}
	if m, ok := raw.(map[string]interface{}); ok {
		if dm, ok2 := m["Data"].(map[string]interface{}); ok2 {
			return []map[string]interface{}{dm}, nil
		}
		if dm, ok2 := m["data"].(map[string]interface{}); ok2 {
			return []map[string]interface{}{dm}, nil
		}
		if da, ok2 := m["Data"].([]interface{}); ok2 {
			return ifaceSliceToMapSlice(da)
		}
		if da, ok2 := m["data"].([]interface{}); ok2 {
			return ifaceSliceToMapSlice(da)
		}
		if _, ok := m["Id"]; ok || m["id"] != nil || m["NumeroContrato"] != nil || m["numero_contrato"] != nil {
			return []map[string]interface{}{m}, nil
		}
		return nil, fmt.Errorf("respuesta sin campo Data válida en %s", url)
	}
	if arr, ok := raw.([]interface{}); ok {
		return ifaceSliceToMapSlice(arr)
	}
	return nil, fmt.Errorf("formato de respuesta no reconocido en %s", url)
}

func GetObjectFromURL(url string) (map[string]interface{}, error) {
	var raw interface{}
	if err := request.GetJson(url, &raw); err != nil {
		return nil, err
	}
	if m, ok := raw.(map[string]interface{}); ok {
		if d, ok2 := m["Data"].(map[string]interface{}); ok2 {
			return d, nil
		}
		if d, ok2 := m["data"].(map[string]interface{}); ok2 {
			return d, nil
		}
		return m, nil
	}
	return nil, fmt.Errorf("formato no soportado en %s", url)
}

func ifaceSliceToMapSlice(i interface{}) ([]map[string]interface{}, error) {
	arr, ok := i.([]interface{})
	if !ok {
		return nil, fmt.Errorf("tipo Data inesperado")
	}
	var out []map[string]interface{}
	for _, it := range arr {
		if m, ok := it.(map[string]interface{}); ok {
			out = append(out, m)
		}
	}
	return out, nil
}
