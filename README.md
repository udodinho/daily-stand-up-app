# Daily-Stand-up-App

### Usage

1. Clone repo
2. Run the command on your terminal
```go
go mod tidy
``` 

### Create Update/Checkin (POST)

```
localhost:8080/checkin
```
Payload
```
{
    "work_done_yesterday": "I am done with a feature",
    "work_to_do_today": "Going to implement anoda",
    "blocked_by": "Josh",
    "breakaway": "I have made some research",
    "sprint": "one",
    "status": "within standup",
    "task_id": "avg-322",
    "date": "2024-04-06"
}
```

### Get Checkins/Updates or with filter(weekStart{2024-02-02 00:00:00}, weekEnd{2024-02-07 00:00:00}, sprint, date, owner, date) (GET)
Endpoint
```
localhost:8080/checkin?owner=James
```

## Tests
Testing is done using the GoMock framework. The ``gomock`` package and the ``mockgen``code generation tool are used for this purpose.
Installation can be done using the following commands:
```
go get go.uber.org/mock/gomock@latest
go install go.uber.org/mock/gomock
mockgen -source=domain/repository.go -destination=test/mock_db.go -package=test
```
run all the test files using:
```testing
go test -v ./...
```