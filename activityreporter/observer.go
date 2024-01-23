package activityreporter

type Observer interface {
	OnNotify(notification string)
}
