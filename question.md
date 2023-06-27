# Translating Viaq logs to JSON

NOTE: the OTEL data model definitions are being moved from https://opentelemetry.io/docs to separate repos:
- https://github.com/open-telemetry/opentelemetry-specification
- https://github.com/open-telemetry/semantic-conventions
Make sure to refer to documents in the new repos.

References:
- [Example JSON frmat for logs](https://github.com/open-telemetry/opentelemetry-proto/blob/main/examples/logs.json)
- [OTEL Log Record Definition](https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/logs/data-model.md)
- [Examples of other translations](https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/logs/data-model-appendix.md)
- [Resource shostnameemantic conventions](https://github.com/open-telemetry/semantic-conventions/blob/main/specification/resource/semantic_conventions/README.md)


# Notes

- Translate @timestamp into numeric Unix nano time
  ``` go
	  timestamp, err := time.Parse(time.RFC3339Nano, "2022-10-20T15:11:30.764362932Z")
	  fmt.Println(timestamp.UnixNano(), err)
  ```
- Split container_image into name and tag  (on last :)


Note: Missing attributes in OTEL, added equivalent attributes.
- pod_ip
- pod_owner
- container_id
- namespace_id
- namespace_labels
- annotations
- labels

# For Later
- [Conventions for k8s Events](https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/logs/event-api.md)
- Consider resourceLogs batch encoding? Resource first then sequence of logs, mixing logs from many resources...



Right now stuff
- ignore eventdata, Need OpenShift data but no OTEL equivalent,
- decode annotations stuff into a map instead of a pre defined struct by using https://github.com/openshift/cluster-logging-operator/blob/master/test/helpers/types/types.go
- there's alot of non equavalient ones, so loko at viaq semantics and find the otel equivalent (DON'T WORRY ABOUT THE ORDER)

Things to add
- collector
- EventData (skip for now)
- Pipelinemetadata
- openshiftmeta
- appliecationlog (kinda?)
- infracontainerlog (kinda?)
- pipelinemetadata in viaqcommon
- viaqindexname in viaqcommon
- viaqmsgid in viaqcommon
- openshift in viaqcommon
- JournalLog
- T struct
- U struct
- systemd struct
- infralog
- linuxauditlog
- AuditLinux
- OVNAuditLog
- AuditLogCommon
- EventRouterLog
- User
- ObjectRef
- ResponseStatus
- Annotations
- OpenshiftAuditLog
- K8sAuditLog
- AuditLog
- AllLog



List of things missing in types.go
 - File
  - 					// NewNamespaceLabels: NewNamespaceLabels{
					// 	KubernetesIoMetadataName:                            log.Kubernetes.NamespaceLabels.KubernetesIoMetadataName,
					// 	OlmOperatorgroupUidD5ae8d2e99f34020998d9fc74c33faeb: log.Kubernetes.NamespaceLabels.OlmOperatorgroupUidD5ae8d2e99f34020998d9fc74c33faeb,
					// 	OpenshiftIoClusterMonitoring:                        log.Kubernetes.NamespaceLabels.OpenshiftIoClusterMonitoring,
					// 	PodSecurityKubernetesIoAudit:                        log.Kubernetes.NamespaceLabels.PodSecurityKubernetesIoAudit,
					// 	PodSecurityKubernetesIoEnforce:                      log.Kubernetes.NamespaceLabels.PodSecurityKubernetesIoEnforce,
					// 	PodSecurityKubernetesIoWarn:                         log.Kubernetes.NamespaceLabels.PodSecurityKubernetesIoWarn,
					// },
					



Questions:
    Is the change time function fine
    Ask about severityTextToNumber
    There's no File line in types.go
    Should the key be hidden if it doesn't exist (would have to make it a pointer then)
    Changed Sequence string `json:"sequence,omitempty"` from OptionalInt
    applicationlog
    infracontainerlog
    changed Timing in Linux audit log into a string, is that okay
    There's 2 different Annotations, New annotations and kuberenetes annotations
    Should I format AuditLogcomomn exactly like k8s
    containerlog in AllLog, should the entire format be the same exact as types.go or should it be like how it was before
    
    
    
    