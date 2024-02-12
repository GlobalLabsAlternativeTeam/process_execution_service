import { GrpcOptions, Transport } from '@nestjs/microservices';
import { join } from 'path';

export const grpcOptions: GrpcOptions = {
  transport: Transport.GRPC,
  options: {
    package: ['treatment'],
    protoPath: ['./treatment.proto'].map((path) => join(__dirname, path)),
  },
};
