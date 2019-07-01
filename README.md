
# query_report

- Easy measurement of query results.
- Summarize the results of regularly executed SQL in CSV.
- Output specific column as a header and specific column as a value in CSV format.

## example

```sql
SELECT user_id, try_count FROM user
```

### 1st exec

| user_id | try_count |
| --- | --- |
| 1 | 0 |
| 2 | 1 |
| 3 | 9 |

### 2nd exec, (Run after 5 seconds. set `interval option`) 

| user_id | try_count |
| --- | --- |
| 1 | 0 |
| 2 | 1 |
| 3 | 12 |

### Output

| datetime | 1 | 2 | 3 |
| --- | --- | --- | --- |
| 2019/07/01 20:18:05 | 0 | 1 | 9 |
| 2019/07/01 20:22:05 | 0 | 1 | 12 |

### example command

```sh
query_report -user=$DB_USER -pass=$DB_PASS -host=$DB_HOST -db=$DB_NAME -query="SELECT user_id, try_count FROM user" -key=user_id -value=try_count -interval=5
```

### parameter

- all require

| parameter | content |
| --- | --- |
| user | Database user |
| pass | Password of database user |
| host | Database host |
| db | Use database | 
| query | Query to execute |
| key | Column use as header |
| value | Column use as value |
| interval | Execution interval (sec) | 
