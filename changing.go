package main

import (
	"encoding/json"
	"fmt"
)

type Log struct {
	Timestamp  string     `json:"@timestamp" yaml:"@timestamp"`
	File       string     `json:"file" yaml:"file"`
	Hostname   string     `json:"hostname" yaml:"hostname"`
	Kubernetes Kubernetes `json:"kubernetes"`
	Level      string     `json:"level"`
	LogType    string     `json:"log_type"`
	Message    string     `json:"message"`
}

type Kubernetes struct {
	Annotations     Annotations
	ContainerId     string          `json:"container_id"`
	ContainerImage  string          `json:"container_image"`
	ContainerName   string          `json:"container_name"`
	Labels          Labels          `json:"labels"`
	NamespaceLabels NamespaceLabels `json:"namespace_labels"`
	NamespaceName   string          `json:"namespace_name"`
	PodId           string          `json:"pod_id"`
	PodIp           string          `json:"pod_ip"`
	PodName         string          `json:"pod_name"`
}

type Annotations struct {
	K8sOvnOrgPodNetworks         string `json:"k8s.ovn.org/pod-networks"`
	K8sV1CniCncfIoNetworkStatus  string `json:"k8s.v1.cni.cncf.io/network-status"`
	K8sV1CniCncfIoNetworksStatus string `json:"k8s.v1.cni.cncf.io/networks-status"`
	OpenshiftIoScc               string `json:"openshift.io/scc"`
}

type Labels struct {
	Run string `json:"run"`
}

type NamespaceLabels struct {
	KubernetesIoMetadataName            string `json:"kubernetes.io/metadata.name"`
	PodSecurityKubernetesIoAudit        string `json:"pod-security.kubernetes.io/audit"`
	PodSecurityKubernetesIoAuditVersion string `json:"pod-security.kubernetes.io/audit-version"`
	PodSecurityKubernetesIoWarn         string `json:"pod-security.kubernetes.io/warn"`
	PodSecurityKubernetesIoWarnVersion  string `json:"pod-security.kubernetes.io/warn-version"`
}

type newLog struct {
	Timestamp      string   `json:"timeUnixNano"`
	SeverityText   string   `json:"severityText"`
	SeverityNumber string   `json:"severityNumber"`
	Body           Body     `json:"body"`
	Resource       Resource `json:"resource"`
}
type Body struct {
	Stringvalue string `json:"stringValue"`
}

type Resource struct {
	Log  Log  `json:"log"`
	Host Host `json:"host"`
}

func main() {
	newLogJson := `{
		"timeUnixNano": 1666278690764362932,
		"severityText": "error",
		"severityNumber": 17,
		"body": {
		"stringValue": "E1020 15:11:30.764269       1 timeout.go:137] post-timeout activity - time-elapsed: 1.321595349s, GET \"/readyz\" result: <nil>"
		},
		"resource": {
		"log": {
			"file": {
			"path": "/var/log/pods/openshift-oauth-apiserver_apiserver-b477bc494-hmj4r_998cef46-bccd-437e-9727-4d6389436885/oauth-apiserver/6.log"
			}
		},
		"host": {
			"name": "oscar7"
		},
		"container": {
			"name": "oauth-apiserver",
			"id": "cri-o://ae8d9c8e46defc9ff1df9a35076e04e7bd6be2d9056500570bf9e3f85e9d6885",
			"image": {
			"name": "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256",
			"tag":":ed90fd557cc619f98a99bc8c552ee7b8a8ee369a3a2cdf2f9a4727878d2d049e"
			}
		},
		"k8s": {
			"pod": {
			"name": "apiserver-b477bc494-hmj4r",
			"uid": "998cef46-bccd-437e-9727-4d6389436885",
			"ip": "10.128.0.33",
			"owner": "ReplicaSet/apiserver-b477bc494",
			"annotations": {
				"k8s.ovn.org/pod-networks": "{\"default\":{\"ip_addresses\":[\"10.128.0.33/23\"],\"mac_address\":\"0a:58:0a:80:00:21\",\"gateway_ips\":[\"10.128.0.1\"],\"ip_address\":\"10.128.0.33/23\",\"gateway_ip\":\"10.128.0.1\"}}",
				"k8s.v1.cni.cncf.io/network-status": "[{\n    \"name\": \"ovn-kubernetes\",\n    \"interface\": \"eth0\",\n    \"ips\": [\n        \"10.128.0.33\"\n    ],\n    \"mac\": \"0a:58:0a:80:00:21\",\n    \"default\": true,\n    \"dns\": {}\n}]",
				"k8s.v1.cni.cncf.io/networks-status": "[{\n    \"name\": \"ovn-kubernetes\",\n    \"interface\": \"eth0\",\n    \"ips\": [\n        \"10.128.0.33\"\n    ],\n    \"mac\": \"0a:58:0a:80:00:21\",\n    \"default\": true,\n    \"dns\": {}\n}]",
				"openshift.io/scc": "privileged",
				"operator.openshift.io/dep-openshift-oauth-apiserver.etcd-client.secret": "OFllOQ==",
				"operator.openshift.io/dep-openshift-oauth-apiserver.etcd-serving-ca.configmap": "f1B6eQ=="
			},
			"labels": {
				"apiserver": "true",
				"app": "openshift-oauth-apiserver",
				"oauth-apiserver-anti-affinity": "true",
				"pod-template-hash": "b477bc494",
				"revision": "2"
			}
		},
			"namespace": {
			"name": "openshift-oauth-apiserver",
			"labels": {
				"kubernetes.io/metadata.name": "openshift-oauth-apiserver",
				"olm.operatorgroup.uid/d5ae8d2e-99f3-4020-998d-9fc74c33faeb": "",
				"openshift.io/cluster-monitoring": "true",
				"pod-security.kubernetes.io/audit": "privileged",
				"pod-security.kubernetes.io/enforce": "privileged",
				"pod-security.kubernetes.io/warn": "privileged"
			}
			}
		},
		"attributes": [
			{
			"key": "log_type",
			"value": "infrastructure"
		}
		]
		}
	}`

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

	targetData := newLog{
		Timestamp: log.Timestamp,
	}

	outputJSON, _ := json.Marshal(targetData)
	fmt.Println(string(outputJSON))
}
