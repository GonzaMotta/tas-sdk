package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/GonzaMotta/tas-sdk/utils/helpers"
	"github.com/GonzaMotta/tas-sdk/utils/types"
	"github.com/joho/godotenv"
)

func main() {
	helpers.CustomLog("success", "Iniciando SDK")
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	helpers.CustomLog("success", "Creando auditoria")
	auditID := CreateAudit()

	log.Println("auditoria id : ", auditID)

	SaveTraceV2("Desarrollo servicio de prueba skd", "Datos de prueba", auditID)

	helpers.CustomLog("success", "Sdk inicio con éxito")
}

func CreateAudit() string {
	helpers.CustomLog("info", "Creando audit")

	url := os.Getenv("API_URL")

	response, err := http.Post(url+"/audit", "application/json", nil)

	if err != nil {
		helpers.CustomLog("error", "Error al crear audit")
	}

	defer response.Body.Close()

	auditResponse := types.CreateAuditResponse{}

	err = json.NewDecoder(response.Body).Decode(&auditResponse)

	if err != nil {
		helpers.CustomLog("error", "Error al decodificar el body")
	}

	log.Println(auditResponse.ID)

	return auditResponse.ID
}

func SaveTrace(trace *types.Trace) string {
	helpers.CustomLog("info", "Guardando trace")

	url := os.Getenv("API_URL")

	jsonData, err := json.Marshal(trace)
	if err != nil {
		log.Fatal("Error al convertir a JSON:", err)
	}

	response, err := http.Post(url+"/trace", "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		helpers.CustomLog("error", "Error al guardar trace")
	}

	defer response.Body.Close()

	traceResponse := types.CreateTraceResponse{}

	err = json.NewDecoder(response.Body).Decode(&traceResponse)

	if err != nil {
		helpers.CustomLog("error", "Error al decodificar el body")
	}

	log.Println(traceResponse.ID)
	return traceResponse.ID

}

func SaveTraceV2(serviceName string, data string, auditID string) string {
	helpers.CustomLog("info", "Guardando trace")

	url := os.Getenv("API_URL")

	trace := types.Trace{
		ServiceName: serviceName,
		Data:        data,
		Audit: types.HelperAudit{
			ID: auditID,
		},
	}

	jsonData, err := json.Marshal(trace)
	if err != nil {
		log.Fatal("Error al convertir a JSON:", err)
	}

	response, err := http.Post(url+"/trace", "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		helpers.CustomLog("error", "Error al guardar trace")
	}

	defer response.Body.Close()

	traceResponse := types.CreateTraceResponse{}

	err = json.NewDecoder(response.Body).Decode(&traceResponse)

	if err != nil {
		helpers.CustomLog("error", "Error al decodificar el body")
	}

	log.Println(traceResponse.ID)
	return traceResponse.ID

}
