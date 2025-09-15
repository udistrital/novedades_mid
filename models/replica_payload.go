package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type PreliquidacionReplica struct {
	Activo            bool        `json:"Activo"`
	Cdp               int         `json:"Cdp"`
	Completo          bool        `json:"Completo"`
	DependenciaId     int         `json:"DependenciaId"`
	Desagregado       interface{} `json:"Desagregado"`
	Documento         string      `json:"Documento"`
	FechaCreacion     string      `json:"FechaCreacion"`
	FechaFin          string      `json:"FechaFin"`
	FechaInicio       string      `json:"FechaInicio"`
	FechaModificacion string      `json:"FechaModificacion"`
	Id                int         `json:"Id"`
	NombreCompleto    string      `json:"NombreCompleto"`
	NumeroContrato    string      `json:"NumeroContrato"`
	NumeroSemanas     int         `json:"NumeroSemanas"`
	PersonaId         int         `json:"PersonaId"`
	ProyectoId        int         `json:"ProyectoId"`
	Resolucion        string      `json:"Resolucion"`
	ResolucionId      int         `json:"ResolucionId"`
	Rp                int         `json:"Rp"`
	TipoNominaId      int         `json:"TipoNominaId"`
	Unico             bool        `json:"Unico"`
	Vacaciones        int         `json:"Vacaciones"`
	ValorContrato     float64     `json:"ValorContrato"`
	Vigencia          int         `json:"Vigencia"`
}

// --- Helpers nuevos: fijar a 12:00 y formatear RFC3339 conservando zona ---
func atNoon(t time.Time) time.Time {
	if t.IsZero() {
		return t
	}
	return time.Date(t.Year(), t.Month(), t.Day(), 12, 0, 0, 0, t.Location())
}

func formatRFC3339LocalNoon(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return atNoon(t).Format(time.RFC3339) // sin UTC(): mantiene offset local
}

func NewPreliquidacionReplicaFromContrato(contrato map[string]interface{}, fIni, fFin time.Time) PreliquidacionReplica {
	return PreliquidacionReplica{
		Activo:            toBool(pick(contrato, "Activo", "activo")),
		Cdp:               toInt(pick(contrato, "Cdp", "cdp")),
		Completo:          false,
		DependenciaId:     toInt(pick(contrato, "DependenciaId", "dependencia_id")),
		Desagregado:       nil,
		Documento:         toString(pick(contrato, "Documento", "documento")),
		FechaCreacion:     toString(pick(contrato, "FechaCreacion", "fecha_creacion")),
		FechaFin:          formatRFC3339LocalNoon(fFin), // <-- 12:00 local
		FechaInicio:       formatRFC3339LocalNoon(fIni), // <-- 12:00 local
		FechaModificacion: toString(pick(contrato, "FechaModificacion", "fecha_modificacion")),
		Id:                toInt(pick(contrato, "Id", "id")),
		NombreCompleto:    toString(pick(contrato, "NombreCompleto", "nombre_completo")),
		NumeroContrato:    toString(pick(contrato, "NumeroContrato", "numero_contrato")),
		NumeroSemanas:     toInt(pick(contrato, "NumeroSemanas", "numero_semanas")),
		PersonaId:         toInt(pick(contrato, "PersonaId", "persona_id")),
		ProyectoId:        toInt(pick(contrato, "ProyectoId", "proyecto_id")),
		Resolucion:        toString(pick(contrato, "Resolucion", "resolucion")),
		ResolucionId:      toInt(pick(contrato, "ResolucionId", "resolucion_id")),
		Rp:                toInt(pick(contrato, "Rp", "rp")),
		TipoNominaId:      toInt(pick(contrato, "TipoNominaId", "tipo_nomina_id")),
		Unico:             toBool(pick(contrato, "Unico", "unico")),
		Vacaciones:        toInt(pick(contrato, "Vacaciones", "vacaciones")),
		ValorContrato:     toFloat64(pick(contrato, "ValorContrato", "valor_contrato")),
		Vigencia:          toInt(pick(contrato, "Vigencia", "vigencia")),
	}
}

func (p PreliquidacionReplica) AsMap() map[string]interface{} {
	b, _ := json.Marshal(p)
	var out map[string]interface{}
	_ = json.Unmarshal(b, &out)
	return out
}

func pick(m map[string]interface{}, keys ...string) interface{} {
	for _, k := range keys {
		if v, ok := m[k]; ok {
			if s, ok2 := v.(string); ok2 {
				if strings.TrimSpace(s) != "" {
					return s
				}
				continue
			}
			if v != nil {
				return v
			}
		}
	}
	return nil
}

func toString(v interface{}) string {
	if v == nil {
		return ""
	}
	switch t := v.(type) {
	case string:
		return t
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64)
	case int:
		return strconv.Itoa(t)
	case int64:
		return strconv.FormatInt(t, 10)
	case bool:
		if t {
			return "true"
		}
		return "false"
	default:
		return fmt.Sprintf("%v", t)
	}
}

func toInt(v interface{}) int {
	switch t := v.(type) {
	case nil:
		return 0
	case int:
		return t
	case int64:
		return int(t)
	case float64:
		return int(t)
	case string:
		i, _ := strconv.Atoi(strings.TrimSpace(t))
		return i
	default:
		return 0
	}
}

func toFloat64(v interface{}) float64 {
	switch t := v.(type) {
	case nil:
		return 0
	case float64:
		return t
	case int:
		return float64(t)
	case int64:
		return float64(t)
	case string:
		f, _ := strconv.ParseFloat(strings.TrimSpace(t), 64)
		return f
	default:
		return 0
	}
}

func toBool(v interface{}) bool {
	switch t := v.(type) {
	case bool:
		return t
	case string:
		s := strings.TrimSpace(strings.ToLower(t))
		return s == "true" || s == "1" || s == "t" || s == "yes" || s == "y"
	case float64:
		return t != 0
	case int:
		return t != 0
	case int64:
		return t != 0
	default:
		return false
	}
}
