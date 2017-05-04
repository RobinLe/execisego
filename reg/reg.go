package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

// LoadDefaultCluster load the default cluster data
func main() {
	b, _ := ioutil.ReadFile("/etc/kubernetes/env.sh")

	re := regexp.MustCompile("CLUSTER_NAME=.*")
	clusterName := re.FindString(string(b))
	clusterName = getValue(clusterName, "CLUSTER_NAME")
	fmt.Println(clusterName)

	re = regexp.MustCompile("MASTER_IP=.*")
	masterIP := re.FindString(string(b))
	masterIP = getValue(masterIP, "MASTER_IP")
	fmt.Println(masterIP)

	re = regexp.MustCompile("SERVICE_IP_RANGE=.*")
	serviceIPRange := re.FindString(string(b))
	serviceIPRange = getValue(serviceIPRange, "SERVICE_IP_RANGE")
	fmt.Println(serviceIPRange)

	re = regexp.MustCompile("CLUSTER_SERVICE_IP=.*")
	clusterServiceIP := re.FindString(string(b))
	clusterServiceIP = getValue(clusterServiceIP, "CLUSTER_SERVICE_IP")
	fmt.Println(clusterServiceIP)

	re = regexp.MustCompile("DNS_SERVICE_IP=.*")
	dnsServiceIP := re.FindString(string(b))
	dnsServiceIP = getValue(dnsServiceIP, "DNS_SERVICE_IP")
	fmt.Println(dnsServiceIP)

	re = regexp.MustCompile("POD_NETWORK=.*")
	podNetwork := re.FindString(string(b))
	podNetwork = getValue(podNetwork, "POD_NETWORK")
	fmt.Println(podNetwork)
}

func getValue(rawStr, compareStr string) string {
	reSet := regexp.MustCompile("{.*:")
	reDefault := regexp.MustCompile("-.*}")

	setValue := reSet.FindString(rawStr)
	defalutValue := reDefault.FindString(rawStr)
	if setValue[1:len(setValue)-1] != compareStr {
		return setValue[1 : len(setValue)-1]
	}
	return defalutValue[1 : len(defalutValue)-1]
}
