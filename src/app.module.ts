import { Module } from '@nestjs/common';

// Modules
import { StatusModule } from './status/status.module';
import { TreatmentModule } from './treatment/treatment.module';
import { PrismaModule } from './prisma/prisma.module';

@Module({
  imports: [PrismaModule, StatusModule, TreatmentModule],
  controllers: [],
  providers: [],
})
export class AppModule {}
