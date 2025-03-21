package tools

import (
	"fmt"
	"time"
)

func FechaMysql() string {
	t := time.Now()
	// ESTO ES IMPORTANTE SE FORZA PARA QUE SI O SI LO FORMATEE EN 2 DIGITOS POR FORMATO CORRECTO EN MYSQL PARA EVITAR ERROR
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

// esto es muy importante para el motor de mysql en golang para que se formatee, muy importante
