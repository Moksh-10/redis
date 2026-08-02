package core

import "github.com/Moksh-10/redis/config"

func evictFirst() {
	for k := range store {
		delete(store, k)
		return
	}
}

func evict() {
	switch config.EvictionStrategy {
	case "simple-first" :
		evictFirst()
	}
}