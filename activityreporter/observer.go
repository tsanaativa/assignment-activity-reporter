package activityreporter

type Observer interface {
	OnNotify(s Subject)
}