package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"cloud.google.com/go/firestore"
)

type ShardChat struct {
	Shard      int
	CreateTime time.Time
}

func ShardPlayground(ctx context.Context) error {
	fs, err := firestore.NewClient(ctx, "firestore-native-sinmetal")
	if err != nil {
		return err
	}

	errCh := make(chan error, 10)
	go func() {
		t := time.NewTicker(3 * time.Second)
		for {
			select {
			case <-t.C:
				e := ShardChat{
					Shard:      rand.Intn(3),
					CreateTime: time.Now(),
				}
				_, _, err := fs.Collection("ShardChat").Add(ctx, &e)
				if err != nil {
					errCh <- err
				}
			}
		}
	}()

	// Inequlity filterとORDER BYには別のFieldを指定できない。従ってShardはEqulity Filterを指定して、Shardの数だけSyncを作成する。
	// Shardが異なるものはCreateTimeでソートされないので、自分でソートする必要がある。
	// You cannot specify different fields for Inequlity filter and ORDER BY. Therefore, `Shard` specifies Equlity Filter and creates Sync for the number of Shards.
	// Items with different Shards are not sorted by CreateTime, so you have to sort them yourself.
	for _, shard := range []int{0, 1, 2, 3} {
		shard := shard
		go func(shard int) {
			iter := fs.Collection("ShardChat").Where("Shard", "==", shard).OrderBy("CreateTime", firestore.Desc).Snapshots(ctx)
			defer iter.Stop()
			for {
				docIter, err := iter.Next()
				if err != nil {
					errCh <- err
				}
				for _, dc := range docIter.Changes {
					fmt.Printf("ID:%v, Shard:%v, CreateTime:%v\n", dc.Doc.Ref.ID, shard, dc.Doc.CreateTime)
				}
			}
		}(shard)
	}

	err = <-errCh
	if err != nil {
		return err
	}

	return nil
}
