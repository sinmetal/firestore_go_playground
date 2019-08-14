package main

import (
	"context"
	"fmt"
	"time"
)

type Row struct {
	User      string
	Favos     []string
	CreatedAt time.Time
}

func main() {
	shardPlaygroundStart()
	//fmt.Println("start")
	//
	//project := flag.String("project", "hogeproject", "project")
	//flag.Parse()
	//fmt.Printf("project=%s\n", *project)
	//
	//ctx := context.Background()
	//ds, err := NewDatastoreClient(ctx, *project)
	//if err != nil {
	//	panic(err)
	//}
	//_, err = ds.Put(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fs, err := NewFirestoreClient(ctx, *project)
	//if err != nil {
	//	panic(err)
	//}
	//_, err = fs.Set(ctx, "FirestoreRoot1")
	//if err != nil {
	//	panic(err)
	//}
	//_, err = fs.Set(ctx, "FirestoreRoot2")
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("Query1")
	//fs.Query1(ctx, "FirestoreSample")
	//
	//fmt.Println("Query2")
	//fs.Query2(ctx, "SubCol")
	//
	//fmt.Println("Query3")
	//fs.Query3(ctx, "SubCol")
	//
	//fmt.Println("end")
}

func shardPlaygroundStart() {
	ctx := context.Background()
	if err := ShardPlayground(ctx); err != nil {
		fmt.Printf("%+v", err)
	}
}
