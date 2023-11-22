import { Module } from '@nestjs/common';
import { TreatmentController } from './treatment.controller';

@Module({
  imports: [],
  controllers: [TreatmentController],
})
export class TreatmentModule {}
