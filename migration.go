package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"time"
	"github.com/thomaspeugeot/migration/netgraph"
)

func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " Start")
	T := time.Now()

	var	nyc_emitter,paris_receiver,nyc_emitterbis,csdNet,paris_receiverbis netgraph.Csd

	non_ip := netgraph.Protocol{"NON_IP"}
	ip := netgraph.Protocol{"IP"}

	emitter := netgraph.System{"emitter"}
	emitterv1 := netgraph.Sv{"emitterv1", & emitter, []*netgraph.Protocol{&ip, &non_ip}}
	emitterv1_0 := netgraph.Svm{"emitterv1_0", & emitterv1}

	receiver := netgraph.System{"receiver"}
	receiverv1 := netgraph.Sv{"receiverv1", & receiver, []*netgraph.Protocol{&ip, &non_ip}}
	receiverv1_0 := netgraph.Svm{"receiverv1_0", & receiverv1}

	nyc_emitter.Name = "1"
	nyc_emitter.Svm = &emitterv1_0
	
	paris_receiver.Name = "2"
	paris_receiver.Svm = &receiverv1_0

	nyc_emitterbis.Name = "1bis"
	paris_receiverbis.Name = "2bis"
	csdNet.Name = "net"

	edge1_2 := netgraph.Edge{&nyc_emitter,&paris_receiver,&non_ip}
	edge1_1bis := netgraph.Edge{&nyc_emitter,&nyc_emitterbis,&ip}
	edge1bis_net := netgraph.Edge{&nyc_emitterbis,&csdNet,&ip}
	edgeNet_2bis := netgraph.Edge{&csdNet,&paris_receiverbis,&ip}
	edge2bis_2 := netgraph.Edge{&paris_receiverbis,&paris_receiver,&ip}

	path := []*netgraph.Edge{&edge1_1bis,&edge1bis_net,&edgeNet_2bis,&edge2bis_2}

	fmt.Println(edge1_2.From.Name)
	fmt.Println(path[3].From.Name)

	mapCsd := make(map[string]netgraph.Csd)
	mapCsd[nyc_emitter.Name] = nyc_emitter
	mapCsd[paris_receiver.Name] = paris_receiver
	b, _ := json.Marshal( mapCsd)

	ioutil.WriteFile("/tmp/dat1", b, 0644)
	
	TT := time.Since(T)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " Ready; T=" + TT.String())
}
