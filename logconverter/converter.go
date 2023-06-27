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
	NewInfraLog         NewInfraLog         `json:"infra,omitempty"`
	NewSystemd          NewSystemd          `json:"systemd,omitempty"`
	LinuxAuditLog       `json:",inline,omitempty"`
}

type NewDocker struct {
	// ContainerID is the id of the container producing the log
	ContainerID string `json:"container_id"`
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
	NewDocker           NewDocker           `json:"docker,omitempty"`
	K8s                 K8s                 `json:"K8s,omitempty"`
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
	NewSystemd          NewSystemd          `json:"systemd,omitempty"`
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

type LinuxAuditLog struct {
	Hostname         string              `json:"hostname"`
	AuditLinux       AuditLinux          `json:"audit.linux"`
	Message          string              `json:"message,omitempty"`
	PipelineMetadata NewPipelinemetadata `json:"pipeline_metadata"`
	Timestamp        time.Time           `json:"@timestamp"`
	LogType          string              `json:"log_type,omitempty"`
	ViaqIndexName    string              `json:"viaq_index_name"`
	ViaqMsgID        string              `json:"viaq_msg_id"`
	Kubernetes       K8s                 `json:"K8s"`
	Openshift        NewOpenshiftMeta    `json:"openshift"`
	Timing           string              `json:",inline"`
	Level            string              `json:"level,omitempty"`
}

type AuditLinux struct {
	Type     string `json:"type,omitempty"`
	RecordID string `json:"record_id,omitempty"`
}

type OVNAuditLog struct {
	Hostname            string              `json:"hostname"`
	Message             string              `json:"message,omitempty"`
	NewPipelinemetadata NewPipelinemetadata `json:"pipeline_metadata"`
	Timestamp           time.Time           `json:"@timestamp"`
	LogType             string              `json:"log_type,omitempty"`
	ViaqIndexName       string              `json:"viaq_index_name"`
	ViaqMsgID           string              `json:"viaq_msg_id"`
	Kubernetes          K8s                 `json:"K8s"`
	NewOpenshiftMeta    NewOpenshiftMeta    `json:"openshift"`
	Level               string              `json:"level,omitempty"`
}

type AuditLogCommon struct {
	Kind                     string              `json:"kind,omitempty"`
	APIVersion               string              `json:"apiVersion,omitempty"`
	Level                    string              `json:"level,omitempty"`
	AuditID                  string              `json:"auditID,omitempty"`
	Stage                    string              `json:"stage,omitempty"`
	RequestURI               string              `json:"requestURI,omitempty"`
	Verb                     string              `json:"verb,omitempty"`
	User                     NewUser             `json:"user,omitempty"`
	SourceIPs                []string            `json:"sourceIPs,omitempty"`
	UserAgent                string              `json:"userAgent,omitempty"`
	ObjectRef                NewObjectRef        `json:"objectRef,omitempty"`
	ResponseStatus           NewResponseStatus   `json:"responseStatus,omitempty"`
	RequestReceivedTimestamp time.Time           `json:"requestReceivedTimestamp,omitempty"`
	StageTimestamp           time.Time           `json:"stageTimestamp,omitempty"`
	NewAnnotations           NewAnnotations      `json:"annotations,omitempty"`
	Message                  interface{}         `json:"message,omitempty"`
	Hostname                 string              `json:"hostname,omitempty"`
	NewPipelinemetadata      NewPipelinemetadata `json:"pipeline_metadata,omitempty"`
	Timestamp                time.Time           `json:"@timestamp,omitempty"`
	LogType                  string              `json:"log_type,omitempty"`
	ViaqIndexName            string              `json:"viaq_index_name,omitempty"`
	ViaqMsgID                string              `json:"viaq_msg_id,omitempty"`
	Kubernetes               K8s                 `json:"K8s,omitempty"`
	OpenshiftLabels          NewOpenshiftMeta    `json:"openshift,omitempty"`
	Timing                   string              `json:",inline"`
}

type EventRouterLog struct {
	NewDocker        NewDocker           `json:"docker"`
	K8s              K8s                 `json:"K8s"`
	Message          string              `json:"message,omitempty"`
	Level            string              `json:"level"`
	Hostname         string              `json:"hostname,omitempty"`
	PipelineMetadata NewPipelinemetadata `json:"pipeline_metadata"`
	Timestamp        time.Time           `json:"@timestamp"`
	LogType          string              `json:"log_type,omitempty"`
	ViaqIndexName    string              `json:"viaq_index_name"`
	ViaqMsgID        string              `json:"viaq_msg_id"`
	OpenshiftLabels  NewOpenshiftMeta    `json:"openshift"`
	Timing           string              `json:",inline"`
}

type NewUser struct {
	Username string   `json:"username,omitempty"`
	UID      string   `json:"uid,omitempty"`
	Groups   []string `json:"groups,omitempty"`
}

type NewObjectRef struct {
	Resource        string `json:"resource,omitempty"`
	ResourceVersion string `json:"resourceVersion,omitempty"`
	Name            string `json:"name,omitempty"`
	Namespace       string `json:"namespace,omitempty"`
	APIGroup        string `json:"apiGroup,omitempty"`
	APIVersion      string `json:"apiVersion,omitempty"`
	UID             string `json:"uid,omitempty"`
}

type NewResponseStatus struct {
	Code int `json:"code,omitempty"`
}

type NewAnnotationsK8s struct {
	AuthorizationK8SIoDecision string `json:"authorization.k8s.io/decision"`
	AuthorizationK8SIoReason   string `json:"authorization.k8s.io/reason"`
}

type NewOpenshiftAuditLog struct {
	AuditLogCommon
	OpenshiftAuditLevel string `json:"openshift_audit_level,omitempty"`
}

type NewK8sAuditLog struct {
	AuditLogCommon
	K8SAuditLevel string `json:"k8s_audit_level,omitempty"`
}

type AuditLog struct {
	Hostname                 string              `json:"hostname,omitempty"`
	AuditLinux               AuditLinux          `json:"audit.linux,omitempty"`
	Message                  string              `json:"message,omitempty"`
	PipelineMetadata         NewPipelinemetadata `json:"pipeline_metadata"`
	Timestamp                time.Time           `json:"@timestamp,omitempty"`
	Docker                   NewDocker           `json:"docker,omitempty"`
	LogType                  string              `json:"log_type,omitempty"`
	ViaqIndexName            string              `json:"viaq_index_name,omitempty"`
	ViaqMsgID                string              `json:"viaq_msg_id,omitempty"`
	Kubernetes               K8s                 `json:"K8s,omitempty"`
	Kind                     string              `json:"kind,omitempty"`
	APIVersion               string              `json:"apiVersion,omitempty"`
	Level                    string              `json:"level,omitempty"`
	AuditID                  string              `json:"auditID,omitempty"`
	Stage                    string              `json:"stage,omitempty"`
	RequestURI               string              `json:"requestURI,omitempty"`
	Verb                     string              `json:"verb,omitempty"`
	User                     NewUser             `json:"user,omitempty"`
	SourceIPs                []string            `json:"sourceIPs,omitempty"`
	UserAgent                string              `json:"userAgent,omitempty"`
	ObjectRef                NewObjectRef        `json:"objectRef,omitempty"`
	ResponseStatus           NewResponseStatus   `json:"responseStatus,omitempty"`
	RequestReceivedTimestamp time.Time           `json:"requestReceivedTimestamp,omitempty"`
	StageTimestamp           time.Time           `json:"stageTimestamp,omitempty"`
	Annotations              NewAnnotationsK8s   `json:"annotations,omitempty"`
	K8SAuditLevel            string              `json:"k8s_audit_level,omitempty"`
	OpenshiftAuditLevel      string              `json:"openshift_audit_level,omitempty"`
}

type AllLog struct {
	//ContainerLog             `json:",inline,omitempty"`
	STREAMID                 string            `json:"_STREAM_ID,omitempty"`
	SYSTEMDINVOCATIONID      string            `json:"_SYSTEMD_INVOCATION_ID,omitempty"`
	Systemd                  NewSystemd        `json:"systemd,omitempty"`
	AuditLinux               AuditLinux        `json:"audit.linux,omitempty"`
	Kind                     string            `json:"kind,omitempty"`
	APIVersion               string            `json:"apiVersion,omitempty"`
	AuditID                  string            `json:"auditID,omitempty"`
	Stage                    string            `json:"stage,omitempty"`
	RequestURI               string            `json:"requestURI,omitempty"`
	Verb                     string            `json:"verb,omitempty"`
	User                     NewUser           `json:"user,omitempty"`
	SourceIPs                []string          `json:"sourceIPs,omitempty"`
	UserAgent                string            `json:"userAgent,omitempty"`
	ObjectRef                NewObjectRef      `json:"objectRef,omitempty"`
	ResponseStatus           NewResponseStatus `json:"responseStatus,omitempty"`
	RequestReceivedTimestamp time.Time         `json:"requestReceivedTimestamp,omitempty"`
	StageTimestamp           time.Time         `json:"stageTimestamp,omitempty"`
	Annotations              NewAnnotationsK8s `json:"annotations,omitempty"`
	K8SAuditLevel            string            `json:"k8s_audit_level,omitempty"`
	OpenshiftAuditLevel      string            `json:"openshift_audit_level,omitempty"`
}

func ConvertLog(log types.ContainerLog) newLog {
	theNewLog := newLog{
		Timestamp:      changeTime(log.ViaQCommon.Timestamp),
		SeverityText:   log.ViaQCommon.Level,
		SeverityNumber: severityTextToNumber(log.ViaQCommon.Level),
		NewPipelinemetadata: NewPipelinemetadata{
			NewLogCollector: NewLogCollector{
				Ipaddr4:    log.ViaQCommon.PipelineMetadata.Collector.Ipaddr4,
				Name:       log.ViaQCommon.PipelineMetadata.Collector.Name,
				ReceivedAt: convertTimeToString(log.ViaQCommon.PipelineMetadata.Collector.ReceivedAt),
				Version:    log.ViaQCommon.PipelineMetadata.Collector.Version,
			},
		},
		Body: Body{
			Stringvalue: log.ViaQCommon.Message,
		},
		ViaqIndexName: log.ViaQCommon.ViaqIndexName,
		ViaqMsgID:     log.ViaQCommon.ViaqMsgID,
		NewOpenshiftMeta: NewOpenshiftMeta{
			ClusterID:       log.ViaQCommon.Openshift.ClusterID,
			OpenshiftLabels: log.ViaQCommon.Openshift.Labels,
			Sequence:        string(log.ViaQCommon.Openshift.Sequence),
		},
		NewInfraLog: NewInfraLog{
			NewDocker: NewDocker{
				ContainerID: log.Docker.ContainerID,
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

func convertTimeToString(oldTime time.Time) string {
	formattedTime := oldTime.Format("2006-01-02T15:04:05.999999999Z")
	return formattedTime
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
