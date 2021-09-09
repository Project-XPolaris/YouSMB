package rpc

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"yousmb/application"
	"yousmb/smb"
)

var DefaultRPCServer = &RPCServer{}

const (
	ErrorCodeUnknown        = 9999
	ErrorCodeEntityNotFound = 6001
	ErrorCodeEntity         = 2001
)

type RPCServer struct {
	server Server
}

func (l *RPCServer) Run() {
	lis, err := net.Listen("tcp", application.Config.RPC)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	rpcServer := grpc.NewServer()
	l.server = Server{}
	RegisterYouSMBServiceServer(rpcServer, &l.server)
	log.Printf("server listening at %v", lis.Addr())
	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type Server struct {
	UnimplementedYouSMBServiceServer
}

func (s Server) GetConfig(context.Context, *Empty) (*ConfigReply, error) {
	sectionListReply := make([]*Section, 0)
	for _, section := range smb.Current.Sections {
		sectionListReply = append(sectionListReply, &Section{
			Name:   &section.Name,
			Fields: section.Fields,
		})
	}
	return &ConfigReply{Sections: sectionListReply}, nil
}
func (s Server) AddFolderConfig(ctx context.Context, in *AddConfigMessage) (*BaseReply, error) {
	err := smb.Current.AddFolder(in.GetName(), in.GetProperties())
	if err != nil {
		return GenerateErrorReply(err), nil
	}
	return GenerateBaseReply(), nil
}

func (s Server) UpdateFolderConfig(ctx context.Context, in *AddConfigMessage) (*BaseReply, error) {
	err := smb.Current.UpdateFolder(in.GetName(), in.GetProperties())
	if err != nil {
		return GenerateErrorReply(err), nil
	}
	return GenerateBaseReply(), nil
}

func (s Server) RemoveFolderConfig(ctx context.Context, in *RemoveConfigMessage) (*BaseReply, error) {
	err := smb.Current.RemoveFolder(in.GetName())
	if err != nil {
		return GenerateErrorReply(err), nil
	}
	return GenerateBaseReply(), nil
}

func (s Server) AddUser(ctx context.Context, in *AddUserMessage) (*BaseReply, error) {
	err := smb.DefaultUserManager.Create(in.GetUsername(), in.GetPassword())
	if err != nil {
		return GenerateErrorReply(err), nil
	}
	return GenerateBaseReply(), nil
}

func (s Server) RemoveUser(ctx context.Context, in *RemoveUserMessage) (*BaseReply, error) {
	err := smb.DefaultUserManager.RemoveUser(in.GetUsername())
	if err != nil {
		return GenerateErrorReply(err), nil
	}
	return GenerateBaseReply(), nil
}
func (s Server) GetSMBStatus(context.Context, *Empty) (*SMBStatusReply, error) {
	replyStatusProcessList := make([]*StatusProcess, 0)
	for _, statusProcess := range smb.DefaultMonitor.Process {
		replyProcess := &StatusProcess{
			PID:      &statusProcess.PID,
			Username: &statusProcess.Username,
			Group:    &statusProcess.Group,
			Machine:  &statusProcess.Machine,
			Raw:      statusProcess.Raw,
		}
		replyStatusProcessList = append(replyStatusProcessList, replyProcess)
	}
	replyStatusSharesList := make([]*StatusShares, 0)
	for _, statusShares := range smb.DefaultMonitor.StatusShares {
		connectAt := statusShares.ConnectAt.String()
		replyStatusSharesList = append(replyStatusSharesList, &StatusShares{
			Service:   &statusShares.Service,
			PID:       &statusShares.PID,
			Machine:   &statusShares.Machine,
			ConnectAt: &connectAt,
			Raw:       statusShares.Raw,
		})
	}
	success := true
	return &SMBStatusReply{
		Success: &success,
		Process: replyStatusProcessList,
		Shares:  replyStatusSharesList,
	}, nil
}

func (s Server) GetInfo(context.Context, *Empty) (*ServiceInfoReply, error) {
	name := "YouSMB Service"
	status := "running"
	success := true
	return &ServiceInfoReply{
		Name:    &name,
		Status:  &status,
		Success: &success,
	}, nil
}
