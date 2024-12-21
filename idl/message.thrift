namespace go messages

include "base.thrift"

struct MessageInfo {
    1: required string from_uid;
    2: required string to_uid;
    3: required string content;
}

struct InsertMessageRequest {
    1: MessageInfo message;
}

struct InsertMessageResponse {
    1: base.Status base;
}

struct PopMessageRequest {
    1: required string uid;
}

struct PopMessageResponseData {
    1: list<MessageInfo> Items;
}
struct PopMessageResponse {
    1: base.Status base;
    2: PopMessageResponseData data;
}

service MessageService {
    InsertMessageResponse InsertMessage(1: InsertMessageRequest request);
    PopMessageResponse PopMessage(1: PopMessageRequest request);
}