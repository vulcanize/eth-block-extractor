package rocks

import (
	"errors"
	"fmt"

	"github.com/tecbot/gorocksdb"
)

const (
	ColHeadersIndex = 2
	ColExtraIndex   = 4
)

type Reader interface {
	GetBlockHash(key []byte) ([]byte, error)
	GetBlockHeader(key []byte) ([]byte, error)
	OpenDatabaseForReadOnlyColumnFamilies(name string) error
}

type RDBReader struct {
	rdb  *gorocksdb.DB
	cfhs []*gorocksdb.ColumnFamilyHandle
}

func (rdbr *RDBReader) GetBlockHash(key []byte) ([]byte, error) {
	return rdbr.GetFromDB(rdbr.cfhs[ColExtraIndex], key)
}

func (rdbr *RDBReader) GetBlockHeader(key []byte) ([]byte, error) {
	return rdbr.GetFromDB(rdbr.cfhs[ColHeadersIndex], key)
}

func (rdbr *RDBReader) GetFromDB(column *gorocksdb.ColumnFamilyHandle, key []byte) ([]byte, error) {
	readOptions := gorocksdb.NewDefaultReadOptions()
	defer readOptions.Destroy()
	data, err := rdbr.rdb.GetCF(readOptions, column, key)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not read from Rocks DB: %s", err.Error()))
	}
	result := make([]byte, len(data.Data()))
	copy(result, data.Data())
	data.Free()
	return result, nil
}

func (rdbr *RDBReader) OpenDatabaseForReadOnlyColumnFamilies(name string) error {
	options := gorocksdb.NewDefaultOptions()
	defer options.Destroy()
	fams, err := gorocksdb.ListColumnFamilies(options, name)
	var opts []*gorocksdb.Options
	for i := 0; i < len(fams); i++ {
		options := gorocksdb.NewDefaultOptions()
		defer options.Destroy()
		opts = append(opts, options)
	}
	rocksDb, cfhs, err := gorocksdb.OpenDbForReadOnlyColumnFamilies(options, name, fams, opts, false)
	if err != nil {
		return err
	}
	rdbr.rdb = rocksDb
	rdbr.cfhs = cfhs
	return nil
}
