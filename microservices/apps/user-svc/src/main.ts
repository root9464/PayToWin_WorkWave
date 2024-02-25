import { UserSvcModule } from './user-svc.module';
import { NestFactory } from '@nestjs/core';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { join } from 'path';

async function bootstrap() {
  const app = await NestFactory.createMicroservice<MicroserviceOptions>(
    UserSvcModule,
    {
      transport: Transport.GRPC,
      options: {
        url: '0.0.0.0:50050',
        package: 'hello',
        protoPath: join(__dirname, '../hello-svc/proto/hello.proto'),
      }
    }
  )
  app.listen();
}
bootstrap();
