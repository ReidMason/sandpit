package messageHandler

import (
	"log/slog"
	"strconv"
	"time"
	"worker/internal/storage"
)

type MessageHandler struct {
	logger  *slog.Logger
	storage *storage.PostgresStorage
}

func NewMessageHandler(logger *slog.Logger, db *storage.PostgresStorage) *MessageHandler {
	return &MessageHandler{
		logger:  logger,
		storage: db,
	}
}

func (m *MessageHandler) HandleMessage(message string) (bool, error) {
	id, err := strconv.Atoi(message)
	if err != nil {
		m.logger.Warn("Invalid id found in queue", slog.String("message", message))
		return true, nil
	}

	m.logger.Info("Getting dataset")
	processing, data, err := m.storage.GetDataset(id)
	if err != nil {
		m.logger.Error("Failed to get dataset", slog.Any("error", err))
		return true, err
	}

	if processing {
		m.logger.Info("Dataset is already being processed", slog.Int("id", id))
		return true, nil
	}

	defer func() {
		err := m.storage.UpdateData(id, false)
		if err != nil {
			m.logger.Error("Failed to update data", slog.Any("error", err))
		}
	}()

	m.logger.Info("sleeping")
	time.Sleep(1 * time.Second)
	m.logger.Info("woke")

	for _, item := range data {
		album, err := m.storage.GetAlbum(item.AlbumId)
		if err != nil {
			m.logger.Error("Failed to get album", slog.Any("error", err))
			return false, err
		}

		if album.Id == -1 {
			albumId, err := m.storage.CreateAlbum(item.AlbumId, "Album "+strconv.Itoa(item.AlbumId))
			if err != nil {
				m.logger.Error("Failed to create album", slog.Any("error", err))
				return false, err
			}

			album.Id = albumId
		}

		err = m.storage.CreatePhoto(item.Id, item.Title, album.Id)
		if err != nil {
			m.logger.Error("Failed to create photo", slog.Any("error", err))
			return false, err
		}
	}

	m.logger.Info("Finished processing message")
	return true, nil
}
