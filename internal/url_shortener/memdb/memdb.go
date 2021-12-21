package inmemory

import (
	"github.com/hashicorp/go-memdb"
)

func InitMemDB() (*memdb.MemDB, error) {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"urls": &memdb.TableSchema{
				Name: "urls",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ShortUrl"},
					},
					"longUrl": &memdb.IndexSchema{
						Name:    "longUrl",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "LongUrl"},
					},
					"expiresAt": &memdb.IndexSchema{
						Name: "expiresAt",
						Unique: false,
						Indexer: &memdb.StringFieldIndex{Field: "ExpiresAt"},
					},
				},
			},
		},
	}

	return memdb.NewMemDB(schema)
}