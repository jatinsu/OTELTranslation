package logconverter

import (
	"strconv"
	"strings"
	"time"

	"github.com/openshift/cluster-logging-operator/test/helpers/types"
)

// New json file
type newLog struct {
	Timestamp           string              `json:"timeUnixNano,omitempty"`
	SeverityText        string              `json:"severityText,omitempty"`
	SeverityNumber      string              `json:"severityNumber,omitempty"`
	Body                Body                `json:"body,omitempty"`
	NewPipelinemetadata NewPipelinemetadata `json:"pipeline_metadata,omitempty"`
	ViaqIndexName       string              `json:"viaq_index_name,omitempty"`
	ViaqMsgID           string              `json:"viaq_msg_id,omitempty"`
	NewOpenshiftMeta    NewOpenshiftMeta    `json:"openshift,omitempty"`
	Resource            Resource            `json:"resource,omitempty"`
}
type NewPipelinemetadata struct {
	NewLogCollector NewLogCollector `json:"collector,omitempty"`
}

type NewLogCollector struct {
	Ipaddr4    string `json:"ipaddr4,omitempty"`
	Name       string `json:"name,omitempty"`
	ReceivedAt string `json:"receivedAt,omitempty"`
	Version    string `json:"version,omitempty"`
}

type Body struct {
	Stringvalue string `json:"stringValue,omitempty"`
}

// need to add to some other struct
type NewOpenshiftMeta struct {
	ClusterID string `json:"cluster_id,omitempty"`

	OpenshiftLabels map[string]string `json:"labels,omitempty"`

	Sequence string `json:"sequence,omitempty"`
}

type NewJournalLog struct {
	newLog              `json:",inline,omitempty"`
	STREAMID            string     `json:"_STREAM_ID,omitempty"`
	SYSTEMDINVOCATIONID string     `json:"_SYSTEMD_INVOCATION_ID,omitempty"`
	NewSystemd          NewSystemd `json:"systemd,omitempty"`
}

type NewT struct {
	BOOTID              string `json:"BOOT_ID,omitempty"`
	CAPEFFECTIVE        string `json:"CAP_EFFECTIVE,omitempty"`
	CMDLINE             string `json:"CMDLINE,omitempty"`
	COMM                string `json:"COMM,omitempty"`
	EXE                 string `json:"EXE,omitempty"`
	GID                 string `json:"GID,omitempty"`
	MACHINEID           string `json:"MACHINE_ID,omitempty"`
	PID                 string `json:"PID,omitempty"`
	SELINUXCONTEXT      string `json:"SELINUX_CONTEXT,omitempty"`
	STREAMID            string `json:"STREAM_ID,omitempty"`
	SYSTEMDCGROUP       string `json:"SYSTEMD_CGROUP,omitempty"`
	SYSTEMDINVOCATIONID string `json:"SYSTEMD_INVOCATION_ID,omitempty"`
	SYSTEMDSLICE        string `json:"SYSTEMD_SLICE,omitempty"`
	SYSTEMDUNIT         string `json:"SYSTEMD_UNIT,omitempty"`
	TRANSPORT           string `json:"TRANSPORT,omitempty"`
	UID                 string `json:"UID,omitempty"`
}

type NewU struct {
	SYSLOGIDENTIFIER string `json:"SYSLOG_IDENTIFIER,omitempty"`
}

type NewSystemd struct {
	T NewT `json:"t,omitempty"`
	U NewU `json:"u,omitempty"`
}

type NewInfraLog struct {
	Docker              Docker              `json:"docker,omitempty"`
	Kubernetes          Kubernetes          `json:"kubernetes,omitempty"`
	Message             string              `json:"message,omitempty"`
	Level               string              `json:"level,omitempty"`
	Hostname            string              `json:"hostname,omitempty"`
	PipelineMetadata    NewPipelinemetadata `json:"pipeline_metadata,omitempty"`
	Timestamp           time.Time           `json:"@timestamp,omitempty"`
	LogType             string              `json:"log_type,omitempty"`
	ViaqIndexName       string              `json:"viaq_index_name,omitempty"`
	ViaqMsgID           string              `json:"viaq_msg_id,omitempty"`
	STREAMID            string              `json:"_STREAM_ID,omitempty"`
	SYSTEMDINVOCATIONID string              `json:"_SYSTEMD_INVOCATION_ID,omitempty"`
	Systemd             Systemd             `json:"systemd,omitempty"`
	OpenshiftLabels     NewOpenshiftMeta    `json:"openshift,omitempty"`
}

