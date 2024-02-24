/* eslint-disable prettier/prettier */
import { Controller, Get } from '@nestjs/common';
import { HelloService } from './hello.service';
import { type HelloWorldResponse } from '../../../lib/proto/out/file';
import { GrpcMethod } from '@nestjs/microservices';

@Controller('hello')
export class HelloController {
  constructor(private readonly helloService: HelloService) {}

  @GrpcMethod('GetHello', 'HelloWorld')
  @Get()
  HelloWorld(name: HelloWorldResponse): HelloWorldResponse {
    return this.helloService.sayHello(name);
  }
}
