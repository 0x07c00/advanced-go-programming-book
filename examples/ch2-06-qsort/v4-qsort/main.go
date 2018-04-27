// Copyright © 2018 ChaiShushan <chaishushan{AT}gmail.com>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

/*
#include <stdlib.h>

typedef int (*qsort_cmp_func_t)(const void* a, const void* b);
extern int go_qsort_compare(void* a, void* b);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	values := []int32{42, 9, 101, 95, 27, 25}

	C.qsort(unsafe.Pointer(&values[0]),
		C.size_t(len(values)), C.size_t(unsafe.Sizeof(values[0])),
		(C.qsort_cmp_func_t)(unsafe.Pointer(C.go_qsort_compare)),
	)

	fmt.Println(values)
}

//export go_qsort_compare
func go_qsort_compare(a, b unsafe.Pointer) C.int {
	pa := (*C.int)(a)
	pb := (*C.int)(b)

	return C.int(*pa - *pb)
}
