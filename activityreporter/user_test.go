package activityreporter_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/tsanaativa-vinnera/assignment-activity-reporter/activityreporter"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/tsanaativa-vinnera/assignment-activity-reporter/activityreporter/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/tsanaativa-vinnera/assignment-activity-reporter/customerror"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	t.Run("should return new user", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		username := "Alice"

		//when
		user := activityreporter.NewUser(username, socialGraph)

		//then
		assert.Equal(t, username, user.Username)
	})

	t.Run("should be able to follow user", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		user1 := activityreporter.NewUser("Alice", socialGraph)
		user2 := activityreporter.NewUser("Bob", socialGraph)

		//when
		err := user2.FollowedBy(user1)

		//then
		assert.Nil(t, err)
	})

	t.Run("should return error when user follow themselves", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		username := "Alice"
		user := activityreporter.NewUser(username, socialGraph)

		//when
		err := user.FollowedBy(user)

		//then
		assert.ErrorIs(t, customerror.ErrFollowThemselves, err)
	})

	t.Run("should return error when user is already followed", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		user1 := activityreporter.NewUser("Alice", socialGraph)
		user2 := activityreporter.NewUser("Bob", socialGraph)
		user2.FollowedBy(user1)

		//when
		err := user2.FollowedBy(user1)

		//then
		assert.ErrorIs(t, customerror.ErrAlreadyFollowed, err)
	})

	t.Run("should be able to upload photo", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		username := "Alice"
		user := activityreporter.NewUser(username, socialGraph)

		//when
		err := user.UploadPhoto()

		//then
		assert.Nil(t, err)
	})

	t.Run("should return error when user has already uploaded", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		username := "Alice"
		user := activityreporter.NewUser(username, socialGraph)
		user.UploadPhoto()

		//when
		err := user.UploadPhoto()

		//then
		assert.ErrorIs(t, customerror.ErrAlreadyUploaded, err)
	})

	t.Run("should be able to like photo", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		user1 := activityreporter.NewUser("Alice", socialGraph)
		user2 := activityreporter.NewUser("Bob", socialGraph)
		user2.FollowedBy(user1)
		user2.UploadPhoto()

		//when
		err := user2.LikedPhotoBy(user1)

		//then
		assert.Nil(t, err)
	})

	t.Run("should be able to like their own photo", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		username := "Alice"
		user := activityreporter.NewUser(username, socialGraph)
		user.UploadPhoto()

		//when
		err := user.LikedPhotoBy(user)

		//then
		assert.Nil(t, err)
	})

	t.Run("should return error when like photo but user has already liked", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		user1 := activityreporter.NewUser("Alice", socialGraph)
		user2 := activityreporter.NewUser("Bob", socialGraph)
		user2.FollowedBy(user1)
		user2.UploadPhoto()
		user2.LikedPhotoBy(user1)

		//when
		err := user2.LikedPhotoBy(user1)

		//then
		assert.ErrorIs(t, customerror.ErrAlreadyLiked, err)
	})

	t.Run("should return error when like photo but photo doesn't exist", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		user1 := activityreporter.NewUser("Alice", socialGraph)
		user2 := activityreporter.NewUser("Bob", socialGraph)
		user2.FollowedBy(user1)

		//when
		err := user2.LikedPhotoBy(user1)

		//then
		assert.Equal(t, customerror.ErrPhotoDoesntExist(user2.Username, user1 == user2), err)
	})

	t.Run("should return error when like photo but user has not followed", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		user1 := activityreporter.NewUser("Alice", socialGraph)
		user2 := activityreporter.NewUser("Bob", socialGraph)
		user2.UploadPhoto()

		//when
		err := user2.LikedPhotoBy(user1)

		//then
		assert.Equal(t, customerror.ErrUnableToLike(user2.Username), err)
	})

	t.Run("should be able to return the right likes count", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		user1 := activityreporter.NewUser("Alice", socialGraph)
		user2 := activityreporter.NewUser("Bob", socialGraph)
		user2.FollowedBy(user1)
		user2.UploadPhoto()
		user2.LikedPhotoBy(user1)

		//when
		likesCount := user2.LikesCount()

		//then
		assert.Equal(t, 1, likesCount)
	})

	t.Run("should log the right activity when user uploaded photo", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		username := "Alice"
		user := activityreporter.NewUser(username, socialGraph)
		user.UploadPhoto()

		//when
		activities := user.ActivityReport()

		//then
		assert.Equal(t, "You uploaded photo", activities[0])
	})

	t.Run("should get notified when followed user uploaded photo", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		user1 := activityreporter.NewUser("Alice", socialGraph)
		user2 := activityreporter.NewUser("Bob", socialGraph)
		user2.FollowedBy(user1)
		user2.UploadPhoto()

		//when
		activities := user1.ActivityReport()

		//then
		assert.Equal(t, "Bob uploaded photo", activities[0])
	})

	t.Run("should log the right activity when user liked photo", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		user1 := activityreporter.NewUser("Alice", socialGraph)
		user2 := activityreporter.NewUser("Bob", socialGraph)
		user2.UploadPhoto()
		user2.FollowedBy(user1)
		user2.LikedPhotoBy(user1)

		//when
		activities := user1.ActivityReport()

		//then
		assert.Equal(t, "You liked Bob's photo", activities[0])
	})

	t.Run("should log the right activity when user liked their own photo", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		user := activityreporter.NewUser("Bob", socialGraph)
		user.UploadPhoto()
		user.LikedPhotoBy(user)
		expectedLog := []string{"You uploaded photo", "You liked your photo"}

		//when
		activities := user.ActivityReport()

		//then
		assert.Equal(t, expectedLog, activities)
	})

	t.Run("should get notified when followed user liked photo", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		user1 := activityreporter.NewUser("Alice", socialGraph)
		user2 := activityreporter.NewUser("Bob", socialGraph)
		user3 := activityreporter.NewUser("John", socialGraph)
		user2.FollowedBy(user1)
		user1.FollowedBy(user3)
		user2.UploadPhoto()
		user2.LikedPhotoBy(user1)

		//when
		activities := user3.ActivityReport()

		//then
		assert.Equal(t, "Alice liked Bob's photo", activities[0])
	})

	t.Run("should log activity not log notification when followed user liked user's photo", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		user1 := activityreporter.NewUser("Alice", socialGraph)
		user2 := activityreporter.NewUser("Bob", socialGraph)
		user2.FollowedBy(user1)
		user1.FollowedBy(user2)
		user2.UploadPhoto()
		user2.LikedPhotoBy(user1)
		expectedLog := []string{"You uploaded photo", "Alice liked your photo"}

		//when
		activities := user2.ActivityReport()

		//then
		assert.Equal(t, expectedLog, activities)
	})

	t.Run("should not get notified when user followed other user after other user acted", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		user := activityreporter.NewUser("Bob", socialGraph)

		mockObserver := new(mocks.Observer)
		mockObserver.On("OnNotify", "Bob uploaded photo").Return()
		user.UploadPhoto()

		//when
		user.Register(mockObserver)

		//then
		mockObserver.AssertNumberOfCalls(t, "OnNotify", 0)
	})

	t.Run("should not log notification when user followed other user after other user acted", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		user1 := activityreporter.NewUser("Alice", socialGraph)
		user2 := activityreporter.NewUser("Bob", socialGraph)
		user3 := activityreporter.NewUser("John", socialGraph)
		user3.FollowedBy(user2)
		user2.UploadPhoto()
		user3.UploadPhoto()
		user3.LikedPhotoBy(user2)
		user2.FollowedBy(user1)

		//when
		activities := user1.ActivityReport()

		//then
		assert.Equal(t, 0, len(activities))
	})
}
