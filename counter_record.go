package sflow

import (
	"encoding/binary"
	"io"
	"unsafe"
)

// GenericInterfaceCounters is a generic switch counters record.
type GenericInterfaceCounters struct {
	Index               uint32
	Type                uint32
	Speed               uint64
	Direction           uint32
	Status              uint32
	InOctets            uint64
	InUnicastPackets    uint32
	InMulticastPackets  uint32
	InBroadcastPackets  uint32
	InDiscards          uint32
	InErrors            uint32
	InUnknownProtocols  uint32
	OutOctets           uint64
	OutUnicastPackets   uint32
	OutMulticastPackets uint32
	OutBroadcastPackets uint32
	OutDiscards         uint32
	OutErrors           uint32
	PromiscuousMode     uint32
}

// EthernetCounters is an Ethernet interface counters record.
type EthernetCounters struct {
	AlignmentErrors           uint32
	FCSErrors                 uint32
	SingleCollisionFrames     uint32
	MultipleCollisionFrames   uint32
	SQETestErrors             uint32
	DeferredTransmissions     uint32
	LateCollisions            uint32
	ExcessiveCollisions       uint32
	InternalMACTransmitErrors uint32
	CarrierSenseErrors        uint32
	FrameTooLongs             uint32
	InternalMACReceiveErrors  uint32
	SymbolErrors              uint32
}

// TokenRingCounters is a token ring interface counters record.
type TokenRingCounters struct {
	LineErrors         uint32
	BurstErrors        uint32
	ACErrors           uint32
	AbortTransErrors   uint32
	InternalErrors     uint32
	LostFrameErrors    uint32
	ReceiveCongestions uint32
	FrameCopiedErrors  uint32
	TokenErrors        uint32
	SoftErrors         uint32
	HardErrors         uint32
	SignalLoss         uint32
	TransmitBeacons    uint32
	Recoverys          uint32
	LobeWires          uint32
	Removes            uint32
	Singles            uint32
	FreqErrors         uint32
}

// VgCounters is a BaseVG interface counters record.
type VgCounters struct {
	InHighPriorityFrames    uint32
	InHighPriorityOctets    uint64
	InNormPriorityFrames    uint32
	InNormPriorityOctets    uint64
	InIPMErrors             uint32
	InOversizeFrameErrors   uint32
	InDataErrors            uint32
	InNullAddressedFrames   uint32
	OutHighPriorityFrames   uint32
	OutHighPriorityOctets   uint64
	TransitionIntoTrainings uint32
	HCInHighPriorityOctets  uint64
	HCInNormPriorityOctets  uint64
	HCOutHighPriorityOctets uint64
}

// VlanCounters is a VLAN counters record.
type VlanCounters struct {
	ID               uint32
	Octets           uint64
	UnicastPackets   uint32
	MulticastPackets uint32
	BroadcastPackets uint32
	Discards         uint32
}

// ProcessorCounters is a switch processor counters record.
type ProcessorCounters struct {
	CPU5s       uint32
	CPU1m       uint32
	CPU5m       uint32
	TotalMemory uint64
	FreeMemory  uint64
}

// HostCPUCounters is a host CPU counters record.
type HostCPUCounters struct {
	Load1m           float32
	Load5m           float32
	Load15m          float32
	ProcessesRunning uint32
	ProcessesTotal   uint32
	NumCPU           uint32
	SpeedCPU         uint32
	Uptime           uint32

	CPUUser         uint32
	CPUNice         uint32
	CPUSys          uint32
	CPUIdle         uint32
	CPUWio          uint32
	CPUIntr         uint32
	CPUSoftIntr     uint32
	Interrupts      uint32
	ContextSwitches uint32

	CPUSteal     uint32
	CPUGuest     uint32
	CPUGuestNice uint32
}

// HostMemoryCounters is a host memory counters record.
type HostMemoryCounters struct {
	Total     uint64
	Free      uint64
	Shared    uint64
	Buffers   uint64
	Cached    uint64
	SwapTotal uint64
	SwapFree  uint64

	PageIn  uint32
	PageOut uint32
	SwapIn  uint32
	SwapOut uint32
}

// HostDiskCounters is a host disk counters record.
type HostDiskCounters struct {
	Total          uint64
	Free           uint64
	MaxUsedPercent float32
	Reads          uint32
	BytesRead      uint64
	ReadTime       uint32
	Writes         uint32
	BytesWritten   uint64
	WriteTime      uint32
}

// HostNetCounters is a host network counters record.
type HostNetCounters struct {
	BytesIn   uint64
	PacketsIn uint32
	ErrorsIn  uint32
	DropsIn   uint32

	BytesOut   uint64
	PacketsOut uint32
	ErrorsOut  uint32
	DropsOut   uint32
}

