package repository

import (
	"context"
	"log"

	"github.com/go/qualityWater/src/models"
)

type Repository interface {
	InsertIotDevice(ctx context.Context, Gateway *models.IotDevice) (interface{}, error)
	GetIotDeviceById(ctx context.Context, id string) (*models.IotDevice, error)
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func InsertGateway(ctx context.Context, iotdevice *models.IotDevice) (interface{}, error) {
	return implementation.InsertIotDevice(ctx, iotdevice)
}

func GetGatewayById(ctx context.Context, id string) (*models.IotDevice, error) {
	log.Print("id: ", id)
	return implementation.GetIotDeviceById(ctx, id)
}
