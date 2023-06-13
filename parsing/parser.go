package main

import (
	"encoding/json"
	"fmt"
)

type Log struct {
	Timestamp  string `json:"@timestamp"`
	File       string
	Hostname   string
	Kubernetes Kubernetes
	Level      string
	LogType    string `json:"log_type"`
	Message    string
}

type Kubernetes struct {
	Annotations     Annotations
	ContainerId     string `json:"container_id"`
	ContainerImage  string `json:"container_image"`
	ContainerName   string `json:"container_name"`
	Labels          Labels
	NamespaceLabels NamespaceLabels `json:"namespace_labels"`
	NamespaceName   string          `json:"namespace_name"`
	PodId           string          `json:"pod_id"`
	PodIp           string          `json:"pod_ip"`
	PodName         string          `json:"pod_name"`
}

type Annotations struct {
	K8sOvnOrgPodNetworks K8sOvnOrgPodNetworks `json:"k8s.ovn.org/pod-networks"`
	K8sV1CniCncfIoNetworkStatus K8sV1CniCncfIoNetworkStatus `json:"k8s.v1.cni.cncf.io/network-status"`
	K8sV1CniCncfIoNetworksStatus K8sV1CniCncfIoNetworksStatus `json:"k8s.v1.cni.cncf.io/networks-status"`
	OpenshiftIoScc               string `json:"openshift.io/scc"`
}

type K8sOvnOrgPodNetworks struct {
	Default Default
}

type Default struct {
	IpAddresses []string `json:"ip_addresses"`
	MacAddress  string   `json:"mac_address"`
	GatewayIps  []string `json:"gateway_ips"`
	IpAddress   string   `json:"ip_address"`
	GatewayIp   string   `json:"gateway_ip"`
}

type K8sV1CniCncfIoNetworkStatus struct {
	Name      string   `json:"name"`
	Interface string   `json:"interface"`
	Ips       []string `json:"ips"`
	Mac       string   `json:"mac"`
	Default   bool     `json:"default"`
	Dns       Dns      `json:"dns"`
}


// need too add the dns struct and find examples of it!
type Dns struct {
}

type K8sV1CniCncfIoNetworksStatus struct {
	Name      string   `json:"name"`
	Interface string   `json:"interface"`
	Ips       []string `json:"ips"`
	Mac       string   `json:"mac"`
	Default   bool     `json:"default"`
	Dns       Dns      `json:"dns"`
}




type Labels struct {
	Run string
}

type NamespaceLabels struct {
	KubernetesIoMetadataName            string `json:"kubernetes.io/metadata.name"`
	PodSecurityKubernetesIoAudit        string `json:"pod-security.kubernetes.io/audit"`
	PodSecurityKubernetesIoAuditVersion string `json:"pod-security.kubernetes.io/audit-version"`
	PodSecurityKubernetesIoWarn         string `json:"pod-security.kubernetes.io/warn"`
	PodSecurityKubernetesIoWarnVersion  string `json:"pod-security.kubernetes.io/warn-version"`
}

func main() {

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

	// this initilizes the struct with the json data, and prints out whatever you want
	var log Log
	json.Unmarshal([]byte(logJson), &log)
	fmt.Println(log.Kubernetes.Annotations.OpenshiftIoScc)
}
