using msgpb = MsgPb;

class App {
    static int Main(string[] args) {
        print(msgpb::EMessageCode.CodeSuccess);

        print(msgpb::Code.Types.Enum(0));
        print(msgpb::Code.Types.Enum.Success);

        print(msgpb::MsgId.Types.Enum(0));
        print(msgpb::MsgId.Types.Enum.LoginRequest);
        print(msgpb::MsgCode.Types.Enum(0));
        print(msgpb::MsgCode.Types.Enum.Success);

        print(msgpb::Msg.Types.Id(0));
        print(msgpb::Msg.Types.Id.LoginRequest);
        print(msgpb::Msg.Types.Code.Success);


        print(msgpb::ResourceDescriptor.Types.Style.Unspecified);
        print(msgpb::ResourceDescriptor.Types.Style.DeclarativeFriendly);

        print(msgpb::ChangeType.Create);

        return 0;
    }
}