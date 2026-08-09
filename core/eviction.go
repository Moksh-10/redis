package core

import (
	"time"
	"github.com/Moksh-10/redis/config"
)

func evictFirst() {
	for k := range store {
		Del(k)
		return
	}
}

func evictAllkeysRandom() {
	evictCount := int64(config.EvictionRatio * float64(config.KeysLimit))
	for k := range store {
		Del(k)
		evictCount--
		if evictCount <= 0 {
			break
		}
	}
}

func getCurrentClock() uint32 {
	return uint32(time.Now().Unix()) & 0x00FFFFFF
}

func getIdleTime(lastAccessedAT uint32) uint32 {
	c := getCurrentClock()
	if c >= lastAccessedAT {
		return c - lastAccessedAT
	}
	return (0x00FFFFFF - lastAccessedAT) + c
}

func populateEvictionPool() {
	sampleSize := 5
	for k := range store {
		ePool.Push(k, store[k].LastAccessedAT)
		sampleSize--
		if sampleSize == 0 {
			break
		}
	}
}

func evistAllkeysLRU() {
	populateEvictionPool()
	evictCount := int16(config.EvictionRatio * float64(config.KeysLimit))
	for i := 0; i < int(evictCount) && len(ePool.pool) > 0; i++ {
		item := ePool.Pop()
		if item == nil {
			return
		}
		Del(item.key)
	}
}

func evict() {
	switch config.EvictionStrategy {
	case "simple-first" :
		evictFirst()
	case "allkeys-random":
		evictAllkeysRandom()
	case "allkeys-lru":
		evistAllkeysLRU()
	}
}