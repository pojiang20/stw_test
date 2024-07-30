package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 独立进行stw
	//tick := time.Tick(1 * time.Second)
	//var r runtime.MemStats
	//go func() {
	//	for {
	//		<-tick
	//		runtime.ReadMemStats(&r)
	//	}
	//}()

	// 构建http服务
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "hi")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
