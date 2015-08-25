package sflow

import (
	"io/ioutil"
	"os"
	"testing"
)

func BenchmarkFlow1Sample(b *testing.B) {
	f, err := os.Open("_test/flow_sample.dump")
	if err != nil {
		b.Fatal(err)
	}

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		b.Fatal(err)
	}

	//d := NewDecoder(f)

	for i := 0; i < b.N; i++ {
		DecodeDatagram(bs)
	}

	//f.Close()
}

func BenchmarkCounterSample(b *testing.B) {
	f, err := os.Open("_test/counter_sample.dump")
	if err != nil {
		b.Fatal(err)
	}

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		b.Fatal(err)
	}

	//d := NewDecoder(f)

	for i := 0; i < b.N; i++ {
		DecodeDatagram(bs)
	}

	//f.Close()
}
