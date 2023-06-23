# otel-translation

A small Go Project to parse and translate ViaQ JSON log files into OTEL YAML files.

## Todo
- [x] collector (made up, couldn't find equavalient)
- [ ] EventData (skip for now)
- [x] Pipelinemetadata (Similar to collector, so no equivalent)
- [x] openshiftmeta (couldn't find equavalient)
- [ ] appliecationlog (kinda?)
- [ ] infracontainerlog (kinda?)
- [x] pipelinemetadata in viaqcommon
- [x] viaqindexname in viaqcommon (couldn't find equivelent, but found it in amazon sdk called index_name)
- [x] viaqmsgid in viaqcommon
- [x] openshift in viaqcommon
- [x] JournalLog (couldn't find equivelent)
- [x] T struct (couldn't find equivelent)
- [x] U struct (couldn't find equivelent)
- [ ] systemd struct
- [ ] infralog
- [ ] linuxauditlog
- [ ] AuditLinux
- [ ] OVNAuditLog
- [ ] AuditLogCommon
- [ ] EventRouterLog
- [ ] User
- [ ] ObjectRef
- [ ] ResponseStatus
- [ ] Annotations
- [ ] OpenshiftAuditLog
- [ ] K8sAuditLog
- [ ] AuditLog
- [ ] AllLog