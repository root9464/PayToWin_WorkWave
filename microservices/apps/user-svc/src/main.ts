import { NestFactory } from '@nestjs/core';
import { UserSvcModule } from './user-svc.module';

async function bootstrap() {
  const app = await NestFactory.create(UserSvcModule);
  await app.listen(3000);
}
bootstrap();
