package link

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/dgraph-io/badger/v2"
)

// Link Object
type Link struct {
	Hash      string    `json:"hash"`
	UserEmail string    `json:"user_email"`
	Orginal   string    `json:"orginal"`
	Visited   int       `json:"visited"`
	CreatedAt time.Time `json:"created_at"`
	ExpireAt  time.Time `json:"expire_at"`
}

func (l Link) Shorten(domain string) string {
	return domain + "/" + l.Hash
}

func Create(txn *badger.Txn, orginal, userEmail string, expire time.Time) (*Link, error) {
	now := time.Now()
	l := &Link{
		Hash:      Encode(uint64(now.UnixNano())),
		Orginal:   orginal,
		UserEmail: userEmail,
		ExpireAt:  expire,
		CreatedAt: now,
	}
	// check hash is unique
	if _, err := Load(txn, l.Hash, false); err == nil {
		return Create(txn, orginal, userEmail, expire)
	}
	if err := l.Save(txn); err != nil {
		return nil, err
	}
	return l, nil
}

func (l *Link) Save(txn *badger.Txn) error {
	data, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return txn.SetEntry(&badger.Entry{
		Key:       []byte(l.Hash),
		Value:     data,
		ExpiresAt: uint64(l.ExpireAt.Unix()),
	})
}

func Load(txn *badger.Txn, hash string, checkExpires bool) (*Link, error) {
	l := new(Link)
	item, err := txn.Get([]byte(hash))
	if err != nil {
		return nil, err
	}
	if item.IsDeletedOrExpired() && checkExpires {
		return nil, errors.New("not matched")
	}
	var data []byte
	result, err := item.ValueCopy(data)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(result, l); err != nil {
		return nil, err
	}
	return l, nil
}
