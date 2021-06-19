package services

import (
	"rest_api/app/knn"
	"rest_api/app/models"
	"rest_api/app/utils"
	"strconv"
	"strings"
)

func TrainingService(parameters models.Parameters) (best_K int, best_accuracy float64, time_elapsed string) {
	lines, err := utils.GetFileByUrl(utils.Url_data)

	if err != nil {
		print(err)
	}

	personas := CleanData(lines)
	best_K, best_accuracy, time_elapsed = knn.TrainingKNN(parameters.Epochs, parameters.ParallelProcs, personas)
	return best_K, best_accuracy, time_elapsed
}

func CleanData(lines []string) (data []utils.PersonaEncuestada) {
	data = make([]utils.PersonaEncuestada, 0)
	for i, line := range lines {
		attributes := strings.Split(line, "|")
		if i != 0 {
			var ingreso_monetario float64
			var condicion_laboral string

			if attributes[67] == "" {
				ingreso_monetario = 0
			} else {
				ingreso_monetario, _ = strconv.ParseFloat(attributes[67], 64)
			}

			if attributes[92] == "1" {
				condicion_laboral = "Empleado"
			} else if attributes[92] == "2" {
				condicion_laboral = "Desempleado Abierto"
			} else {
				condicion_laboral = "Desempleado Oculto"
			}

			sexo, _ := strconv.ParseFloat(attributes[14], 64)
			edad, _ := strconv.ParseFloat(attributes[15], 64)
			etnia, _ := strconv.ParseFloat(attributes[90], 64)
			nivel_educativo, _ := strconv.ParseFloat(attributes[16], 64)
			ultimo_cargo, _ := strconv.ParseFloat(attributes[47], 64)
			frecuencia_pago, _ := strconv.ParseFloat(attributes[66], 64)
			seguro_salud, _ := strconv.ParseFloat(attributes[82], 64)

			persona := utils.PersonaEncuestada{
				Data:  []float64{sexo, edad, etnia, nivel_educativo, ultimo_cargo, frecuencia_pago, ingreso_monetario, seguro_salud},
				Class: condicion_laboral,
			}
			data = append(data, persona)
		}
	}

	return data
}
