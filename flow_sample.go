package sflow

import "encoding/binary"

const (
	// flow_data	0	1	sampled_header	sFlow Version 5
	// flow_data	0	2	sampled_ethernet	sFlow Version 5
	// flow_data	0	3	sampled_ipv4	sFlow Version 5
	// flow_data	0	4	sampled_ipv6	sFlow Version 5
	// flow_data	0	1001	extended_switch	sFlow Version 5
	// flow_data	0	1002	extended_router	sFlow Version 5
	// flow_data	0	1003	extended_gateway	sFlow Version 5
	// flow_data	0	1004	extended_user	sFlow Version 5
	// flow_data	0	1005	extended_url (deprecated)	sFlow Version 5
	// flow_data	0	1006	extended_mpls	sFlow Version 5
	// flow_data	0	1007	extended_nat	sFlow Version 5
	// flow_data	0	1008	extended_mpls_tunnel	sFlow Version 5
	// flow_data	0	1009	extended_mpls_vc	sFlow Version 5
	// flow_data	0	1010	extended_mpls_FTN	sFlow Version 5
	// flow_data	0	1011	extended_mpls_LDP_FEC	sFlow Version 5
	// flow_data	0	1012	extended_vlantunnel	sFlow Version 5
	// flow_data	0	1013	extended_80211_payload	sFlow 802.11 Structures
	// flow_data	0	1014	extended_80211_rx	sFlow 802.11 Structures
	// flow_data	0	1015	extended_80211_tx	sFlow 802.11 Structures
	// flow_data	0	1016	extended_80211_aggregation	sFlow 802.11 Structures
	// flow_data	0	1017	extended_openflow_v1 (deprecated)	sFlow OpenFlow Structures
	// flow_data	0	1018	extended_fc	sFlow, CEE and FCoE
	// flow_data	0	1019	extended_queue_length	sFlow for queue length monitoring
	// flow_data	0	1020	extended_nat_port	sFlow Port NAT Structures
	// flow_data	0	1021	extended_L2_tunnel_egress	sFlow Tunnel Structures
	// flow_data	0	1022	extended_L2_tunnel_ingress	sFlow Tunnel Structures
	// flow_data	0	1023	extended_ipv4_tunnel_egress	sFlow Tunnel Structures
	// flow_data	0	1024	extended_ipv4_tunnel_ingress	sFlow Tunnel Structures
	// flow_data	0	1025	extended_ipv6_tunnel_egress	sFlow Tunnel Structures
	// flow_data	0	1026	extended_ipv6_tunnel_ingress	sFlow Tunnel Structures
	// flow_data	0	1027	extended_decapsulate_egress	sFlow Tunnel Structures
	// flow_data	0	1028	extended_decapsulate_ingress	sFlow Tunnel Structures
	// flow_data	0	1029	extended_vni_egress	sFlow Tunnel Structures
	// flow_data	0	1030	extended_vni_ingress	sFlow Tunnel Structures
	// flow_data	0	1031	extended_ib_lrh	sFlow InfiniBand Structures
	// flow_data	0	1032	extended_ib_grh	sFlow InfiniBand Structures
	// flow_data	0	1033	extended_ib_brh	sFlow InfiniBand Structures
	// flow_data	0	2000	transaction	Host performance statistics
	// flow_data	0	2001	extended_nfs_storage_transaction	Host performance statistics
	// flow_data	0	2002	extensed_scsi_storage_transaction	Host performance statistics
	// flow_data	0	2003	extended_http_transaction	Host performance statistics
	// flow_data	0	2100	extended_socket_ipv4	sFlow Host Structures
	// flow_data	0	2101	extended_socket_ipv6	sFlow Host Structures
	// flow_data	0	2102	extended_proxy_socket_ipv4	sFlow HTTP Structures
	// flow_data	0	2103	extended_proxy_socket_ipv6 	sFlow HTTP Structures
	// flow_data	0	2200	memcached_operation	sFlow Memcache Structures
	// flow_data	0	2201	http_request (deprecated)	sFlow for HTTP
	// flow_data	0	2202	app_operation	sFlow Application Structures
	// flow_data	0	2203	app_parent_context	sFlow Application Structures
	// flow_data	0	2204	app_initiator	sFlow Application Structures
	// flow_data	0	2205	app_target	sFlow Application Structures
	// flow_data	0	2206	http_request	sFlow HTTP Structures
	// flow_data	0	2207	extended_proxy_request	sFlow HTTP Structures
	// flow_data	0	2208	extended_nav_timing	Navigation Timing

	TypeRawPacketFlowRecord     = 1
	TypeEthernetFrameFlowRecord = 2
	TypeIpv4FlowRecord          = 3
	TypeIpv6FlowRecord          = 4

	TypeExtendedSwitchFlowRecord     = 1001
	TypeExtendedRouterFlowRecord     = 1002
	TypeExtendedGatewayFlowRecord    = 1003
	TypeExtendedUserFlowRecord       = 1004
	TypeExtendedUrlFlowRecord        = 1005
	TypeExtendedMlpsFlowRecord       = 1006
	TypeExtendedNatFlowRecord        = 1007
	TypeExtendedMlpsTunnelFlowRecord = 1008
	TypeExtendedMlpsVcFlowRecord     = 1009
	TypeExtendedMlpsFecFlowRecord    = 1010
	TypeExtendedMlpsLvpFecFlowRecord = 1011
	TypeExtendedVlanFlowRecord       = 1012
)

type FlowSample struct {
	SequenceNum      uint32
	SourceIdType     byte
	SourceIdIndexVal uint32 // NOTE: this is 3 bytes in the datagram
	SamplingRate     uint32
	SamplePool       uint32
	Drops            uint32
	Input            uint32
	Output           uint32
	NumRecords       uint32
	//Records          []Record
}

// SampleType returns the type of sFlow sample.
func (s *FlowSample) SampleType() int {
	return TypeFlowSample
}

// func (s *FlowSample) GetRecords() []Record {
// 	return s.Records
// }

func DecodeFlowSample(bs []byte) ([]byte, *FlowSample, error) {
	if len(bs) < 32 {
		return nil, nil, ErrInvalidSliceLength
	}

	s := &FlowSample{}
	s.SequenceNum = binary.BigEndian.Uint32(bs)
	s.SourceIdType = bs[4]
	s.SourceIdIndexVal = uint32(bs[7]) | uint32(bs[6]<<8) | uint32(bs[5]<<16)
	s.SamplingRate = binary.BigEndian.Uint32(bs[8:])
	s.SamplePool = binary.BigEndian.Uint32(bs[12:])
	s.Drops = binary.BigEndian.Uint32(bs[16:])
	s.Input = binary.BigEndian.Uint32(bs[20:])
	s.Output = binary.BigEndian.Uint32(bs[24:])
	s.NumRecords = binary.BigEndian.Uint32(bs[28:])

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

	return bs[32:], s, nil
}
