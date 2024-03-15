package main

import (
	"fmt"
	"runtime"
	"test/script"
)

func main() {
	runtime.GOMAXPROCS(2)
	fmt.Println(script.PlusOne())

	// oapi.NewServer()
	// t := map[string]interface{}{
	// 	"nbudi": "nggnng",
	// 	"tono":  48,
	// }
	// byt, _ := json.MarshalIndent(t, "", "")
	// hrank.CheckMagazine([]string{"give", "me", "one", "grand", "today", "night"}, []string{"give", "one", "grand", "grand", "today"})
	// script.Convertarraybytstring()
	// fmt.Println(string(byt))
	// hrank.ExecReversedLinkedlist()
	// redis.Client=
	// other.LockFeature()

	// script.ChannelMain()

}
