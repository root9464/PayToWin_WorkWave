import { Injectable } from '@nestjs/common';
import { type HelloWorldResponse } from '../../../lib/proto/out/file';

@Injectable()
export class HelloService {
  sayHello(name: HelloWorldResponse): HelloWorldResponse {
    return {message: name.message}
  }
}
