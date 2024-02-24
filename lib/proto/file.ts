/* eslint-disable */
import { GrpcMethod, GrpcStreamMethod } from "@nestjs/microservices";
import { Observable } from "rxjs";

export const protobufPackage = "file";

export interface HelloWorldResponse {
  message: string;
}

export const FILE_PACKAGE_NAME = "file";

export interface GetHelloClient {
  helloWorld(request: HelloWorldResponse): Observable<HelloWorldResponse>;
}

export interface GetHelloController {
  helloWorld(
    request: HelloWorldResponse,
  ): Promise<HelloWorldResponse> | Observable<HelloWorldResponse> | HelloWorldResponse;
}

export function GetHelloControllerMethods() {
  return function (constructor: Function) {
    const grpcMethods: string[] = ["helloWorld"];
    for (const method of grpcMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
      GrpcMethod("GetHello", method)(constructor.prototype[method], method, descriptor);
    }
    const grpcStreamMethods: string[] = [];
    for (const method of grpcStreamMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
      GrpcStreamMethod("GetHello", method)(constructor.prototype[method], method, descriptor);
    }
  };
}

export const GET_HELLO_SERVICE_NAME = "GetHello";
