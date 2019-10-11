package main

import (
	"testing"
	"os"
)

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, sep := "", ""
    		for _, arg := range os.Args[1:] {
        		s += sep + arg
        		sep = " "
   	 	}
	}
}
