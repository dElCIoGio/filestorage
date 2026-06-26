package app

import (
	"github.com/dElCIoGio/filestorage/internal/modules/video-upload/application"
	"github.com/dElCIoGio/filestorage/internal/platform/config"
	"github.com/dElCIoGio/filestorage/internal/platform/events"
	"github.com/dElCIoGio/filestorage/internal/platform/workerspool"
	"log/slog"
)

type VideoUploadJob workerspool.Job[application.VideoUploadData]
type VideoUploadWorkerPool = workerspool.WorkerPool[application.VideoUploadData]

type FileStorageApp struct {
	Config                config.Config
	Logger                *slog.Logger
	Bus                   *events.Bus
	VideoUploadWorkerPool *VideoUploadWorkerPool
}

type Dependencies struct {
	cfg    config.Config
	logger *slog.Logger
	bus    events.Bus

	VideoUploadWorkerPool *VideoUploadWorkerPool
}

func New(deps *Dependencies) *FileStorageApp {
	return &FileStorageApp{
		// core
		Config: deps.cfg,
		Logger: deps.logger,
		Bus:    &deps.bus,

		// modules

		// workers
		VideoUploadWorkerPool: deps.VideoUploadWorkerPool,

		// others
	}
}
