import * as grpcWeb from 'grpc-web';

import * as ping_pb from './ping_pb';


export class PingServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  echo(
    request: ping_pb.MessageInput,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: ping_pb.MessageOutput) => void
  ): grpcWeb.ClientReadableStream<ping_pb.MessageOutput>;

  now(
    request: ping_pb.Empty,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: ping_pb.Time) => void
  ): grpcWeb.ClientReadableStream<ping_pb.Time>;

}

export class PingServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  echo(
    request: ping_pb.MessageInput,
    metadata?: grpcWeb.Metadata
  ): Promise<ping_pb.MessageOutput>;

  now(
    request: ping_pb.Empty,
    metadata?: grpcWeb.Metadata
  ): Promise<ping_pb.Time>;

}

