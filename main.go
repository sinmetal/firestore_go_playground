package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go/types"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"io/ioutil"
)

type FirestoreListenForm struct {
	AddTarget    Target `json:"addTarget"`
	RemoveTarget int    `json:"removeTarget"`
}

type Target struct {
	TargetID int         `json:"targetId"`
	Once     bool        `json:"once"`
	Query    QueryTarget `json:"query"`
	//	Documents
	ResumeToken string `json:"resumeToken"`
	ReadTime    string `json:"readTime"`
}

type QueryTarget struct {
	Parent          string          `json:"parent"`
	StructuredQuery StructuredQuery `json:"structuredQuery"`
}

type StructuredQuery struct {
	Select Projection         `json:"select"`
	From   CollectionSelector `json:"from"`
	Where  Filter             `json:"where"`
}

type Projection struct {
	Fields []FieldReference `json:"fields"`
}

type FieldReference struct {
	FieldPath string `json:"fieldPath"`
}

type CollectionSelector struct {
	CollectionID   string `json:"collectionId"`
	AllDescendants bool   `json:"allDescendants"`
}

type Filter struct {
	CompositeFilter `json:"compositeFilter"`
}

type CompositeFilter struct {
	OP      string   `json:"op"`
	Filters []Filter `json:"filters"`
}

type Operator struct {
	Field FieldReference
}

type FieldFilter struct {
	Field FieldReference
	OP    string
	Value
}

type Value struct {
	NullValue      types.Nil
	BooleanValue   bool
	IntegerValue   int64
	DoubleValue    float64
	TimestampValue string
	StringValue    string
	BytesValue     string
	ReferenceValue string
	ArrayValue     []Value
}

type ArrayValue struct {
	Values []Value
}

func main() {
	c := context.Background()

	form := FirestoreListenForm{
		AddTarget: Target{
			TargetID: 1,
			Query: QueryTarget{
				Parent: "projects/sinmetal-firestore/databases/sinmetal-firestore",
			},
		},
	}
	j, err := json.Marshal(form)
	if err != nil {
		log.Fatal(err)
	}

	client, err := getClient(c, "https://www.googleapis.com/auth/cloud-platform", "https://www.googleapis.com/auth/datastore")
	if err != nil {
		log.Fatal(err)
	}
	r, err := client.Post("https://firestore.googleapis.com/v1beta1/projects/sinmetal-firestore/databases/default/:listen", "application/json", bytes.NewReader(j))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response Code = %d\n", r.StatusCode)

	//d := json.NewDecoder(r.Body)
	//defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("fatal read from response body. err = %s", err.Error())
	}

	fmt.Printf("response body = %s", b)

	// 人間に読みやすくなってるjsonを1行にするために、1回interface{]に読み込んで、jsonにし直す
	//var any interface{}
	//err = d.Decode(&any)
	//if err != nil {
	//	log.Fatalf("fatal decode from response body. %s", err.Error())
	//}

	//rb, err := json.Marshal(&any)
	//if err != nil {
	//	log.Fatalf("fatal json.Marshal from response body. %s", err.Error())
	//}
	//fmt.Printf("Response Body = %s\n", string(rb))
}

func getClient(c context.Context, scope ...string) (*http.Client, error) {
	client, err := google.DefaultClient(c, scope...)
	if err != nil {
		return nil, err
	}

	return client, nil
}
