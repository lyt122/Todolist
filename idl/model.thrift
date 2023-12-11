namespace go model


struct User {
    1: required i64 uid
    2: required string username
    3: required string password
    4: required string creatAt
    5: required string updateAt
}
struct Task {
    1: required i64 id
    2: required string title
    3: required string content
    4: required i16 status
    5: required string creatAt
    6: required string updateAt
    7: required string startAt
    8: required string endAt
}

struct TaskList{
    1: required list<Task> tasks
    2: required i64 total
}

struct BaseResp{
    1: required i16 code
    2: required string msg
}