package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/go/qualityWater/src/models"
	"github.com/go/qualityWater/src/repository"
	"github.com/go/qualityWater/src/server"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var (
	ErrIotDeviceNotFound = errors.New("iotdevice not found")
)

type UpsertIotDeviceRequest struct {
	Data models.Data `json:"data"`
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func GetIotDeviceByHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		iotdevice, err := repository.GetGatewayById(r.Context(), params["id"])

		if iotdevice == nil {
			http.Error(w, ErrIotDeviceNotFound.Error(), http.StatusNotFound)
			return
		}

		if err != nil {
			log.Println("[error] ", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(iotdevice)
	}
}

func InsertIotDeviceByHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := Response{}

		defer json.NewEncoder(w).Encode(response)
		var iotdevice = UpsertIotDeviceRequest{}

		if err := json.NewDecoder(r.Body).Decode(&iotdevice); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Print("invalid body", err)
			response.Error = err.Error()
			return
		}

		// requiredFields := map[string]string{
		// 	"Temperature":        iotdevice.Data.Temperature,
		// 	"RelativeHumidity":   iotdevice.Data.RelativeHumidity,
		// 	"BarometricPressure": iotdevice.Data.BarometricPressure,
		// 	"RainFlow":           iotdevice.Data.RainFlow,
		// 	"PMtwoPointFive":     iotdevice.Data.PMtwoPointFive,
		// 	"PMTen":              iotdevice.Data.PMTen,
		// 	"CO":                 iotdevice.Data.CO,
		// 	"C2O":                iotdevice.Data.C2O,
		// }

		// var missingFields []string
		// for fieldName, value := range requiredFields {
		// 	if value == "" {
		// 		missingFields = append(missingFields, fieldName)
		// 	}
		// }

		// if len(missingFields) > 0 {
		// 	errorMessage := "Missing fields: " + strings.Join(missingFields, ", ")
		// 	http.Error(w, errorMessage, http.StatusBadRequest)
		// 	log.Print(errorMessage)
		// 	response.Error = errorMessage
		// 	return
		// }

		data := &models.Data{
			Temperature:        iotdevice.Data.Temperature,
			RelativeHumidity:   iotdevice.Data.RelativeHumidity,
			BarometricPressure: iotdevice.Data.BarometricPressure,
			RainFlow:           iotdevice.Data.RainFlow,
			PMtwoPointFive:     iotdevice.Data.PMtwoPointFive,
			PMTen:              iotdevice.Data.PMTen,
			CO:                 iotdevice.Data.CO,
			C2O:                iotdevice.Data.C2O,
		}
		newIotDevice := models.IotDevice{
			Id:   uuid.NewString(),
			Data: *data,
			Hour: time.Now(),
		}

		iotdeviceID, err := repository.InsertGateway(r.Context(), &newIotDevice)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Print("invalid id", err)
			response.Error = err.Error()
			return
		}

		response.Data = newIotDevice.Id
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{
			Data:  newIotDevice.Id,
			Error: "nil",
		})
		log.Print("iotdevice inserted with id ", iotdeviceID)
	}
}
