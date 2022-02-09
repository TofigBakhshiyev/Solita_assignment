### Solita Assigment Backend
- CSV parsing and validation, `implemented`
- Endpoints to fetch data from farms with different granularities (by month, by metric), `implemented`, `by month is not implemented`
- Aggregate calculation endpoints, endpoint which returns monthly averages, min/max and other statistical analysis, `implemented`, `by month is not implemented`
- Add tests, `implemented`
- Input and output validation, `implemented`
- Database connectio, `implemented`

### Requirements
- go verison 1.17.4
- Docker

### Run the project
- ```go test```
- ```go run main.go```

### Server is running
- ```farms/min/:sensortype```
- ```farms/max/:sensortype```

### Requirements for PostgrSQL
```docker run -d -p 5432:5432 --name postgresql_test -e POSTGRES_PASSWORD=test postgres```
```docker exec -it postgresql_test bash```
```psql -U postgres```
```CREATE DATABASE solitadb```

    ```CREATE TABLE farms (ID UUID DEFAULT gen_random_uuid ()   PRIMARY KEY, Location VARCHAR ( 50 ) UNIQUE NOT NULL, Datetime timestamp default NULL,
    SensorType VARCHAR ( 50 ) UNIQUE NOT NULL,
    Value float
    );```