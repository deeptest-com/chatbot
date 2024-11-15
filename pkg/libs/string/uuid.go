package _str

import (
	"github.com/oklog/ulid/v2"
	"github.com/satori/go.uuid"
	"math/rand"
	"strings"
	"time"
)

func Uuid() string {
	uid := uuid.Must(uuid.NewV4()).String()
	return strings.Replace(uid, "-", "", -1)
}

func UuidWithSep() string {
	uid := uuid.Must(uuid.NewV4()).String()
	return uid
}

func Ulid() string {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	rand, _ := ulid.New(ms, entropy)

	ret := strings.ToLower(rand.String())
	ret = strings.Replace(ret, "-", "", -1)

	return ret
}
