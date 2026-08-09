package core

import "github.com/Moksh-10/redis/config"

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

func evict() {
	switch config.EvictionStrategy {
	case "simple-first" :
		evictFirst()
	case "allkeys-random":
		evictAllkeysRandom()
	}
}