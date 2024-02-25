import { HelloWorldResponse } from '../proto/hello';
import { Injectable } from '@nestjs/common';

@Injectable()
export class HelloSvcService {
  getHello(name: HelloWorldResponse): HelloWorldResponse {
    return {message: name.message};
  }
}
