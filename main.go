package main

import "log"

func main() {
	ct := Counter{Value: 0}
	log.Println(ct.Value)
	ct.Increment()
	log.Println(ct.Value)
	incr(&ct)
	log.Println(ct.Value)
}
