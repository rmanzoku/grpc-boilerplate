/**
 * @fileoverview gRPC-Web generated client stub for healthcheck
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as healthcheck_pb from './healthcheck_pb';


export class HealthcheckServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoHealthcheck = new grpcWeb.MethodDescriptor(
    '/healthcheck.HealthcheckService/Healthcheck',
    grpcWeb.MethodType.UNARY,
    healthcheck_pb.Empty,
    healthcheck_pb.MessageOutput,
    (request: healthcheck_pb.Empty) => {
      return request.serializeBinary();
    },
    healthcheck_pb.MessageOutput.deserializeBinary
  );

  healthcheck(
    request: healthcheck_pb.Empty,
    metadata: grpcWeb.Metadata | null): Promise<healthcheck_pb.MessageOutput>;

  healthcheck(
    request: healthcheck_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: healthcheck_pb.MessageOutput) => void): grpcWeb.ClientReadableStream<healthcheck_pb.MessageOutput>;

  healthcheck(
    request: healthcheck_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: healthcheck_pb.MessageOutput) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/healthcheck.HealthcheckService/Healthcheck',
        request,
        metadata || {},
        this.methodInfoHealthcheck,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/healthcheck.HealthcheckService/Healthcheck',
    request,
    metadata || {},
    this.methodInfoHealthcheck);
  }

}