var (
	genericInterfaceCountersSize = int(unsafe.Sizeof(GenericInterfaceCounters{}))
	ethernetCountersSize         = int(unsafe.Sizeof(EthernetCounters{}))
	tokenRingCountersSize        = int(unsafe.Sizeof(TokenRingCounters{}))
	vgCountersSize               = int(unsafe.Sizeof(VgCounters{}))
	vlanCountersSize             = int(unsafe.Sizeof(VlanCounters{}))
	processorCountersSize        = int(unsafe.Sizeof(ProcessorCounters{}))
	hostCPUCountersSize          = 68
	hostMemoryCountersSize       = int(unsafe.Sizeof(HostMemoryCounters{}))
	hostDiskCountersSize         = 52
	hostNetCountersSize          = 40
)

// RecordType returns the type of counter record.
func (c *GenericInterfaceCounters) RecordType() int {
	return TypeGenericInterfaceCountersRecord
}

func decodeGenericInterfaceCountersRecord(bs []byte) ([]byte, *GenericInterfaceCounters, error) {
	if len(bs) < genericInterfaceCountersSize {
		return nil, nil, ErrDecodingRecord
	}
	c := &GenericInterfaceCounters{}
	fields := []interface{}{
		&c.Index,
		&c.Type,
		&c.Speed,
		&c.Direction,
		&c.Status,
		&c.InOctets,
		&c.InUnicastPackets,
		&c.InMulticastPackets,
		&c.InBroadcastPackets,
		&c.InDiscards,
		&c.InErrors,
		&c.InUnknownProtocols,
		&c.OutOctets,
		&c.OutUnicastPackets,
		&c.OutMulticastPackets,
		&c.OutBroadcastPackets,
		&c.OutDiscards,
		&c.OutErrors,
		&c.PromiscuousMode,
	}

	return bs[genericInterfaceCountersSize:], c, readFields(bs, fields)
}

func (c *GenericInterfaceCounters) encode(w io.Writer) error {
	var err error

	err = binary.Write(w, binary.BigEndian, uint32(c.RecordType()))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, uint32(genericInterfaceCountersSize))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, c)
	return err
}

// RecordType returns the type of counter record.
func (c *EthernetCounters) RecordType() int {
	return TypeEthernetCountersRecord
}

func decodeEthernetCountersRecord(bs []byte) ([]byte, *EthernetCounters, error) {
	if len(bs) < ethernetCountersSize {
		return nil, nil, ErrDecodingRecord
	}

	c := &EthernetCounters{}
	fields := []interface{}{
		&c.AlignmentErrors,
		&c.FCSErrors,
		&c.SingleCollisionFrames,
		&c.MultipleCollisionFrames,
		&c.SQETestErrors,
		&c.DeferredTransmissions,
		&c.LateCollisions,
		&c.ExcessiveCollisions,
		&c.InternalMACTransmitErrors,
		&c.CarrierSenseErrors,
		&c.FrameTooLongs,
		&c.InternalMACReceiveErrors,
		&c.SymbolErrors,
	}

	return bs[ethernetCountersSize:], c, readFields(bs, fields)
}

func (c *EthernetCounters) encode(w io.Writer) error {
	var err error

	err = binary.Write(w, binary.BigEndian, uint32(c.RecordType()))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, uint32(ethernetCountersSize))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, c)
	return err
}

// RecordType returns the type of counter record.
func (c *TokenRingCounters) RecordType() int {
	return TypeTokenRingCountersRecord
}

func decodeTokenRingCountersRecord(bs []byte) ([]byte, *TokenRingCounters, error) {
	if len(bs) < tokenRingCountersSize {
		return nil, nil, ErrDecodingRecord
	}

	c := &TokenRingCounters{}
	fields := []interface{}{
		&c.LineErrors,
		&c.BurstErrors,
		&c.ACErrors,
		&c.AbortTransErrors,
		&c.InternalErrors,
		&c.LostFrameErrors,
		&c.ReceiveCongestions,
		&c.FrameCopiedErrors,
		&c.TokenErrors,
		&c.SoftErrors,
		&c.HardErrors,
		&c.SignalLoss,
		&c.TransmitBeacons,
		&c.Recoverys,
		&c.LobeWires,
		&c.Removes,
		&c.Singles,
		&c.FreqErrors,
	}

	return bs[tokenRingCountersSize:], c, readFields(bs, fields)
}

func (c *TokenRingCounters) encode(w io.Writer) error {
	var err error

	err = binary.Write(w, binary.BigEndian, uint32(c.RecordType()))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, uint32(tokenRingCountersSize))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, c)
	return err
}

