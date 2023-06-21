package logconverter
import (
	"strconv"
	"strings"
	"time"
)

type Log struct {
	Timestamp  string     `json:"@timestamp"`
	File       string     `json:"file"`
	Hostname   string     `json:"hostname"`
	Kubernetes Kubernetes `json:"kubernetes"`
	Level      string     `json:"level"`
	LogType    string     `json:"log_type"`
	Message    string     `json:"message"`
}

type Kubernetes struct {
	Annotations     Annotations     `json:"annotations"`
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
	K8sOvnOrgPodNetworks                                  string `json:"k8s.ovn.org/pod-networks"`
	K8sV1CniCncfIoNetworkStatus                           string `json:"k8s.v1.cni.cncf.io/network-status"`
	K8sV1CniCncfIoNetworksStatus                          string `json:"k8s.v1.cni.cncf.io/networks-status"`
	OpenshiftIoScc                                        string `json:"openshift.io/scc"`
	OperatorOpenshiftOauthApiserverEtcdClientSecret       string `json:"operator.openshift.io/dep-openshift-oauth-apiserver.etcd-client.secret"`
	OperatorOpenshiftOauthApiserverEtcdServingCaConfigmap string `json:"operator.openshift.io/dep-openshift-oauth-apiserver.etcd-serving-ca.configmap"`
}

type Labels struct {
	Apiserver                  string `json:"apiserver"`
	App                        string `json:"app"`
	OauthApiserverAntiAffinity string `json:"oauth-apiserver-anti-affinity"`
	PodTemplateHash            string `json:"pod-template-hash"`
	Revision                   string `json:"revision"`
}

type NamespaceLabels struct {
	KubernetesIoMetadataName                            string `json:"kubernetes.io/metadata.name"`
	OlmOperatorgroupUidD5ae8d2e99f34020998d9fc74c33faeb string `json:"olm.operatorgroup.uid/d5ae8d2e-99f3-4020-998d-9fc74c33faeb"`
	OpenshiftIoClusterMonitoring                        string `json:"openshift.io/cluster-monitoring"`
	PodSecurityKubernetesIoAudit                        string `json:"pod-security.kubernetes.io/audit"`
	PodSecurityKubernetesIoEnforce                      string `json:"pod-security.kubernetes.io/enforce"`
	PodSecurityKubernetesIoWarn                         string `json:"pod-security.kubernetes.io/warn"`
}

// New json file
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
	TheLog     TheLog       `json:"log"`
	Host       Host         `json:"host"`
	Container  Container    `json:"container"`
	K8s        K8s          `json:"k8s"`
	Attributes []Attributes `json:"attributes"`
}

type TheLog struct {
	TheFile TheFile `json:"file"`
}

type TheFile struct {
	Path string `json:"path"`
}

type Host struct {
	Name string `json:"name"`
}

type Container struct {
	Name  string `json:"name"`
	Id    string `json:"id"`
	Image Image  `json:"image"`
}

type Image struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type K8s struct {
	Pod       Pod       `json:"pod"`
	Namespace Namespace `json:"namespace"`
}

type Pod struct {
	Name           string         `json:"name"`
	Uid            string         `json:"uid"`
	Ip             string         `json:"ip"`
	Owner          string         `json:"owner"`
	NewAnnotations NewAnnotations `json:"annotations"`
	K8sLabels      K8sLabels      `json:"labels"`
}

type NewAnnotations struct {
	K8sOvnOrgPodNetworks                                  string `json:"k8s.ovn.org/pod-networks"`
	K8sV1CniCncfIoNetworkStatus                           string `json:"k8s.v1.cni.cncf.io/network-status"`
	K8sV1CniCncfIoNetworksStatus                          string `json:"k8s.v1.cni.cncf.io/networks-status"`
	OpenshiftIoScc                                        string `json:"openshift.io/scc"`
	OperatorOpenshiftOauthApiserverEtcdClientSecret       string `json:"operator.openshift.io/dep-openshift-oauth-apiserver.etcd-client.secret"`
	OperatorOpenshiftOauthApiserverEtcdServingCaConfigmap string `json:"operator.openshift.io/dep-openshift-oauth-apiserver.etcd-serving-ca.configmap"`
}

type K8sLabels struct {
	Apiserver                  string `json:"apiserver"`
	App                        string `json:"app"`
	OauthApiserverAntiAffinity string `json:"oauth-apiserver-anti-affinity"`
	PodTemplateHash            string `json:"pod-template-hash"`
	Revision                   string `json:"revision"`
}

type Namespace struct {
	Name               string             `json:"name"`
	NewNamespaceLabels NewNamespaceLabels `json:"labels"`
}

