package app

import (
	"crypto/md5"
	"fmt"
	"sync"
	"time"
)

type AppState struct {
	sync.RWMutex
	Checksum   string
	UpdateTime time.Time
}

func NewAppState() *AppState {
	state := &AppState{}
	state.updateChecksum()
	go state.autoUpdateChecksum()
	return state
}

func (s *AppState) updateChecksum() {
	s.Lock()
	defer s.Unlock()

	timestamp := time.Now().Unix()
	data := fmt.Sprintf("%d", timestamp)
	hash := md5.Sum([]byte(data))
	s.Checksum = fmt.Sprintf("%x", hash)
	s.UpdateTime = time.Now()
}

func (s *AppState) autoUpdateChecksum() {
	for {
		now := time.Now().Unix()
		nextReload := ((now / 1000) + 1) * 1000
		waitDuration := time.Duration(nextReload-now) * time.Second
		
		time.Sleep(waitDuration)
		s.updateChecksum()
	}
}

func (s *AppState) GetChecksum() string {
	s.RLock()
	defer s.RUnlock()
	return s.Checksum
} 