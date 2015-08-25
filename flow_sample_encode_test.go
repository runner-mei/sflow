package sflow

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestDecodeEncodeAndDecodeFlowSample(t *testing.T) {
	f, err := os.Open("_test/flow_sample.dump")
	if err != nil {
		t.Fatal(err)
	}

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	next, dgram, err := DecodeDatagram(bs)
	if err != nil {
		t.Fatal(err)
	}

	if dgram.Version != 5 {
		t.Errorf("Expected datagram version %v, got %v", 5, dgram.Version)
	}

	//if int(dgram.NumSamples) != len(dgram.Samples) {
	//	t.Fatalf("expected NumSamples to be %d, but len(Samples) is %d", dgram.NumSamples, len(dgram.Samples))
	//}

	if dgram.NumSamples != 1 {
		t.Fatalf("expected 1 sample, got %d", dgram.NumSamples)
	}

	next, record_next, dgram_samples, err := DecodeSample(next)
	if err != nil {
		t.Fatal(err)
	}

	sample, ok := dgram_samples.(*FlowSample)
	if !ok {
		t.Fatalf("expected a FlowSample, got %T", dgram.Samples[0])
	}

	// buf := &bytes.Buffer{}

	// err = sample.encode(buf)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// // We need to skip the first 8 bytes. That's the header.
	// var skip [8]byte
	// buf.Read(skip[:])

	// // bytes.Buffer is not an io.ReadSeeker. bytes.Reader is.
	// decodedSample, err := DecodeFlowSample(buf.Bytes())
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// sample, ok = decodedSample.(*FlowSample)
	// if !ok {
	// 	t.Fatalf("expected a FlowSample, got %T", decodedSample)
	// }

	if sample.NumRecords != 2 {
		t.Fatalf("expected 2 records, got %d", sample.NumRecords)
	}

	record_next, record, err := DecodeFlowRecord(record_next)
	if err != nil {
		t.Fatal(err)
	}

	rec, ok := record.(*RawPacketFlowRecord)
	if !ok {
		t.Fatalf("expected a RawPacketFlowRecords, got %T", record)
	}

	if rec.Protocol != 1 {
		t.Errorf("expected Protocol to be 1, got %d", rec.Protocol)
	}

	if rec.FrameLength != 318 {
		t.Errorf("expected FrameLength to be 318, got %d", rec.FrameLength)
	}

	if rec.Stripped != 4 {
		t.Errorf("expected FrameLength to be 4, got %d", rec.Stripped)
	}

	if rec.HeaderSize != 128 {
		t.Errorf("expected FrameLength to be 128, got %d", rec.HeaderSize)
	}
}
