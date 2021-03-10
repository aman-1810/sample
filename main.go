package main

import (
	"fmt"
	"regexp"
	"time"

	memdb "github.com/hashicorp/go-memdb"
	gocache "github.com/patrickmn/go-cache"
)

var cache = gocache.New(-1, -1)

func main() {
	memoryDbImplementation()
	fmt.Println("Normal Matching")
	// fmt.Println(time.Now())
	t1 := time.Now()
	for i := 0; i < 100000; i++ {
		_, _ = regexp.Match("/^[a-zA-Z0-9!@#$%^&*()_+\\-=\\[\\]{};':\\|,.<>\\/?]*$/", []byte("abyyuyiuiuiuiu"))
	}
	fmt.Println(time.Since(t1))
	// re, _ :=
	// fmt.Println(re)
	// fmt.Println(time.Now())

	type P struct {
		X, Y, Z int
		Name    string
	}

	type Q struct {
		X, Y *int32
		Name string
	}

	fmt.Println("With Compiler Matching")
	tx, _ := regexp.Compile("/^[a-zA-Z0-9!@#$%^&*()_+\\-=\\[\\]{};':\\|,.<>\\/?]*$/")
	cache.Set("templateReg", tx, -1)

	t2 := time.Now()
	for i := 0; i < 100000; i++ {
		txCache, _ := cache.Get("templateReg")
		txCacheReg := txCache.(*regexp.Regexp)
		txCacheReg.Match([]byte("abyyuyiuiuiuiu"))
	}
	fmt.Println(time.Since(t2))

	// var network bytes.Buffer        // Stand-in for a network connection
	// enc := gob.NewEncoder(&network) // Will write to network.
	// dec := gob.NewDecoder(&network) // Will read from network.
	// // Encode (send) the value.
	// err := enc.Encode(tx)
	// if err != nil {
	// 	log.Fatal("encode error:", err)
	// }

	// // HERE ARE YOUR BYTES!!!!
	// fmt.Println(network.Bytes())

	// // Decode (receive) the value.
	// var sam *regexp.Regexp
	// err = dec.Decode(&sam)
	// if err != nil {
	// 	log.Fatal("decode error:", err)
	// }
	// fmt.Println(sam.Match())
	// fmt.Println(time.Now())

	// fmt.Println(another.IsPhoneNumber("6556787768", "IN"))

	// fmt.Println(numbers.IsPrime(455))

	// fmt.Println(strings.Reverse("Cucumber"))
}

func memoryDbImplementation() {

	type Template struct {
		ID   int
		Desc interface{}
	}

	// Create a compiled regex
	tx, _ := regexp.Compile("/^[a-zA-Z0-9!@#$%^&*()_+\\-=\\[\\]{};':\\|,.<>\\/?]*$/")

	// Create the DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"template": &memdb.TableSchema{
				Name: "template",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "ID"},
					},
					"desc": &memdb.IndexSchema{
						Name:    "desc",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Desc"},
					},
				},
			},
		},
	}

	// Create a new data base
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	// Create a write transaction
	txn := db.Txn(true)
	var templates []*Template
	for i := 1; i <= 100000; i++ {
		templates = append(templates, &Template{i, tx})
	}

	for _, p := range templates {
		if err := txn.Insert("template", p); err != nil {
			panic(err)
		}
	}

	// Commit the transaction
	txn.Commit()

	// Create read-only transaction
	txn = db.Txn(false)
	defer txn.Abort()

	// List all the people
	// matchExpr, _ := txn.Get("template", "desc")
	// txCacheReg := txCache.(*regexp.Regexp)

	it, err := txn.Get("template", "desc")
	if err != nil {
		panic(err)
	}

	fmt.Println("Memdb Transaction")
	t1 := time.Now()
	var cnt = 0
	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*Template)
		txMemReg := p.Desc.(*regexp.Regexp)
		txMemReg.Match([]byte("abyyuyiuiuiuiu"))
		cnt++
	}

	fmt.Println(time.Since(t1))
	fmt.Println(cnt)

}
