package app

import (
	"github.com/dElCIoGio/filestorage/internal/modules/video-upload/application"
	"github.com/dElCIoGio/filestorage/internal/platform/config"
	"github.com/dElCIoGio/filestorage/internal/platform/events"
	"github.com/dElCIoGio/filestorage/internal/platform/logging"
	"github.com/dElCIoGio/filestorage/internal/platform/workerspool"
)

func NewDev(cfg config.Config) (*FileStorageApp, error) {
	logger := logging.NewLogger()
	bus := events.NewInMemoryBus(logger)

	videoUploadWorkerPool := workerspool.NewWorkerPool[application.VideoUploadData](5)

	app := New(&Dependencies{
		cfg:    cfg,
		logger: logger,
		bus:    bus,

		VideoUploadWorkerPool: videoUploadWorkerPool,
	})

	return app, nil
}
