package main

import (
	"context"
	"flag"
	"fmt"

	"cloud.google.com/go/firestore"
)

func main() {
	fmt.Println("start")

	project := flag.String("project", "hogeproject", "project")
	flag.Parse()
	fmt.Printf("project=%s", *project)

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, *project)
	defer client.Close()
	if err != nil {
		panic(err)
	}

	iter := client.Collection("world-default-player-position").Snapshots(ctx)
	defer iter.Stop()
	for {
		dociter, err := iter.Next()
		if err != nil {
			panic(err)
		}
		dslist, err := dociter.GetAll()
		if err != nil {
			panic(err)
		}
		for _, v := range dslist {
			fmt.Printf("%+v", v.Data())
		}
	}

	fmt.Println("end")
}
