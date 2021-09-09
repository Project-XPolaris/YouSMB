package rpc

import "yousmb/utils"

func GenerateErrorReply(err error) *BaseReply {
	return &BaseReply{Success: utils.GetBoolPointer(false), Reason: utils.GetStringPointer(err.Error())}
}
func GenerateBaseReply() *BaseReply {
	return &BaseReply{Success: utils.GetBoolPointer(true)}
}
