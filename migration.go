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

	non_ip := netgraph.Protocol{"NON_IP"}
	ip := netgraph.Protocol{"IP"}

	emitter := netgraph.System{"emitter"}
	emitterv1 := netgraph.Sv{"emitterv1", & emitter, []*netgraph.Protocol{&ip, &non_ip}}
	emitterv1_0 := netgraph.Svm{"emitterv1_0", & emitterv1}

	csd1.Name = "1"
	csd1.Svm = &emitterv1_0
	csd2.Name = "2"
	csd1bis.Name = "1bis"
	csd2bis.Name = "2bis"
	csdNet.Name = "net"

	edge1_2 := netgraph.Edge{&csd1,&csd2,&non_ip}
	edge1_1bis := netgraph.Edge{&csd1,&csd1bis,&ip}
	edge1bis_net := netgraph.Edge{&csd1bis,&csdNet,&ip}
	edgeNet_2bis := netgraph.Edge{&csdNet,&csd2bis,&ip}
	edge2bis_2 := netgraph.Edge{&csd2bis,&csd2,&ip}

	path := []*netgraph.Edge{&edge1_1bis,&edge1bis_net,&edgeNet_2bis,&edge2bis_2}

	fmt.Println(edge1_2.From.Name)
	fmt.Println(path[3].From.Name)

	TT := time.Since(T)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " Ready; T=" + TT.String())
}
