package model

import "github.com/duolacloud/seata-golang/pkg/base/meta"

// SessionCondition for query GlobalSession
type SessionCondition struct {
	TransactionId      int64
	Xid                string
	Status             meta.GlobalStatus
	Statuses           []meta.GlobalStatus
	OverTimeAliveMills int64
}
