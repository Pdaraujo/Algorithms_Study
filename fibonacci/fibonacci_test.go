package fibonacci

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
	"time"
)

type cenas func (n int64) int64

func TestFibonacci(t *testing.T) {
	timeForFunc(fibonacci)
}

func TestFasterFibonacci(t *testing.T) {
	timeForFunc(fasterFibonacci)
}

func timeForFunc(c cenas) {
	n := 60
	start := time.Now()
	c(int64(n))
	elapsed := time.Since(start)
	fmt.Printf("Function named %s, took %d miliseconds for n = %d \n ", GetFunctionName(c), elapsed, n)
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}