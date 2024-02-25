import { NestFactory } from '@nestjs/core';
import { HelloSvcModule } from './hello-svc.module';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { join } from 'path';

async function bootstrap() {
  const app = await NestFactory.createMicroservice<MicroserviceOptions>(
    HelloSvcModule,
    {
      transport: Transport.GRPC,
      options: {
        url: '0.0.0.0:50051',
        package: 'hello',
        protoPath: join(__dirname, '../hello-svc/proto/hello.proto'),
      }
    }
  )
  app.listen();
}
bootstrap();
