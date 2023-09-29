// @generated by protoc-gen-connect-es v1.0.0 with parameter "target=ts"
// @generated from file todo/v1/todo.proto (package todo.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateToDoRequest, CreateToDoResponse } from "./todo_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service todo.v1.ToDoService
 */
export const ToDoService = {
  typeName: "todo.v1.ToDoService",
  methods: {
    /**
     * @generated from rpc todo.v1.ToDoService.CreateToDo
     */
    createToDo: {
      name: "CreateToDo",
      I: CreateToDoRequest,
      O: CreateToDoResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

