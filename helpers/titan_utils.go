package helpers

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

// --- Nuevo helper: normaliza cualquier time.Time al mediodía (12:00:00.000) ---
func atNoon(t time.Time) time.Time {
	if t.IsZero() {
		return t
	}
	return time.Date(t.Year(), t.Month(), t.Day(), 12, 0, 0, 0, t.Location())
}

func GetRowId(m map[string]interface{}) int {
	switch v := m["Id"].(type) {
	case float64:
		return int(v)
	case string:
		i, _ := strconv.Atoi(v)
		return i
	default:
		return 0
	}
}

func GetBool(m map[string]interface{}, k string) bool {
	if v, ok := m[k]; ok {
		switch t := v.(type) {
		case bool:
			return t
		case string:
			return t == "true" || t == "1"
		case float64:
			return t != 0
		}
	}
	return false
}

func ParseFechaCreacion(m map[string]interface{}) time.Time {
	for _, k := range []string{"FechaCreacion", "fecha_creacion", "CreatedAt"} {
		if v, ok := m[k]; ok {
			if t := ParseTimeAny(v); !t.IsZero() {
				return atNoon(t) // <-- fuerza 12:00:00.000
			}
		}
	}
	if id := GetRowId(m); id > 0 {
		// Si cae por Id como Unix, también se normaliza a 12:00
		return atNoon(time.Unix(int64(id), 0))
	}
	return time.Time{}
}

func PickInitialRow(rows []map[string]interface{}) map[string]interface{} {
	if len(rows) == 0 {
		return nil
	}
	sort.Slice(rows, func(i, j int) bool {
		ti := ParseFechaCreacion(rows[i])
		tj := ParseFechaCreacion(rows[j])
		if ti.Equal(tj) {
			return GetRowId(rows[i]) < GetRowId(rows[j])
		}
		return ti.Before(tj)
	})
	return rows[0]
}

func UpdateRowFechasActivo(row map[string]interface{}, fIni, fFin time.Time, activo bool) map[string]interface{} {
	// Normaliza siempre a 12:00 antes de guardar
	fIni = atNoon(fIni)
	fFin = atNoon(fFin)

	row["FechaInicio"] = fIni.Format(time.RFC3339)
	row["FechaFin"] = fFin.Format(time.RFC3339)
	row["FechaModificacion"] = time.Now().Format("2006-01-02 15:04:05")
	row["Activo"] = activo
	return row
}

func FechasDeRow(row map[string]interface{}) (time.Time, time.Time) {
	var fi, ff time.Time
	for _, k := range []string{"FechaInicio", "fecha_inicio"} {
		if v, ok := row[k]; ok {
			fi = atNoon(ParseTimeAny(v)) // <-- fuerza 12:00
			break
		}
	}
	for _, k := range []string{"FechaFin", "fecha_fin"} {
		if v, ok := row[k]; ok {
			ff = atNoon(ParseTimeAny(v)) // <-- fuerza 12:00
			break
		}
	}
	if ff.IsZero() && !fi.IsZero() {
		ff = fi
	}
	return fi, ff
}

func SetValorContrato(row map[string]interface{}, v float64) {
	row["ValorContrato"] = v
	row["valor_contrato"] = v
}

func ErrorString(err error) string {
	if err == nil {
		return ""
	}
	return fmt.Sprintf("%v", err)
}
