package models

type IotDevice struct {
	Id   string `json:"id"`
	Data Data   `json:"data"`
}

type Data struct {
	Temperature        float64 `json:"temperature"`
	RelativeHumidity   float64 `json:"relativehumidity"`
	BarometricPressure float64 `json:"barometricpressure"`
	RainFlow           float64 `json:"rainflow"`
	PMtwoPointFive     float64 `json:"PM2.5"`
	PMTen              float64 `json:"PM10"`
	CO                 float64 `json:"CO"`
	C2O                float64 `json:"C2O"`
}
