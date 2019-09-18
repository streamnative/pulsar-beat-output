// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

type Compute string

// Enum values for Compute
const (
	ComputeValue       Compute = "VALUE"
	ComputeStandard    Compute = "STANDARD"
	ComputePerformance Compute = "PERFORMANCE"
	ComputePower       Compute = "POWER"
	ComputeGraphics    Compute = "GRAPHICS"
	ComputePowerpro    Compute = "POWERPRO"
	ComputeGraphicspro Compute = "GRAPHICSPRO"
)

func (enum Compute) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Compute) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ConnectionState string

// Enum values for ConnectionState
const (
	ConnectionStateConnected    ConnectionState = "CONNECTED"
	ConnectionStateDisconnected ConnectionState = "DISCONNECTED"
	ConnectionStateUnknown      ConnectionState = "UNKNOWN"
)

func (enum ConnectionState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ConnectionState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DedicatedTenancyModificationStateEnum string

// Enum values for DedicatedTenancyModificationStateEnum
const (
	DedicatedTenancyModificationStateEnumPending   DedicatedTenancyModificationStateEnum = "PENDING"
	DedicatedTenancyModificationStateEnumCompleted DedicatedTenancyModificationStateEnum = "COMPLETED"
	DedicatedTenancyModificationStateEnumFailed    DedicatedTenancyModificationStateEnum = "FAILED"
)

func (enum DedicatedTenancyModificationStateEnum) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DedicatedTenancyModificationStateEnum) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DedicatedTenancySupportEnum string

// Enum values for DedicatedTenancySupportEnum
const (
	DedicatedTenancySupportEnumEnabled DedicatedTenancySupportEnum = "ENABLED"
)

func (enum DedicatedTenancySupportEnum) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DedicatedTenancySupportEnum) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DedicatedTenancySupportResultEnum string

// Enum values for DedicatedTenancySupportResultEnum
const (
	DedicatedTenancySupportResultEnumEnabled  DedicatedTenancySupportResultEnum = "ENABLED"
	DedicatedTenancySupportResultEnumDisabled DedicatedTenancySupportResultEnum = "DISABLED"
)

func (enum DedicatedTenancySupportResultEnum) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DedicatedTenancySupportResultEnum) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ModificationResourceEnum string

// Enum values for ModificationResourceEnum
const (
	ModificationResourceEnumRootVolume  ModificationResourceEnum = "ROOT_VOLUME"
	ModificationResourceEnumUserVolume  ModificationResourceEnum = "USER_VOLUME"
	ModificationResourceEnumComputeType ModificationResourceEnum = "COMPUTE_TYPE"
)

func (enum ModificationResourceEnum) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ModificationResourceEnum) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ModificationStateEnum string

// Enum values for ModificationStateEnum
const (
	ModificationStateEnumUpdateInitiated  ModificationStateEnum = "UPDATE_INITIATED"
	ModificationStateEnumUpdateInProgress ModificationStateEnum = "UPDATE_IN_PROGRESS"
)

func (enum ModificationStateEnum) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ModificationStateEnum) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OperatingSystemType string

// Enum values for OperatingSystemType
const (
	OperatingSystemTypeWindows OperatingSystemType = "WINDOWS"
	OperatingSystemTypeLinux   OperatingSystemType = "LINUX"
)

func (enum OperatingSystemType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OperatingSystemType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ReconnectEnum string

// Enum values for ReconnectEnum
const (
	ReconnectEnumEnabled  ReconnectEnum = "ENABLED"
	ReconnectEnumDisabled ReconnectEnum = "DISABLED"
)

func (enum ReconnectEnum) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ReconnectEnum) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RunningMode string

// Enum values for RunningMode
const (
	RunningModeAutoStop RunningMode = "AUTO_STOP"
	RunningModeAlwaysOn RunningMode = "ALWAYS_ON"
)

