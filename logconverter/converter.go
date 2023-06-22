package logconverter

import (
	"strconv"
	"strings"
	"time"

	"github.com/openshift/cluster-logging-operator/test/helpers/types"
)

// New json file
type newLog struct {
	Timestamp      string   `json:"timeUnixNano,omitempty"`
	SeverityText   string   `json:"severityText,omitempty"`
	SeverityNumber string   `json:"severityNumber,omitempty"`
	Body           Body     `json:"body,omitempty"`
	NewLogCollector NewLogCollector	`json:"collector,omitempty"`
	Resource       Resource `json:"resource,omitempty"`
}

type NewLogCollector struct {
	Ipaddr4 string `json:"ipaddr4,omitempty"`
	Name    string `json:"name,omitempty"`
	ReceivedAt string `json:"receivedAt,omitempty"`
	Version string `json:"version,omitempty"`
}

type Body struct {
	Stringvalue string `json:"stringValue,omitempty"`
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

// var ErrParse = errors.New("logs could not be parsed")

// // ContainerLog
// type ContainerLog struct {
// 	ViaQCommon `json:",inline,omitempty"`

// 	// +optional
// 	// +deprecated
// 	Docker Docker `json:"docker,omitempty"`

// 	// The Kubernetes-specific metadata
// 	Kubernetes Kubernetes `json:"kubernetes,omitempty"`

// 	// Original log entry as a structured object.
// 	//
// 	//Example:
// 	// `{"pid":21631,"ppid":21618,"worker":0,"message":"starting fluentd worker pid=21631 ppid=21618 worker=0"}`
// 	//
// 	// This field may be present if the forwarder was configured to parse structured JSON logs.
// 	// If the original log entry was a valid structured log, this field will contain an equivalent JSON structure.
// 	// Otherwise this field will be empty or absent, and the `message` field will contain the original log message.
// 	// The `structured` field includes the same sub-fields as the original log message.
// 	// +optional
// 	Structured map[string]interface{} `json:"structured,omitempty"`
// }

// type Docker struct {

// 	// ContainerID is the id of the container producing the log
// 	ContainerID string `json:"container_id"`
// }

// type Kubernetes struct {

// 	// Annotations associated with the Kubernetes pod
// 	// +optional
// 	Annotations map[string]string `json:"annotations,omitempty"`

// 	// ContainerName of the the pod container that produced the log
// 	ContainerName string `json:"container_name,omitempty"`

// 	//NamespaceName where the pod is deployed
// 	NamespaceName string `json:"namespace_name,omitempty"`

// 	// PodName is the name of the pod
// 	PodName string `json:"pod_name,omitempty"`

// 	// +optional
// 	ContainerID string `json:"container_id,omitempty"`
// 	// +optional
// 	ContainerImage string `json:"container_image,omitempty"`
// 	// +optional
// 	ContainerImageID string `json:"container_image_id,omitempty"`

// 	//PodID is the unique uuid of the pod
// 	// +optional
// 	PodID string `json:"pod_id,omitempty"`

// 	// +docgen:ignore
// 	// +optional
// 	PodIP string `json:"pod_ip,omitempty"`

// 	//Host is the kubernetes node name that hosts the pod
// 	// +optional
// 	Host string `json:"host,omitempty"`

// 	//MasterURL is the url to the apiserver
// 	// +deprecated
// 	MasterURL string `json:"master_url,omitempty"`

// 	//NamespaceID is the unique uuid of the namespace
// 	// +optional
// 	NamespaceID string `json:"namespace_id,omitempty"`

// 	//FlatLabels is an array of the pod labels joined as key=value
// 	// +optional
// 	// +deprecated
// 	// +docgen:type=array
// 	FlatLabels []string `json:"flat_labels,omitempty"`

// 	// Labels present on the Pod at time the log was generated
// 	// +optional
// 	Labels map[string]string `json:"labels,omitempty"`

// 	// +docgen:ignore
// 	// +deprecated
// 	// +optional
// 	OrphanedNamespace string `json:"orphaned_namespace,omitempty"`

// 	// NamespaceLabels are the labels present on the pod namespace
// 	// +optional
// 	NamespaceLabels map[string]string `json:"namespace_labels,omitempty"`
// }

// type Collector struct {
// 	//Ipaddr4 is the ipV4 address of the collector
// 	//+optional
// 	Ipaddr4 string `json:"ipaddr4,omitempty"`

// 	//+deprecated
// 	Inputname string `json:"inputname,omitempty"`

// 	//Name is the implementation of the collector agent
// 	Name string `json:"name,omitempty"`

// 	//ReceivedAt the time the collector received the log entry
// 	ReceivedAt time.Time `json:"received_at,omitempty"`

// 	//Version is collector version information
// 	Version string `json:"version,omitempty"`
// }

// // EventData encodes an eventrouter event and previous event, with a verb for
// // whether the event is created or updated.
// type EventData struct {
// 	Verb     string    `json:"verb"`
// 	Event    *v1.Event `json:"event"`
// 	OldEvent *v1.Event `json:"old_event,omitempty"`
// }

// type PipelineMetadata struct {

// 	//Collector metadata
// 	Collector Collector `json:"collector,omitempty"`
// }

// type OpenshiftMeta struct {

// 	//ClusterID is the unique id of the cluster where the workload is deployed
// 	ClusterID string `json:"cluster_id,omitempty"`

// 	//Labels is a set of common, static labels that were spec'd for log forwarding
// 	//to be sent with the log Records
// 	//+optional
// 	Labels map[string]string `json:"labels,omitempty"`

// 	//Sequence is increasing id used in conjunction with the timestamp to estblish a linear timeline
// 	//of log records.  This was added as a workaround for logstores that do not have nano-second precision.
// 	Sequence OptionalInt `json:"sequence,omitempty"`
// }

// // Application Logs are container logs from all namespaces except "openshift" and "openshift-*" namespaces
// type ApplicationLog ContainerLog

// // Infrastructure logs are
// // - Journal logs
// // - logs from "openshift" and "openshift-*" namespaces

// // InfraContainerLog
// // InfraContainerLog logs are container logs from "openshift" and "openshift-*" namespaces
// type InfraContainerLog ContainerLog

// type ViaQCommon struct {

// 	// A UTC value that marks when the log payload was created.
// 	//
// 	// If the creation time is not known when the log payload was first collected. The “@” prefix denotes a
// 	// field that is reserved for a particular use.
// 	//
// 	// format:
// 	//
// 	// * yyyy-MM-dd HH:mm:ss,SSSZ
// 	// * yyyy-MM-dd'T'HH:mm:ss.SSSSSSZ
// 	// * yyyy-MM-dd'T'HH:mm:ssZ
// 	// * dateOptionalTime
// 	//
// 	// example: `2015-01-24 14:06:05.071000000 Z`
// 	Timestamp time.Time `json:"@timestamp,omitempty"`

// 	// Original log entry text, UTF-8 encoded
// 	//
// 	// This field may be absent or empty if a non-empty `structured` field is present.
// 	// See the description of `structured` for additional details.
// 	// +optional
// 	Message string `json:"message,omitempty"`

// 	// The normalized log level
// 	//
// 	// The logging level from various sources, including `rsyslog(severitytext property)`, python's logging module, and others.
// 	//        The following values come from link:http://sourceware.org/git/?p=glibc.git;a=blob;f=misc/sys/syslog.h;h=ee01478c4b19a954426a96448577c5a76e6647c0;hb=HEAD#l74[`syslog.h`], and are preceded by their http://sourceware.org/git/?p=glibc.git;a=blob;f=misc/sys/syslog.h;h=ee01478c4b19a954426a96448577c5a76e6647c0;hb=HEAD#l51[numeric equivalents]:
// 	//
// 	//        * `0` = `emerg`, system is unusable.
// 	//        * `1` = `alert`, action must be taken immediately.
// 	//        * `2` = `crit`, critical conditions.
// 	//        * `3` = `err`, error conditions.
// 	//        * `4` = `warn`, warning conditions.
// 	//        * `5` = `notice`, normal but significant condition.
// 	//        * `6` = `info`, informational.
// 	//        * `7` = `debug`, debug-level messages.
// 	//        The two following values are not part of `syslog.h` but are widely used:
// 	//        * `8` = `trace`, trace-level messages, which are more verbose than `debug` messages.
// 	//        * `9` = `unknown`, when the logging system gets a value it doesn't recognize.
// 	//        Map the log levels or priorities of other logging systems to their nearest match in the preceding list. For example, from link:https://docs.python.org/2.7/library/logging.html#logging-levels[python logging], you can match `CRITICAL` with `crit`, `ERROR` with `err`, and so on.
// 	Level string `json:"level,omitempty"`

// 	// The name of the host where this log message originated. In a Kubernetes cluster, this is the same as `kubernetes.host`.
// 	Hostname string `json:"hostname,omitempty"`

// 	// Metadata related to ViaQ log collection pipeline. Everything about log collector, normalizers, mappings goes here.
// 	// Data in this subgroup is forwarded for troubleshooting and tracing purposes.
// 	// +deprecated
// 	PipelineMetadata PipelineMetadata `json:"pipeline_metadata,omitempty"`

// 	//The source type of the log. The `log_type` field may contain one of these strings, or may have additional dot-separated components, for example "infrastructure.container" or "infrastructure.node".
// 	//
// 	// * "application": Container logs generated by user applications running in the cluster, except infrastructure containers.
// 	// * "infrastructure": Node logs (such as syslog or journal logs), and container logs from pods in the openshift*, kube*, or default projects.
// 	// * "audit":
// 	// ** Node logs from auditd (/var/log/audit/audit.log)
// 	// ** Kubernetes and OpenShift apiservers audit logs.
// 	// ** OVN audit logs
// 	//
// 	LogType string `json:"log_type,omitempty"`

// 	// ViaqIndexName used with Elasticsearch 6.x and later, this is a name of a write index alias (e.g. app-write).
// 	//
// 	// The value depends on the log type of this message. Detailed documentation is found at https://github.com/openshift/enhancements/blob/master/enhancements/cluster-logging/cluster-logging-es-rollover-data-design.md#data-model.
// 	// +optional
// 	ViaqIndexName string `json:"viaq_index_name,omitempty"`

// 	// ViaqMessageId is a unique ID assigned to each message. The format is not specified.
// 	//
// 	//It may be a UUID or a Base64 (e.g. 82f13a8e-882a-4344-b103-f0a6f30fd218),
// 	// or some other ASCII value and is used as the `_id` of the document when sending to Elasticsearch. The intended use of this field is that if you use another
// 	// logging store or application other than Elasticsearch, but you still need to correlate data with the data stored
// 	// in Elasticsearch, this field will give you the exact document corresponding to the record.
// 	// +optional
// 	ViaqMsgID string `json:"viaq_msg_id,omitempty"`

// 	// Openshift specific metadata
// 	Openshift OpenshiftMeta `json:"openshift,omitempty"`
// }

// // JournalLog is linux journal logs
// type JournalLog struct {
// 	ViaQCommon          `json:",inline,omitempty"`
// 	STREAMID            string  `json:"_STREAM_ID,omitempty"`
// 	SYSTEMDINVOCATIONID string  `json:"_SYSTEMD_INVOCATION_ID,omitempty"`
// 	Systemd             Systemd `json:"systemd,omitempty"`
// }

// type T struct {
// 	BOOTID              string `json:"BOOT_ID,omitempty"`
// 	CAPEFFECTIVE        string `json:"CAP_EFFECTIVE,omitempty"`
// 	CMDLINE             string `json:"CMDLINE,omitempty"`
// 	COMM                string `json:"COMM,omitempty"`
// 	EXE                 string `json:"EXE,omitempty"`
// 	GID                 string `json:"GID,omitempty"`
// 	MACHINEID           string `json:"MACHINE_ID,omitempty"`
// 	PID                 string `json:"PID,omitempty"`
// 	SELINUXCONTEXT      string `json:"SELINUX_CONTEXT,omitempty"`
// 	STREAMID            string `json:"STREAM_ID,omitempty"`
// 	SYSTEMDCGROUP       string `json:"SYSTEMD_CGROUP,omitempty"`
// 	SYSTEMDINVOCATIONID string `json:"SYSTEMD_INVOCATION_ID,omitempty"`
// 	SYSTEMDSLICE        string `json:"SYSTEMD_SLICE,omitempty"`
// 	SYSTEMDUNIT         string `json:"SYSTEMD_UNIT,omitempty"`
// 	TRANSPORT           string `json:"TRANSPORT,omitempty"`
// 	UID                 string `json:"UID,omitempty"`
// }

// type U struct {
// 	SYSLOGIDENTIFIER string `json:"SYSLOG_IDENTIFIER,omitempty"`
// }

// type Systemd struct {
// 	T T `json:"t,omitempty"`
// 	U U `json:"u,omitempty"`
// }

// // InfraLog is union of JournalLog and InfraContainerLog
// type InfraLog struct {
// 	Docker              Docker           `json:"docker,omitempty"`
// 	Kubernetes          Kubernetes       `json:"kubernetes,omitempty"`
// 	Message             string           `json:"message,omitempty"`
// 	Level               string           `json:"level,omitempty"`
// 	Hostname            string           `json:"hostname,omitempty"`
// 	PipelineMetadata    PipelineMetadata `json:"pipeline_metadata,omitempty"`
// 	Timestamp           time.Time        `json:"@timestamp,omitempty"`
// 	LogType             string           `json:"log_type,omitempty"`
// 	ViaqIndexName       string           `json:"viaq_index_name,omitempty"`
// 	ViaqMsgID           string           `json:"viaq_msg_id,omitempty"`
// 	STREAMID            string           `json:"_STREAM_ID,omitempty"`
// 	SYSTEMDINVOCATIONID string           `json:"_SYSTEMD_INVOCATION_ID,omitempty"`
// 	Systemd             Systemd          `json:"systemd,omitempty"`
// 	OpenshiftLabels     OpenshiftMeta    `json:"openshift,omitempty"`
// }

// /*
// Audit logs are
//  - Audit logs generated by linux
//  - Audit logs generated by kubernetes
//  - Audit logs generated by openshift
//  - Audit logs generated by Openshift virtual network
// */

// // LinuxAuditLog is generated by linux operating system
// type LinuxAuditLog struct {
// 	Hostname         string           `json:"hostname"`
// 	AuditLinux       AuditLinux       `json:"audit.linux"`
// 	Message          string           `json:"message,omitempty"`
// 	PipelineMetadata PipelineMetadata `json:"pipeline_metadata"`
// 	Timestamp        time.Time        `json:"@timestamp"`
// 	LogType          string           `json:"log_type,omitempty"`
// 	ViaqIndexName    string           `json:"viaq_index_name"`
// 	ViaqMsgID        string           `json:"viaq_msg_id"`
// 	Kubernetes       Kubernetes       `json:"kubernetes"`
// 	Openshift        OpenshiftMeta    `json:"openshift"`
// 	Timing           `json:",inline"`
// 	Level            string `json:"level,omitempty"`
// }

// type AuditLinux struct {
// 	Type     string `json:"type,omitempty"`
// 	RecordID string `json:"record_id,omitempty"`
// }

// // OVN Audit log
// type OVNAuditLog struct {
// 	Hostname         string           `json:"hostname"`
// 	Message          string           `json:"message,omitempty"`
// 	PipelineMetadata PipelineMetadata `json:"pipeline_metadata"`
// 	Timestamp        time.Time        `json:"@timestamp"`
// 	LogType          string           `json:"log_type,omitempty"`
// 	ViaqIndexName    string           `json:"viaq_index_name"`
// 	ViaqMsgID        string           `json:"viaq_msg_id"`
// 	Kubernetes       Kubernetes       `json:"kubernetes"`
// 	Openshift        OpenshiftMeta    `json:"openshift"`
// 	Level            string           `json:"level,omitempty"`
// }

// // AuditLogCommon is common to k8s and openshift auditlogs
// type AuditLogCommon struct {
// 	Kind                     string           `json:"kind,omitempty"`
// 	APIVersion               string           `json:"apiVersion,omitempty"`
// 	Level                    string           `json:"level,omitempty"`
// 	AuditID                  string           `json:"auditID,omitempty"`
// 	Stage                    string           `json:"stage,omitempty"`
// 	RequestURI               string           `json:"requestURI,omitempty"`
// 	Verb                     string           `json:"verb,omitempty"`
// 	User                     User             `json:"user,omitempty"`
// 	SourceIPs                []string         `json:"sourceIPs,omitempty"`
// 	UserAgent                string           `json:"userAgent,omitempty"`
// 	ObjectRef                ObjectRef        `json:"objectRef,omitempty"`
// 	ResponseStatus           ResponseStatus   `json:"responseStatus,omitempty"`
// 	RequestReceivedTimestamp time.Time        `json:"requestReceivedTimestamp,omitempty"`
// 	StageTimestamp           time.Time        `json:"stageTimestamp,omitempty"`
// 	Annotations              Annotations      `json:"annotations,omitempty"`
// 	Message                  interface{}      `json:"message,omitempty"`
// 	Hostname                 string           `json:"hostname,omitempty"`
// 	PipelineMetadata         PipelineMetadata `json:"pipeline_metadata,omitempty"`
// 	Timestamp                time.Time        `json:"@timestamp,omitempty"`
// 	LogType                  string           `json:"log_type,omitempty"`
// 	ViaqIndexName            string           `json:"viaq_index_name,omitempty"`
// 	ViaqMsgID                string           `json:"viaq_msg_id,omitempty"`
// 	Kubernetes               Kubernetes       `json:"kubernetes,omitempty"`
// 	OpenshiftLabels          OpenshiftMeta    `json:"openshift,omitempty"`
// 	Timing                   `json:",inline"`
// }

// // EventRouterLog is generated by event router
// type EventRouterLog struct {
// 	Docker           Docker           `json:"docker"`
// 	Kubernetes       Kubernetes       `json:"kubernetes"`
// 	Message          string           `json:"message,omitempty"`
// 	Level            string           `json:"level"`
// 	Hostname         string           `json:"hostname,omitempty"`
// 	PipelineMetadata PipelineMetadata `json:"pipeline_metadata"`
// 	Timestamp        time.Time        `json:"@timestamp"`
// 	LogType          string           `json:"log_type,omitempty"`
// 	ViaqIndexName    string           `json:"viaq_index_name"`
// 	ViaqMsgID        string           `json:"viaq_msg_id"`
// 	OpenshiftLabels  OpenshiftMeta    `json:"openshift"`
// 	Timing           `json:",inline"`
// }

// type User struct {
// 	Username string   `json:"username,omitempty"`
// 	UID      string   `json:"uid,omitempty"`
// 	Groups   []string `json:"groups,omitempty"`
// }
// type ObjectRef struct {
// 	Resource        string `json:"resource,omitempty"`
// 	ResourceVersion string `json:"resourceVersion,omitempty"`
// 	Name            string `json:"name,omitempty"`
// 	Namespace       string `json:"namespace,omitempty"`
// 	APIGroup        string `json:"apiGroup,omitempty"`
// 	APIVersion      string `json:"apiVersion,omitempty"`
// 	UID             string `json:"uid,omitempty"`
// }
// type ResponseStatus struct {
// 	Code int `json:"code,omitempty"`
// }
// type Annotations struct {
// 	AuthorizationK8SIoDecision string `json:"authorization.k8s.io/decision"`
// 	AuthorizationK8SIoReason   string `json:"authorization.k8s.io/reason"`
// }

// // OpenshiftAuditLog is audit log generated by openshift-apiserver
// type OpenshiftAuditLog struct {
// 	AuditLogCommon
// 	OpenshiftAuditLevel string `json:"openshift_audit_level,omitempty"`
// }

// // K8sAuditLog is audit logs generated by kube-apiserver
// type K8sAuditLog struct {
// 	AuditLogCommon
// 	K8SAuditLevel string `json:"k8s_audit_level,omitempty"`
// }

// // AuditLog is a union of LinuxAudit, K8sAudit, OpenshiftAudit logs
// type AuditLog struct {
// 	Hostname                 string           `json:"hostname,omitempty"`
// 	AuditLinux               AuditLinux       `json:"audit.linux,omitempty"`
// 	Message                  string           `json:"message,omitempty"`
// 	PipelineMetadata         PipelineMetadata `json:"pipeline_metadata"`
// 	Timestamp                time.Time        `json:"@timestamp,omitempty"`
// 	Docker                   Docker           `json:"docker,omitempty"`
// 	LogType                  string           `json:"log_type,omitempty"`
// 	ViaqIndexName            string           `json:"viaq_index_name,omitempty"`
// 	ViaqMsgID                string           `json:"viaq_msg_id,omitempty"`
// 	Kubernetes               Kubernetes       `json:"kubernetes,omitempty"`
// 	Kind                     string           `json:"kind,omitempty"`
// 	APIVersion               string           `json:"apiVersion,omitempty"`
// 	Level                    string           `json:"level,omitempty"`
// 	AuditID                  string           `json:"auditID,omitempty"`
// 	Stage                    string           `json:"stage,omitempty"`
// 	RequestURI               string           `json:"requestURI,omitempty"`
// 	Verb                     string           `json:"verb,omitempty"`
// 	User                     User             `json:"user,omitempty"`
// 	SourceIPs                []string         `json:"sourceIPs,omitempty"`
// 	UserAgent                string           `json:"userAgent,omitempty"`
// 	ObjectRef                ObjectRef        `json:"objectRef,omitempty"`
// 	ResponseStatus           ResponseStatus   `json:"responseStatus,omitempty"`
// 	RequestReceivedTimestamp time.Time        `json:"requestReceivedTimestamp,omitempty"`
// 	StageTimestamp           time.Time        `json:"stageTimestamp,omitempty"`
// 	Annotations              Annotations      `json:"annotations,omitempty"`
// 	K8SAuditLevel            string           `json:"k8s_audit_level,omitempty"`
// 	OpenshiftAuditLevel      string           `json:"openshift_audit_level,omitempty"`
// }

// // AllLog is a union of all log types
// type AllLog struct {
// 	ContainerLog             `json:",inline,omitempty"`
// 	STREAMID                 string         `json:"_STREAM_ID,omitempty"`
// 	SYSTEMDINVOCATIONID      string         `json:"_SYSTEMD_INVOCATION_ID,omitempty"`
// 	Systemd                  Systemd        `json:"systemd,omitempty"`
// 	AuditLinux               AuditLinux     `json:"audit.linux,omitempty"`
// 	Kind                     string         `json:"kind,omitempty"`
// 	APIVersion               string         `json:"apiVersion,omitempty"`
// 	AuditID                  string         `json:"auditID,omitempty"`
// 	Stage                    string         `json:"stage,omitempty"`
// 	RequestURI               string         `json:"requestURI,omitempty"`
// 	Verb                     string         `json:"verb,omitempty"`
// 	User                     User           `json:"user,omitempty"`
// 	SourceIPs                []string       `json:"sourceIPs,omitempty"`
// 	UserAgent                string         `json:"userAgent,omitempty"`
// 	ObjectRef                ObjectRef      `json:"objectRef,omitempty"`
// 	ResponseStatus           ResponseStatus `json:"responseStatus,omitempty"`
// 	RequestReceivedTimestamp time.Time      `json:"requestReceivedTimestamp,omitempty"`
// 	StageTimestamp           time.Time      `json:"stageTimestamp,omitempty"`
// 	Annotations              Annotations    `json:"annotations,omitempty"`
// 	K8SAuditLevel            string         `json:"k8s_audit_level,omitempty"`
// 	OpenshiftAuditLevel      string         `json:"openshift_audit_level,omitempty"`
// }