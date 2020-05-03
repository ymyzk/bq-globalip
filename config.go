package main

import (
	"errors"
	"flag"
	"strings"
)

type Config struct {
	BigQueryProject string
	BigQueryDataset string
	BigQueryTable   string
}

func ParseOptions() (*Config, error) {
	bigquery := flag.String("bigquery", "", "BigQuery (project:dataset.table)")
	flag.Parse()

	if *bigquery == "" {
		return nil, errors.New("bigquery option is required")
	}

	bigqueryComponents := strings.FieldsFunc(*bigquery, func(r rune) bool {
		return r == ':' || r == '.'
	})
	if len(bigqueryComponents) != 3 {
		return nil, errors.New("invalid bigquery argument format")
	}

	return &Config{
		BigQueryProject: bigqueryComponents[0],
		BigQueryDataset: bigqueryComponents[1],
		BigQueryTable:   bigqueryComponents[2],
	}, nil
}
