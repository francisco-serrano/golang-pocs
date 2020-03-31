package memory

import (
	"fmt"
	"runtime"
)

func Run() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)

	fmt.Println("KBs of Allocated Heap Objects", ms.Alloc/1000)
	fmt.Println("Cumulative KBs Allocated for Heap Objects", ms.TotalAlloc/1000)
	fmt.Println("Total KBs of Memory Obtained from the OS", ms.Sys/1000)
}
