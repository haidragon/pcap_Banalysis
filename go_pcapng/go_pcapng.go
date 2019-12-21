package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	_ "net"  //没用到
	"strings"
)
var (
	pcapFile string ="test.pcapng"
	handle   *pcap.Handle
	err      error
)

//打印数据包
func test_dump_packet()  {
	// 打开数据包
	handle, err = pcap.OpenOffline(pcapFile)
	if err != nil { log.Fatal(err) }
	defer handle.Close()

	// Loop through packets in file
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}
}
//设置过滤 tcp and addr 112.80.248.75  https://biot.com/capstats/bpf.html
// https://staight.github.io/2018/07/25/BPF%E8%BF%87%E6%BB%A4%E8%A7%84%E5%88%99/
func test_dump_packet_set_filter_addr()  {
	// 打开数据包
	handle, err = pcap.OpenOffline(pcapFile)
	if err != nil { log.Fatal(err) }
	defer handle.Close()
	// Set filter
	var filter string = "tcp and host 112.80.248.75"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}
	// Loop through packets in file
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		//fmt.Println(packet)
		for Data :=range packet.Data(){
			fmt.Print(Data)
		}
		fmt.Println()
	}
}
//https://github.com/jilios/security/blob/c7dce2180dba16bc1cb9c84cfd51652ac7e7f401/packinject/decodeLayer.go
//https://github.com/emrechan/golang/blob/3d954b36e0ec3e7635d4d762b8f37f87435e2f1d/packetIpLister/main.go
//https://github.com/yamatcha/usingGopacket/tree/90cd05c963e372e45dbee2e82025b97f09f65e57/practice
//https://github.com/yamatcha/usingGopacket/blob/90cd05c963e372e45dbee2e82025b97f09f65e57/practice/decoding_packet.go
//https://github.com/yamatcha/usingGopacket/blob/90cd05c963e372e45dbee2e82025b97f09f65e57/practice/display_layer.go
//https://github.com/cipepser/goPacketCapture/blob/e644eb59f6abb7db3748577ef025c9eb43a72b03/src/main/decoding_packet.go
type Ethernet struct {
	error  error  //能同名
}
func printPacketInfo(packet gopacket.Packet) {
	// Let's see if the packet is an ethernet packet
	//	LayerTypeEthernet  = gopacket.RegisterLayerType(17, gopacket.LayerTypeMetadata{Name: "Ethernet", Decoder: gopacket.DecodeFunc(decodeEthernet)})
	//https://github.com/google/gopacket/blob/d3896096dccd01f640f8c1470c311feef18d2c64/layers/enums.go
	//https://github.com/ShturminMaxim/CSharpStudyProjects/blob/689412efa73bd42010971fc554260e585112da60/Reflexion/Reflexion/Program.cs
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	if ethernetLayer != nil {
		fmt.Println("Ethernet")
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)//断言 类型转换
		fmt.Println("Source MAC: ", ethernetPacket.SrcMAC)
		fmt.Println("Destination MAC: ", ethernetPacket.DstMAC)
		// Ethernet type is typically IPv4 but could be APR or other
		//https://github.com/google/gopacket/blob/d3896096dccd01f640f8c1470c311feef18d2c64/layers/enums.go
		fmt.Println("Ethernet type: ", ethernetPacket.EthernetType)
	}

	// Let's see if the packet is IP (even through the ether type told us)
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		fmt.Println("IPv4 layer detected. ")
		ip, _ := ipLayer.(*layers.IPv4)
		// IP layer variables:
		// Version (Ether 4 or 6)
		// IHL (IP Header Length in 32-bit words)
		// TOS, Length, Id, Flags, FragOffset, TTL, Protocol (TCP?),
		// Checksum, SrcIp, DstIp
		fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
		fmt.Println("Protocol: ", ip.Protocol)
		fmt.Println()
	}

	// Let's see if the packet is TCP
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		fmt.Println("TCP layer detected")
		tcp, _ := tcpLayer.(*layers.TCP)
		// TCP layer variables:
		// SrcPort, DstPort, Seq, Ack, DataOffset, Window, Checksum,
		// Urgent
		// Bool flags: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR, NS
		fmt.Printf("From port %d to %d\n", tcp.SrcPort, tcp.DstPort)
		fmt.Println("Sequence number: ", tcp.Seq)
		fmt.Println()
	}

	// Iterate over all layers, printing out each layer type
	fmt.Println("All packet layers:")
	for _, layer := range packet.Layers() {
		fmt.Println("- ", layer.LayerType())
	}

	// When iterating through packet.Layer() above,
	// if it lists Payload layer then that is the same as
	// this applicationLayer. applicationLayer contains the payload
	applicationLayer := packet.ApplicationLayer()
	if applicationLayer != nil {
		fmt.Println("Application layer/Payload found.")
		fmt.Printf("%s\n", applicationLayer.Payload())

		// Search for a string inside the payload
		if strings.Contains(string(applicationLayer.Payload()), "HTTP") {
			fmt.Println("HTTP found!")
		}
	}

	// Check for errors
	if err := packet.ErrorLayer(); err != nil {
		fmt.Println("Error decoding some part of the packet: ", err)
	}
}
func test_dump_packet_layers_info(){
	handle, err = pcap.OpenOffline(pcapFile)
	if err != nil { log.Fatal(err) }
	defer handle.Close()

	// Loop through packets in file
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		printPacketInfo(packet)
	}
}
//https://github.com/google/gopacket/blob/master/layers/layertypes.go
//https://github.com/google/gopacket/blob/0ad7f2610e344e58c1c95e2adda5c3258da8e97b/layers/ip4.go
//type IPv4 struct {
//	BaseLayer
//	Version    uint8
//	IHL        uint8
//	TOS        uint8
//	Length     uint16
//	Id         uint16
//	Flags      IPv4Flag
//	FragOffset uint16
//	TTL        uint8
//	Protocol   IPProtocol
//	Checksum   uint16
//	SrcIP      net.IP
//	DstIP      net.IP
//	Options    []IPv4Option
//	Padding    []byte
//}
func test_dump_packet_all_ip(){
	handle, err = pcap.OpenOffline(pcapFile)
	if err != nil { log.Fatal(err) }
	defer handle.Close()

	// Loop through packets in file
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		if ipLayer != nil {
			//fmt.Println("IPv4 layer detected. ")
			IPv4, _ := ipLayer.(*layers.IPv4)
			// IP layer variables:
			// Version (Ether 4 or 6)
			// IHL (IP Header Length in 32-bit words)
			// TOS, Length, Id, Flags, FragOffset, TTL, Protocol (TCP?),
			// Checksum, SrcIp, DstIp
			//fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
			//fmt.Println("Protocol: ", ip.Protocol)
			//fmt.Println()
			fmt.Println(IPv4.SrcIP)
			fmt.Println(IPv4.DstIP)
		}
	}
}
//https://github.com/google/gopacket/blob/d3896096dccd01f640f8c1470c311feef18d2c64/layers/dns.go
//type DNS struct {
//	BaseLayer
//
//	// Header fields
//	ID     uint16
//	QR     bool
//	OpCode DNSOpCode
//
//	AA bool  // Authoritative answer
//	TC bool  // Truncated
//	RD bool  // Recursion desired
//	RA bool  // Recursion available
//	Z  uint8 // Reserved for future use
//
//	ResponseCode DNSResponseCode
//	QDCount      uint16 // Number of questions to expect
//	ANCount      uint16 // Number of answers to expect
//	NSCount      uint16 // Number of authorities to expect
//	ARCount      uint16 // Number of additional records to expect
//
//	// Entries
//	Questions   []DNSQuestion
//	Answers     []DNSResourceRecord
//	Authorities []DNSResourceRecord
//	Additionals []DNSResourceRecord
//
//	// buffer for doing name decoding.  We use a single reusable buffer to avoid
//	// name decoding on a single object via multiple DecodeFromBytes calls
//	// requiring constant allocation of small byte slices.
//	buffer []byte
//}
func test_dump_packet_all_DNS(){
	handle, err = pcap.OpenOffline(pcapFile)
	if err != nil { log.Fatal(err) }
	defer handle.Close()

	// Loop through packets in file
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		ipLayer := packet.Layer(layers.LayerTypeDNS)
		if ipLayer != nil {
			//fmt.Println("IPv4 layer detected. ")
			dns, _ := ipLayer.(*layers.DNS)
			// IP layer variables:
			// Version (Ether 4 or 6)
			// IHL (IP Header Length in 32-bit words)
			// TOS, Length, Id, Flags, FragOffset, TTL, Protocol (TCP?),
			// Checksum, SrcIp, DstIp
			//fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
			//fmt.Println("Protocol: ", ip.Protocol)
			//fmt.Println()
			fmt.Println("dnsID:", dns.ID)
			fmt.Println("answers:", dns.ANCount)
			fmt.Println("是否是回应包：", dns.QR) //false查询、true回应
			fmt.Println("Queries:", string(dns.Questions[0].Name))

			//fmt.Println(dns.DstIP)
		}
	}
}
//https://github.com/google/gopacket/blob/master/layers/layertypes.go
//https://github.com/google/gopacket/blob/d3896096dccd01f640f8c1470c311feef18d2c64/layers/ethernet.go
//type Ethernet struct {
//	BaseLayer
//	SrcMAC, DstMAC net.HardwareAddr
//	EthernetType   EthernetType
//	// Length is only set if a length field exists within this header.  Ethernet
//	// headers follow two different standards, one that uses an EthernetType, the
//	// other which defines a length the follows with a LLC header (802.3).  If the
//	// former is the case, we set EthernetType and Length stays 0.  In the latter
//	// case, we set Length and EthernetType = EthernetTypeLLC.
//	Length uint16
//}
func test_dump_packet_all_MAC()  {
	handle, err = pcap.OpenOffline(pcapFile)
	if err != nil { log.Fatal(err) }
	defer handle.Close()
	// Loop through packets in file
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
		if ethernetLayer != nil {
			ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)//断言 类型转换
			fmt.Println(ethernetPacket.SrcMAC)
			fmt.Println(ethernetPacket.DstMAC)
		}
	}
}
func main() {
	test_dump_packet_all_MAC()
}


