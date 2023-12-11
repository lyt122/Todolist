namespace go task


include "model.thrift"

struct CreateTaskRequest{
    1: required string title
    2: required string content
    3: required i64 startAt
    4: required i64 endAt
}

struct CreateTaskResponse{
    1: model.BaseResp resp
    2: model.Task data
}
struct UpdateTaskResponse{
    1: model.BaseResp resp
}
struct UpdateTaskRequest{
    1: required i64 id
    2: required i16 status
}
struct DeleteTaskResponse{
    1: model.BaseResp resp
}
struct DeleteTaskRequest{
    1: required i64 id
}
struct DeleteTaskByStatusRequest{
    1: required i16 status
}
struct DeleteTaskByStatusResponse{
    1: model.BaseResp resp
}


struct QueryTaskListByStatusRequest{
    1: required i64 pageSize
    2: required i64 pageNum
    3: required i16 status
}

struct QueryTaskListByStatusResponse{
    1: model.BaseResp resp
    2: model.TaskList data
}

struct QueryTaskListByKeywordRequest{
    1: required i64 pageSize
    2: required i64 pageNum
    3: required string keyword
}

struct QueryTaskListByKeywordResponse{
    1: model.BaseResp resp
    2: model.TaskList data
}


service TaskService{
    CreateTaskResponse CreateTask(1: CreateTaskRequest req)(api.post="/task/create")
    QueryTaskListByStatusResponse QueryTaskByStatus(1: QueryTaskListByStatusRequest req)(api.get="/task/query/status")
    QueryTaskListByKeywordResponse QueryTaskByKeyword(1: QueryTaskListByKeywordRequest req)(api.get="/task/query/keyword")
    UpdateTaskResponse UpdateTask(1: UpdateTaskRequest req)(api.put="/task/update")
    DeleteTaskResponse DeleteOneTask(1: DeleteTaskRequest req)(api.delete="/task/delete")
    DeleteTaskByStatusResponse DeleteTaskByStatus(1: DeleteTaskByStatusRequest req)(api.delete="/task/delete/status")
}

