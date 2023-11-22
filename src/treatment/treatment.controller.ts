import { Controller } from '@nestjs/common';
import { GrpcMethod, GrpcStreamMethod } from '@nestjs/microservices';
import { Observable, Subject } from 'rxjs';

// Interfaces
// import { InstanceById } from './interfaces/instance-by-id.interface';
import { Instance } from './interfaces/instance.interface';
import { InstanceStatus } from './interfaces/instance-status.interface';
import { CreateInstanceReq } from './interfaces/create-instance-req.interface';

@Controller('treatment')
export class TreatmentController {
  constructor() {}

  // InstanceService ==========================================================

  @GrpcMethod('InstanceService')
  createInstance(data: CreateInstanceReq): Instance {
    const now = new Date().toISOString();

    const createdInstance: Instance = {
      id: 1,
      status: InstanceStatus.INSTANCE_STATUS_RUNNING,
      schemaInstanceId: 'dummySchemaInstanceId',
      startedAt: now,
      finishedAt: now,
      deletedAt: now,
      ...data,
    };

    return createdInstance;
  }

  // Examples =================================================================

  // @GrpcMethod('InstanceService')
  // findOne(data: InstanceById): Instance {
  //   return this.items.find(({ id }) => id === data.id);
  // }

  // @GrpcStreamMethod('InstanceService')
  // findMany(data$: Observable<InstanceById>): Observable<Instance> {
  //   const instance$ = new Subject<Instance>();

  //   const onNext = (instanceById: InstanceById) => {
  //     const item = this.items.find(({ id }) => id === instanceById.id);
  //     instance$.next(item);
  //   };
  //   const onComplete = () => instance$.complete();
  //   data$.subscribe({
  //     next: onNext,
  //     complete: onComplete,
  //   });

  //   return instance$.asObservable();
  // }
}
