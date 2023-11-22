import { NestFactory } from '@nestjs/core';
import { MicroserviceOptions } from '@nestjs/microservices';
import { Logger } from '@nestjs/common';
import { AppModule } from './app.module';
import { grpcOptions } from './grpc.options';

async function bootstrap() {
  // const app = await NestFactory.createMicroservice<MicroserviceOptions>(
  //   AppModule,
  //   grpcOptions,
  // );
  // await app.listen();
  // Logger.log(`Application is running`);

  const app = await NestFactory.create(AppModule);
  app.connectMicroservice<MicroserviceOptions>(grpcOptions);

  await app.startAllMicroservices();
  await app.listen(3000);
  Logger.log(`Application is running on: ${await app.getUrl()}`);
}

bootstrap();
