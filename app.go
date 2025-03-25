package main

import (
	"context"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

// App struct
type App struct {
	ctx context.Context
}

func OpenDb() *bolt.DB {
	dbname := "mydb.db"
	db, err := bolt.Open(dbname, 0755, nil)
	if err != nil {
		log.Fatalf("Error opening db: %s", err.Error())
	}
	fmt.Println("** DB opened **")
	return db
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetValue(bucket string, key string) string {
	var keyvalue string
	fmt.Printf("\n==> Get value...")
	db := OpenDb()
	defer db.Close()

	fmt.Printf("\nLooking up value for key: [%s] in [%s]", key, bucket)

	// if err := db.Update(func(tx *bolt.Tx) error {
	// 	_, err := tx.CreateBucketIfNotExists([]byte(bucket))
	// 	if err != nil {
	// 		log.Fatalf("Error: CreateBucketIfNotExists: [%s]\n", err.Error())
	// 	}
	// 	return err
	// }); err != nil {
	// 	log.Fatalf("***>> Error creating bucket: [%s] ==> %s\n", bucket, err.Error())
	// }
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		fmt.Println("bucket: ", b)
		keyvalue = string(b.Get([]byte(key)))
		fmt.Println("\nValue of key [", key, "] = [", keyvalue, "]")
		return nil
	}); err != nil {
		log.Fatalf("\n*** Error looking up key: %s ==> %s\n", key, err.Error())
	}
	return keyvalue
}

func (a *App) AddKeyValue(bucket string, key string, value string) {
	db := OpenDb()
	if err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		err := b.Put([]byte(key), []byte(value))
		if err != nil {
			fmt.Printf("??? Error adding key, value: %s", err.Error())
		}
		return err
	}); err != nil {
		log.Fatalf("<<< Error adding key=[%s], value=[%s] ==> %s >>>\n", key, value, err.Error())
	}
	fmt.Printf("\nKey: %s Value: %s Added", key, value)
}

func (a *App) CreateBucket(bucket string) bool {
	returnval := true
	OpenDb()

	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			fmt.Printf("error creating bucket: %s", err.Error())
			returnval = false
		}
		if err := tx.Commit(); err != nil {
			fmt.Printf("**** error commiting bucket: %s ****", err.Error())
			returnval = false
		}
		fmt.Println("==> create bucket: ", b)
		return nil
	})
	fmt.Println("created bucket ", bucket)
	return returnval
}
