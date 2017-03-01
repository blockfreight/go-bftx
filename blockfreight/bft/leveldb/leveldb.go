package leveldb

import (
	"fmt"
	//"log"
	
	"github.com/syndtr/goleveldb/leveldb"
)

func HandleError(err error, place string){
	if err != nil {
        //log.Fatal(place)
        //log.Fatal(err)
        fmt.Println(place, err)
    }
}

func OpenDB(db_path string) (db *leveldb.DB, err error) {
	fmt.Println("Creating / Opening leveldb db...")
	db, err = leveldb.OpenFile(db_path, nil)
	return db, err

}

func CloseDB(db *leveldb.DB) {
	db.Close()	//Use defer
}

func InsertBFTX(key string, value string, db *leveldb.DB) (error){
	//return err
	//fmt.Println("key", key)
	return db.Put([]byte(key), []byte(value), nil)
}

func Iterate(db *leveldb.DB) (n int, err error){
	iter := db.NewIterator(nil, nil)
	n = 0
	for iter.Next() {
	    /*key := iter.Key()
	    value := iter.Value()
		fmt.Println("\nKey: "+string(key)+"\nValue: "+string(value)+"\n")*/
		n+=1
	}
	iter.Release()
	return n, iter.Error()
}