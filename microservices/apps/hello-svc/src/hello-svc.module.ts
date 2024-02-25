import { Module } from '@nestjs/common';
import { HelloSvcController } from './hello-svc.controller';
import { HelloSvcService } from './hello-svc.service';

@Module({
  imports: [],
  controllers: [HelloSvcController],
  providers: [HelloSvcService],
})
export class HelloSvcModule {}
