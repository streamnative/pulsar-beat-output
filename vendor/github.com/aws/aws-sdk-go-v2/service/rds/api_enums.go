// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

type ActivityStreamMode string

// Enum values for ActivityStreamMode
const (
	ActivityStreamModeSync  ActivityStreamMode = "sync"
	ActivityStreamModeAsync ActivityStreamMode = "async"
)

func (enum ActivityStreamMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ActivityStreamMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ActivityStreamStatus string

// Enum values for ActivityStreamStatus
const (
	ActivityStreamStatusStopped  ActivityStreamStatus = "stopped"
	ActivityStreamStatusStarting ActivityStreamStatus = "starting"
	ActivityStreamStatusStarted  ActivityStreamStatus = "started"
	ActivityStreamStatusStopping ActivityStreamStatus = "stopping"
)

func (enum ActivityStreamStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ActivityStreamStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ApplyMethod string

// Enum values for ApplyMethod
const (
	ApplyMethodImmediate     ApplyMethod = "immediate"
	ApplyMethodPendingReboot ApplyMethod = "pending-reboot"
)

func (enum ApplyMethod) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ApplyMethod) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SourceType string

// Enum values for SourceType
const (
	SourceTypeDbInstance        SourceType = "db-instance"
	SourceTypeDbParameterGroup  SourceType = "db-parameter-group"
	SourceTypeDbSecurityGroup   SourceType = "db-security-group"
	SourceTypeDbSnapshot        SourceType = "db-snapshot"
	SourceTypeDbCluster         SourceType = "db-cluster"
	SourceTypeDbClusterSnapshot SourceType = "db-cluster-snapshot"
)

func (enum SourceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SourceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
