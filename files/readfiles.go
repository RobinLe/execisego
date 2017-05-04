package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	ca, _ := ioutil.ReadFile("/etc/kubernetes/pki/ca.pem")
	fmt.Println(string(ca))

	b, _ := ioutil.ReadFile("/etc/kubernetes/env.sh")

	re := regexp.MustCompile("CLUSTER_NAME=.*")
	clusterName := re.FindString(string(b))
	clusterName = clusterName[14 : len(clusterName)-1]
	fmt.Println(clusterName)

	re = regexp.MustCompile("MASTER_IP=.*")
	masterIP := re.FindString(string(b))[10:]
	fmt.Println(masterIP)

	re = regexp.MustCompile("SERVICE_IP_RANGE=.*")
	serviceIPRange := re.FindString(string(b))[17:]
	fmt.Println(serviceIPRange)

	re = regexp.MustCompile("CLUSTER_SERVICE_IP=.*")
	clusterServiceIP := re.FindString(string(b))[19:]
	fmt.Println(clusterServiceIP)

	re = regexp.MustCompile("DNS_SERVICE_IP=.*")
	dnsServiceIP := re.FindString(string(b))[15:]
	fmt.Println(dnsServiceIP)

	re = regexp.MustCompile("POD_NETWORK=.*")
	podNetwork := re.FindString(string(b))[12:]
	fmt.Println(podNetwork)
}
