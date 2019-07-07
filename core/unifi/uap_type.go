package unifi

import (
	"encoding/json"
	"time"
)

// UAP represents all the data from the Ubiquiti Controller for a Unifi Access Point.
type UAP struct {
	/* This was auto generated and then slowly edited by hand
	   to get all the data types right and graphable.
	*/
	ID           string   `json:"_id"`
	Adopted      FlexBool `json:"adopted"`
	AntennaTable []struct {
		Default   FlexBool `json:"default"`
		ID        FlexInt  `json:"id"`
		Name      string   `json:"name"`
		Wifi0Gain FlexInt  `json:"wifi0_gain"`
		Wifi1Gain FlexInt  `json:"wifi1_gain"`
	} `json:"antenna_table"`
	BandsteeringMode string `json:"bandsteering_mode,omitempty"`
	BoardRev         int    `json:"board_rev"`
	Cfgversion       string `json:"cfgversion"`
	ConfigNetwork    struct {
		Type string `json:"type"`
		IP   string `json:"ip"`
	} `json:"config_network"`
	CountrycodeTable []int `json:"countrycode_table"`
	EthernetTable    []struct {
		Mac     string  `json:"mac"`
		NumPort FlexInt `json:"num_port"`
		Name    string  `json:"name"`
	} `json:"ethernet_table"`
	FwCaps              int      `json:"fw_caps"`
	HasEth1             FlexBool `json:"has_eth1"`
	HasSpeaker          FlexBool `json:"has_speaker"`
	InformIP            string   `json:"inform_ip"`
	InformURL           string   `json:"inform_url"`
	IP                  string   `json:"ip"`
	LedOverride         string   `json:"led_override"`
	Mac                 string   `json:"mac"`
	MeshStaVapEnabled   FlexBool `json:"mesh_sta_vap_enabled"`
	Model               string   `json:"model"`
	Name                string   `json:"name"`
	OutdoorModeOverride string   `json:"outdoor_mode_override"`
	PortTable           []struct {
		PortIdx    FlexInt  `json:"port_idx"`
		OpMode     string   `json:"op_mode"`
		PortconfID string   `json:"portconf_id"`
		AttrNoEdit FlexBool `json:"attr_no_edit,omitempty"`
		Media      string   `json:"media"`
		Name       string   `json:"name"`
		PoeCaps    FlexInt  `json:"poe_caps"`
		PortPoe    FlexBool `json:"port_poe"`
		TxBytesR   FlexInt  `json:"tx_bytes-r"`
		RxBytesR   FlexInt  `json:"rx_bytes-r"`
		BytesR     FlexInt  `json:"bytes-r"`
		PortDelta  struct {
			TimeDelta FlexInt `json:"time_delta"`
		} `json:"port_delta"`
		Enable       FlexBool `json:"enable"`
		Masked       FlexBool `json:"masked"`
		AggregatedBy FlexBool `json:"aggregated_by"`
	} `json:"port_table"`
	RadioTable []struct {
		Radio              string   `json:"radio"`
		Name               string   `json:"name"`
		Channel            FlexInt  `json:"channel"`
		Ht                 string   `json:"ht"`
		TxPowerMode        string   `json:"tx_power_mode"`
		TxPower            FlexInt  `json:"tx_power"`
		MaxTxpower         FlexInt  `json:"max_txpower"`
		MinTxpower         FlexInt  `json:"min_txpower"`
		Nss                FlexInt  `json:"nss"`
		MinRssiEnabled     FlexBool `json:"min_rssi_enabled"`
		BuiltinAntenna     FlexBool `json:"builtin_antenna"`
		BuiltinAntGain     FlexInt  `json:"builtin_ant_gain"`
		CurrentAntennaGain FlexInt  `json:"current_antenna_gain"`
		RadioCaps          FlexInt  `json:"radio_caps"`
		WlangroupID        string   `json:"wlangroup_id"`
		Is11Ac             FlexBool `json:"is_11ac,omitempty"`
		HasDfs             FlexBool `json:"has_dfs,omitempty"`
		HasFccdfs          FlexBool `json:"has_fccdfs,omitempty"`
	} `json:"radio_table"`
	ScanRadioTable        []interface{} `json:"scan_radio_table"`
	Serial                string        `json:"serial"`
	SiteID                string        `json:"site_id"`
	SiteName              string        `json:"-"`
	Type                  string        `json:"type"`
	Version               string        `json:"version"`
	VwireTable            []interface{} `json:"vwire_table"`
	WifiCaps              int           `json:"wifi_caps"`
	WlangroupIDNa         string        `json:"wlangroup_id_na"`
	WlangroupIDNg         string        `json:"wlangroup_id_ng"`
	RequiredVersion       string        `json:"required_version"`
	HwCaps                int           `json:"hw_caps"`
	Unsupported           FlexBool      `json:"unsupported"`
	UnsupportedReason     FlexInt       `json:"unsupported_reason"`
	SysErrorCaps          int           `json:"sys_error_caps"`
	HasFan                FlexBool      `json:"has_fan"`
	HasTemperature        FlexBool      `json:"has_temperature"`
	DeviceID              string        `json:"device_id"`
	State                 int           `json:"state"`
	LastSeen              FlexInt       `json:"last_seen"`
	Upgradable            FlexBool      `json:"upgradable"`
	AdoptableWhenUpgraded FlexBool      `json:"adoptable_when_upgraded"`
	Rollupgrade           FlexBool      `json:"rollupgrade"`
	KnownCfgversion       string        `json:"known_cfgversion"`
	Uptime                FlexInt       `json:"uptime"`
	UUptime               FlexInt       `json:"_uptime"`
	Locating              FlexBool      `json:"locating"`
	ConnectRequestIP      string        `json:"connect_request_ip"`
	ConnectRequestPort    string        `json:"connect_request_port"`
	SysStats              struct {
		Loadavg1  float64 `json:"loadavg_1,string"`
		Loadavg15 float64 `json:"loadavg_15,string"`
		Loadavg5  float64 `json:"loadavg_5,string"`
		MemBuffer FlexInt `json:"mem_buffer"`
		MemTotal  FlexInt `json:"mem_total"`
		MemUsed   FlexInt `json:"mem_used"`
	} `json:"sys_stats"`
	SystemStats struct {
		CPU    float64 `json:"cpu,string"`
		Mem    float64 `json:"mem,string"`
		Uptime float64 `json:"uptime,string"`
	} `json:"system-stats"`
	SSHSessionTable  []interface{} `json:"ssh_session_table"`
	Scanning         FlexBool      `json:"scanning"`
	SpectrumScanning FlexBool      `json:"spectrum_scanning"`
	GuestToken       string        `json:"guest_token"`
	Meshv3PeerMac    string        `json:"meshv3_peer_mac"`
	Satisfaction     FlexInt       `json:"satisfaction"`
	Isolated         FlexBool      `json:"isolated"`
	RadioTableStats  []struct {
		Name        string      `json:"name"`
		Channel     FlexInt     `json:"channel"`
		Radio       string      `json:"radio"`
		AstTxto     interface{} `json:"ast_txto"`
		AstCst      interface{} `json:"ast_cst"`
		AstBeXmit   FlexInt     `json:"ast_be_xmit"`
		CuTotal     FlexInt     `json:"cu_total"`
		CuSelfRx    FlexInt     `json:"cu_self_rx"`
		CuSelfTx    FlexInt     `json:"cu_self_tx"`
		Gain        FlexInt     `json:"gain"`
		State       string      `json:"state"`
		Extchannel  FlexInt     `json:"extchannel"`
		TxPower     FlexInt     `json:"tx_power"`
		TxPackets   FlexInt     `json:"tx_packets"`
		TxRetries   FlexInt     `json:"tx_retries"`
		NumSta      FlexInt     `json:"num_sta"`
		GuestNumSta FlexInt     `json:"guest-num_sta"`
		UserNumSta  FlexInt     `json:"user-num_sta"`
	} `json:"radio_table_stats"`
	Uplink struct {
		FullDuplex       FlexBool `json:"full_duplex"`
		IP               string   `json:"ip"`
		Mac              string   `json:"mac"`
		MaxVlan          int      `json:"max_vlan"`
		Name             string   `json:"name"`
		Netmask          string   `json:"netmask"`
		NumPort          int      `json:"num_port"`
		RxBytes          FlexInt  `json:"rx_bytes"`
		RxDropped        FlexInt  `json:"rx_dropped"`
		RxErrors         FlexInt  `json:"rx_errors"`
		RxMulticast      FlexInt  `json:"rx_multicast"`
		RxPackets        FlexInt  `json:"rx_packets"`
		Speed            FlexInt  `json:"speed"`
		TxBytes          FlexInt  `json:"tx_bytes"`
		TxDropped        FlexInt  `json:"tx_dropped"`
		TxErrors         FlexInt  `json:"tx_errors"`
		TxPackets        FlexInt  `json:"tx_packets"`
		Up               FlexBool `json:"up"`
		MaxSpeed         FlexInt  `json:"max_speed"`
		Type             string   `json:"type"`
		TxBytesR         FlexInt  `json:"tx_bytes-r"`
		RxBytesR         FlexInt  `json:"rx_bytes-r"`
		UplinkMac        string   `json:"uplink_mac"`
		UplinkRemotePort int      `json:"uplink_remote_port"`
	} `json:"uplink"`
	VapTable []struct {
		AnomaliesBarChart struct {
			HighTCPLatency    FlexInt `json:"high_tcp_latency"`
			HighTCPPacketLoss FlexInt `json:"high_tcp_packet_loss"`
			HighWifiLatency   FlexInt `json:"high_wifi_latency"`
			HighWifiRetries   FlexInt `json:"high_wifi_retries"`
			LowPhyRate        FlexInt `json:"low_phy_rate"`
			PoorStreamEff     FlexInt `json:"poor_stream_eff"`
			SleepyClient      FlexInt `json:"sleepy_client"`
			StaArpTimeout     FlexInt `json:"sta_arp_timeout"`
			StaDNSTimeout     FlexInt `json:"sta_dns_timeout"`
			StaIPTimeout      FlexInt `json:"sta_ip_timeout"`
			WeakSignal        FlexInt `json:"weak_signal"`
		} `json:"anomalies_bar_chart"`
		AnomaliesBarChartNow struct {
			HighTCPLatency    FlexInt `json:"high_tcp_latency"`
			HighTCPPacketLoss FlexInt `json:"high_tcp_packet_loss"`
			HighWifiLatency   FlexInt `json:"high_wifi_latency"`
			HighWifiRetries   FlexInt `json:"high_wifi_retries"`
			LowPhyRate        FlexInt `json:"low_phy_rate"`
			PoorStreamEff     FlexInt `json:"poor_stream_eff"`
			SleepyClient      FlexInt `json:"sleepy_client"`
			StaArpTimeout     FlexInt `json:"sta_arp_timeout"`
			StaDNSTimeout     FlexInt `json:"sta_dns_timeout"`
			StaIPTimeout      FlexInt `json:"sta_ip_timeout"`
			WeakSignal        FlexInt `json:"weak_signal"`
		} `json:"anomalies_bar_chart_now"`
		AvgClientSignal     FlexInt `json:"avg_client_signal"`
		Bssid               string  `json:"bssid"`
		Ccq                 int     `json:"ccq"`
		Channel             int     `json:"channel"`
		Essid               string  `json:"essid"`
		Extchannel          int     `json:"extchannel"`
		ID                  string  `json:"id"`
		MacFilterRejections int     `json:"mac_filter_rejections"`
		Name                string  `json:"name"`
		NumSatisfactionSta  FlexInt `json:"num_satisfaction_sta"`
		NumSta              int     `json:"num_sta"`
		Radio               string  `json:"radio"`
		RadioName           string  `json:"radio_name"`
		ReasonsBarChart     struct {
			PhyRate       FlexInt `json:"phy_rate"`
			Signal        FlexInt `json:"signal"`
			SleepyClient  FlexInt `json:"sleepy_client"`
			StaArpTimeout FlexInt `json:"sta_arp_timeout"`
			StaDNSTimeout FlexInt `json:"sta_dns_timeout"`
			StaIPTimeout  FlexInt `json:"sta_ip_timeout"`
			StreamEff     FlexInt `json:"stream_eff"`
			TCPLatency    FlexInt `json:"tcp_latency"`
			TCPPacketLoss FlexInt `json:"tcp_packet_loss"`
			WifiLatency   FlexInt `json:"wifi_latency"`
			WifiRetries   FlexInt `json:"wifi_retries"`
		} `json:"reasons_bar_chart"`
		ReasonsBarChartNow struct {
			PhyRate       FlexInt `json:"phy_rate"`
			Signal        FlexInt `json:"signal"`
			SleepyClient  FlexInt `json:"sleepy_client"`
			StaArpTimeout FlexInt `json:"sta_arp_timeout"`
			StaDNSTimeout FlexInt `json:"sta_dns_timeout"`
			StaIPTimeout  FlexInt `json:"sta_ip_timeout"`
			StreamEff     FlexInt `json:"stream_eff"`
			TCPLatency    FlexInt `json:"tcp_latency"`
			TCPPacketLoss FlexInt `json:"tcp_packet_loss"`
			WifiLatency   FlexInt `json:"wifi_latency"`
			WifiRetries   FlexInt `json:"wifi_retries"`
		} `json:"reasons_bar_chart_now"`
		RxBytes    FlexInt `json:"rx_bytes"`
		RxCrypts   FlexInt `json:"rx_crypts"`
		RxDropped  FlexInt `json:"rx_dropped"`
		RxErrors   FlexInt `json:"rx_errors"`
		RxFrags    FlexInt `json:"rx_frags"`
		RxNwids    FlexInt `json:"rx_nwids"`
		RxPackets  FlexInt `json:"rx_packets"`
		RxTCPStats struct {
			Goodbytes FlexInt `json:"goodbytes"`
			LatAvg    FlexInt `json:"lat_avg"`
			LatMax    FlexInt `json:"lat_max"`
			LatMin    FlexInt `json:"lat_min"`
			Stalls    FlexInt `json:"stalls"`
		} `json:"rx_tcp_stats"`
		Satisfaction      FlexInt `json:"satisfaction"`
		SatisfactionNow   FlexInt `json:"satisfaction_now"`
		State             string  `json:"state"`
		TxBytes           FlexInt `json:"tx_bytes"`
		TxCombinedRetries FlexInt `json:"tx_combined_retries"`
		TxDataMpduBytes   FlexInt `json:"tx_data_mpdu_bytes"`
		TxDropped         FlexInt `json:"tx_dropped"`
		TxErrors          FlexInt `json:"tx_errors"`
		TxPackets         FlexInt `json:"tx_packets"`
		TxPower           FlexInt `json:"tx_power"`
		TxRetries         FlexInt `json:"tx_retries"`
		TxRtsRetries      FlexInt `json:"tx_rts_retries"`
		TxSuccess         FlexInt `json:"tx_success"`
		TxTCPStats        struct {
			Goodbytes FlexInt `json:"goodbytes"`
			LatAvg    FlexInt `json:"lat_avg"`
			LatMax    FlexInt `json:"lat_max"`
			LatMin    FlexInt `json:"lat_min"`
			Stalls    FlexInt `json:"stalls"`
		} `json:"tx_tcp_stats"`
		TxTotal          FlexInt  `json:"tx_total"`
		Up               FlexBool `json:"up"`
		Usage            string   `json:"usage"`
		WifiTxAttempts   FlexInt  `json:"wifi_tx_attempts"`
		WifiTxDropped    FlexInt  `json:"wifi_tx_dropped"`
		WifiTxLatencyMov struct {
			Avg        FlexInt `json:"avg"`
			Max        FlexInt `json:"max"`
			Min        FlexInt `json:"min"`
			Total      FlexInt `json:"total"`
			TotalCount FlexInt `json:"total_count"`
		} `json:"wifi_tx_latency_mov"`
		T          string      `json:"t"`
		WlanconfID string      `json:"wlanconf_id"`
		IsGuest    FlexBool    `json:"is_guest"`
		IsWep      FlexBool    `json:"is_wep"`
		ApMac      string      `json:"ap_mac"`
		MapID      interface{} `json:"map_id"`
		SiteID     string      `json:"site_id"`
	} `json:"vap_table"`
	DownlinkTable []interface{} `json:"downlink_table"`
	VwireVapTable []interface{} `json:"vwire_vap_table"`
	BytesD        FlexInt       `json:"bytes-d"`
	TxBytesD      FlexInt       `json:"tx_bytes-d"`
	RxBytesD      FlexInt       `json:"rx_bytes-d"`
	BytesR        FlexInt       `json:"bytes-r"`
	LastUplink    struct {
		UplinkMac        string `json:"uplink_mac"`
		UplinkRemotePort int    `json:"uplink_remote_port"`
	} `json:"last_uplink"`
	Stat          *UAPStat      `json:"stat"`
	TxBytes       FlexInt       `json:"tx_bytes"`
	RxBytes       FlexInt       `json:"rx_bytes"`
	Bytes         FlexInt       `json:"bytes"`
	VwireEnabled  FlexBool      `json:"vwireEnabled"`
	UplinkTable   []interface{} `json:"uplink_table"`
	NumSta        int           `json:"num_sta"`
	UserNumSta    int           `json:"user-num_sta"`
	GuestNumSta   int           `json:"guest-num_sta"`
	TwoPhaseAdopt FlexBool      `json:"two_phase_adopt,omitempty"`
}

