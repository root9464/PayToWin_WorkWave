import { Controller, Get } from '@nestjs/common';
import { HelloSvcService } from './hello-svc.service';

import { HelloWorldResponse } from '../../../proto/hello/hello';
import { GrpcMethod } from '@nestjs/microservices';

@Controller()
export class HelloSvcController {
  constructor(private readonly helloSvcService: HelloSvcService) {}

  @GrpcMethod('GetHello', 'HelloWorld')
  @Get()
  getHello(name: HelloWorldResponse): HelloWorldResponse {
    return this.helloSvcService.getHello(name);
  }
}
