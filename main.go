package sdkTas

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"reflect"

	"github.com/GonzaMotta/tas-sdk/utils/helpers"
	"github.com/GonzaMotta/tas-sdk/utils/types"
)

var url string = "http://localhost:3010"

var globalAuditId types.CreateAuditResponse

func CreateAudit() (string, error) {

	response, err := http.Post(url+"/audit", "application/json", nil)

	if err != nil {
		helpers.CustomLog("error", "Error al crear audit")
		return "", err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&globalAuditId)

	if err != nil {
		helpers.CustomLog("error", "Error al decodificar el body")

		return "", err
	}

	return globalAuditId.ID, nil
}

func SaveTrace(serviceName string, data any, auditId ...string) (string, error) {

	if rv := reflect.ValueOf(data); rv.Kind() != reflect.String {
		data, _ = json.Marshal(data)
		data = string(data.([]byte))
	}

	trace := types.Trace{
		ServiceName: serviceName,
		Data:        data.(string),
		Audit:       globalAuditId.ID,
	}

	if len(auditId) > 0 {
		trace.Audit = auditId[0]
	}

	if trace.Audit == "" && globalAuditId.ID == "" {
		auditId, err := CreateAudit()
		if err != nil {
			helpers.CustomLog("error", "Error al crear audit")
		}

		trace.Audit = auditId
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
