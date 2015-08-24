package sflow

import "encoding/binary"

type ExpandedFlowSample struct {
	SequenceNum      uint32
	SourceIdType     uint32
	SourceIdIndexVal uint32 // NOTE: this is 3 bytes in the datagram
	SamplingRate     uint32
	SamplePool       uint32
	Drops            uint32
	InputFormat      uint32
	InputValue       uint32
	OutputFormat     uint32
	OutputValue      uint32
	NumRecords       uint32
	//Records          []Record
}

// SampleType returns the type of sFlow sample.
func (s *ExpandedFlowSample) SampleType() int {
	return TypeExpandedFlowSample
}

// func (s *ExpandedFlowSample) GetRecords() []Record {
// 	return s.Records
// }

func DecodeExpandedFlowSample(bs []byte) ([]byte, *ExpandedFlowSample, error) {
	if len(bs) < 44 {
		return nil, nil, ErrInvalidSliceLength
	}

	s := &ExpandedFlowSample{}
	s.SequenceNum = binary.BigEndian.Uint32(bs)
	s.SourceIdType = binary.BigEndian.Uint32(bs[4:])
	s.SourceIdIndexVal = binary.BigEndian.Uint32(bs[8:])
	s.SamplingRate = binary.BigEndian.Uint32(bs[12:])
	s.SamplePool = binary.BigEndian.Uint32(bs[16:])
	s.Drops = binary.BigEndian.Uint32(bs[20:])
	s.InputFormat = binary.BigEndian.Uint32(bs[24:])
	s.InputValue = binary.BigEndian.Uint32(bs[28:])
	s.OutputFormat = binary.BigEndian.Uint32(bs[32:])
	s.OutputValue = binary.BigEndian.Uint32(bs[36:])
	s.NumRecords = binary.BigEndian.Uint32(bs[40:])

	// for i := uint32(0); i < s.numRecords; i++ {
	// 	rec, err := decodeFlowRecord(r)
	// 	if nil != err {
	// 		fmt.Println(i)
	// 		return nil, err
	// 	}
	// 	if nil != rec {
	// 		s.Records = append(s.Records, rec)
	// 	}
	// }

	return bs[44:], s, nil
}