// RecordType returns the type of counter record.
func (c *VgCounters) RecordType() int {
	return TypeVgCountersRecord
}

func decodeVgCountersRecord(bs []byte) ([]byte, *VgCounters, error) {
	if len(bs) < vgCountersSize {
		return nil, nil, ErrDecodingRecord
	}

	c := &VgCounters{}
	fields := []interface{}{
		&c.InHighPriorityFrames,
		&c.InHighPriorityOctets,
		&c.InNormPriorityFrames,
		&c.InNormPriorityOctets,
		&c.InIPMErrors,
		&c.InOversizeFrameErrors,
		&c.InDataErrors,
		&c.InNullAddressedFrames,
		&c.OutHighPriorityFrames,
		&c.OutHighPriorityOctets,
		&c.TransitionIntoTrainings,
		&c.HCInHighPriorityOctets,
		&c.HCInNormPriorityOctets,
		&c.HCOutHighPriorityOctets,
	}

	return bs[vgCountersSize:], c, readFields(bs, fields)
}

func (c *VgCounters) encode(w io.Writer) error {
	var err error

	err = binary.Write(w, binary.BigEndian, uint32(c.RecordType()))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, uint32(vgCountersSize))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, c)
	return err
}

// RecordType returns the type of counter record.
func (c *VlanCounters) RecordType() int {
	return TypeVlanCountersRecord
}

func decodeVlanCountersRecord(bs []byte) ([]byte, *VlanCounters, error) {
	if len(bs) < vlanCountersSize {
		return nil, nil, ErrDecodingRecord
	}

	c := &VlanCounters{}
	fields := []interface{}{
		&c.ID,
		&c.Octets,
		&c.UnicastPackets,
		&c.MulticastPackets,
		&c.BroadcastPackets,
		&c.Discards,
	}

	return bs[vlanCountersSize:], c, readFields(bs, fields)
}

func (c *VlanCounters) encode(w io.Writer) error {
	var err error

	err = binary.Write(w, binary.BigEndian, uint32(c.RecordType()))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, uint32(vlanCountersSize))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, c)
	return err
}

// RecordType returns the type of counter record.
func (c *ProcessorCounters) RecordType() int {
	return TypeProcessorCountersRecord
}

func decodeProcessorCountersRecord(bs []byte) ([]byte, *ProcessorCounters, error) {
	if len(bs) < processorCountersSize {
		return nil, nil, ErrDecodingRecord
	}

	c := &ProcessorCounters{}
	fields := []interface{}{
		&c.CPU5s,
		&c.CPU1m,
		&c.CPU5m,
		&c.TotalMemory,
		&c.FreeMemory,
	}

	return bs[processorCountersSize:], c, readFields(bs, fields)
}

func (c *ProcessorCounters) encode(w io.Writer) error {
	var err error

	err = binary.Write(w, binary.BigEndian, uint32(c.RecordType()))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, uint32(processorCountersSize))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, c)
	return err
}

// RecordType returns the type of counter record.
func (c *HostCPUCounters) RecordType() int {
	return TypeHostCPUCountersRecord
}

func decodeHostCPUCountersRecord(bs []byte) ([]byte, *HostCPUCounters, error) {
	if len(bs) != 68 && len(bs) != 80 {
		return nil, nil, ErrDecodingRecord
	}

	c := &HostCPUCounters{}
	fields := []interface{}{
		&c.Load1m,
		&c.Load5m,
		&c.Load15m,
		&c.ProcessesRunning,
		&c.ProcessesTotal,
		&c.NumCPU,
		&c.SpeedCPU,
		&c.Uptime,
		&c.CPUUser,
		&c.CPUNice,
		&c.CPUSys,
		&c.CPUIdle,
		&c.CPUWio,
		&c.CPUIntr,
		&c.CPUSoftIntr,
		&c.Interrupts,
		&c.ContextSwitches,
		&c.CPUSteal,
		&c.CPUGuest,
		&c.CPUGuestNice,
	}

	return bs[hostCPUCountersSize:], c, readFields(bs, fields)
}

func (c *HostCPUCounters) encode(w io.Writer) error {
	var err error

	err = binary.Write(w, binary.BigEndian, uint32(c.RecordType()))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, uint32(hostCPUCountersSize))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, c)
	return err
}

// RecordType returns the type of counter record.
func (c *HostMemoryCounters) RecordType() int {
	return TypeHostMemoryCountersRecord
}

