package service

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/takutakahashi/agent-connect/gen/proto"
	"github.com/takutakahashi/agent-connect/gen/proto/protoconnect"
	remoteagent "github.com/takutakahashi/agent-connect/pkg/agent"
)

// RemoteAgentServer はRemoteAgentServiceの実装です
type RemoteAgentServer struct {
	factory remoteagent.AgentFactory
}

// NewRemoteAgentServer は新しいRemoteAgentServerを作成します
func NewRemoteAgentServer(factory remoteagent.AgentFactory) *RemoteAgentServer {
	return &RemoteAgentServer{factory: factory}
}

// ExecuteTask はタスクを実行するメソッドです
func (s *RemoteAgentServer) ExecuteTask(
	ctx context.Context,
	req *connect.Request[proto.ExecuteTaskRequest],
) (*connect.Response[proto.ExecuteTaskResponse], error) {
	slog.Info("ExecuteTask", "request", req.Msg)

	// リクエストからsession_idを取得
	agent := s.factory.NewAgent(req.Msg.SessionId)
	out, err := agent.Execute(ctx, req.Msg.Instruction)
	if err != nil {
		return nil, err
	}
	// レスポンスの作成
	res := &proto.ExecuteTaskResponse{
		SessionId: req.Msg.SessionId,
		Stdout:    out,
		Stderr:    "",
		Success:   true,
	}

	return connect.NewResponse(res), nil
}

// Ping はサーバーの状態を確認するメソッドです
func (s *RemoteAgentServer) Ping(
	ctx context.Context,
	req *connect.Request[proto.PingRequest],
) (*connect.Response[proto.PingResponse], error) {
	slog.Info("Ping received")

	res := &proto.PingResponse{
		Status: "OK",
	}

	return connect.NewResponse(res), nil
}

// Handler はRemoteAgentServiceのHTTPハンドラを返します
func (s *RemoteAgentServer) Handler() (string, http.Handler) {
	return protoconnect.NewRemoteAgentServiceHandler(s)
}
