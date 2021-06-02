package main

import (
	"fmt"
	"proto-example/go/codepb"
	"proto-example/go/msgpb"
)

func main() {
	fmt.Println(msgpb.EMessageCode_CODE_Success)

	fmt.Println(msgpb.Code_Enum(0))
	fmt.Println(msgpb.Code_Success)

	fmt.Println(msgpb.MsgId_Enum(0))
	fmt.Println(msgpb.MsgId_LoginRequest)
	fmt.Println(msgpb.MsgCode_Enum(0))
	fmt.Println(msgpb.MsgCode_Success)

	fmt.Println(msgpb.Msg_Id(0))
	fmt.Println(msgpb.Msg_IdLoginRequest)
	fmt.Println(msgpb.Msg_CodeSuccess)

	fmt.Println(msgpb.ResourceDescriptor_DECLARATIVE_FRIENDLY)

	fmt.Println(msgpb.ChangeType_CREATE)

	fmt.Println(codepb.Code_Success)

}
