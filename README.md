# cube
cube tool for coding

# install

go install -v github.com/rockburning/cube/cmd/cube@latest

# Usage

```
Usage:
cube rpc -s=YourService -m=CreaetUser,DeleteUser
```

# For help

```
cube help
```

# 生成示例
cube rpc -s=YourService -m=CreaetUser,DeleteUsers -req=Req

```protobuf
service YourService {
  rpc CreaetUser (CreaetUserReq) returns (CreaetUserReply){
    option (google.api.http) = {
    };
  };
  rpc DeleteUsers (DeleteUsersReq) returns (DeleteUsersReply){
    option (google.api.http) = {
    };
  };

}

message CreaetUserReq{

}

message CreaetUserReply{

}

message DeleteUsersReq{

}

message DeleteUsersReply{

}

```