package sdkTas

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/GonzaMotta/tas-sdk/utils/helpers"
	"github.com/GonzaMotta/tas-sdk/utils/types"
)

var url string = "http://localhost:3010"

/*
func init() {
	helpers.CustomLog("success", "Iniciando SDK")
	err := godotenv.Load(".env")

	url := os.Getenv("API_URL")

	log.Println("url", url)

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	helpers.CustomLog("success", "Sdk inicio con éxito")
}

*/

func CreateAudit() string {

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

func SaveStringTrace(serviceName string, data string, auditID string) string {

	trace := types.Trace{
		ServiceName: serviceName,
		Data:        data,
		Audit:       auditID,
		TypeValue:   "string",
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

	return traceResponse.ID

}

func SaveObjectTrace(serviceName string, data interface{}, auditID string) string {

	jsonString, err := json.Marshal(data)

	if err != nil {
		log.Fatal("Error al convertir a JSON:", err)
	}

	trace := types.Trace{
		ServiceName: serviceName,
		Data:        string(jsonString),
		Audit:       auditID,
		TypeValue:   "objecto",
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

	return traceResponse.ID

}
