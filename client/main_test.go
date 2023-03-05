package main

import (
	"flag"
	"strings"
	"testing"

	"go-micro.dev/v4/logger"
)

func BenchmarkCallMonolith(b *testing.B) {
	benchtime := flag.Lookup("test.benchtime").Value.String()
	if b.N == 1 && strings.HasSuffix(benchtime, "x") && benchtime != "1x" {
		logger.Infof("Skipping the sub-benchmark discovery run for -benchtime=Nx")
		return
	}

	for n := 0; n < b.N; n++ {
		logger.Infof("Monolith Benchmark run %v out of %v", n+1, b.N)
		callMonolith()
		logger.Infof("Monolith Benchmark run %v out of %v finished", n+1, b.N)
	}
}

func BenchmarkCallMicroservices(b *testing.B) {
	benchtime := flag.Lookup("test.benchtime").Value.String()
	if b.N == 1 && strings.HasSuffix(benchtime, "x") && benchtime != "1x" {
		logger.Infof("Skipping the sub-benchmark discovery run for -benchtime=Nx")
		return
	}

	for n := 0; n < b.N; n++ {
		logger.Infof("Microservices Benchmark run %v out of %v", n+1, b.N)
		callMicroservices()
		logger.Infof("Microservices Benchmark run %v out of %v finished", n+1, b.N)
	}
}
