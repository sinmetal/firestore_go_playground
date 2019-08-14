package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type FirestoreClient struct {
	C *firestore.Client
}

func NewFirestoreClient(ctx context.Context, projectID string) (*FirestoreClient, error) {
	c, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil, err
	}
	return &FirestoreClient{
		C: c,
	}, nil
}

func (c *FirestoreClient) Set(ctx context.Context, collection string) (*firestore.WriteResult, error) {
	ref := c.C.Collection(collection).Doc("ParentDoc").Collection("SubCol").Doc("Sample")
	return ref.Set(ctx, &Row{
		User:      "sinmetal",
		Favos:     []string{"Datastore", "Firestore"},
		CreatedAt: time.Now(),
	})
}

func (c *FirestoreClient) Query1(ctx context.Context, collection string) {
	it := c.C.CollectionGroup(collection).WherePath([]string{
		"ParentDoc",
	}, "==", "ParentDoc").Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s: %s\n", doc.Ref.Parent.Path, doc.Ref.ID)
	}
}

func (c *FirestoreClient) Query2(ctx context.Context, collection string) {
	it := c.C.CollectionGroup(collection).Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s: %s\n", doc.Ref.Parent.Path, doc.Ref.ID)
	}
}

func (c *FirestoreClient) Query3(ctx context.Context, collection string) {
	it := c.C.CollectionGroup(collection).Where("User", "==", "sinmetal").Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s: %s\n", doc.Ref.Parent.Path, doc.Ref.ID)
	}
}

//client, err := firestore.NewClient(ctx, *project)
//defer client.Close()
//if err != nil {
//	panic(err)
//}
//
//iter := client.Collection("world-default-player-position").Snapshots(ctx)
//defer iter.Stop()
//for {
//	dociter, err := iter.Next()
//	if err != nil {
//		panic(err)
//	}
//	dslist, err := dociter.GetAll()
//	if err != nil {
//		panic(err)
//	}
//	for _, v := range dslist {
//		fmt.Printf("%+v", v.Data())
//	}
//}
