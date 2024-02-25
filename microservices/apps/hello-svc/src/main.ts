import { NestFactory } from '@nestjs/core';
import { HelloSvcModule } from './hello-svc.module';

async function bootstrap() {
  const app = await NestFactory.create(HelloSvcModule);
  await app.listen(3000);
}
bootstrap();
