// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=ts"
// @generated from file proto/remote.proto (package remote, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { ExecuteTaskRequest, ExecuteTaskResponse, PingRequest, PingResponse } from "./remote_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service remote.RemoteAgentService
 */
export const RemoteAgentService = {
  typeName: "remote.RemoteAgentService",
  methods: {
    /**
     * @generated from rpc remote.RemoteAgentService.ExecuteTask
     */
    executeTask: {
      name: "ExecuteTask",
      I: ExecuteTaskRequest,
      O: ExecuteTaskResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc remote.RemoteAgentService.Ping
     */
    ping: {
      name: "Ping",
      I: PingRequest,
      O: PingResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

