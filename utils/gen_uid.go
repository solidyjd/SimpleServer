package utils

import (
	uuid "github.com/satori/go.uuid"
)

func GenUid() string {
	u1,_ := uuid.NewV4()
	//tcp.ConId = u1.String()
	Log.Infof("UUIDv4: %s\n", u1)
	return u1.String()
}
