package main

import (
	pborman "github.com/pborman/uuid"
	satori "github.com/satori/go.uuid"
	"testing"
)

var (
	pbormanUUID = pborman.NewRandom()
	satoriUUID  = satori.NewV4()
	UUIDstring  = satoriUUID.String()
)

func BenchmarkSatoriToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = satoriUUID.String()
	}
}

func BenchmarkPbormanToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = pbormanUUID.String()
	}
}

func BenchmarkSatoriFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = satori.FromString(UUIDstring)
	}
}

func BenchmarkPbormanFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pborman.Parse(UUIDstring)
	}
}
