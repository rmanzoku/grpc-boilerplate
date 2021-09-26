/**
 * @fileoverview gRPC-Web generated client stub for ping
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_api_annotations_pb = require('./google/api/annotations_pb.js')
const proto = {};
proto.ping = require('./ping_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.ping.PingServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.ping.PingServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.ping.MessageInput,
 *   !proto.ping.MessageOutput>}
 */
const methodDescriptor_PingService_Echo = new grpc.web.MethodDescriptor(
  '/ping.PingService/Echo',
  grpc.web.MethodType.UNARY,
  proto.ping.MessageInput,
  proto.ping.MessageOutput,
  /**
   * @param {!proto.ping.MessageInput} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ping.MessageOutput.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.ping.MessageInput,
 *   !proto.ping.MessageOutput>}
 */
const methodInfo_PingService_Echo = new grpc.web.AbstractClientBase.MethodInfo(
  proto.ping.MessageOutput,
  /**
   * @param {!proto.ping.MessageInput} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ping.MessageOutput.deserializeBinary
);


/**
 * @param {!proto.ping.MessageInput} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.ping.MessageOutput)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.ping.MessageOutput>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ping.PingServiceClient.prototype.echo =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ping.PingService/Echo',
      request,
      metadata || {},
      methodDescriptor_PingService_Echo,
      callback);
};


/**
 * @param {!proto.ping.MessageInput} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.ping.MessageOutput>}
 *     Promise that resolves to the response
 */
proto.ping.PingServicePromiseClient.prototype.echo =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ping.PingService/Echo',
      request,
      metadata || {},
      methodDescriptor_PingService_Echo);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.ping.Empty,
 *   !proto.ping.Time>}
 */
const methodDescriptor_PingService_Now = new grpc.web.MethodDescriptor(
  '/ping.PingService/Now',
  grpc.web.MethodType.UNARY,
  proto.ping.Empty,
  proto.ping.Time,
  /**
   * @param {!proto.ping.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ping.Time.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.ping.Empty,
 *   !proto.ping.Time>}
 */
const methodInfo_PingService_Now = new grpc.web.AbstractClientBase.MethodInfo(
  proto.ping.Time,
  /**
   * @param {!proto.ping.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ping.Time.deserializeBinary
);


/**
 * @param {!proto.ping.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.ping.Time)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.ping.Time>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ping.PingServiceClient.prototype.now =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ping.PingService/Now',
      request,
      metadata || {},
      methodDescriptor_PingService_Now,
      callback);
};


/**
 * @param {!proto.ping.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.ping.Time>}
 *     Promise that resolves to the response
 */
proto.ping.PingServicePromiseClient.prototype.now =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ping.PingService/Now',
      request,
      metadata || {},
      methodDescriptor_PingService_Now);
};


module.exports = proto.ping;

