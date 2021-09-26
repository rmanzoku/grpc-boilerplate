import * as jspb from 'google-protobuf'

import * as google_api_annotations_pb from './google/api/annotations_pb';


export class MessageInput extends jspb.Message {
  getMsg(): string;
  setMsg(value: string): MessageInput;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MessageInput.AsObject;
  static toObject(includeInstance: boolean, msg: MessageInput): MessageInput.AsObject;
  static serializeBinaryToWriter(message: MessageInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MessageInput;
  static deserializeBinaryFromReader(message: MessageInput, reader: jspb.BinaryReader): MessageInput;
}

export namespace MessageInput {
  export type AsObject = {
    msg: string,
  }
}

export class MessageOutput extends jspb.Message {
  getMsg(): string;
  setMsg(value: string): MessageOutput;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MessageOutput.AsObject;
  static toObject(includeInstance: boolean, msg: MessageOutput): MessageOutput.AsObject;
  static serializeBinaryToWriter(message: MessageOutput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MessageOutput;
  static deserializeBinaryFromReader(message: MessageOutput, reader: jspb.BinaryReader): MessageOutput;
}

export namespace MessageOutput {
  export type AsObject = {
    msg: string,
  }
}

export class Time extends jspb.Message {
  getT(): number;
  setT(value: number): Time;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Time.AsObject;
  static toObject(includeInstance: boolean, msg: Time): Time.AsObject;
  static serializeBinaryToWriter(message: Time, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Time;
  static deserializeBinaryFromReader(message: Time, reader: jspb.BinaryReader): Time;
}

export namespace Time {
  export type AsObject = {
    t: number,
  }
}

export class Empty extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Empty.AsObject;
  static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
  static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Empty;
  static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
  export type AsObject = {
  }
}

