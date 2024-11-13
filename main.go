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

func CreateAudit() (string, error) {

	response, err := http.Post(url+"/audit", "application/json", nil)

	if err != nil {
		helpers.CustomLog("error", "Error al crear audit")
	}

	defer response.Body.Close()

	auditResponse := types.CreateAuditResponse{}

	err = json.NewDecoder(response.Body).Decode(&auditResponse)

	if err != nil {
		helpers.CustomLog("error", "Error al decodificar el body")

		return "", err
	}

	log.Println(auditResponse.ID)

	return auditResponse.ID, nil
}

func SaveStringTrace(serviceName string, data string, auditID string) (string, error) {

	trace := types.Trace{
		ServiceName: serviceName,
		Data:        data,
		Audit:       auditID,
		TypeValue:   "string",
	}

	jsonData, err := json.Marshal(trace)
	if err != nil {
		log.Fatal("Error al convertir a JSON:", err)

		return "", err
	}

	response, err := http.Post(url+"/trace", "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		helpers.CustomLog("error", "Error al guardar trace")

		return "", err
	}

	defer response.Body.Close()

	traceResponse := types.CreateTraceResponse{}

	err = json.NewDecoder(response.Body).Decode(&traceResponse)

	if err != nil {
		helpers.CustomLog("error", "Error al decodificar el body")

		return "", err
	}

	return traceResponse.ID, nil

}

func SaveObjectTrace(serviceName string, data interface{}, auditID string) (string, error) {

	jsonString, err := json.Marshal(data)

	if err != nil {
		log.Fatal("Error al convertir a JSON:", err)

		return "", err
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

		return "", err
	}

	response, err := http.Post(url+"/trace", "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		helpers.CustomLog("error", "Error al guardar trace")

		return "", err
	}

	defer response.Body.Close()

	traceResponse := types.CreateTraceResponse{}

	err = json.NewDecoder(response.Body).Decode(&traceResponse)

	if err != nil {
		helpers.CustomLog("error", "Error al decodificar el body")

		return "", err
	}

	return traceResponse.ID, nil

}
