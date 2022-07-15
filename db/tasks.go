package db

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key   uint
	Value string
}

func InitDB() error {
	var err error
	db, err = bolt.Open("tasks.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func CreateTask(task string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id, _ := b.NextSequence()
		key := itob(int(id))
		return b.Put(key, []byte(task))
	})
}

func ReadTasklist() ([]Task, error) {
	var task []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		if b == nil {
			return fmt.Errorf("bucket does not exist")
		}
		b.ForEach(func(k, v []byte) error {
			//			fmt.Println("key=", btoi(k), "value=", string(v))
			task = append(task, Task{
				Key:   uint(btoi(k)),
				Value: string(v),
			})
			return nil
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return task, nil
}

func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(taskBucket))
		return bucket.Delete(itob(key))
	})
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(v []byte) int64 {
	return int64(binary.BigEndian.Uint64(v))
}
