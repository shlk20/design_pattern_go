package strategy

import "testing"

func TestStrategy(t *testing.T) {
	lfu := &Lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")
	// cache.get("a")
	cache.add("c", "3")

	lru := &Lru{}
	cache.setEvictionAlgo(lru)
	cache.add("d", "4")

	fifo := &Fifo{}
	cache.setEvictionAlgo(fifo)
	cache.add("e", "5")
}
