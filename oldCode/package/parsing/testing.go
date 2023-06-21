package main

import (
	"encoding/json"
	"fmt"
	"parser"
)

func main() {
	var log Log
	json.Unmarshal([]byte(logJson), &log)
	fmt.Printf("The timestamp is: %s\n", log.Timestamp)
	fmt.Printf("The file is: %s\n", log.File)
	fmt.Printf("The hostname is: %s\n", log.Hostname)
	fmt.Printf("The kubernetes container id is: %s\n", log.Kubernetes.ContainerId)
	fmt.Printf("The kubernetes container image is: %s\n", log.Kubernetes.ContainerImage)
	fmt.Printf("The kubernetes container name is: %s\n", log.Kubernetes.ContainerName)
	fmt.Printf("The kubernetes namespace name is: %s\n", log.Kubernetes.NamespaceName)
	fmt.Printf("The kubernetes pod id is: %s\n", log.Kubernetes.PodId)
	fmt.Printf("The kubernetes pod ip is: %s\n", log.Kubernetes.PodIp)
	fmt.Printf("The kubernetes pod name is: %s\n", log.Kubernetes.PodName)
	fmt.Printf("The kubernetes annotation k8s.ovn.org/pod-networks is: %s\n", log.Kubernetes.Annotations.K8sOvnOrgPodNetworks)
	fmt.Printf("The kubernetes annotation k8s.v1.cni.cncf.io/network-status is: %s\n", log.Kubernetes.Annotations.K8sV1CniCncfIoNetworkStatus)
	fmt.Printf("The kubernetes annotation k8s.v1.cni.cncf.io/networks-status is: %s\n", log.Kubernetes.Annotations.K8sV1CniCncfIoNetworksStatus)
	fmt.Printf("The kubernetes annotation openshift.io/scc is: %s\n", log.Kubernetes.Annotations.OpenshiftIoScc)
	fmt.Printf("The kubernetes label run is: %s\n", log.Kubernetes.Labels.Run)
	fmt.Printf("The kubernetes namespace label kubernetes.io/metadata.name is: %s\n", log.Kubernetes.NamespaceLabels.KubernetesIoMetadataName)
	fmt.Printf("The kubernetes namespace label pod-security.kubernetes.io/audit is: %s\n", log.Kubernetes.NamespaceLabels.PodSecurityKubernetesIoAudit)
	fmt.Printf("The kubernetes namespace label pod-security.kubernetes.io/audit-version is: %s\n", log.Kubernetes.NamespaceLabels.PodSecurityKubernetesIoAuditVersion)
	fmt.Printf("The kubernetes namespace label pod-security.kubernetes.io/warn is: %s\n", log.Kubernetes.NamespaceLabels.PodSecurityKubernetesIoWarn)
	fmt.Printf("The kubernetes namespace label pod-security.kubernetes.io/warn-version is: %s\n", log.Kubernetes.NamespaceLabels.PodSecurityKubernetesIoWarnVersion)
	fmt.Printf("The log level is: %s\n", log.Level)
	fmt.Printf("The log type is: %s\n", log.LogType)
	fmt.Printf("The log message is: %s\n", log.Message)
}
