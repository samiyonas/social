package store

import (
    "database/sql"
)

type Storage struct {
    User *UsersStore
    Post *PostsStore
}

func NewStorage(db *sql.DB) Storage {
    store := Storage{
        User: &UsersStore{db},
        Post: &PostsStore{db},
    }

    return store
}
