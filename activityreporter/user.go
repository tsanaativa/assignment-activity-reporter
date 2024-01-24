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

	if !u.IsFollowedBy(*following) {
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

		u.socialGraph.AddToTrending(u)

		return nil
	}
	return customerror.ErrAlreadyUploaded
}

func (u *User) LikedPhotoBy(liker *User) error {
	if u.hasUploadedPhoto {

		if u.IsFollowedBy(*liker) || u.isEqualTo(*liker) {

			if !u.isAlreadyLikedBy(*liker) {
				u.likedByList = append(u.likedByList, *liker)

				likerStr := liker.Username
				if u.isEqualTo(*liker) {
					likerStr = "You"

				} else {
					likerLog := fmt.Sprintf("You liked %s's photo", u.Username)
					liker.logActivity(likerLog)
				}

				log := fmt.Sprintf("%s liked your photo", likerStr)
				u.logActivity(log)

				if !liker.IsFollowedBy(*u) {
					notification := fmt.Sprintf("%s liked %s's photo", liker.Username, u.Username)
					liker.Notify(notification)
				}

				u.socialGraph.UpdateTrending(u)

				return nil
			}

			return customerror.ErrAlreadyLiked
		}

		return customerror.ErrUnableToLike(u.Username)
	}

	return customerror.ErrPhotoDoesntExist(u.Username, u.isEqualTo(*liker))
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

func (u *User) IsFollowedBy(followed User) bool {
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

func (u *User) ActivityLog() []string {
	return u.activityReport
}

func (u *User) Register(observer Observer) {
	u.followerList = append(u.followerList, observer)
}

func (u *User) Notify(notification string) {
	for _, observer := range u.followerList {
		observer.OnNotify(notification)
	}
}

func (u *User) OnNotify(notification string) {
	u.logActivity(notification)
}
