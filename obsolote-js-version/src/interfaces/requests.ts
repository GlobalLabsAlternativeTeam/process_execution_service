import { Schema } from './objects';

export interface CreateInstanceReq {
  schema: Schema;
  doctorId: number;
  patientId: number;
}

export interface InstanceById {
  id: number;
}
