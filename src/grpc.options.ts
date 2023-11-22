import { GrpcOptions, Transport } from '@nestjs/microservices';
import { join } from 'path';

export const grpcOptions: GrpcOptions = {
  transport: Transport.GRPC,
  options: {
    package: ['hero'],
    protoPath: ['./hero/hero.proto'].map((path) => join(__dirname, path)),
  },
};
