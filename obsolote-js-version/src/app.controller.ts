import { Controller } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
// import { Observable, Subject } from 'rxjs';
import { Instance } from '@prisma/client';
import { CreateInstanceReq } from '~/interfaces';
import { PrismaService } from '~/prisma/prisma.service';

@Controller()
export class AppController {
  constructor(private prisma: PrismaService) {}

  @GrpcMethod('TreatmentService')
  async createInstance(req: CreateInstanceReq): Promise<Instance> {
    const createdInstance = await this.prisma.instance.create({
      data: {
        schemaInstanceId: 'dummySchemaInstanceId',
        startedAt: new Date().toISOString(),
        ...req,
      },
    });

    return createdInstance;
  }

  // Examples =================================================================

  // @GrpcMethod('TreatmentService')
  // findOne(data: InstanceById): Instance {
  //   return this.items.find(({ id }) => id === data.id);
  // }

  // @GrpcStreamMethod('TreatmentService')
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
