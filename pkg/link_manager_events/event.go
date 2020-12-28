package link_manager_events

import (
	om "github.com/zhangminghui6106/delinkcious-0.5/pkg/object_model"
)

type Event struct {
	EventType om.EventTypeEnum
	Username  string
	Link      *om.Link
}
