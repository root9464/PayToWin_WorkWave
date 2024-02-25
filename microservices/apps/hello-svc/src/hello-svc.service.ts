import { Injectable } from '@nestjs/common';

@Injectable()
export class HelloSvcService {
  getHello(): string {
    return 'Hello World!';
  }
}
