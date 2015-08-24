package sflow

import (
	"encoding/binary"
	"errors"
)

var ErrInvalidSliceLength = errors.New("sflow: invalid slice length")
var ErrUnsupportedIpVersion = errors.New("sflow: unsupported datagram version")
var ErrUnsupportedDatagramVersion = errors.New("sflow: unsupported datagram version")

func DecodeDatagram(bs []byte) ([]byte, *Datagram, error) {
	if len(bs) < 28 {
		return nil, nil, ErrInvalidSliceLength
	}

	// Decode headers first
	dgram := &Datagram{}

	dgram.Version = binary.BigEndian.Uint32(bs)

	//if dgram.Version != 5 {
	//	return nil, ErrUnsupportedDatagramVersion
	//}
	dgram.IpVersion = binary.BigEndian.Uint32(bs[4:])

	//err = binary.Read(d.reader, binary.BigEndian, &)
	//if err != nil {
	//	return nil, err
	//}

	ipLen := 4
	switch dgram.IpVersion {
	case 1: // IP_V4
		ipLen = 4
	case 2: // IP_V6
		ipLen = 16
		if len(bs) < 40 {
			return nil, nil, ErrInvalidSliceLength
		}
	case 0: // UNKNOWN
		ipLen = 0
	default:
		return nil, nil, ErrUnsupportedIpVersion
	}
	if ipLen > 0 {
		dgram.IpAddress = bs[8 : 8+ipLen]
	}
	dgram.SubAgentId = binary.BigEndian.Uint32(bs[8+ipLen:])
	dgram.SequenceNumber = binary.BigEndian.Uint32(bs[12+ipLen:])
	dgram.Uptime = binary.BigEndian.Uint32(bs[16+ipLen:])
	dgram.NumSamples = binary.BigEndian.Uint32(bs[20+ipLen:])
	return bs[24+ipLen:], dgram, nil
}
