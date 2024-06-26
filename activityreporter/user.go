package activityreporter

import (
	"fmt"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/tsanaativa-vinnera/assignment-activity-reporter/customerror"
)

type User struct {
	Username string

	hasUploadedPhoto bool
	likedByList      []User

	followerList   []Observer
	activityReport []string

	socialGraph *SocialGraph
}

func NewUser(Username string, socialGraph *SocialGraph) *User {
	return &User{
		Username:    Username,
		socialGraph: socialGraph,
	}
}

func (u *User) FollowedBy(following *User) error {
	if u.Username == following.Username {
		return customerror.ErrFollowThemselves
	}

	if !u.isFollowedBy(*following) {
		u.Register(following)
		return nil
	}

	return customerror.ErrAlreadyFollowed
}

func (u *User) UploadPhoto() error {
	if !u.hasUploadedPhoto {
		u.hasUploadedPhoto = true

		u.logActivity("You uploaded photo")

		notification := fmt.Sprintf("%s uploaded photo", u.Username)
		u.Notify(notification)

		return nil
	}
	return customerror.ErrAlreadyUploaded
}

func (u *User) LikedPhotoBy(liker *User) error {
	if u.hasUploadedPhoto {

		if u.isFollowedBy(*liker) || u.isEqualTo(*liker) {

			if !u.isAlreadyLikedBy(*liker) {
				u.likedByList = append(u.likedByList, *liker)

				u.updateLikeActivity(liker)

				u.socialGraph.UpdateTrending(u)

				return nil
			}

			return customerror.ErrAlreadyLiked
		}

		return customerror.ErrUnableToLike(u.Username)
	}

	return customerror.ErrPhotoDoesntExist(u.Username, u.isEqualTo(*liker))
}

func (u *User) updateLikeActivity(liker *User) {
	if u.isEqualTo(*liker) {
		log := "You liked your photo"
		liker.logActivity(log)

	} else {
		likerLog := fmt.Sprintf("You liked %s's photo", u.Username)
		liker.logActivity(likerLog)

		log := fmt.Sprintf("%s liked your photo", liker.Username)
		u.logActivity(log)
	}

	notification := fmt.Sprintf("%s liked %s's photo", liker.Username, u.Username)
	liker.Notify(notification)
}

func (u *User) isEqualTo(otherUser User) bool {
	return u.Username == otherUser.Username
}

func (u *User) isAlreadyLikedBy(otherUser User) bool {
	for _, v := range u.likedByList {
		if v.isEqualTo(otherUser) {
			return true
		}
	}
	return false
}

func (u *User) LikesCount() int {
	return len(u.likedByList)
}

func (u *User) isFollowedBy(followed User) bool {
	for _, v := range u.followerList {
		user := v.(*User)

		if user.Username == followed.Username {
			return true
		}
	}
	return false
}

func (u *User) logActivity(log string) {
	u.activityReport = append(u.activityReport, log)
}

func (u *User) ActivityReport() []string {
	return u.activityReport
}

func (u *User) Register(observer Observer) {
	u.followerList = append(u.followerList, observer)
}

func (u *User) Notify(notification string) {
	notifSlice := strings.Fields(notification)
	isLike := notifSlice[1] == "liked" && notifSlice[3] == "photo"

	var liked string
	if isLike {
		liked = strings.Split(notifSlice[2], "'")[0]
	}

	for _, observer := range u.followerList {
		obvUser := observer.(*User)

		if !isLike || !(obvUser.isFollowedBy(*u) && liked == obvUser.Username) {
			observer.OnNotify(notification)
		}

	}
}

func (u *User) OnNotify(notification string) {
	u.logActivity(notification)
}
