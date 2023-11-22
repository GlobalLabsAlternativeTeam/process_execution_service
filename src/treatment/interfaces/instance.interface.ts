import { InstanceStatus } from './instance-status.interface';

export interface Instance {
  id: number;
  doctorId: number;
  patientId: number;
  status: InstanceStatus;
  schemaInstanceId: string;
  startedAt: string;
  finishedAt: string;
  deletedAt: string;
}
