package gotesting_test

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"testing"
)

func BenchmarkSHA1(b *testing.B) {
	data := []byte("Hello world how are you?")
	b.StartTimer() //timer is started after the data is allocated
	//b.N is the iteration that we define earlier
	for i := 0; i < b.N; i++ {
		sha1.Sum(data)
	}
}

func BenchmarkSHA256(b *testing.B) {
	data := []byte("Hello world how are you?")
	b.StartTimer() //timer is started after the data is allocated
	//b.N is the iteration that we define earlier
	for i := 0; i < b.N; i++ {
		sha256.Sum256(data)
	}
}

func BenchmarkSHA512(b *testing.B) {
	data := []byte("Hello world how are you?")
	b.StartTimer() //timer is started after the data is allocated
	//b.N is the iteration that we define earlier
	for i := 0; i < b.N; i++ {
		sha512.Sum512(data)
	}
}
