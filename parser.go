package main

import (
	"encoding/json"
	"fmt"
)

type Log struct {
	Timestamp string
	File string
	Hostname string
}

func main(){
	
	logJson := `{
		"@timestamp": "2022-10-20T14:53:47.653917399Z",
		"file": "/var/log/pods/test-log-generator_test-log-generator_cbc621a0-6b87-43bf-843c-ffeed2f1207d/test-log-generator/15.log",
		"hostname": "oscar7",
		"kubernetes": {
		  "annotations": {
			"k8s.ovn.org/pod-networks": "{\"default\":{\"ip_addresses\":[\"10.128.0.193/23\"],\"mac_address\":\"0a:58:0a:80:00:c1\",\"gateway_ips\":[\"10.128.0.1\"],\"ip_address\":\"10.128.0.193/23\",\"gateway_ip\":\"10.128.0.1\"}}",
			"k8s.v1.cni.cncf.io/network-status": "[{\n    \"name\": \"ovn-kubernetes\",\n    \"interface\": \"eth0\",\n    \"ips\": [\n        \"10.128.0.193\"\n    ],\n    \"mac\": \"0a:58:0a:80:00:c1\",\n    \"default\": true,\n    \"dns\": {}\n}]",
			"k8s.v1.cni.cncf.io/networks-status": "[{\n    \"name\": \"ovn-kubernetes\",\n    \"interface\": \"eth0\",\n    \"ips\": [\n        \"10.128.0.193\"\n    ],\n    \"mac\": \"0a:58:0a:80:00:c1\",\n    \"default\": true,\n    \"dns\": {}\n}]",
			"openshift.io/scc": "privileged"
		  },
		  "container_id": "cri-o://21c253ee55a9714ef21398f63db8b2adb56fc3eae2dcdf481002f924c0639113",
		  "container_image": "quay.io/rojacob/cluster-logging-load-client:0.0.1-db25b80",
		  "container_name": "test-log-generator",
		  "labels": {
			"run": "test-log-generator"
		  },
		  "namespace_labels": {
			"kubernetes.io/metadata.name": "test-log-generator",
			"pod-security.kubernetes.io/audit": "restricted",
			"pod-security.kubernetes.io/audit-version": "v1.24",
			"pod-security.kubernetes.io/warn": "restricted",
			"pod-security.kubernetes.io/warn-version": "v1.24"
		  },
		  "namespace_name": "test-log-generator",
		  "pod_id": "cbc621a0-6b87-43bf-843c-ffeed2f1207d",
		  "pod_ip": "10.128.0.193",
		  "pod_name": "test-log-generator"
		},
		"level": "default",
		"log_type": "application",
		"message": "Use \"logger [command] --help\" for more information about a command."
	  }`



	
	var log Log
	json.Unmarshal([]byte(logJson), &log)
	fmt.Println(log.Hostname)
}

