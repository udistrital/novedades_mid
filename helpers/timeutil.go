// helpers/timeutil.go
package helpers

import (
	"strings"
	"time"
)

func ParseTimeAny(v interface{}) time.Time {
	switch x := v.(type) {
	case string:
		s := strings.TrimSpace(x)
		if s == "" {
			return time.Time{}
		}
		if i := strings.Index(s, " +"); i > 0 {
			s = s[:i]
		}
		formats := []string{
			time.RFC3339,
			"2006-01-02 15:04:05.999999",
			"2006-01-02 15:04:05",
			"2006-01-02",
			"02/01/2006",
			"02-01-2006",
			"2006/01/02",
			"2006-01-02T15:04:05.000Z",
			"2006-01-02T15:04:05.000-07:00",
		}
		for _, f := range formats {
			if t, err := time.Parse(f, s); err == nil {
				return t
			}
		}
	case map[string]interface{}:
		if v, ok := x["Time"]; ok {
			return ParseTimeAny(v)
		}
	}
	return time.Time{}
}

func MesesEntre(inicio, fin time.Time) [][2]int {
	if fin.Before(inicio) {
		inicio, fin = fin, inicio
	}
	inicio = time.Date(inicio.Year(), time.Month(inicio.Month()), 1, 0, 0, 0, 0, time.UTC)
	fin = time.Date(fin.Year(), time.Month(fin.Month()), 1, 0, 0, 0, 0, time.UTC)

	var res [][2]int
	cur := inicio
	for !cur.After(fin) {
		res = append(res, [2]int{int(cur.Month()), cur.Year()})
		cur = cur.AddDate(0, 1, 0)
	}
	return res
}

func MesesEntreSafe(inicio, fin time.Time) [][2]int {
	if inicio.IsZero() || fin.IsZero() {
		return [][2]int{}
	}
	return MesesEntre(inicio, fin)
}

func MonthBoundsClipped(ano, mes int, globalIni, globalFin time.Time) (time.Time, time.Time) {
	start := time.Date(ano, time.Month(mes), 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0).Add(-time.Second)
	if start.Before(globalIni) {
		start = globalIni
	}
	if end.After(globalFin) {
		end = globalFin
	}
	return start, end
}
