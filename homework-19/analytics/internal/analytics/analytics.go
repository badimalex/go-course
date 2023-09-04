package analytics

import (
	"sync"
	"time"
)

type Service struct {
	mu           sync.Mutex
	total        int
	unique       map[string]struct{}
	longest      string
	shortest     string
	history      map[string]time.Time
	totalLinkLen int
}

func New() *Service {
	return &Service{
		unique:  make(map[string]struct{}),
		history: make(map[string]time.Time),
	}
}

func (s *Service) AddInfo(orig, short string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.total++

	if _, exists := s.unique[orig]; !exists {
		s.unique[orig] = struct{}{}
	}

	if len(orig) > len(s.longest) {
		s.longest = orig
	}
	if len(s.shortest) == 0 || len(orig) < len(s.shortest) {
		s.shortest = orig
	}

	s.totalLinkLen += len(orig)
	s.history[short] = time.Now()
}

func (s *Service) Stats() map[string]interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()

	avgLen := 0
	if s.total > 0 {
		avgLen = s.totalLinkLen / s.total
	}

	return map[string]interface{}{
		"total":    s.total,
		"unique":   len(s.unique),
		"avg_len":  avgLen,
		"longest":  s.longest,
		"shortest": s.shortest,
		"history":  s.history,
	}
}
