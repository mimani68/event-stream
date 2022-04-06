package in_memory_db

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"zarinworld.ir/event/pkg/log_handler"
)

var db []DatabaseDocument

type DatabaseDocument struct {
	Id    string `json:"id,omitempty"`
	Key   string `json:"key"`
	Value string `json:"value"`
	Time  string `json:"time"`
}

func Set(key string, object interface{}) (success bool, err error) {
	a := DatabaseDocument{}
	a.Id = uuid.New().String()
	a.Key = key
	a.Value = object.(string)
	a.Time = time.Now().Format(time.RFC3339)
	db = append(db, a)
	return true, nil
}

func Get(key string) (data string, err error) {
	err = errors.New("key " + key + " was not exists.")
	for _, record := range db {
		if record.Key == key {
			data = record.Value
			err = nil
		}
	}
	return data, err
}

func KeyGenerator(source string, label string, trx string) string {
	if trx == "" {
		log_handler.LoggerF("Generating Key for inMemoryDb Failed beacse trx==\"\"")
	}
	if source == "" {
		log_handler.LoggerF("Generating Key for inMemoryDb Failed beacse source==\"\"")
	}
	return source + "-" + label + "-" + trx
}