type Resource struct {
	TheLog     TheLog       `json:"log,omitempty"`
	Host       Host         `json:"host,omitempty"`
	Container  Container    `json:"container,omitempty"`
	K8s        K8s          `json:"k8s,omitempty"`
	Attributes []Attributes `json:"attributes,omitempty"`
}

type TheLog struct {
	TheFile TheFile `json:"file,omitempty"`
}

type TheFile struct {
	Path string `json:"path,omitempty"`
}

type Host struct {
	Name string `json:"name,omitempty"`
}

type Container struct {
	Name  string `json:"name,omitempty"`
	Id    string `json:"id,omitempty"`
	Image Image  `json:"image,omitempty"`
}

type Image struct {
	Name string `json:"name,omitempty"`
	Tag  string `json:"tag,omitempty"`
}

type K8s struct {
	Pod       Pod       `json:"pod,omitempty"`
	Namespace Namespace `json:"namespace,omitempty"`
}

type Pod struct {
	Name           string         `json:"name,omitempty"`
	Uid            string         `json:"uid,omitempty"`
	Ip             string         `json:"ip,omitempty"`
	Owner          string         `json:"owner,omitempty"`
	NewAnnotations NewAnnotations `json:"annotations,omitempty"`
	K8sLabels      K8sLabels      `json:"labels,omitempty"`
}

type NewAnnotations map[string]string

type K8sLabels map[string]string

type Namespace struct {
	Name               string             `json:"name,omitempty"`
	NewNamespaceLabels NewNamespaceLabels `json:"labels,omitempty"`
}

type NewNamespaceLabels map[string]string

type Attributes struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

func ConvertLog(log types.ContainerLog) newLog {
	theNewLog := newLog{
		Timestamp:      changeTime(log.ViaQCommon.Timestamp),
		SeverityText:   log.ViaQCommon.Level,
		SeverityNumber: severityTextToNumber(log.ViaQCommon.Level),
		Body: Body{
			Stringvalue: log.ViaQCommon.Message,
		},
		Resource: Resource{
			TheLog: TheLog{
				// TheFile: TheFile{
				// 	Path: log.File,
				// },
			},
			Host: Host{
				Name: log.Hostname,
			},
			Container: Container{
				Name: log.Kubernetes.ContainerName,
				Id:   log.Kubernetes.ContainerID,
				Image: Image{
					Name: imageSpliceBefore(log.Kubernetes.ContainerImage),
					Tag:  imageSplice(log.Kubernetes.ContainerImage),
				},
			},
			K8s: K8s{
				Pod: Pod{
					Name:           log.Kubernetes.PodName,
					Uid:            log.Kubernetes.PodID,
					Ip:             log.Kubernetes.PodIP,
					Owner:          log.Kubernetes.Host,
					NewAnnotations: NewAnnotations(log.Kubernetes.Annotations),
					// THIS
					K8sLabels: K8sLabels(log.Kubernetes.Labels),
				},
				Namespace: Namespace{
					Name: log.Kubernetes.NamespaceName,
					// THIS
					NewNamespaceLabels: NewNamespaceLabels(log.Kubernetes.NamespaceLabels),
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
func changeTime(oldTime time.Time) string {
	formattedTime := oldTime.Format("2006-01-02T15:04:05.999999999Z")
	parsedTime, _ := time.Parse(time.RFC3339Nano, formattedTime)
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

func severityTextToNumber(severityText string) string {
	switch severityText {
	case "trace":
		return "8"
	case "debug":
		return "7"
	case "info":
		return "6"
	case "notice":
		return "5"
	case "warn":
		return "4"
	case "err":
		return "3"
	case "crit":
		return "2"
	case "alert":
		return "1"
	case "emerg":
		return "0"
	default:
		return "9"
	}
}
