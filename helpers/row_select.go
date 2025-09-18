package helpers

import (
	"sort"
	"strconv"
	"strings"
	"time"
)

// ---- API que usa anular_novedad_service.go ----

// IsAdicionOProrroga: ajusta aquí los códigos si cambian en tu negocio.
// 8 = Adición, 1 = Prórroga (según tu código actual).
func IsAdicionOProrroga(tipo int) bool { return tipo == 8 || tipo == 1 }

// PickRowPreferExtraRP:
// - Si hay filas con RP distinto al RP "base" (contrato original), retorna la más reciente entre esas.
// - Si no hay RP distinto, retorna la más reciente de todas.
// - Fallback implícito: si solo hay 1, retorna esa.
func PickRowPreferExtraRP(rows []map[string]interface{}) map[string]interface{} {
	if len(rows) == 0 {
		return nil
	}
	base := pickInitialRowLocal(rows) // más antigua (original)
	baseRP := rowGetIntKeys(base, "Rp", "rp")

	var extras []map[string]interface{}
	for _, r := range rows {
		if rowGetIntKeys(r, "Rp", "rp") != baseRP {
			extras = append(extras, r)
		}
	}

	// 1) Preferir la más reciente con RP distinto
	if len(extras) > 0 {
		sortRowsByRecientes(extras)
		return extras[0]
	}

	// 2) Si no hay RP distinto, tomar la más reciente de todas
	all := append([]map[string]interface{}(nil), rows...) // copia para no mutar
	sortRowsByRecientes(all)
	return all[0]
}

// ---- Helpers internos (solo para este archivo) ----

func rowGetIntKeys(m map[string]interface{}, keys ...string) int {
	for _, k := range keys {
		if v, ok := m[k]; ok {
			switch t := v.(type) {
			case int:
				return t
			case int64:
				return int(t)
			case float64:
				return int(t)
			case string:
				i, _ := strconv.Atoi(strings.TrimSpace(t))
				return i
			}
		}
	}
	return 0
}

// Ordena DESC por FechaCreacion; si empata, DESC por Id (más nuevo primero)
func sortRowsByRecientes(rows []map[string]interface{}) {
	sort.Slice(rows, func(i, j int) bool {
		ti := rowParseFechaCreacion(rows[i])
		tj := rowParseFechaCreacion(rows[j])
		if ti.Equal(tj) {
			return rowGetID(rows[i]) > rowGetID(rows[j])
		}
		return ti.After(tj)
	})
}

// La más antigua por fecha (y si empata, por Id asc). Emula el comportamiento previo.
func pickInitialRowLocal(rows []map[string]interface{}) map[string]interface{} {
	cp := append([]map[string]interface{}(nil), rows...)
	sort.Slice(cp, func(i, j int) bool {
		ti := rowParseFechaCreacion(cp[i])
		tj := rowParseFechaCreacion(cp[j])
		if ti.Equal(tj) {
			return rowGetID(cp[i]) < rowGetID(cp[j])
		}
		return ti.Before(tj)
	})
	return cp[0]
}

func rowGetID(m map[string]interface{}) int {
	if v, ok := m["Id"]; ok {
		switch t := v.(type) {
		case int:
			return t
		case int64:
			return int(t)
		case float64:
			return int(t)
		case string:
			i, _ := strconv.Atoi(strings.TrimSpace(t))
			return i
		}
	}
	return 0
}

func rowParseFechaCreacion(m map[string]interface{}) time.Time {
	// intenta varias llaves comunes
	for _, k := range []string{"FechaCreacion", "fecha_creacion", "CreatedAt"} {
		if v, ok := m[k]; ok {
			if t := parseTimeLoose(v); !t.IsZero() {
				return t
			}
		}
	}
	// fallback: usa el Id como epoch (si es numérico tipo "segundos")
	if id := rowGetID(m); id > 0 {
		return time.Unix(int64(id), 0)
	}
	return time.Time{}
}

// parseTimeLoose intenta algunos formatos comunes además de RFC3339
func parseTimeLoose(v interface{}) time.Time {
	switch x := v.(type) {
	case time.Time:
		return x
	case string:
		s := strings.TrimSpace(x)
		if s == "" {
			return time.Time{}
		}
		// intentos en cascada
		if t, err := time.Parse(time.RFC3339, s); err == nil {
			return t
		}
		if t, err := time.Parse("2006-01-02 15:04:05", s); err == nil {
			return t
		}
		if t, err := time.Parse("2006-01-02", s); err == nil {
			return t
		}
		return time.Time{}
	case float64:
		return time.Unix(int64(x), 0)
	case int:
		return time.Unix(int64(x), 0)
	case int64:
		return time.Unix(x, 0)
	default:
		return time.Time{}
	}
}
