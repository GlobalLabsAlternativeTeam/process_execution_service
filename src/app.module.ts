import { Module } from '@nestjs/common';

// Modules
import { StatusModule } from './status/status.module';
import { PrismaModule } from './prisma/prisma.module';
import { AppController } from './app.controller';

@Module({
  imports: [PrismaModule, StatusModule],
  controllers: [AppController],
  providers: [],
})
export class AppModule {}
