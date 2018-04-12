package rocks

import (
	"errors"
	"fmt"

	"github.com/tecbot/gorocksdb"
)

type Reader interface {
	GetBlockHeader(key []byte) ([]byte, error)
	OpenDatabaseForReadOnlyColumnFamilies(name string) error
}

// wraps gorocksdb calls
type RDBReader struct {
	rdb  *gorocksdb.DB
	cfhs []*gorocksdb.ColumnFamilyHandle
}

func (rdbr *RDBReader) GetBlockHeader(key []byte) ([]byte, error) {
	readOptions := gorocksdb.NewDefaultReadOptions()
	defer readOptions.Destroy()
	data, err := rdbr.rdb.GetCF(readOptions, rdbr.cfhs[2], key)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not read from Rocks DB: %s", err.Error()))
	}
	toReturn := make([]byte, len(data.Data()))
	copy(toReturn, data.Data())
	data.Free()
	return toReturn, nil
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
