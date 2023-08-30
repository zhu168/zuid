package zuid

import (
	"encoding/binary"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

const (
	maxWorkerID  = 1023
	maxSequence  = 4095
	workerIDBits = 10
	sequenceBits = 12
)

var Epoch = uint64(1682908800000) // 2023-08-30 00:00:00 UTC timestamp

type ZUID struct {
	mu       sync.Mutex
	workerID uint64
	lastTime uint64
	sequence uint64
}

func NewZUID(workerID uint64) (*ZUID, error) {
	if workerID > maxWorkerID {
		return nil, fmt.Errorf("worker ID must be between 0 and %d", maxWorkerID)
	}
	return &ZUID{
		workerID: workerID,
	}, nil
}

func (s *ZUID) nextID() (uint64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	currentTime := uint64(time.Now().UnixNano()) / uint64(time.Millisecond)

	if currentTime < s.lastTime {
		return 0, fmt.Errorf("clock is moving backwards, waiting until %d", s.lastTime)
	}

	if currentTime == s.lastTime {
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			for currentTime <= s.lastTime {
				currentTime = uint64(time.Now().UnixNano()) / uint64(time.Millisecond)
			}
		}
	} else {
		s.sequence = 0
	}

	s.lastTime = currentTime

	ID := ((currentTime - Epoch) << (workerIDBits + sequenceBits)) |
		(s.workerID << sequenceBits) |
		s.sequence

	return ID, nil
}

func (s *ZUID) NextID() ([]byte, error) {
	sfid, err := s.nextID()
	if err != nil {
		return nil, err
	}
	var sfbyte8 [8]byte
	binary.BigEndian.PutUint64(sfbyte8[:], sfid)
	sfslice := sfbyte8[:]
	// Generate a UUID
	u, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	// Get the last 64 bits of the UUID and convert it to a hexadecimal string
	uuidbyte16 := ([16]byte(u))
	uuidslice := uuidbyte16[8:]
	data := append(sfslice, uuidslice...)

	return data, nil
}

func (s *ZUID) NextIDString() (string, error) {
	bs, err := s.NextID()
	if err != nil {
		return "", err
	}
	str := fmt.Sprintf("%x", bs)
	return str, nil
}
func (s *ZUID) NextIDSimple() string {
	str, err := s.NextIDString()
	if err != nil {
		return ""
	}
	return str
}
