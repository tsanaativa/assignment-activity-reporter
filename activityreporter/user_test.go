package activityreporter_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/tsanaativa-vinnera/assignment-activity-reporter/activityreporter"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/tsanaativa-vinnera/assignment-activity-reporter/customerror"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	t.Run("should return new user", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		username := "Alice"

		//whem
		user := activityreporter.NewUser(username, &socialGraph)

		//then
		assert.Equal(t, user.Username, username)
	})

	t.Run("should be able to follow user", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		user1 := activityreporter.NewUser("Alice", &socialGraph)
		user2 := activityreporter.NewUser("Bob", &socialGraph)

		//when
		err := user2.FollowedBy(user1)

		//then
		assert.Nil(t, err)
	})

	t.Run("should return error when user follow themselves", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		username := "Alice"
		user := activityreporter.NewUser(username, &socialGraph)

		//when
		err := user.FollowedBy(user)

		//then
		assert.ErrorIs(t, err, customerror.ErrFollowThemselves)
	})

	t.Run("should return error when user is already followed", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		user1 := activityreporter.NewUser("Alice", &socialGraph)
		user2 := activityreporter.NewUser("Bob", &socialGraph)
		user2.FollowedBy(user1)

		//when
		err := user2.FollowedBy(user1)

		//then
		assert.ErrorIs(t, err, customerror.ErrAlreadyFollowed)
	})

	t.Run("should be able to upload photo", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		username := "Alice"
		user := activityreporter.NewUser(username, &socialGraph)

		//when
		err := user.UploadPhoto()

		//then
		assert.Nil(t, err)
	})

	t.Run("should return error when user has already uploaded", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		username := "Alice"
		user := activityreporter.NewUser(username, &socialGraph)
		user.UploadPhoto()

		//when
		err := user.UploadPhoto()

		//then
		assert.ErrorIs(t, err, customerror.ErrAlreadyUploaded)
	})
}
