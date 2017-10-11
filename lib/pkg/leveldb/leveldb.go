// File: ./blockfreight/lib/leveldb/leveldb.go
// Summary: Application code for Blockfreight™ | The blockchain of global freight.
// License: MIT License
// Company: Blockfreight, Inc.
// Author: Julian Nunez, Neil Tran, Julian Smith, Gian Felipe & contributors
// Site: https://blockfreight.com
// Support: <support@blockfreight.com>

// Copyright © 2017 Blockfreight, Inc. All Rights Reserved.

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

// Package leveldb provides some useful functions to work with LevelDB.
// It has common database functions as OpenDB, CloseDB, Insert and Iterate.
package leveldb

import (
	// =======================
	// Golang Standard library
	// =======================
	"encoding/json" // Implements encoding and decoding of JSON as defined in RFC 4627.
	"errors"        // Implements functions to manipulate errors.

	// ====================
	// Third-party packages
	// ====================
	"github.com/syndtr/goleveldb/leveldb" // Implementation of the LevelDB key/value database in the Go programming language.

	// ======================
	// Blockfreight™ packages
	// ======================
	"github.com/blockfreight/go-bftx/lib/app/bf_tx" // Defines the Blockfreight™ Transaction (BF_TX) transaction standard and provides some useful functions to work with the BF_TX.
)

var dbPath = "bft-db" //Folder name where is going to be the LevelDB

// OpenDB is a function that receives the path of the DB, creates or opens that DB and return ir with a possible error if that occurred.
func OpenDB(dbPath string) (db *leveldb.DB, err error) {
	db, err = leveldb.OpenFile(dbPath, nil)
	return db, err

}

// CloseDB is a function that receives a DB pointer that closes the connection to DB.
func CloseDB(db *leveldb.DB) {
	db.Close()
}

// InsertBFTX is a function that receives the key and value strings to insert a tuple in determined DB, the final parameter. As result, it returns a true or false bool.
func InsertBFTX(key string, value string, db *leveldb.DB) error {
	return db.Put([]byte(key), []byte(value), nil)
}

// Total is a function that returns the total of BF_TX stored in the DB.
func Total() (int, error) {
	db, err := OpenDB(dbPath)
	defer CloseDB(db)
	if err != nil {
		return 0, err
	}

	iter := db.NewIterator(nil, nil)
	n := 0
	for iter.Next() {
		n += 1
	}
	iter.Release()
	return n, iter.Error()
}

// RecordOnDB is a function that receives the content of the BF_RX JSON to insert it into the DB and return true or false according to the result.
func RecordOnDB(id string, json string) error {
	db, err := OpenDB(dbPath)
	defer CloseDB(db)
	if err != nil {
		return err
	}
	err = InsertBFTX(id, json, db)
	if err != nil {
		return err
	}
	return nil
}

// GetBfTx is a function that receives a bf_tx id, and returns the BF_TX if it exists.
func GetBfTx(id string) (bf_tx.BF_TX, error) {
	var bftx bf_tx.BF_TX
	db, err := OpenDB(dbPath)
	defer CloseDB(db)
	if err != nil {
		return bftx, err
	}

	data, err := db.Get([]byte(id), nil)
	if err != nil {
		if err.Error() == "leveldb: not found" {
			return bftx, errors.New("LevelDB Get function: BF_TX not found.")
		}
		return bftx, errors.New("LevelDB Get function: " + err.Error())
	}

	json.Unmarshal(data, &bftx)
	return bftx, nil
}

// Verify is a function that receives a content and look for a BF_TX that has the same content.
func Verify(jcontent string) ([]byte, error) {
	var bftx bf_tx.BF_TX
	db, err := OpenDB(dbPath)
	defer CloseDB(db)
	if err != nil {
		return nil, err
	}

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()

		// Get a BF_TX by id
		json.Unmarshal(value, &bftx)

		// Reinitialize the BF_TX
		bftx = bf_tx.Reinitialize(bftx)

		// Get the BF_TX old_content in string format
		content, err := bf_tx.BF_TXContent(bftx)
		if err != nil {
			return nil, err
		}

		if jcontent == content {
			iter.Release()
			//strconv.Atoi(string(buf))
			return key, nil
		}
	}
	iter.Release()

	return nil, iter.Error()
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
