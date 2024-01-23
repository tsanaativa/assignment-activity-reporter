package activityreporter

import (
	"fmt"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/tsanaativa-vinnera/assignment-activity-reporter/customerror"
)

type User struct {
	Username string

	hasUploadedPhoto bool
	likedByList      []User

	followerList   []Observer
	activityReport []string
}

func NewUser(Username string) User {
	return User{
		Username: Username,
	}
}

func (u *User) Follow(followed User) error {
	if u.Username == followed.Username {
		return customerror.ErrFollowThemselves
	}

	if !u.IsFollowing(followed) {
		followed.Register(u)
		return nil
	}

	return customerror.ErrAlreadyFollowed
}

func (u *User) UploadPhoto() {
	u.hasUploadedPhoto = true

	u.logActivity("You uploaded photo")

	notification := fmt.Sprintf("%s uploaded photo", u.Username)
	u.Notify(notification)
}

func (u *User) LikePhoto(liked User) error {
	if u.IsFollowing(liked) {
		liked.likedByList = append(liked.likedByList, *u)

		log := fmt.Sprintf("You liked %s's photo", liked.Username)
		u.logActivity(log)

		notification := fmt.Sprintf("%s liked %s's photo", u.Username, liked.Username)
		u.Notify(notification)

		return nil
	}
	return customerror.ErrUnableToLike(liked.Username)
}

func (u *User) LikesCount() int {
	return len(u.likedByList)
}

func (u *User) IsFollowing(followed User) bool {
	var observerU Observer = u
	for _, v := range followed.followerList {
		if v == observerU {
			return true
		}
	}
	return false
}

func (u *User) logActivity(log string) {
	u.activityReport = append(u.activityReport, log)
}

func (u *User) ActivityLog() []string {
	return u.activityReport
}

func (u *User) Register(observer Observer) {
	u.followerList = append(u.followerList, observer)
}

func (u *User) Notify(notification string) {
	for _, observer := range u.followerList {
		observer.OnNotify(u, notification)
	}
}

func (u *User) OnNotify(subject Subject, notification string) {
	u.activityReport = append(u.activityReport, notification)
}
