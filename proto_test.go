package test

import (
	"testing"
	"time"

	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
)

func BenchmarkGolangProto(b *testing.B) {
	a := A{}
	for i := 0; i < 10000; i++ {
		a.Bs = append(a.Bs, &B{
			X: 1000000,
			Y: 1000000,
			Z: 1000000,
		})
	}

	data, err := proto.Marshal(&a)
	if err != nil {
		panic(err)
	}

	b.SetParallelism(1000)
	b.ResetTimer()
	maxLatency := time.Duration(0)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			before := time.Now()
			pa := new(A)
			if err := proto.Unmarshal(data, pa); err != nil {
				panic(err)
			}
			latency := time.Since(before)
			if latency > maxLatency {
				maxLatency = latency
			}
		}
	})

	b.Log(maxLatency)
}

func BenchmarkGoGoProto(b *testing.B) {
	a := AA{}
	for i := 0; i < 10000; i++ {
		a.Bs = append(a.Bs, &BB{
			X: 1000000,
			Y: 1000000,
			Z: 1000000,
		})
	}

	data, err := gogoproto.Marshal(&a)
	if err != nil {
		panic(err)
	}

	b.SetParallelism(1000)
	b.ResetTimer()
	maxLatency := time.Duration(0)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			before := time.Now()
			pa := new(AA)
			if err := gogoproto.Unmarshal(data, pa); err != nil {
				panic(err)
			}
			latency := time.Since(before)
			if latency > maxLatency {
				maxLatency = latency
			}
		}
	})

	b.Log(maxLatency)
}
