// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var todo_pb = require('./todo_pb.js');

function serialize_todo_Task(arg) {
  if (!(arg instanceof todo_pb.Task)) {
    throw new Error('Expected argument of type todo.Task');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_todo_Task(buffer_arg) {
  return todo_pb.Task.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todo_TaskList(arg) {
  if (!(arg instanceof todo_pb.TaskList)) {
    throw new Error('Expected argument of type todo.TaskList');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_todo_TaskList(buffer_arg) {
  return todo_pb.TaskList.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todo_Text(arg) {
  if (!(arg instanceof todo_pb.Text)) {
    throw new Error('Expected argument of type todo.Text');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_todo_Text(buffer_arg) {
  return todo_pb.Text.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todo_Void(arg) {
  if (!(arg instanceof todo_pb.Void)) {
    throw new Error('Expected argument of type todo.Void');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_todo_Void(buffer_arg) {
  return todo_pb.Void.deserializeBinary(new Uint8Array(buffer_arg));
}


var TasksService = exports.TasksService = {
  list: {
    path: '/todo.Tasks/List',
    requestStream: false,
    responseStream: false,
    requestType: todo_pb.Void,
    responseType: todo_pb.TaskList,
    requestSerialize: serialize_todo_Void,
    requestDeserialize: deserialize_todo_Void,
    responseSerialize: serialize_todo_TaskList,
    responseDeserialize: deserialize_todo_TaskList,
  },
  add: {
    path: '/todo.Tasks/Add',
    requestStream: false,
    responseStream: false,
    requestType: todo_pb.Text,
    responseType: todo_pb.Task,
    requestSerialize: serialize_todo_Text,
    requestDeserialize: deserialize_todo_Text,
    responseSerialize: serialize_todo_Task,
    responseDeserialize: deserialize_todo_Task,
  },
};

exports.TasksClient = grpc.makeGenericClientConstructor(TasksService);
