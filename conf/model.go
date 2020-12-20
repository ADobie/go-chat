package conf

type uid = string

const archiveSize = 20
const chanSize = 10

const msgJoin = "[加入房间]"
const msgLeave = "[离开房间]"
const msgTyping = "[正在输入]"

const(
	EventTypeJoin = "event-join"
	EventTypeMsg = "event-message"
	EventTypeLeave = "event-leave"
	EventTypeBroadcast = "event-broadcast"
	EventTypeSystem = "event-system"
	EventTypeImage  = "event-image"
)





