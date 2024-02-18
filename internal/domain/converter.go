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
		TreatmentId:     t.TreatmentID,
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
		CreatedAt:   ConvertTimestampFromTime(p.CreatedAt),
		UpdatedAt:   ConvertTimestampFromTime(p.UpdatedAt),
		DeletedAt:   ConvertTimestampFromTime(p.DeletedAt),
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
	return &process_execution_service.Task{
		Id:          int64(t.ID),
		Level:       int32(t.Level),
		Name:        t.Name,
		Status:      ConvertTaskStatus(t.Status),
		BlockedBy:   t.BlockedBy,
		Responsible: t.Responsible,
		TimeLimit:   t.TimeLimit,
		Children:    TasksToGRPC(t.Children),
		Comment:     wrapperspb.String(t.Comment.Value),
	}
}

func ProtoToSchema(proto *process_execution_service.Schema) Schema {
	// Convert process_execution_service.Schema to domain.Schema
	return Schema{
		SchemaID:   proto.SchemaId,
		AuthorID:   proto.AuthorId,
		SchemaName: proto.SchemaName,
		CreatedAt:  proto.CreatedAt.AsTime(),
		UpdatedAt:  proto.UpdatedAt.AsTime(),
		DeletedAt:  proto.DeletedAt.AsTime(),
		Tasks:      ProtoToTasks(proto.Tasks),
	}
}

func ProtoToTasks(protoTasks []*process_execution_service.Task) []Task {
	tasks := make([]Task, len(protoTasks))
	for i, protoTask := range protoTasks {
		var comment *struct {
			Value string `json:"value"`
		}
		if protoTask.Comment != nil {
			val := protoTask.Comment.GetValue() // Extract string value from wrapperspb.StringValue
			comment = &struct {
				Value string `json:"value"`
			}{Value: val}
		}
		tasks[i] = Task{
			ID:          int(protoTask.Id),
			Level:       int(protoTask.Level),
			Name:        protoTask.Name,
			Status:      ConvertProroToTaskStatus(protoTask.Status),
			BlockedBy:   protoTask.BlockedBy,
			Responsible: protoTask.Responsible,
			TimeLimit:   protoTask.TimeLimit,
			Children:    ProtoToTasks(protoTask.Children),
			Comment:     *comment,
		}
	}
	return tasks
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

func ConvertTimestampFromTime(t time.Time) *timestamp.Timestamp {
	return &timestamp.Timestamp{
		Seconds: int64(t.Second()),
		Nanos:   int32(t.Nanosecond()),
	}
}
func ConvertTaskStatus(status string) process_execution_service.TaskStatus {
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

func ConvertProroToTaskStatus(status process_execution_service.TaskStatus) string {
	switch status.String() {
	case "TASK_STATUS_NOT_STARTED":
		return "NOT_STARTED"
	case "TASK_STATUS_IN_PROGRESS":
		return "IN_PROGRESS"
	case "TASK_STATUS_BLOCKED":
		return "BLOCKED"
	case "TASK_STATUS_DONE":
		return "DONE"
	default:
		return "STATUS_UNSPECIFIED"
	}
}
