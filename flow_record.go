package sflow

import "encoding/binary"

// RawPacketFlow is a raw Ethernet header flow record.
type RawPacketFlowRecord struct {
	Protocol    uint32
	FrameLength uint32
	Stripped    uint32
	HeaderSize  uint32
	Header      []byte
}

// RecordType returns the type of flow record.
func (f *RawPacketFlowRecord) RecordType() int {
	return TypeRawPacketFlowRecord
}

func DecodeRawPacketFlowRecord(bs []byte) ([]byte, *RawPacketFlowRecord, error) {
	if len(bs) < 16 {
		return nil, nil, ErrInvalidSliceLength
	}

	f := &RawPacketFlowRecord{}
	f.Protocol = binary.BigEndian.Uint32(bs)
	f.FrameLength = binary.BigEndian.Uint32(bs[4:])
	f.Stripped = binary.BigEndian.Uint32(bs[8:])
	f.HeaderSize = binary.BigEndian.Uint32(bs[12:])

	if f.HeaderSize > 0 {
		padding := (4 - f.HeaderSize) % 4
		if padding < 0 {
			padding += 4
		}
		f.Header = make([]byte, f.HeaderSize+padding)
		copy(f.Header, bs[16:16+f.HeaderSize])
		// We need to consume the padded length,
		// but len(Header) should still be HeaderSize.
		f.Header = f.Header[:f.HeaderSize]
	}
	return bs[16+f.HeaderSize:], f, nil
}

// func (f RawPacketFlowRecord) encode(w io.Writer) error {
// 	var err error

// 	err = binary.Write(w, binary.BigEndian, uint32(f.RecordType()))
// 	if err != nil {
// 		return err
// 	}

// 	// We need to calculate encoded size of the record.
// 	encodedRecordLength := uint32(4 * 4) // 4 32-bit records

// 	// Add the length of the header padded to a multiple of 4 bytes.
// 	padding := (4 - f.HeaderSize) % 4
// 	if padding < 0 {
// 		padding += 4
// 	}

// 	encodedRecordLength += f.HeaderSize + padding

// 	err = binary.Write(w, binary.BigEndian, encodedRecordLength)
// 	if err != nil {
// 		return err
// 	}

// 	err = binary.Write(w, binary.BigEndian, f.Protocol)
// 	if err != nil {
// 		return err
// 	}

// 	err = binary.Write(w, binary.BigEndian, f.FrameLength)
// 	if err != nil {
// 		return err
// 	}

// 	err = binary.Write(w, binary.BigEndian, f.Stripped)
// 	if err != nil {
// 		return err
// 	}

// 	err = binary.Write(w, binary.BigEndian, f.HeaderSize)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = w.Write(append(f.Header, make([]byte, padding)...))

// 	return err
// }

// ExtendedSwitchFlow is an extended switch flow record.
// type ExtendedSwitchFlow struct {
// 	SourceVlan          uint32
// 	SourcePriority      uint32
// 	DestinationVlan     uint32
// 	DestinationPriority uint32
// }

// // RecordType returns the type of flow record.
// func (f ExtendedSwitchFlow) RecordType() int {
// 	return TypeExtendedSwitchFlowRecord
// }

// func DecodeExtendedSwitchFlow(bs []byte) ([]byte, *ExtendedSwitchFlow, error) {
// 	f := ExtendedSwitchFlow{}
// 	err := binary.Read(r, binary.BigEndian, &f)
// 	return f, err
// }

// func (f ExtendedSwitchFlow) encode(w io.Writer) error {
// 	var err error

// 	err = binary.Write(w, binary.BigEndian, uint32(f.RecordType()))
// 	if err != nil {
// 		return err
// 	}

// 	encodedRecordLength := uint32(4 * 4) // 4 32-bit records

// 	err = binary.Write(w, binary.BigEndian, encodedRecordLength)
// 	if err != nil {
// 		return err
// 	}

// 	return binary.Write(w, binary.BigEndian, f)
// }

type EthernetFrameFlowRecord struct {
	Length uint32 /* The length of the MAC packet received on the
	   network, excluding lower layer encapsulations
	   and framing bits but including FCS octets */
	SrcMac [6]byte /* Source MAC address */
	DstMac [6]byte /* Destination MAC address */
	Type   uint32  /* Ethernet packet type */
}

func (f *EthernetFrameFlowRecord) RecordType() int {
	return TypeEthernetFrameFlowRecord
}

