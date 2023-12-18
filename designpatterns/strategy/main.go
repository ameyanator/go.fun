package main

func main() {
	lru := &Lru{}
	cache := initCache(lru)

	cache.add("ameya", "sinha")
	cache.add("arya", "akhilesh")
	cache.add("akhilesh", "kumar")
	fifo := &Fifo{}

	cache.setEvictionAlgo(fifo)
	cache.add("manisha", "sinha")
}