// UAPStat holds the "stat" data for an access point.
// This is split out because of a JSON data format change from 5.10 to 5.11.
type UAPStat struct {
	*ap
}
type ap struct {
	SiteID                        string    `json:"site_id"`
	O                             string    `json:"o"`
	Oid                           string    `json:"oid"`
	Ap                            string    `json:"ap"`
	Time                          FlexInt   `json:"time"`
	Datetime                      time.Time `json:"datetime"`
	GuestWifi0RxPackets           FlexInt   `json:"guest-wifi0-rx_packets"`
	GuestWifi1RxPackets           FlexInt   `json:"guest-wifi1-rx_packets"`
	UserWifi1RxPackets            FlexInt   `json:"user-wifi1-rx_packets"`
	UserWifi0RxPackets            FlexInt   `json:"user-wifi0-rx_packets"`
	UserRxPackets                 FlexInt   `json:"user-rx_packets"`
	GuestRxPackets                FlexInt   `json:"guest-rx_packets"`
	Wifi0RxPackets                FlexInt   `json:"wifi0-rx_packets"`
	Wifi1RxPackets                FlexInt   `json:"wifi1-rx_packets"`
	RxPackets                     FlexInt   `json:"rx_packets"`
	GuestWifi0RxBytes             FlexInt   `json:"guest-wifi0-rx_bytes"`
	GuestWifi1RxBytes             FlexInt   `json:"guest-wifi1-rx_bytes"`
	UserWifi1RxBytes              FlexInt   `json:"user-wifi1-rx_bytes"`
	UserWifi0RxBytes              FlexInt   `json:"user-wifi0-rx_bytes"`
	UserRxBytes                   FlexInt   `json:"user-rx_bytes"`
	GuestRxBytes                  FlexInt   `json:"guest-rx_bytes"`
	Wifi0RxBytes                  FlexInt   `json:"wifi0-rx_bytes"`
	Wifi1RxBytes                  FlexInt   `json:"wifi1-rx_bytes"`
	RxBytes                       FlexInt   `json:"rx_bytes"`
	GuestWifi0RxErrors            FlexInt   `json:"guest-wifi0-rx_errors"`
	GuestWifi1RxErrors            FlexInt   `json:"guest-wifi1-rx_errors"`
	UserWifi1RxErrors             FlexInt   `json:"user-wifi1-rx_errors"`
	UserWifi0RxErrors             FlexInt   `json:"user-wifi0-rx_errors"`
	UserRxErrors                  FlexInt   `json:"user-rx_errors"`
	GuestRxErrors                 FlexInt   `json:"guest-rx_errors"`
	Wifi0RxErrors                 FlexInt   `json:"wifi0-rx_errors"`
	Wifi1RxErrors                 FlexInt   `json:"wifi1-rx_errors"`
	RxErrors                      FlexInt   `json:"rx_errors"`
	GuestWifi0RxDropped           FlexInt   `json:"guest-wifi0-rx_dropped"`
	GuestWifi1RxDropped           FlexInt   `json:"guest-wifi1-rx_dropped"`
	UserWifi1RxDropped            FlexInt   `json:"user-wifi1-rx_dropped"`
	UserWifi0RxDropped            FlexInt   `json:"user-wifi0-rx_dropped"`
	UserRxDropped                 FlexInt   `json:"user-rx_dropped"`
	GuestRxDropped                FlexInt   `json:"guest-rx_dropped"`
	Wifi0RxDropped                FlexInt   `json:"wifi0-rx_dropped"`
	Wifi1RxDropped                FlexInt   `json:"wifi1-rx_dropped"`
	RxDropped                     FlexInt   `json:"rx_dropped"`
	GuestWifi0RxCrypts            FlexInt   `json:"guest-wifi0-rx_crypts"`
	GuestWifi1RxCrypts            FlexInt   `json:"guest-wifi1-rx_crypts"`
	UserWifi1RxCrypts             FlexInt   `json:"user-wifi1-rx_crypts"`
	UserWifi0RxCrypts             FlexInt   `json:"user-wifi0-rx_crypts"`
	UserRxCrypts                  FlexInt   `json:"user-rx_crypts"`
	GuestRxCrypts                 FlexInt   `json:"guest-rx_crypts"`
	Wifi0RxCrypts                 FlexInt   `json:"wifi0-rx_crypts"`
	Wifi1RxCrypts                 FlexInt   `json:"wifi1-rx_crypts"`
	RxCrypts                      FlexInt   `json:"rx_crypts"`
	GuestWifi0RxFrags             FlexInt   `json:"guest-wifi0-rx_frags"`
	GuestWifi1RxFrags             FlexInt   `json:"guest-wifi1-rx_frags"`
	UserWifi1RxFrags              FlexInt   `json:"user-wifi1-rx_frags"`
	UserWifi0RxFrags              FlexInt   `json:"user-wifi0-rx_frags"`
	UserRxFrags                   FlexInt   `json:"user-rx_frags"`
	GuestRxFrags                  FlexInt   `json:"guest-rx_frags"`
	Wifi0RxFrags                  FlexInt   `json:"wifi0-rx_frags"`
	Wifi1RxFrags                  FlexInt   `json:"wifi1-rx_frags"`
	RxFrags                       FlexInt   `json:"rx_frags"`
	GuestWifi0TxPackets           FlexInt   `json:"guest-wifi0-tx_packets"`
	GuestWifi1TxPackets           FlexInt   `json:"guest-wifi1-tx_packets"`
	UserWifi1TxPackets            FlexInt   `json:"user-wifi1-tx_packets"`
	UserWifi0TxPackets            FlexInt   `json:"user-wifi0-tx_packets"`
	UserTxPackets                 FlexInt   `json:"user-tx_packets"`
	GuestTxPackets                FlexInt   `json:"guest-tx_packets"`
	Wifi0TxPackets                FlexInt   `json:"wifi0-tx_packets"`
	Wifi1TxPackets                FlexInt   `json:"wifi1-tx_packets"`
	TxPackets                     FlexInt   `json:"tx_packets"`
	GuestWifi0TxBytes             FlexInt   `json:"guest-wifi0-tx_bytes"`
	GuestWifi1TxBytes             FlexInt   `json:"guest-wifi1-tx_bytes"`
	UserWifi1TxBytes              FlexInt   `json:"user-wifi1-tx_bytes"`
	UserWifi0TxBytes              FlexInt   `json:"user-wifi0-tx_bytes"`
	UserTxBytes                   FlexInt   `json:"user-tx_bytes"`
	GuestTxBytes                  FlexInt   `json:"guest-tx_bytes"`
	Wifi0TxBytes                  FlexInt   `json:"wifi0-tx_bytes"`
	Wifi1TxBytes                  FlexInt   `json:"wifi1-tx_bytes"`
	TxBytes                       FlexInt   `json:"tx_bytes"`
	GuestWifi0TxErrors            FlexInt   `json:"guest-wifi0-tx_errors"`
	GuestWifi1TxErrors            FlexInt   `json:"guest-wifi1-tx_errors"`
	UserWifi1TxErrors             FlexInt   `json:"user-wifi1-tx_errors"`
	UserWifi0TxErrors             FlexInt   `json:"user-wifi0-tx_errors"`
	UserTxErrors                  FlexInt   `json:"user-tx_errors"`
	GuestTxErrors                 FlexInt   `json:"guest-tx_errors"`
	Wifi0TxErrors                 FlexInt   `json:"wifi0-tx_errors"`
	Wifi1TxErrors                 FlexInt   `json:"wifi1-tx_errors"`
	TxErrors                      FlexInt   `json:"tx_errors"`
	GuestWifi0TxDropped           FlexInt   `json:"guest-wifi0-tx_dropped"`
	GuestWifi1TxDropped           FlexInt   `json:"guest-wifi1-tx_dropped"`
	UserWifi1TxDropped            FlexInt   `json:"user-wifi1-tx_dropped"`
	UserWifi0TxDropped            FlexInt   `json:"user-wifi0-tx_dropped"`
	UserTxDropped                 FlexInt   `json:"user-tx_dropped"`
	GuestTxDropped                FlexInt   `json:"guest-tx_dropped"`
	Wifi0TxDropped                FlexInt   `json:"wifi0-tx_dropped"`
	Wifi1TxDropped                FlexInt   `json:"wifi1-tx_dropped"`
	TxDropped                     FlexInt   `json:"tx_dropped"`
	GuestWifi0TxRetries           FlexInt   `json:"guest-wifi0-tx_retries"`
	GuestWifi1TxRetries           FlexInt   `json:"guest-wifi1-tx_retries"`
	UserWifi1TxRetries            FlexInt   `json:"user-wifi1-tx_retries"`
	UserWifi0TxRetries            FlexInt   `json:"user-wifi0-tx_retries"`
	UserTxRetries                 FlexInt   `json:"user-tx_retries"`
	GuestTxRetries                FlexInt   `json:"guest-tx_retries"`
	Wifi0TxRetries                FlexInt   `json:"wifi0-tx_retries"`
	Wifi1TxRetries                FlexInt   `json:"wifi1-tx_retries"`
	TxRetries                     FlexInt   `json:"tx_retries"`
	GuestWifi0MacFilterRejections FlexInt   `json:"guest-wifi0-mac_filter_rejections"`
	GuestWifi1MacFilterRejections FlexInt   `json:"guest-wifi1-mac_filter_rejections"`
	UserWifi1MacFilterRejections  FlexInt   `json:"user-wifi1-mac_filter_rejections"`
	UserWifi0MacFilterRejections  FlexInt   `json:"user-wifi0-mac_filter_rejections"`
	UserMacFilterRejections       FlexInt   `json:"user-mac_filter_rejections"`
	GuestMacFilterRejections      FlexInt   `json:"guest-mac_filter_rejections"`
	Wifi0MacFilterRejections      FlexInt   `json:"wifi0-mac_filter_rejections"`
	Wifi1MacFilterRejections      FlexInt   `json:"wifi1-mac_filter_rejections"`
	MacFilterRejections           FlexInt   `json:"mac_filter_rejections"`
	GuestWifi0WifiTxAttempts      FlexInt   `json:"guest-wifi0-wifi_tx_attempts"`
	GuestWifi1WifiTxAttempts      FlexInt   `json:"guest-wifi1-wifi_tx_attempts"`
	UserWifi1WifiTxAttempts       FlexInt   `json:"user-wifi1-wifi_tx_attempts"`
	UserWifi0WifiTxAttempts       FlexInt   `json:"user-wifi0-wifi_tx_attempts"`
	UserWifiTxAttempts            FlexInt   `json:"user-wifi_tx_attempts"`
	GuestWifiTxAttempts           FlexInt   `json:"guest-wifi_tx_attempts"`
	Wifi0WifiTxAttempts           FlexInt   `json:"wifi0-wifi_tx_attempts"`
	Wifi1WifiTxAttempts           FlexInt   `json:"wifi1-wifi_tx_attempts"`
	WifiTxAttempts                FlexInt   `json:"wifi_tx_attempts"`
	GuestWifi0WifiTxDropped       FlexInt   `json:"guest-wifi0-wifi_tx_dropped"`
	GuestWifi1WifiTxDropped       FlexInt   `json:"guest-wifi1-wifi_tx_dropped"`
	UserWifi1WifiTxDropped        FlexInt   `json:"user-wifi1-wifi_tx_dropped"`
	UserWifi0WifiTxDropped        FlexInt   `json:"user-wifi0-wifi_tx_dropped"`
	UserWifiTxDropped             FlexInt   `json:"user-wifi_tx_dropped"`
	GuestWifiTxDropped            FlexInt   `json:"guest-wifi_tx_dropped"`
	Wifi0WifiTxDropped            FlexInt   `json:"wifi0-wifi_tx_dropped"`
	Wifi1WifiTxDropped            FlexInt   `json:"wifi1-wifi_tx_dropped"`
	WifiTxDropped                 FlexInt   `json:"wifi_tx_dropped"`
	Bytes                         FlexInt   `json:"bytes"`
	Duration                      FlexInt   `json:"duration"`
}

// UnmarshalJSON unmarshalls 5.10 or 5.11 formatted Access Point Stat data.
func (v *UAPStat) UnmarshalJSON(data []byte) error {
	var n struct {
		ap `json:"ap"`
	}
	v.ap = &n.ap
	err := json.Unmarshal(data, v.ap) // controller version 5.10.
	if err != nil {
		return json.Unmarshal(data, &n) // controller version 5.11.
	}
	return nil
}
