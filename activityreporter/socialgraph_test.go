package activityreporter_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/tsanaativa-vinnera/assignment-activity-reporter/activityreporter"
	"github.com/stretchr/testify/assert"
)

func TestSocialGraph(t *testing.T) {
	t.Run("should return new social graph", func(t *testing.T) {
		//when
		socialGraph := activityreporter.NewSocialGraph()

		//then
		assert.NotNil(t, socialGraph)
	})

	t.Run("should be able add new user", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		username := "Alice"
		user := activityreporter.NewUser(username, &socialGraph)

		//when
		socialGraph.AddNewUser(user)

		//then
		_, isExist := socialGraph.IsUserExist(user.Username)
		assert.True(t, isExist)
	})

	t.Run("should return true when isuserexist and user exist", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		username := "Alice"
		user := activityreporter.NewUser(username, &socialGraph)

		//when
		socialGraph.AddNewUser(user)

		//then
		_, isExist := socialGraph.IsUserExist(user.Username)
		assert.True(t, isExist)
	})

	t.Run("should return false when isuserexist and user is not added", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		username := "Alice"

		//when
		user := activityreporter.NewUser(username, &socialGraph)

		//then
		_, isExist := socialGraph.IsUserExist(user.Username)
		assert.False(t, isExist)
	})

	t.Run("should be able to add user to trending list", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		username := "Alice"
		user := activityreporter.NewUser(username, &socialGraph)
		socialGraph.AddNewUser(user)

		//when
		user.UploadPhoto()

		//then
		assert.Equal(t, user, socialGraph.Trending()[0])
	})

	t.Run("should be able to get the right trending list", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		user1 := activityreporter.NewUser("Alice", &socialGraph)
		user2 := activityreporter.NewUser("Bob", &socialGraph)
		socialGraph.AddNewUser(user1)
		socialGraph.AddNewUser(user2)
		user1.UploadPhoto()
		user2.UploadPhoto()
		user2.FollowedBy(user1)
		user2.LikedPhotoBy(user1)

		//when
		firstUser := socialGraph.Trending()[0]

		//then
		assert.Equal(t, "Bob", firstUser.Username)
	})

	t.Run("should be able to get 3 trending users when users are more than 3", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		user1 := activityreporter.NewUser("Alice", &socialGraph)
		user2 := activityreporter.NewUser("Bob", &socialGraph)
		user3 := activityreporter.NewUser("John", &socialGraph)
		user4 := activityreporter.NewUser("Bill", &socialGraph)
		socialGraph.AddNewUser(user1)
		socialGraph.AddNewUser(user2)
		socialGraph.AddNewUser(user3)
		socialGraph.AddNewUser(user4)

		//when
		user1.UploadPhoto()
		user2.UploadPhoto()
		user3.UploadPhoto()
		user4.UploadPhoto()

		//then
		assert.Equal(t, 3, len(socialGraph.Trending()))
	})

	t.Run("should be able to update trending list correctly", func(t *testing.T) {
		//given
		socialGraph := activityreporter.NewSocialGraph()
		user1 := activityreporter.NewUser("Alice", &socialGraph)
		user2 := activityreporter.NewUser("Bob", &socialGraph)
		socialGraph.AddNewUser(user1)
		socialGraph.AddNewUser(user2)
		user1.UploadPhoto()
		user2.UploadPhoto()
		user2.FollowedBy(user1)

		//when
		user2.LikedPhotoBy(user1)

		//then
		assert.Equal(t, user2, socialGraph.Trending()[0])
	})

}
