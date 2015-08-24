package sflow

import (
	"encoding/binary"
	"errors"
)

const (
	// opaque	enterprise	format	struct	reference
	// sample_data	0	1	flow_sample	sFlow Version 5
	// sample_data	0	2	counter_sample	sFlow Version 5
	// sample_data	0	3	flow_sample_expanded	sFlow Version 5
	// sample_data	0	4	counter_sample_expanded	sFlow Version 5
	TypeFlowSample            = 1
	TypeCounterSample         = 2
	TypeExpandedFlowSample    = 3
	TypeExpandedCounterSample = 4
)

var (
	ErrUnknownSampleType = errors.New("sflow: Unknown sample type")
)

type Sample interface {
	SampleType() int
	//	GetRecords() []Record
	//encode(w io.Writer) error
}

func DecodeSample(bs []byte) ([]byte, []byte, Sample, error) {
	if len(bs) < 8 {
		return nil, nil, nil, ErrInvalidSliceLength
	}
	format := binary.BigEndian.Uint32(bs)
	length := binary.BigEndian.Uint32(bs[4:])

	if len(bs) < 8+int(length) {
		return nil, nil, nil, ErrInvalidSliceLength
	}

	switch format {
	case TypeFlowSample:
		next, s, e := DecodeFlowSample(bs[8 : 8+length])
		return bs[8+length:], next, s, e
	case TypeCounterSample:
		next, s, e := DecodeCounterSample(bs[8 : 8+length])
		return bs[8+length:], next, s, e
	case TypeExpandedFlowSample:
		next, s, e := DecodeExpandedFlowSample(bs[8 : 8+length])
		return bs[8+length:], next, s, e
	case TypeExpandedCounterSample:
		next, s, e := DecodeExpandedCounterSample(bs[8 : 8+length])
		return bs[8+length:], next, s, e
	default:
		return nil, nil, nil, ErrUnknownSampleType
	}
}
