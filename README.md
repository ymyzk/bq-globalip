# bq-globalip
Record the current global IPv4 address to a BigQuery table.

## Preparation
Create a BigQuery table with the following schema:
```json
[
  {
    "name": "time",
    "type": "TIMESTAMP",
    "mode": "REQUIRED"
  },
  {
    "name": "address",
    "type": "STRING",
    "mode": "REQUIRED"
  }
]
```

## Usage
```shell
$ export GOOGLE_APPLICATION_CREDENTIALS=<your key>.json
$ bq-globalip -bigquery <project>:<dataset>.<table>
```

If you want to periodically record an address, please run `bq-globalip` with cron or a scheduler you prefer.
