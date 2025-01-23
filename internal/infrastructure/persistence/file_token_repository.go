package persistence

import (
	"cursor-api/internal/domain/model"
	"encoding/json"
	"os"
	"sync"
)

type FileTokenRepository struct {
	filePath string
	mutex    sync.RWMutex
}

func NewFileTokenRepository(filePath string) *FileTokenRepository {
	return &FileTokenRepository{
		filePath: filePath,
	}
}

func (r *FileTokenRepository) GetAll() ([]model.TokenInfo, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	data, err := os.ReadFile(r.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.TokenInfo{}, nil
		}
		return nil, err
	}

	var tokens []model.TokenInfo
	if err := json.Unmarshal(data, &tokens); err != nil {
		return nil, err
	}

	return tokens, nil
}

func (r *FileTokenRepository) SaveAll(tokens []model.TokenInfo) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	data, err := json.MarshalIndent(tokens, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.filePath, data, 0644)
}

func (r *FileTokenRepository) Reload() ([]model.TokenInfo, error) {
	return r.GetAll()
} 