func DecodeEthernetFrameFlowRecord(bs []byte) ([]byte, *EthernetFrameFlowRecord, error) {
	if len(bs) < 20 {
		return nil, nil, ErrInvalidSliceLength
	}

	var res = &EthernetFrameFlowRecord{}
	res.Length = binary.BigEndian.Uint32(bs)
	copy(res.SrcMac[:], bs[4:10])
	copy(res.DstMac[:], bs[10:16])
	res.Type = binary.BigEndian.Uint32(bs[16:])
	return bs[20:], res, nil
}

type Ipv4FlowRecord struct {
	Length   uint32 /* The length of the IP packet excluding lower layer encapsulations */
	Protocol uint32 /* IP Protocol type (for example, TCP = 6, UDP = 17) */
	SrcAddr  uint32 /* Source IP Address */
	DstAddr  uint32 /* Destination IP Address */
	SrcPort  uint32 /* TCP/UDP source port number or equivalent */
	DstPort  uint32 /* TCP/UDP destination port number or equivalent */
	TcpFlags uint32 /* TCP flags */
	Tos      uint32
}

func (f *Ipv4FlowRecord) RecordType() int {
	return TypeIpv4FlowRecord
}

func DecodeIpv4FlowRecord(bs []byte) ([]byte, *Ipv4FlowRecord, error) {
	if len(bs) < 32 {
		return nil, nil, ErrInvalidSliceLength
	}

	var res = &Ipv4FlowRecord{}
	res.Length = binary.BigEndian.Uint32(bs)
	res.Protocol = binary.BigEndian.Uint32(bs[4:])
	res.SrcAddr = binary.BigEndian.Uint32(bs[8:])
	res.DstAddr = binary.BigEndian.Uint32(bs[12:])
	res.SrcPort = binary.BigEndian.Uint32(bs[16:])
	res.DstPort = binary.BigEndian.Uint32(bs[20:])
	res.TcpFlags = binary.BigEndian.Uint32(bs[24:])
	res.Tos = binary.BigEndian.Uint32(bs[28:])
	return bs[32:], res, nil
}

type Ipv6FlowRecord struct {
	Length   uint32   /* The length of the IP packet excluding lower layer encapsulations */
	Protocol uint32   /* IP Protocol type (for example, TCP = 6, UDP = 17) */
	SrcAddr  [16]byte /* Source IP Address */
	DstAddr  [16]byte /* Destination IP Address */
	SrcPort  uint32   /* TCP/UDP source port number or equivalent */
	DstPort  uint32   /* TCP/UDP destination port number or equivalent */
	TcpFlags uint32   /* TCP flags */
	Priority uint32
}

func (f *Ipv6FlowRecord) RecordType() int {
	return TypeIpv6FlowRecord
}

func DecodeIpv6FlowRecord(bs []byte) ([]byte, *Ipv6FlowRecord, error) {
	if len(bs) < 56 {
		return nil, nil, ErrInvalidSliceLength
	}

	var res = &Ipv6FlowRecord{}
	res.Length = binary.BigEndian.Uint32(bs)
	res.Protocol = binary.BigEndian.Uint32(bs[4:])
	copy(res.SrcAddr[:], bs[8:24])
	copy(res.DstAddr[:], bs[24:40])
	res.SrcPort = binary.BigEndian.Uint32(bs[40:])
	res.DstPort = binary.BigEndian.Uint32(bs[44:])
	res.TcpFlags = binary.BigEndian.Uint32(bs[48:])
	res.Priority = binary.BigEndian.Uint32(bs[52:])
	return bs[56:], res, nil
}

func decodeFlowRecord(bs []byte) ([]byte, Record, error) {
	if len(bs) < 8 {
		return nil, nil, ErrInvalidSliceLength
	}

	format := binary.BigEndian.Uint32(bs)
	length := binary.BigEndian.Uint32(bs[4:])
	if len(bs) < 8+int(length) {
		return nil, nil, ErrInvalidSliceLength
	}

	switch format {
	case TypeRawPacketFlowRecord:
		_, s, e := DecodeRawPacketFlowRecord(bs[8 : 8+length])
		return bs[8+length:], s, e
	case TypeEthernetFrameFlowRecord:
		_, s, e := DecodeEthernetFrameFlowRecord(bs[8 : 8+length])
		return bs[8+length:], s, e
	case TypeIpv4FlowRecord:
		_, s, e := DecodeIpv4FlowRecord(bs[8 : 8+length])
		return bs[8+length:], s, e
	case TypeIpv6FlowRecord:
		_, s, e := DecodeIpv6FlowRecord(bs[8 : 8+length])
		return bs[8+length:], s, e
	//case TypeExtendedSwitchFlowRecord:
	//	_, s, e := decodedExtendedSwitchFlow(bs[8 : 8+length])
	//	return bs[8+length:], s, e
	default:
		return bs[8+length:], nil, nil
	}
}
