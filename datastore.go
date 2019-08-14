package main

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/datastore"
)

type DatastoreClient struct {
	C *datastore.Client
}

func NewDatastoreClient(ctx context.Context, projectID string) (*DatastoreClient, error) {
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil, err
	}
	return &DatastoreClient{
		C: client,
	}, nil
}

func (c *DatastoreClient) Put(ctx context.Context) (*datastore.Key, error) {
	rk := datastore.NameKey("DatastoreRoot", "Root", nil)

	k := datastore.NameKey("DatastoreSample", "Sample", rk)

	r := &Row{
		User:      "sinmetal",
		Favos:     []string{"Datastore", "Firestore"},
		CreatedAt: time.Now(),
	}
	return c.C.Put(ctx, k, r)
}
