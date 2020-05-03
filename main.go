package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"cloud.google.com/go/bigquery"
)

type Item struct {
	Time    time.Time
	Address string
}

func InsertToBigQuery(ctx context.Context, client *bigquery.Client, config *Config, time time.Time, address net.IP) error {
	uploader := client.Dataset(config.BigQueryDataset).Table(config.BigQueryTable).Uploader()
	items := []*Item{
		{Time: time, Address: address.String()},
	}
	return uploader.Put(ctx, items)
}

func main() {
	ctx := context.Background()

	config, err := ParseOptions()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse command line arguments: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Configuration: %+v\n", config)

	ipClient := NewAWSCheckIPClient()
	now := time.Now()
	addr, err := ipClient.Get(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get a public global IPv4 address: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Obtained the current global IPv4 address %v\n", addr)

	client, err := bigquery.NewClient(ctx, config.BigQueryProject)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create a BigQuery client: %v\n", err)
		os.Exit(1)
	}
	if err := InsertToBigQuery(ctx, client, config, now, addr); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert rows to the BigQuery table: %v\n", err)
		os.Exit(1)
	}
}
