package q3

import "testing"

//性能测试，方法名称必须以Benchmark为前缀，参数为*testing.B
//hou to run:go test -v -bench=. -run=^$
func BenchmarkGetPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetPrimes(1000)
		//b.Log(GetPrimes(1000))
	}
}