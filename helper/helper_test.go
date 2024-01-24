package helper_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/tsanaativa-vinnera/assignment-activity-reporter/customerror"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/tsanaativa-vinnera/assignment-activity-reporter/helper"
	"github.com/stretchr/testify/assert"
)

func TestHelper(t *testing.T) {

	t.Run("should be able to handle setup when input is valid", func(t *testing.T) {
		//given
		input := "Alice follows Bob"

		//when
		err := helper.HandleSetup(input)

		//then
		assert.Nil(t, err)
	})

	t.Run("should return error invalid keyword when handling setup but input is invalid", func(t *testing.T) {
		//given
		input := "Alicefollowsbob"

		//when
		err := helper.HandleSetup(input)

		//then
		assert.ErrorIs(t, err, customerror.ErrInvalidKeyword)
	})

	t.Run("should return error invalid keyword when handling setup but keyword is invalid", func(t *testing.T) {
		//given
		input := "Alice folows bob"

		//when
		err := helper.HandleSetup(input)

		//then
		assert.ErrorIs(t, err, customerror.ErrInvalidKeyword)
	})

	t.Run("should return error follow themselves when user follow themselves", func(t *testing.T) {
		//given
		input := "Alice follows Alice"

		//when
		err := helper.HandleSetup(input)

		//then
		assert.ErrorIs(t, err, customerror.ErrFollowThemselves)
	})

	t.Run("should return error already followed when user is already followed", func(t *testing.T) {
		//given
		input := "Alice follows Bob"
		helper.HandleSetup(input)

		//when
		err := helper.HandleSetup(input)

		//then
		assert.ErrorIs(t, err, customerror.ErrAlreadyFollowed)
	})

	t.Run("should be able to handle action upload when input is valid", func(t *testing.T) {
		//given
		input := "Alice uploaded photo"
		helper.HandleSetup("Alice follows Bob")

		//when
		err := helper.HandleAction(input)

		//then
		assert.Nil(t, err)
	})

	t.Run("should return error unknown user when handling upload but user does not exist", func(t *testing.T) {
		//given
		input := "x uploaded photo"

		//when
		err := helper.HandleAction(input)

		//then
		assert.Equal(t, err, customerror.ErrUnknownUser("x"))
	})

	t.Run("should return error already uploaded when handling upload but user has uploaded", func(t *testing.T) {
		//given
		helper.HandleSetup("Alice follows Bob")
		input := "Alice uploaded photo"
		helper.HandleAction(input)

		//when
		err := helper.HandleAction(input)

		//then
		assert.ErrorIs(t, err, customerror.ErrAlreadyUploaded)
	})

	t.Run("should return error invalid keyword when handling upload but keyword is invalid", func(t *testing.T) {
		//given
		input := "Alice upload photo"
		helper.HandleSetup("Alice follows Bob")

		//when
		err := helper.HandleAction(input)

		//then
		assert.ErrorIs(t, err, customerror.ErrInvalidKeyword)
	})

	t.Run("should be able to handle action like when input is valid", func(t *testing.T) {
		//given
		helper.HandleSetup("Alice follows Bob")
		helper.HandleAction("Bob uploaded photo")
		input := "Alice likes Bob photo"

		//when
		err := helper.HandleAction(input)

		//then
		assert.Nil(t, err)
	})

	t.Run("should be able to handle action like their own photo when input is valid", func(t *testing.T) {
		//given
		helper.HandleSetup("Alice follows Bob")
		helper.HandleAction("Alice uploaded photo")
		input := "Alice likes Alice photo"

		//when
		err := helper.HandleAction(input)

		//then
		assert.Nil(t, err)
	})

	t.Run("should return error when handling like but liked photo does not exist", func(t *testing.T) {
		//given
		helper.HandleSetup("Alice follows Billy")
		input := "Alice likes Billy photo"
		isMyself := false

		//when
		err := helper.HandleAction(input)

		//then
		assert.Equal(t, err, customerror.ErrPhotoDoesntExist("Billy", isMyself))
	})

	t.Run("should return error when handling like their own photo but photo does not exist", func(t *testing.T) {
		//given
		helper.HandleSetup("Tiara follows Bob")
		input := "Tiara likes Tiara photo"
		isMyself := true

		//when
		err := helper.HandleAction(input)

		//then
		assert.Equal(t, err, customerror.ErrPhotoDoesntExist("Tiara", isMyself))
	})

	t.Run("should return error when handling like but user has not followed", func(t *testing.T) {
		//given
		helper.HandleSetup("Alice follows Bob")
		helper.HandleAction("Alice uploaded photo")
		input := "Bob likes Alice photo"

		//when
		err := helper.HandleAction(input)

		//then
		assert.Equal(t, err, customerror.ErrUnableToLike("Alice"))
	})

	t.Run("should return error when handling like but user already liked", func(t *testing.T) {
		//given
		helper.HandleSetup("Alice follows Bob")
		helper.HandleAction("Bob uploaded photo")
		input := "Alice likes Bob photo"
		helper.HandleAction(input)

		//when
		err := helper.HandleAction(input)

		//then
		assert.Equal(t, err, customerror.ErrAlreadyLiked)
	})

	t.Run("should return error invalid keyword when handling like but keyword is invalid", func(t *testing.T) {
		//given
		helper.HandleSetup("Alice follows Bob")
		helper.HandleAction("Bob uploaded photo")
		input := "Alice like Bob photo"

		//when
		err := helper.HandleAction(input)

		//then
		assert.ErrorIs(t, err, customerror.ErrInvalidKeyword)
	})

	t.Run("should return error unknown user when handling like but user does not exist", func(t *testing.T) {
		//given
		helper.HandleSetup("Alice follows Bob")
		helper.HandleAction("Bob uploaded photo")
		input := "John likes Bob photo"

		//when
		err := helper.HandleAction(input)

		//then
		assert.Equal(t, err, customerror.ErrUnknownUser("John"))
	})

	t.Run("should return error unknown user when handling like but liked user does not exist", func(t *testing.T) {
		//given
		helper.HandleSetup("Alice follows Bob")
		helper.HandleAction("Bob uploaded photo")
		input := "Alice likes John photo"

		//when
		err := helper.HandleAction(input)

		//then
		assert.Equal(t, err, customerror.ErrUnknownUser("John"))
	})

	t.Run("should be able to handle display when input is valid", func(t *testing.T) {
		//given
		helper.HandleSetup("Alice follows Bob")
		helper.HandleAction("Alice uploaded photo")
		input := "Alice"

		//when
		err := helper.HandleDisplay(input)

		//then
		assert.Nil(t, err)
	})

	t.Run("should return error unknown user when handling display but user does not exist", func(t *testing.T) {
		//given
		input := "Mai"

		//when
		err := helper.HandleDisplay(input)

		//then
		assert.Equal(t, err, customerror.ErrUnknownUser("Mai"))
	})

	t.Run("should be able to handle trending", func(t *testing.T) {
		//given
		helper.HandleSetup("Ricky follows June")
		helper.HandleSetup("Ricky follows Michael")
		helper.HandleSetup("Ricky follows Yun")

		helper.HandleSetup("Sarah follows June")
		helper.HandleSetup("Sarah follows Michael")
		helper.HandleSetup("Sarah follows Yun")

		helper.HandleSetup("Yun follows Michael")

		helper.HandleAction("Ricky uploaded photo")
		helper.HandleAction("June uploaded photo")
		helper.HandleAction("Michael uploaded photo")
		helper.HandleAction("Yun uploaded photo")

		helper.HandleAction("Ricky likes June photo")
		helper.HandleAction("Sarah likes June photo")
		helper.HandleAction("June likes June photo")

		helper.HandleAction("Ricky likes Yun photo")
		helper.HandleAction("Yun likes Yun photo")
		helper.HandleAction("Sarah likes Yun photo")

		helper.HandleAction("Ricky likes Michael photo")
		helper.HandleAction("Sarah likes Michael photo")
		helper.HandleAction("Yun likes Michael photo")
		helper.HandleAction("Michael likes Michael photo")

		expectedTrending := []string{"Michael", "June", "Yun"}

		//when
		trending := helper.HandleTrending()
		actualTrending := []string{trending[0].Username, trending[1].Username, trending[2].Username}

		//then
		assert.Equal(t, expectedTrending, actualTrending)
	})

}
