package utils

import (
	"fmt"
	"strings"
	"time"
)

// CustomTime adalah alias untuk time.Time
type CustomTime time.Time

// UnmarshalJSON mengurai string JSON dengan format kustom "DD-MM-YYYY HH:mm"
func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	if s == "" || s == "null" {
		return nil
	}

	// Tentukan lokasi waktu (contoh: WIB)
	wibLocation := time.FixedZone("WIB", 7*60*60)

	// Parsing string menggunakan layout "02-01-2006 15:04"
	t, err := time.ParseInLocation("02-01-2006 15:04", s, wibLocation)
	if err != nil {
		return fmt.Errorf("invalid date format: %s", s)
	}
	*ct = CustomTime(t)
	return nil
}