func decodeHostMemoryCountersRecord(bs []byte) ([]byte, *HostMemoryCounters, error) {
	if len(bs) < hostMemoryCountersSize {
		return nil, nil, ErrDecodingRecord
	}

	c := &HostMemoryCounters{}
	fields := []interface{}{
		&c.Total,
		&c.Free,
		&c.Shared,
		&c.Buffers,
		&c.Cached,
		&c.SwapTotal,
		&c.SwapFree,
		&c.PageIn,
		&c.PageOut,
		&c.SwapIn,
		&c.SwapOut,
	}

	return bs[hostMemoryCountersSize:], c, readFields(bs, fields)
}

func (c *HostMemoryCounters) encode(w io.Writer) error {
	var err error

	err = binary.Write(w, binary.BigEndian, uint32(c.RecordType()))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, uint32(hostMemoryCountersSize))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, c)
	return err
}

// RecordType returns the type of counter record.
func (c *HostDiskCounters) RecordType() int {
	return TypeHostDiskCountersRecord
}

func decodeHostDiskCountersRecord(bs []byte) ([]byte, *HostDiskCounters, error) {
	if len(bs) < hostDiskCountersSize {
		return nil, nil, ErrDecodingRecord
	}

	c := &HostDiskCounters{}
	fields := []interface{}{
		&c.Total,
		&c.Free,
		&c.MaxUsedPercent,
		&c.Reads,
		&c.BytesRead,
		&c.ReadTime,
		&c.Writes,
		&c.BytesWritten,
		&c.WriteTime,
	}

	return bs[hostDiskCountersSize:], c, readFields(bs, fields)
}

func (c *HostDiskCounters) encode(w io.Writer) error {
	var err error

	err = binary.Write(w, binary.BigEndian, uint32(c.RecordType()))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, uint32(hostDiskCountersSize))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, c)
	return err
}

// RecordType returns the type of counter record.
func (c *HostNetCounters) RecordType() int {
	return TypeHostNetCountersRecord
}

func decodeHostNetCountersRecord(bs []byte) ([]byte, *HostNetCounters, error) {
	if len(bs) < hostNetCountersSize {
		return nil, nil, ErrDecodingRecord
	}

	c := &HostNetCounters{}

	fields := []interface{}{
		&c.BytesIn,
		&c.PacketsIn,
		&c.ErrorsIn,
		&c.DropsIn,
		&c.BytesOut,
		&c.PacketsOut,
		&c.ErrorsOut,
		&c.DropsOut,
	}

	return bs[hostNetCountersSize:], c, readFields(bs, fields)
}

func (c *HostNetCounters) encode(w io.Writer) error {
	var err error

	err = binary.Write(w, binary.BigEndian, uint32(c.RecordType()))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, uint32(hostNetCountersSize))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, c)
	return err
}

func DecodeCounterRecord(bs []byte) ([]byte, Record, error) {
	if len(bs) < 8 {
		return nil, nil, ErrInvalidSliceLength
	}

	format := binary.BigEndian.Uint32(bs)
	length := binary.BigEndian.Uint32(bs[4:])
	if len(bs) < 8+int(length) {
		return nil, nil, ErrInvalidSliceLength
	}

	switch format {
	case TypeGenericInterfaceCountersRecord:
		_, s, e := decodeGenericInterfaceCountersRecord(bs[8 : 8+length])
		return bs[8+length:], s, e
	case TypeEthernetCountersRecord:
		_, s, e := decodeEthernetCountersRecord(bs[8 : 8+length])
		return bs[8+length:], s, e
	case TypeTokenRingCountersRecord:
		_, s, e := decodeTokenRingCountersRecord(bs[8 : 8+length])
		return bs[8+length:], s, e
	case TypeVgCountersRecord:
		_, s, e := decodeVgCountersRecord(bs[8 : 8+length])
		return bs[8+length:], s, e
	case TypeVlanCountersRecord:
		_, s, e := decodeVlanCountersRecord(bs[8 : 8+length])
		return bs[8+length:], s, e
	case TypeProcessorCountersRecord:
		_, s, e := decodeProcessorCountersRecord(bs[8 : 8+length])
		return bs[8+length:], s, e
	case TypeHostCPUCountersRecord:
		_, s, e := decodeHostCPUCountersRecord(bs[8 : 8+length])
		return bs[8+length:], s, e
	case TypeHostMemoryCountersRecord:
		_, s, e := decodeHostMemoryCountersRecord(bs[8 : 8+length])
		return bs[8+length:], s, e
	case TypeHostDiskCountersRecord:
		_, s, e := decodeHostDiskCountersRecord(bs[8 : 8+length])
		return bs[8+length:], s, e
	case TypeHostNetCountersRecord:
		_, s, e := decodeHostNetCountersRecord(bs[8 : 8+length])
		return bs[8+length:], s, e
	default:
		return bs[8+length:], nil, nil
	}
}
