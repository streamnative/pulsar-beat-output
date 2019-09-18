// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package budgets

// The type of a budget. It must be one of the following types:
//
// COST, USAGE, RI_UTILIZATION, or RI_COVERAGE.
type BudgetType string

// Enum values for BudgetType
const (
	BudgetTypeUsage         BudgetType = "USAGE"
	BudgetTypeCost          BudgetType = "COST"
	BudgetTypeRiUtilization BudgetType = "RI_UTILIZATION"
	BudgetTypeRiCoverage    BudgetType = "RI_COVERAGE"
)

func (enum BudgetType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum BudgetType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The comparison operator of a notification. Currently the service supports
// the following operators:
//
// GREATER_THAN, LESS_THAN, EQUAL_TO
type ComparisonOperator string

// Enum values for ComparisonOperator
const (
	ComparisonOperatorGreaterThan ComparisonOperator = "GREATER_THAN"
	ComparisonOperatorLessThan    ComparisonOperator = "LESS_THAN"
	ComparisonOperatorEqualTo     ComparisonOperator = "EQUAL_TO"
)

func (enum ComparisonOperator) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ComparisonOperator) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type NotificationState string

// Enum values for NotificationState
const (
	NotificationStateOk    NotificationState = "OK"
	NotificationStateAlarm NotificationState = "ALARM"
)

func (enum NotificationState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum NotificationState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The type of a notification. It must be ACTUAL or FORECASTED.
type NotificationType string

// Enum values for NotificationType
const (
	NotificationTypeActual     NotificationType = "ACTUAL"
	NotificationTypeForecasted NotificationType = "FORECASTED"
)

func (enum NotificationType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum NotificationType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The subscription type of the subscriber. It can be SMS or EMAIL.
type SubscriptionType string

// Enum values for SubscriptionType
const (
	SubscriptionTypeSns   SubscriptionType = "SNS"
	SubscriptionTypeEmail SubscriptionType = "EMAIL"
)

func (enum SubscriptionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SubscriptionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The type of threshold for a notification. It can be PERCENTAGE or ABSOLUTE_VALUE.
type ThresholdType string

// Enum values for ThresholdType
const (
	ThresholdTypePercentage    ThresholdType = "PERCENTAGE"
	ThresholdTypeAbsoluteValue ThresholdType = "ABSOLUTE_VALUE"
)

func (enum ThresholdType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ThresholdType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The time unit of the budget, such as MONTHLY or QUARTERLY.
type TimeUnit string

// Enum values for TimeUnit
const (
	TimeUnitDaily     TimeUnit = "DAILY"
	TimeUnitMonthly   TimeUnit = "MONTHLY"
	TimeUnitQuarterly TimeUnit = "QUARTERLY"
	TimeUnitAnnually  TimeUnit = "ANNUALLY"
)

func (enum TimeUnit) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TimeUnit) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
