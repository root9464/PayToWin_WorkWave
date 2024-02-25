import { Controller, Get } from '@nestjs/common';
import { HelloSvcService } from './hello-svc.service';
import { HelloWorldResponse } from '../proto/hello';

@Controller()
export class HelloSvcController {
  constructor(private readonly helloSvcService: HelloSvcService) {}

  @Get()
  getHello(name: HelloWorldResponse): HelloWorldResponse {
    return this.helloSvcService.getHello(name);
  }
}
