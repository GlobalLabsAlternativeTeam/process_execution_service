// internal/domain/converter.go

package domain

import (
	"time"

	process_execution_service "server/proto"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TreatmentToGRPC(t *Treatment) *process_execution_service.Treatment {
	return &process_execution_service.Treatment{
		TreatmentId:     t.TreatmentdID,
		DoctorId:        t.DoctorID,
		PatientId:       t.PatientID,
		Status:          convertTreatmentStatus(t.Status),
		PatternInstance: PatternInstanceToGRPC(&t.PatternInstance),
		StartedAt:       convertTimestampFromString(t.StartedAt),
		FinishedAt:      convertTimestampFromString(t.FinishedAt),
		DeletedAt:       convertTimestampFromString(t.DeletedAt),
	}
}

func PatternInstanceToGRPC(p *PatternInstance) *process_execution_service.PatternInstance {
	return &process_execution_service.PatternInstance{
		InstanceId:  p.SchemaInstanceID,
		Status:      convertPatternInstanceStatus(p.PatternInstanceStatus),
		PatternId:   p.SchemaID,
		AuthorId:    p.AuthorID,
		PatternName: p.SchemaName,
		CreatedAt:   convertTimestampFromTime(p.CreatedAt),
		UpdatedAt:   convertTimestampFromTime(p.UpdatedAt),
		DeletedAt:   convertTimestampFromTime(p.DeletedAt),
		Tasks:       TasksToGRPC(p.Tasks),
	}
}

func TasksToGRPC(tasks []Task) []*process_execution_service.Task {
	var grpcTasks []*process_execution_service.Task
	for _, t := range tasks {
		grpcTask := TaskToGRPC(&t)
		grpcTasks = append(grpcTasks, grpcTask)
	}
	return grpcTasks
}

func TaskToGRPC(t *Task) *process_execution_service.Task {
	var blockedBy []int64
	for _, v := range t.BlockedBy {
		if intValue, ok := v.(int); ok {
			blockedBy = append(blockedBy, int64(intValue))
		} else if int64Value, ok := v.(int64); ok {
			blockedBy = append(blockedBy, int64Value)
		}
		// Add more type conversions if needed for other cases
	}
	return &process_execution_service.Task{
		Id:          int64(t.ID),
		Level:       int32(t.Level),
		Name:        t.Name,
		Status:      convertTaskStatus(t.Status),
		BlockedBy:   blockedBy,
		Responsible: t.Responsible,
		TimeLimit:   t.TimeLimit,
		Children:    TasksToGRPC(t.Children),
		Comment:     wrapperspb.String(t.Comment.Value),
	}
}

func convertPatternInstanceStatus(status string) process_execution_service.PatternInstanceStatus {
	switch status {
	case "NOT_STARTED":
		return process_execution_service.PatternInstanceStatus_PATTERN_INSTANCE_STATUS_NOT_STARTED
	case "RUNNING":
		return process_execution_service.PatternInstanceStatus_PATTERN_INSTANCE_STATUS_STATUS_RUNNING
	case "BLOCKED":
		return process_execution_service.PatternInstanceStatus_PATTERN_INSTANCE_STATUS_STATUS_BLOCKED
	case "COMPLETED":
		return process_execution_service.PatternInstanceStatus_PATTERN_INSTANCE_STATUS_STATUS_COMPLETED
	case "CANCELLED":
		return process_execution_service.PatternInstanceStatus_PATTERN_INSTANCE_STATUS_STATUS_CANCELLED
	default:
		return process_execution_service.PatternInstanceStatus_PATTERN_INSTANCE_STATUS_STATUS_UNSPECIFIED
	}
}

func convertTreatmentStatus(status string) process_execution_service.TreatmentStatus {
	switch status {
	case "RUNNING":
		return process_execution_service.TreatmentStatus_TREATMENT_STATUS_RUNNING
	case "BLOCKED":
		return process_execution_service.TreatmentStatus_TREATMENT_STATUS_BLOCKED
	case "COMPLETED":
		return process_execution_service.TreatmentStatus_TREATMENT_STATUS_COMPLETED
	case "CANCELLED":
		return process_execution_service.TreatmentStatus_TREATMENT_STATUS_CANCELLED
	default:
		return process_execution_service.TreatmentStatus_TREATMENT_STATUS_UNSPECIFIED
	}
}

func convertTimestampFromString(t string) *timestamp.Timestamp {
	// Implement conversion from string to Timestamp as needed
	// Example implementation:
	parsedTime, _ := time.Parse(time.RFC3339, t)
	return &timestamp.Timestamp{
		Seconds: int64(parsedTime.Second()),
		Nanos:   int32(parsedTime.Nanosecond()),
	}
}

func convertTimestampFromTime(t time.Time) *timestamp.Timestamp {
	return &timestamp.Timestamp{
		Seconds: int64(t.Second()),
		Nanos:   int32(t.Nanosecond()),
	}
}
func convertTaskStatus(status string) process_execution_service.TaskStatus {
	switch status {
	case "NOT_STARTED":
		return process_execution_service.TaskStatus_TASK_STATUS_NOT_STARTED
	case "IN_PROGRESS":
		return process_execution_service.TaskStatus_TASK_STATUS_IN_PROGRESS
	case "BLOCKED":
		return process_execution_service.TaskStatus_TASK_STATUS_BLOCKED
	case "DONE":
		return process_execution_service.TaskStatus_TASK_STATUS_DONE
	default:
		return process_execution_service.TaskStatus_TASK_STATUS_UNSPECIFIED
	}
}
