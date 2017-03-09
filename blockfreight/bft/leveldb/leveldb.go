// File: ./blockfreight/leveldb/leveldb.go
// Summary: Application code for Blockfreight™ | The blockchain of global freight.
// License: MIT License
// Company: Blockfreight, Inc.
// Author: Julian Nunez, Neil Tran, Julian Smith & contributors
// Site: https://blockfreight.com
// Support: <support@blockfreight.com>

// Copyright 2017 Blockfreight, Inc.

// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// =================================================================================================================================================
// =================================================================================================================================================
//
// BBBBBBBBBBBb     lll                                kkk             ffff                         iii                  hhh            ttt
// BBBB``````BBBB   lll                                kkk            fff                           ```                  hhh            ttt
// BBBB      BBBB   lll      oooooo        ccccccc     kkk    kkkk  fffffff  rrr  rrr    eeeee      iii     gggggg ggg   hhh  hhhhh   tttttttt
// BBBBBBBBBBBB     lll    ooo    oooo    ccc    ccc   kkk   kkk    fffffff  rrrrrrrr eee    eeee   iii   gggg   ggggg   hhhh   hhhh  tttttttt
// BBBBBBBBBBBBBB   lll   ooo      ooo   ccc           kkkkkkk        fff    rrrr    eeeeeeeeeeeee  iii  gggg      ggg   hhh     hhh    ttt
// BBBB       BBB   lll   ooo      ooo   ccc           kkkk kkkk      fff    rrr     eeeeeeeeeeeee  iii   ggg      ggg   hhh     hhh    ttt
// BBBB      BBBB   lll   oooo    oooo   cccc    ccc   kkk   kkkk     fff    rrr      eee      eee  iii    ggg    gggg   hhh     hhh    tttt    ....
// BBBBBBBBBBBBB    lll     oooooooo       ccccccc     kkk     kkkk   fff    rrr       eeeeeeeee    iii     gggggg ggg   hhh     hhh     ttttt  ....
//                                                                                                        ggg      ggg
//   Blockfreight™ | The blockchain of global freight.                                                      ggggggggg
//
// =================================================================================================================================================
// =================================================================================================================================================

//Package leveldb provides some useful functions to work with any topic associated with LevelDB.
//It has common database functions as OpenDB, CloseDB, Insert and Iterate.
package leveldb

import (
	"fmt"
	"encoding/json"
	"strconv"

	"github.com/blockfreight/blockfreight-alpha/blockfreight/bft/bf_tx"
	"github.com/davecgh/go-spew/spew"
	"github.com/syndtr/goleveldb/leveldb"
)

var db_path string = "bft-db"

//HandleError is a function that receives an error and a name of the procedure where that error ocurred, and print a clear error message.
func HandleError(err error, place string) {
	if err != nil {
		//log.Fatal(place)
		//log.Fatal(err)
		fmt.Println(place, err)
	}
}

//OpenDB is a function that receives the path of the DB, creates or opens that DB and return ir with a possible error if that ocurred.
func OpenDB(db_path string) (db *leveldb.DB, err error) {
	fmt.Println("Creating / Opening leveldb db...")
	db, err = leveldb.OpenFile(db_path, nil)
	return db, err

}

//CloseDB is a function that receives a DB pointer that closes the connection to DB.
func CloseDB(db *leveldb.DB) {
	db.Close() //Use defer
}

//InsertBF_TX is a function that receives the key and value strings to insert a tuple in determined DB, the final parameter. As result, it returns a true or false bool. 
func InsertBF_TX(key string, value string, db *leveldb.DB) error {
	//return err
	//fmt.Println("key", key)
	return db.Put([]byte(key), []byte(value), nil)
}

//Iterate is a function that receives a DB pointer and check all single tuples in that DB. It returns, the total of tuples were found and an error if that happened.
func Iterate(db *leveldb.DB) (n int, err error) {
	iter := db.NewIterator(nil, nil)
	n = 0
	for iter.Next() {
		/*key := iter.Key()
		    value := iter.Value()
			fmt.Println("\nKey: "+string(key)+"\nValue: "+string(value)+"\n")*/
		n += 1
	}
	iter.Release()
	return n, iter.Error()
}

//RecordOnDB is a function that receives the content of the BF_RX JSON to insert it into the DB and return true or false according to the result.
func RecordOnDB( /*id string, */ json string) bool { //TODO: Check the id
	db, err := OpenDB(db_path)
	defer CloseDB(db)

	//Get the number of bf_tx on DB
	var n int
	n, err = Iterate(db)

	HandleError(err, "Create or Open Database")
	//fmt.Println("Database created / open on "+db_path)

	err = InsertBF_TX(strconv.Itoa(n+1), json, db)
	//err = InsertBF_TX(id, json, db)    //TODO: Check the id

	//Iteration
	n, err = Iterate(db)
	HandleError(err, "Iteration")
	fmt.Println("Total: " + strconv.Itoa(n))

	return true
}

//GetBfTx is a function that receives a bf_tx id, searches that bf_tx and returns its content.
func GetBfTx(id string) string{
	db, err := OpenDB(db_path)
	defer CloseDB(db)

	data, err := db.Get([]byte(id), nil)
	HandleError(err, "GetBfTx")
	var bf_tx bf_tx.BF_TX
	json.Unmarshal(data, &bf_tx)
	spew.Dump(bf_tx)
	
	return "Ok!"
}

// =================================================
// Blockfreight™ | The blockchain of global freight.
// =================================================

// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBB                    BBBBBBBBBBBBBBBBBBB
// BBBBBBB                       BBBBBBBBBBBBBBBB
// BBBBBBB                        BBBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBB        BBBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBB        BBBBBBBBBBBBBBB
// BBBBBBB       BBBBBBB         BBBBBBBBBBBBBBBB
// BBBBBBB                     BBBBBBBBBBBBBBBBBB
// BBBBBBB                        BBBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBBB        BBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBBBB       BBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBBB        BBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBB        BBB       BBBBB
// BBBBBBB                       BBBB       BBBBB
// BBBBBBB                    BBBBBBB       BBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB

// ==================================================
// Blockfreight™ | The blockchain for global freight.
// ==================================================