type NewNamespaceLabels struct {
	KubernetesIoMetadataName                            string `json:"kubernetes.io/metadata.name"`
	OlmOperatorgroupUidD5ae8d2e99f34020998d9fc74c33faeb string `json:"olm.operatorgroup.uid/d5ae8d2e-99f3-4020-998d-9fc74c33faeb"`
	OpenshiftIoClusterMonitoring                        string `json:"openshift.io/cluster-monitoring"`
	PodSecurityKubernetesIoAudit                        string `json:"pod-security.kubernetes.io/audit"`
	PodSecurityKubernetesIoEnforce                      string `json:"pod-security.kubernetes.io/enforce"`
	PodSecurityKubernetesIoWarn                         string `json:"pod-security.kubernetes.io/warn"`
}

type Attributes struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func ConvertLog(log Log) newLog{
	theNewLog := newLog{
		Timestamp:      changeTime(log.Timestamp),
		SeverityText:   log.Level,
		SeverityNumber: "17",
		Body: Body{
			Stringvalue: log.Message,
		},
		Resource: Resource{
			TheLog: TheLog{
				TheFile: TheFile{
					Path: log.File,
				},
			},
			Host: Host{
				Name: log.Hostname,
			},
			Container: Container{
				Name: log.Kubernetes.ContainerName,
				Id:   log.Kubernetes.ContainerId,
				Image: Image{
					Name: imageSpliceBefore(log.Kubernetes.ContainerImage),
					Tag:  imageSplice(log.Kubernetes.ContainerImage),
				},
			},
			K8s: K8s{
				Pod: Pod{
					Name:  log.Kubernetes.PodName,
					Uid:   log.Kubernetes.PodId,
					Ip:    log.Kubernetes.PodIp,
					Owner: log.Kubernetes.Annotations.K8sOvnOrgPodNetworks,
					NewAnnotations: NewAnnotations{
						K8sOvnOrgPodNetworks:                                  log.Kubernetes.Annotations.K8sOvnOrgPodNetworks,
						K8sV1CniCncfIoNetworkStatus:                           log.Kubernetes.Annotations.K8sV1CniCncfIoNetworkStatus,
						K8sV1CniCncfIoNetworksStatus:                          log.Kubernetes.Annotations.K8sV1CniCncfIoNetworksStatus,
						OpenshiftIoScc:                                        log.Kubernetes.Annotations.OpenshiftIoScc,
						OperatorOpenshiftOauthApiserverEtcdClientSecret:       log.Kubernetes.Annotations.OperatorOpenshiftOauthApiserverEtcdClientSecret,
						OperatorOpenshiftOauthApiserverEtcdServingCaConfigmap: log.Kubernetes.Annotations.OperatorOpenshiftOauthApiserverEtcdServingCaConfigmap,
					},
					K8sLabels: K8sLabels{
						Apiserver:                  log.Kubernetes.Labels.Apiserver,
						App:                        log.Kubernetes.NamespaceLabels.KubernetesIoMetadataName,
						OauthApiserverAntiAffinity: log.Kubernetes.Labels.OauthApiserverAntiAffinity,
						PodTemplateHash:            log.Kubernetes.Labels.PodTemplateHash,
						Revision:                   log.Kubernetes.Labels.Revision,
					},
				},
				Namespace: Namespace{
					Name: log.Kubernetes.NamespaceName,
					NewNamespaceLabels: NewNamespaceLabels{
						KubernetesIoMetadataName:                            log.Kubernetes.NamespaceLabels.KubernetesIoMetadataName,
						OlmOperatorgroupUidD5ae8d2e99f34020998d9fc74c33faeb: log.Kubernetes.NamespaceLabels.OlmOperatorgroupUidD5ae8d2e99f34020998d9fc74c33faeb,
						OpenshiftIoClusterMonitoring:                        log.Kubernetes.NamespaceLabels.OpenshiftIoClusterMonitoring,
						PodSecurityKubernetesIoAudit:                        log.Kubernetes.NamespaceLabels.PodSecurityKubernetesIoAudit,
						PodSecurityKubernetesIoEnforce:                      log.Kubernetes.NamespaceLabels.PodSecurityKubernetesIoEnforce,
						PodSecurityKubernetesIoWarn:                         log.Kubernetes.NamespaceLabels.PodSecurityKubernetesIoWarn,
					},
				},
			},
			Attributes: []Attributes{
				{
					Key:   "log_type",
					Value: log.LogType,
				},
			},
		},
	}
	return theNewLog
}
func changeTime(oldTime string) string {
	parsedTime, _ := time.Parse(time.RFC3339Nano, oldTime)
	unixNanoTime := parsedTime.UnixNano()
	return strconv.Itoa(int(unixNanoTime))
}

func imageSpliceBefore(oldImage string) string {
	before := strings.LastIndex(oldImage, ":")
	return oldImage[:before] // everything up to the last : is
}

func imageSplice(oldImage string) string {
	newImage := oldImage[strings.Index(oldImage, ":"):]
	return newImage
}