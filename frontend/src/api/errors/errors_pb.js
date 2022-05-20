// source: errors/errors.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var google_protobuf_descriptor_pb = require('google-protobuf/google/protobuf/descriptor_pb.js');
goog.object.extend(proto, google_protobuf_descriptor_pb);
goog.exportSymbol('proto.errors.code', null, global);
goog.exportSymbol('proto.errors.defaultCode', null, global);

/**
 * A tuple of {field number, class constructor} for the extension
 * field named `defaultCode`.
 * @type {!jspb.ExtensionFieldInfo<number>}
 */
proto.errors.defaultCode = new jspb.ExtensionFieldInfo(
    1108,
    {defaultCode: 0},
    null,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         null),
    0);

google_protobuf_descriptor_pb.EnumOptions.extensionsBinary[1108] = new jspb.ExtensionFieldBinaryInfo(
    proto.errors.defaultCode,
    jspb.BinaryReader.prototype.readInt32,
    jspb.BinaryWriter.prototype.writeInt32,
    undefined,
    undefined,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.EnumOptions.extensions[1108] = proto.errors.defaultCode;


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `code`.
 * @type {!jspb.ExtensionFieldInfo<number>}
 */
proto.errors.code = new jspb.ExtensionFieldInfo(
    1109,
    {code: 0},
    null,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         null),
    0);

google_protobuf_descriptor_pb.EnumValueOptions.extensionsBinary[1109] = new jspb.ExtensionFieldBinaryInfo(
    proto.errors.code,
    jspb.BinaryReader.prototype.readInt32,
    jspb.BinaryWriter.prototype.writeInt32,
    undefined,
    undefined,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.EnumValueOptions.extensions[1109] = proto.errors.code;

goog.object.extend(exports, proto.errors);
