package main

import (
	"fmt"

	store_cache "forThundercore/policy"
)

func main() {
	c := store_cache.NewCache(256)
	c_fifo := store_cache.NewFifoCache(256)

	//assign key value
	c.Set("a", "123")
	c_fifo.Set("a", "qwerr")

	//Get normal cache item:a
	result1, ok1 := c.Get("a")
	fmt.Println(result1, ok1) // 123 true

	//Get fifo cache item:a
	result2, ok2 := c_fifo.Get("a")
	fmt.Println(result2, ok2) // qwerr false

	//pop fifo cache top item
	result3, ok3 := c_fifo.Pop()
	fmt.Println(result3, ok3) // qwerr true

	//pop fifo cache top item
	result4, ok4 := c_fifo.Pop()
	fmt.Println(result4, ok4) // false

	//Get normal cache item:a
	result5, ok5 := c.Get("a")
	fmt.Println(result5, ok5) // 123 true
}
