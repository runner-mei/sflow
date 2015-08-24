package sflow

import "encoding/binary"

type ExpandedCounterSample struct {
	SequenceNum      uint32
	SourceIdType     uint32
	SourceIdIndexVal uint32 // NOTE: this is 3 bytes in the datagram
	NumRecords       uint32
	//Records          []Record
}

// SampleType returns the type of sFlow sample.
func (s *ExpandedCounterSample) SampleType() int {
	return TypeExpandedCounterSample
}

// func (s *ExpandedCounterSample) GetRecords() []Record {
// 	return s.Records
// }

func DecodeExpandedCounterSample(bs []byte) ([]byte, *ExpandedCounterSample, error) {
	if len(bs) < 16 {
		return nil, nil, ErrInvalidSliceLength
	}

	s := &ExpandedCounterSample{}
	s.SequenceNum = binary.BigEndian.Uint32(bs)
	s.SourceIdType = binary.BigEndian.Uint32(bs[4:])
	s.SourceIdIndexVal = binary.BigEndian.Uint32(bs[8:])
	s.NumRecords = binary.BigEndian.Uint32(bs[12:])

	// for i := uint32(0); i < s.numRecords; i++ {
	// 	rec, err := decodeCounterRecord(r)
	// 	if nil != err {
	// 		return nil, err
	// 	}
	// 	if nil != rec {
	// 		s.Records = append(s.Records, rec)
	// 	}
	// }

	return bs[16:], s, nil
}

// func (s *ExpandedCounterSample) encode(w io.Writer) error {
// 	var err error

// 	// We first need to encode the records.
// 	buf := &bytes.Buffer{}

// 	for _, rec := range s.Records {
// 		err = rec.encode(buf)
// 		if err != nil {
// 			return ErrEncodingRecord
// 		}
// 	}

// 	// Fields
// 	encodedSampleSize := uint32(4 + 1 + 3 + 4)

// 	// Encoded records
// 	encodedSampleSize += uint32(buf.Len())

// 	err = binary.Write(w, binary.BigEndian, uint32(s.SampleType()))
// 	if err != nil {
// 		return err
// 	}

// 	err = binary.Write(w, binary.BigEndian, encodedSampleSize)
// 	if err != nil {
// 		return err
// 	}

// 	err = binary.Write(w, binary.BigEndian, s.SequenceNum)
// 	if err != nil {
// 		return err
// 	}

// 	err = binary.Write(w, binary.BigEndian, uint32(s.SourceIdType)|s.SourceIdIndexVal<<24)
// 	if err != nil {
// 		return err
// 	}

// 	err = binary.Write(w, binary.BigEndian, uint32(len(s.Records)))
// 	if err != nil {
// 		return err
// 	}

// 	_, err = io.Copy(w, buf)

// 	return err
// }
