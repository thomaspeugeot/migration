package main

import (
	"fmt"
	"time"
	"github.com/thomaspeugeot/migration/netgraph"
)

func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " Start")
	T := time.Now()

	var	csd1,csd2,csd1bis,csdNet,csd2bis netgraph.Csd

	csd1.Name = "1"
	csd2.Name = "2"
	csd1bis.Name = "1bis"
	csd2bis.Name = "2bis"
	csdNet.Name = "net"

	non_ip := netgraph.Protocol{"NON_IP"}
	ip := netgraph.Protocol{"IP"}
	edge1_2 := netgraph.CsdEdge{&csd1,&csd2,&non_ip}
	edge1_1bis := netgraph.CsdEdge{&csd1,&csd1bis,&ip}
	edge1bis_net := netgraph.CsdEdge{&csd1bis,&csdNet,&ip}
	edgeNet_2bis := netgraph.CsdEdge{&csdNet,&csd2bis,&ip}
	edge2bis_2 := netgraph.CsdEdge{&csd2bis,&csd2,&ip}

	path := [4]*netgraph.CsdEdge{&edge1_1bis,&edge1bis_net,&edgeNet_2bis,&edge2bis_2}

	fmt.Println(edge1_2.A.Name)
	fmt.Println(path[1].A.Name)

	TT := time.Since(T)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " Ready; T=" + TT.String())
}
