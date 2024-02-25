import { Controller, Get } from '@nestjs/common';
import { HelloSvcService } from './hello-svc.service';

@Controller()
export class HelloSvcController {
  constructor(private readonly helloSvcService: HelloSvcService) {}

  @Get()
  getHello(): string {
    return this.helloSvcService.getHello();
  }
}
