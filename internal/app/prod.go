package app

import (
	"github.com/dElCIoGio/filestorage/internal/modules/signing"
	signingService "github.com/dElCIoGio/filestorage/internal/modules/signing/service"
	"github.com/dElCIoGio/filestorage/internal/modules/video-upload/application"
	"github.com/dElCIoGio/filestorage/internal/platform/config"
	"github.com/dElCIoGio/filestorage/internal/platform/events"
	"github.com/dElCIoGio/filestorage/internal/platform/logging"
	"github.com/dElCIoGio/filestorage/internal/platform/workerspool"
)

func NewProd(cfg config.Config) (*FileStorageApp, error) {
	logger := logging.NewLogger()
	bus := events.NewInMemoryBus(logger)

	videoUploadWorkerPool := workerspool.NewWorkerPool[application.VideoUploadData](3, 30)

	service := signingService.NewFakeService()
	signingModule := signing.NewModule(service)

	app := New(&Dependencies{
		cfg:    cfg,
		logger: logger,
		bus:    bus,

		VideoUploadWorkerPool: videoUploadWorkerPool,
		SigningModule:         *signingModule,
	})

	return app, nil
}
