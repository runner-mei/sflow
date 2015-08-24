package sflow

import "encoding/binary"

const (
	// counter_data	0	1	if_counters	sFlow Version 5
	// counter_data	0	2	ethernet_counters	sFlow Version 5
	// counter_data	0	3	tokenring_counters	sFlow Version 5
	// counter_data	0	4	vg_counters	sFlow Version 5
	// counter_data	0	5	vlan_counters	sFlow Version 5
	// counter_data	0	6	ieee80211_counters	sFlow 802.11 Structures
	// counter_data	0	7	lag_port_stats	sFlow LAG Counters Structure
	// counter_data	0	8	slow_path_counts	Fast path / slow path
	// counter_data	0	9	ib_counters	sFlow InfiniBand Structures
	// counter_data	0	1001	processor	sFlow Version 5
	// counter_data	0	1002	radio_utilization	sFlow 802.11 Structures
	// counter-data	0	1003	queue_length	sFlow for queue length monitoring
	// counter-data	0	1004	of_port	sFlow OpenFlow Structures
	// counter-data	0	1005	port_name	sFlow OpenFlow Structures
	// counter data	0	2000	host_descr	sFlow Host Structures
	// counter_data	0	2001	host_adapters	sFlow Host Structures
	// counter_data	0	2002	host_parent	sFlow Host Structures
	// counter_data	0	2003	host_cpu	sFlow Host Structures
	// counter_data	0	2004	host_memory	sFlow Host Structures
	// counter_data	0	2005	host_disk_io	sFlow Host Structures
	// counter_data	0	2006	host_net_io	sFlow Host Structures
	// counter_data	0	2007	mib2_ip_group	sFlow Host TCP/IP Counters
	// counter_data	0	2008	mib2_icmp_group	sFlow Host TCP/IP Counters
	// counter_data	0	2009	mib2_tcp_group	sFlow Host TCP/IP Counters
	// counter_data	0	2010	mib2_udp_group	sFlow Host TCP/IP Counters
	// counter_data	0	2100	virt_node	sFlow Host Structures
	// counter_data	0	2101	virt_cpu	sFlow Host Structures
	// counter_data	0	2102	virt_memory	sFlow Host Structures
	// counter_data	0	2103	virt_disk_io	sFlow Host Structures
	// counter_data	0	2104	virt_net_io	sFlow Host Structures
	// counter_data	0	2105	jmx_runtime	sFlow Java Virtual Machine Structures
	// counter_data	0	2106	jmx_statistics	sFlow Java Virtual Machine Structures
	// counter_data	0	2200	memcached_counters (deprecated)	sFlow for memcached
	// counter_data	0	2201	http_counters	sFlow HTTP Structures
	// counter_data	0	2202	app_operations	sFlow Application Structures
	// counter_data	0	2203	app_resources	sFlow Application Structures
	// counter_data	0	2204	memcache_counters	sFlow Memcache Structures
	// counter_data	0	2206	app_workers	sFlow Application Structures
	// counter_data	0	2207	ovs_dp_stats	Open vSwitch performance monitoring
	// counter_data	0	3000	energy	Energy management
	// counter_data	0	3001	temperature	Energy management
	// counter_data	0	3002	humidity	Energy management
	// counter_data	0	3003	fans	Energy management
	// counter_data	4413	3	hw_tables	sFlow Broadcom Switch ASIC Table Utilization Structures
	// counter_data	5703	1	nvidia_gpu	sFlow NVML GPU Structures

	TypeGenericInterfaceCountersRecord = 1
	TypeEthernetCountersRecord         = 2
	TypeTokenRingCountersRecord        = 3
	TypeVgCountersRecord               = 4
	TypeVlanCountersRecord             = 5

	TypeProcessorCountersRecord  = 1001
	TypeHostCPUCountersRecord    = 2003
	TypeHostMemoryCountersRecord = 2004
	TypeHostDiskCountersRecord   = 2005
	TypeHostNetCountersRecord    = 2006

	// Custom (Enterprise) types
	TypeApplicationCountersRecord = (1)<<12 + 1
)

type CounterSample struct {
	SequenceNum      uint32
	SourceIdType     byte
	SourceIdIndexVal uint32 // NOTE: this is 3 bytes in the datagram
	NumRecords       uint32
	//Records          []Record
}

// SampleType returns the type of sFlow sample.
func (s *CounterSample) SampleType() int {
	return TypeCounterSample
}

// func (s *CounterSample) GetRecords() []Record {
// 	return s.Records
// }

func DecodeCounterSample(bs []byte) ([]byte, *CounterSample, error) {
	if len(bs) < 12 {
		return nil, nil, ErrInvalidSliceLength
	}

	s := &CounterSample{}
	s.SequenceNum = binary.BigEndian.Uint32(bs)
	s.SourceIdType = bs[4]
	s.SourceIdIndexVal = uint32(bs[7]) | uint32(bs[6]<<8) | uint32(bs[5]<<16)
	s.NumRecords = binary.BigEndian.Uint32(bs[8:])

	// for i := uint32(0); i < s.numRecords; i++ {
	// 	rec, err := decodeCounterRecord(r)
	// 	if nil != err {
	// 		return nil, err
	// 	}
	// 	if nil != rec {
	// 		s.Records = append(s.Records, rec)
	// 	}
	// }

	return bs[12:], s, nil
}

// func (s *CounterSample) encode(w io.Writer) error {
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
