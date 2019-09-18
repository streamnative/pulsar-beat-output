// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package batch

type ArrayJobDependency string

// Enum values for ArrayJobDependency
const (
	ArrayJobDependencyNToN       ArrayJobDependency = "N_TO_N"
	ArrayJobDependencySequential ArrayJobDependency = "SEQUENTIAL"
)

func (enum ArrayJobDependency) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ArrayJobDependency) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CEState string

// Enum values for CEState
const (
	CEStateEnabled  CEState = "ENABLED"
	CEStateDisabled CEState = "DISABLED"
)

func (enum CEState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CEState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CEStatus string

// Enum values for CEStatus
const (
	CEStatusCreating CEStatus = "CREATING"
	CEStatusUpdating CEStatus = "UPDATING"
	CEStatusDeleting CEStatus = "DELETING"
	CEStatusDeleted  CEStatus = "DELETED"
	CEStatusValid    CEStatus = "VALID"
	CEStatusInvalid  CEStatus = "INVALID"
)

func (enum CEStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CEStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CEType string

// Enum values for CEType
const (
	CETypeManaged   CEType = "MANAGED"
	CETypeUnmanaged CEType = "UNMANAGED"
)

func (enum CEType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CEType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CRType string

// Enum values for CRType
const (
	CRTypeEc2  CRType = "EC2"
	CRTypeSpot CRType = "SPOT"
)

func (enum CRType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CRType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DeviceCgroupPermission string

// Enum values for DeviceCgroupPermission
const (
	DeviceCgroupPermissionRead  DeviceCgroupPermission = "READ"
	DeviceCgroupPermissionWrite DeviceCgroupPermission = "WRITE"
	DeviceCgroupPermissionMknod DeviceCgroupPermission = "MKNOD"
)

func (enum DeviceCgroupPermission) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeviceCgroupPermission) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type JQState string

// Enum values for JQState
const (
	JQStateEnabled  JQState = "ENABLED"
	JQStateDisabled JQState = "DISABLED"
)

func (enum JQState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum JQState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type JQStatus string

// Enum values for JQStatus
const (
	JQStatusCreating JQStatus = "CREATING"
	JQStatusUpdating JQStatus = "UPDATING"
	JQStatusDeleting JQStatus = "DELETING"
	JQStatusDeleted  JQStatus = "DELETED"
	JQStatusValid    JQStatus = "VALID"
	JQStatusInvalid  JQStatus = "INVALID"
)

func (enum JQStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum JQStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type JobDefinitionType string

// Enum values for JobDefinitionType
const (
	JobDefinitionTypeContainer JobDefinitionType = "container"
	JobDefinitionTypeMultinode JobDefinitionType = "multinode"
)

func (enum JobDefinitionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum JobDefinitionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusSubmitted JobStatus = "SUBMITTED"
	JobStatusPending   JobStatus = "PENDING"
	JobStatusRunnable  JobStatus = "RUNNABLE"
	JobStatusStarting  JobStatus = "STARTING"
	JobStatusRunning   JobStatus = "RUNNING"
	JobStatusSucceeded JobStatus = "SUCCEEDED"
	JobStatusFailed    JobStatus = "FAILED"
)

func (enum JobStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum JobStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeGpu ResourceType = "GPU"
)

func (enum ResourceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResourceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
