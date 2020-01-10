package vt

// #cgo CFLAGS: -O3
// #define LZ4_DISABLE_DEPRECATE_WARNINGS
// #include "src/lz4.h"
// #include "src/lz4.c"
import "C"

import (
	"fmt"
	"github.com/w4n97q/vtlib/src"
)

//Show print vtlib name
func Show() {
	fmt.Println("i'm " + src.Name)
	fmt.Println("i'm vtlib")
}
