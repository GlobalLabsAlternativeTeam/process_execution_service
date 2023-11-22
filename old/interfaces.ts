// Components

export interface Schema {
  schema_id: number;
  author_id: number;
  schema_name: string;
  created_at?: Date;
  updated_at?: Date;
  deleted_at?: Date;
  tasks: Task[];
}

export interface Task {
  id: number;
  level: number;
  name: string;
  status: TaskStatus;
  blocked_by: number[];
  responsible?: number;
  time_limit?: number;
  children: Task[];
  comment?: string;
}

export interface Instance {
  instance_id: number;
  doctor_id: number;
  patient_id: number;
  status: InstanceStatus;
  schema_instance_id: string;
  started_at?: Date;
  finished_at?: Date;
  deleted_at?: Date;
}

export interface SchemaInstance {
  schema_instance_id: string;
  schema_reference: Schema;
}

// Status

export enum TaskStatus {
  TASK_STATUS_UNSPECIFIED = 0,
  TASK_STATUS_NOT_STARTED = 1,
  TASK_STATUS_IN_PROGRESS = 2,
  TASK_STATUS_BLOCKED = 3,
  TASK_STATUS_DONE = 4,
}

export interface TaskStatusGroup {
  task_id: number;
  task_status: TaskStatus;
}

export enum InstanceStatus {
  INSTANCE_STATUS_UNSPECIFIED = 0,
  INSTANCE_STATUS_RUNNING = 1,
  INSTANCE_STATUS_BLOCKED = 2,
  INSTANCE_STATUS_COMPLETED = 3,
  INSTANCE_STATUS_CANCELLED = 4,
}

// Requests

export interface CreateInstanceReq {
  schema: Schema;
  patient_id: number;
  doctor_id: number;
}

export interface GetInstancesByPatientReq {
  patient_id: number;
}

export interface GetInstanceByIdReq {
  instance_id: string;
}

export interface CompleteTaskReq {
  instance_id: string;
  task_id: number[];
}

export interface GetTaskReq {
  instance_id: number;
  task_id: number;
}

export interface GetSchemaInstanceReq {
  schema_instance_id: string;
}

export interface InstanceById {
  id: number;
}

export interface DummyInstance {
  instance_id: number;
  doctor_id: number;
  patient_id: number;
  schema_instance_id: string;
}
