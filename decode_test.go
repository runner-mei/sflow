package sflow

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestDecodeGenericEthernetCounterSample(t *testing.T) {
	f, err := os.Open("_test/counter_sample.dump")
	if err != nil {
		t.Fatal(err)
	}
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	//d := NewDecoder(f)

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

	next, record_next, dgram_sample, err := DecodeSample(next)
	if err != nil {
		t.Fatal(err)
	}

	sample, ok := dgram_sample.(*CounterSample)
	if !ok {
		t.Fatalf("expected a CounterSample, got %T", dgram_sample)
	}

	if sample.NumRecords != 2 {
		t.Fatalf("expected 2 records, got %d", sample.NumRecords)
	}

	record_next, record, err := DecodeCounterRecord(record_next)
	if err != nil {
		t.Fatal(err)
	}

	ethCounters, ok := record.(*EthernetCounters)
	if !ok {
		t.Fatalf("expected a EthernetCounters record, got %T", record)
	}

	expectedEthCountersRec := EthernetCounters{}
	if *ethCounters != expectedEthCountersRec {
		t.Errorf("expected\n%#v, got\n%#v", expectedEthCountersRec, ethCounters)
	}

	record_next, record, err = DecodeCounterRecord(record_next)
	if err != nil {
		t.Fatal(err)
	}

	genericInterfaceCounters, ok := record.(*GenericInterfaceCounters)
	if !ok {
		t.Fatalf("expected a GenericInterfaceCounters record, got %T", record)
	}

	expectedGenericInterfaceCounters := GenericInterfaceCounters{
		Index:               9,
		Type:                6,
		Speed:               100000000,
		Direction:           1,
		Status:              3,
		InOctets:            79282473,
		InUnicastPackets:    329128,
		InMulticastPackets:  0,
		InBroadcastPackets:  1493,
		InDiscards:          0,
		InErrors:            0,
		InUnknownProtocols:  0,
		OutOctets:           764247430,
		OutUnicastPackets:   9470970,
		OutMulticastPackets: 780342,
		OutBroadcastPackets: 877721,
		OutDiscards:         0,
		OutErrors:           0,
		PromiscuousMode:     1,
	}

	if *genericInterfaceCounters != expectedGenericInterfaceCounters {
		t.Errorf("expected\n%#v, got\n%#v", expectedGenericInterfaceCounters, genericInterfaceCounters)
	}
}

func TestDecodeHostCounters(t *testing.T) {
	f, err := os.Open("_test/host_sample.dump")
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

	next, record_next, dgram_sample, err := DecodeSample(next)
	if err != nil {
		t.Fatal(err)
	}

	sample, ok := dgram_sample.(*CounterSample)
	if !ok {
		t.Fatalf("expected a CounterSample, got %T", dgram_sample)
	}

	if sample.NumRecords != 6 {
		t.Fatalf("expected 4 records, got %d", sample.NumRecords, sample)
	}

	//var record Record
	for i := uint32(0); i < sample.NumRecords; i++ {
		record_next, _, err = DecodeCounterRecord(record_next)
		if err != nil {
			t.Fatal(i, err)
		}
	}

	// TODO: check values
}

func TestDecodeFlow1(t *testing.T) {
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

	// if int(dgram.NumSamples) != len(dgram.Samples) {
	// 	t.Fatalf("expected NumSamples to be %d, but len(Samples) is %d", dgram.NumSamples, len(dgram.Samples))
	// }

	if dgram.NumSamples != 1 {
		t.Fatalf("expected 1 sample, got %d", dgram.NumSamples)
	}

	next, record_next, dgram_sample, err := DecodeSample(next)
	if err != nil {
		t.Fatal(err)
	}

	sample, ok := dgram_sample.(*FlowSample)
	if !ok {
		t.Fatalf("expected a FlowSample, got %T", dgram_sample)
	}

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
