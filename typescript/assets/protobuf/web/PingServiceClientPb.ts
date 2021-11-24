/**
 * @fileoverview gRPC-Web generated client stub for ping
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as ping_pb from './ping_pb';


export class PingServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'binary';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoEcho = new grpcWeb.AbstractClientBase.MethodInfo(
    ping_pb.MessageOutput,
    (request: ping_pb.MessageInput) => {
      return request.serializeBinary();
    },
    ping_pb.MessageOutput.deserializeBinary
  );

  echo(
    request: ping_pb.MessageInput,
    metadata: grpcWeb.Metadata | null): Promise<ping_pb.MessageOutput>;

  echo(
    request: ping_pb.MessageInput,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: ping_pb.MessageOutput) => void): grpcWeb.ClientReadableStream<ping_pb.MessageOutput>;

  echo(
    request: ping_pb.MessageInput,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: ping_pb.MessageOutput) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/ping.PingService/Echo',
        request,
        metadata || {},
        this.methodInfoEcho,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/ping.PingService/Echo',
    request,
    metadata || {},
    this.methodInfoEcho);
  }

  methodInfoNow = new grpcWeb.AbstractClientBase.MethodInfo(
    ping_pb.Time,
    (request: ping_pb.Empty) => {
      return request.serializeBinary();
    },
    ping_pb.Time.deserializeBinary
  );

  now(
    request: ping_pb.Empty,
    metadata: grpcWeb.Metadata | null): Promise<ping_pb.Time>;

  now(
    request: ping_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: ping_pb.Time) => void): grpcWeb.ClientReadableStream<ping_pb.Time>;

  now(
    request: ping_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: ping_pb.Time) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/ping.PingService/Now',
        request,
        metadata || {},
        this.methodInfoNow,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/ping.PingService/Now',
    request,
    metadata || {},
    this.methodInfoNow);
  }

  methodInfoDBSelect = new grpcWeb.AbstractClientBase.MethodInfo(
    ping_pb.MessageOutput,
    (request: ping_pb.Empty) => {
      return request.serializeBinary();
    },
    ping_pb.MessageOutput.deserializeBinary
  );

  dBSelect(
    request: ping_pb.Empty,
    metadata: grpcWeb.Metadata | null): Promise<ping_pb.MessageOutput>;

  dBSelect(
    request: ping_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: ping_pb.MessageOutput) => void): grpcWeb.ClientReadableStream<ping_pb.MessageOutput>;

  dBSelect(
    request: ping_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: ping_pb.MessageOutput) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/ping.PingService/DBSelect',
        request,
        metadata || {},
        this.methodInfoDBSelect,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/ping.PingService/DBSelect',
    request,
    metadata || {},
    this.methodInfoDBSelect);
  }

  methodInfoKMSSign = new grpcWeb.AbstractClientBase.MethodInfo(
    ping_pb.MessageOutput,
    (request: ping_pb.MessageInput) => {
      return request.serializeBinary();
    },
    ping_pb.MessageOutput.deserializeBinary
  );

  kMSSign(
    request: ping_pb.MessageInput,
    metadata: grpcWeb.Metadata | null): Promise<ping_pb.MessageOutput>;

  kMSSign(
    request: ping_pb.MessageInput,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: ping_pb.MessageOutput) => void): grpcWeb.ClientReadableStream<ping_pb.MessageOutput>;

  kMSSign(
    request: ping_pb.MessageInput,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: ping_pb.MessageOutput) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/ping.PingService/KMSSign',
        request,
        metadata || {},
        this.methodInfoKMSSign,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/ping.PingService/KMSSign',
    request,
    metadata || {},
    this.methodInfoKMSSign);
  }

}