func (enum RunningMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RunningMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TargetWorkspaceState string

// Enum values for TargetWorkspaceState
const (
	TargetWorkspaceStateAvailable        TargetWorkspaceState = "AVAILABLE"
	TargetWorkspaceStateAdminMaintenance TargetWorkspaceState = "ADMIN_MAINTENANCE"
)

func (enum TargetWorkspaceState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TargetWorkspaceState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type WorkspaceDirectoryState string

// Enum values for WorkspaceDirectoryState
const (
	WorkspaceDirectoryStateRegistering   WorkspaceDirectoryState = "REGISTERING"
	WorkspaceDirectoryStateRegistered    WorkspaceDirectoryState = "REGISTERED"
	WorkspaceDirectoryStateDeregistering WorkspaceDirectoryState = "DEREGISTERING"
	WorkspaceDirectoryStateDeregistered  WorkspaceDirectoryState = "DEREGISTERED"
	WorkspaceDirectoryStateError         WorkspaceDirectoryState = "ERROR"
)

func (enum WorkspaceDirectoryState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum WorkspaceDirectoryState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type WorkspaceDirectoryType string

// Enum values for WorkspaceDirectoryType
const (
	WorkspaceDirectoryTypeSimpleAd    WorkspaceDirectoryType = "SIMPLE_AD"
	WorkspaceDirectoryTypeAdConnector WorkspaceDirectoryType = "AD_CONNECTOR"
)

func (enum WorkspaceDirectoryType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum WorkspaceDirectoryType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type WorkspaceImageIngestionProcess string

// Enum values for WorkspaceImageIngestionProcess
const (
	WorkspaceImageIngestionProcessByolRegular     WorkspaceImageIngestionProcess = "BYOL_REGULAR"
	WorkspaceImageIngestionProcessByolGraphics    WorkspaceImageIngestionProcess = "BYOL_GRAPHICS"
	WorkspaceImageIngestionProcessByolGraphicspro WorkspaceImageIngestionProcess = "BYOL_GRAPHICSPRO"
)

func (enum WorkspaceImageIngestionProcess) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum WorkspaceImageIngestionProcess) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type WorkspaceImageRequiredTenancy string

// Enum values for WorkspaceImageRequiredTenancy
const (
	WorkspaceImageRequiredTenancyDefault   WorkspaceImageRequiredTenancy = "DEFAULT"
	WorkspaceImageRequiredTenancyDedicated WorkspaceImageRequiredTenancy = "DEDICATED"
)

func (enum WorkspaceImageRequiredTenancy) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum WorkspaceImageRequiredTenancy) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type WorkspaceImageState string

// Enum values for WorkspaceImageState
const (
	WorkspaceImageStateAvailable WorkspaceImageState = "AVAILABLE"
	WorkspaceImageStatePending   WorkspaceImageState = "PENDING"
	WorkspaceImageStateError     WorkspaceImageState = "ERROR"
)

func (enum WorkspaceImageState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum WorkspaceImageState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type WorkspaceState string

// Enum values for WorkspaceState
const (
	WorkspaceStatePending          WorkspaceState = "PENDING"
	WorkspaceStateAvailable        WorkspaceState = "AVAILABLE"
	WorkspaceStateImpaired         WorkspaceState = "IMPAIRED"
	WorkspaceStateUnhealthy        WorkspaceState = "UNHEALTHY"
	WorkspaceStateRebooting        WorkspaceState = "REBOOTING"
	WorkspaceStateStarting         WorkspaceState = "STARTING"
	WorkspaceStateRebuilding       WorkspaceState = "REBUILDING"
	WorkspaceStateMaintenance      WorkspaceState = "MAINTENANCE"
	WorkspaceStateAdminMaintenance WorkspaceState = "ADMIN_MAINTENANCE"
	WorkspaceStateTerminating      WorkspaceState = "TERMINATING"
	WorkspaceStateTerminated       WorkspaceState = "TERMINATED"
	WorkspaceStateSuspended        WorkspaceState = "SUSPENDED"
	WorkspaceStateUpdating         WorkspaceState = "UPDATING"
	WorkspaceStateStopping         WorkspaceState = "STOPPING"
	WorkspaceStateStopped          WorkspaceState = "STOPPED"
	WorkspaceStateError            WorkspaceState = "ERROR"
)

func (enum WorkspaceState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum WorkspaceState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
