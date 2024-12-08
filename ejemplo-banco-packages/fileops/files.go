package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Mayúscula inicial para exportar la función
func WriteFloatInFile(fileName string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

// Mayúscula inicial para exportar la función
func ReadFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		WriteFloatInFile(fileName, 1000)
		return 1000, errors.New("no existe el archivo")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		WriteFloatInFile(fileName, 1000)
		return 1000, errors.New("el dato no se puede convertir a número")
	}

	return value, nil
}
