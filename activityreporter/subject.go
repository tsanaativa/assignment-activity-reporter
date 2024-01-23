package activityreporter

type Subject interface {
	Register(observer Observer)
	Notify(notification string)
}
