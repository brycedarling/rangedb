// Code generated by go generate; DO NOT EDIT.
// This file was generated at
// 2021-01-15 23:18:16.599746 -0800 PST m=+0.001241195
package chat

import "github.com/inklabs/rangedb"

func BindEvents(binder rangedb.EventBinder) {
	binder.Bind(
		&RoomWasOnBoarded{},
		&RoomWasJoined{},
		&MessageWasSentToRoom{},
		&PrivateMessageWasSentToRoom{},
		&UserWasRemovedFromRoom{},
		&UserWasBannedFromRoom{},
		&UserWasOnBoarded{},
		&UserWasWarned{},
	)
}
