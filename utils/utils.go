package utils

import (
	"sync"
	"time"
)

//Utils ...
type Utils interface {
	TimeNow() time.Time
}

type utils struct{}

var (
	utilsInstance Utils
	utilsOnce     sync.Once
)

//GetUtils ...
func GetUtils() Utils {
	utilsOnce.Do(func() {
		if utilsInstance == nil {
			utilsInstance = &utils{}
		}
	})
	return utilsInstance
}

//SetUtils ...
func SetUtils(instance Utils) {
	utilsInstance = instance
}

func (u *utils) TimeNow() time.Time {
	return time.Now()
}
