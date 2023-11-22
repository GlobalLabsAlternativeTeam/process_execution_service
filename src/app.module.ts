import { Module } from '@nestjs/common';

// Modules
import { StatusModule } from './status/status.module';
import { TreatmentModule } from './treatment/treatment.module';

@Module({
  imports: [StatusModule, TreatmentModule],
  controllers: [],
  providers: [],
})
export class AppModule {}
