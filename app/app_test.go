package app_test

import (
	"bufio"
	"strings"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/tsanaativa-vinnera/assignment-activity-reporter/app"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/tsanaativa-vinnera/assignment-activity-reporter/customerror"
	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {

	t.Run("should be able to handle setup", func(t *testing.T) {
		//given
		input := "1\nAlice follows Bob\n5\n"
		buf := strings.NewReader(input)
		scanner := bufio.NewScanner(buf)

		//when
		err := app.NewApp().RunApp(scanner)

		//then
		assert.Nil(t, err)
	})

	t.Run("should return error invalid keyword when handling setup but input is invalid", func(t *testing.T) {
		//given
		input := "1\nAlice Grace follows Bob\n5\n"
		buf := strings.NewReader(input)
		scanner := bufio.NewScanner(buf)

		//when
		err := app.NewApp().RunApp(scanner)

		//then
		assert.ErrorIs(t, err, customerror.ErrInvalidKeyword)
	})

	t.Run("should return error invalid keyword when handling setup but keyword is invalid", func(t *testing.T) {
		//given
		input := "1\nAlice flows Bob\n5\n"
		buf := strings.NewReader(input)
		scanner := bufio.NewScanner(buf)

		//when
		err := app.NewApp().RunApp(scanner)

		//then
		assert.ErrorIs(t, err, customerror.ErrInvalidKeyword)
	})

	t.Run("should return error follow themselves when user follow themselves", func(t *testing.T) {
		//given
		input := "1\nAlice follows Alice\n5\n"
		buf := strings.NewReader(input)
		scanner := bufio.NewScanner(buf)

		//when
		err := app.NewApp().RunApp(scanner)

		//then
		assert.ErrorIs(t, err, customerror.ErrFollowThemselves)
	})

	t.Run("should return error already followed when user is already followed", func(t *testing.T) {
		//given
		input := "1\nAlice follows Bob\n1\nAlice follows Bob\n5\n"
		buf := strings.NewReader(input)
		scanner := bufio.NewScanner(buf)

		//when
		err := app.NewApp().RunApp(scanner)

		//then
		assert.ErrorIs(t, err, customerror.ErrAlreadyFollowed)
	})

	t.Run("should be able to handle action upload when input is valid", func(t *testing.T) {
		//given
		input := "1\nAlice follows Bob\n2\nAlice uploaded photo\n5\n"
		buf := strings.NewReader(input)
		scanner := bufio.NewScanner(buf)

		//when
		err := app.NewApp().RunApp(scanner)

		//then
		assert.Nil(t, err)
	})

	t.Run("should return invalid keyword when handle action but input is invalid", func(t *testing.T) {
		//given
		input := "1\nAlice follows Bob\n2\nAlice Grace likes Bobby photo\n5\n"
		buf := strings.NewReader(input)
		scanner := bufio.NewScanner(buf)

		//when
		err := app.NewApp().RunApp(scanner)

		//then
		assert.ErrorIs(t, err, customerror.ErrInvalidKeyword)
	})

	t.Run("should return error unknown user when handling upload but user does not exist", func(t *testing.T) {
		//given
		input := "2\nx uploaded photo\n5\n"
		buf := strings.NewReader(input)
		scanner := bufio.NewScanner(buf)

		//when
		err := app.NewApp().RunApp(scanner)

		//then
		assert.Equal(t, err, customerror.ErrUnknownUser("x"))
	})

	t.Run("should return error already uploaded when handling upload but user has uploaded", func(t *testing.T) {
		//given
		input := "1\nAlice follows Bob\n2\nAlice uploaded photo\n2\nAlice uploaded photo\n5\n"
		buf := strings.NewReader(input)
		scanner := bufio.NewScanner(buf)

		//when
		err := app.NewApp().RunApp(scanner)

		//then
		assert.ErrorIs(t, err, customerror.ErrAlreadyUploaded)
	})

	t.Run("should return error invalid keyword when handling upload but keyword is invalid", func(t *testing.T) {
		//given
		input := "1\nAlice follows Bob\n2\nAlice upload photo\n5\n"
		buf := strings.NewReader(input)
		scanner := bufio.NewScanner(buf)

		//when
		err := app.NewApp().RunApp(scanner)

		//then
		assert.ErrorIs(t, err, customerror.ErrInvalidKeyword)
	})

	t.Run("should be able to handle action like when input is valid", func(t *testing.T) {
		//given
		input := "1\nAlice follows Bob\n2\nBob uploaded photo\n2\nAlice likes Bob photo\n5\n"
		buf := strings.NewReader(input)
		scanner := bufio.NewScanner(buf)

		//when
		err := app.NewApp().RunApp(scanner)

		//then
		assert.Nil(t, err)
	})

	t.Run("should be able to handle action like their own photo when input is valid", func(t *testing.T) {
		//given
		input := "1\nAlice follows Bob\n2\nAlice uploaded photo\n2\nAlice likes Alice photo\n5\n"
		buf := strings.NewReader(input)
		scanner := bufio.NewScanner(buf)

		//when
		err := app.NewApp().RunApp(scanner)

		//then
		assert.Nil(t, err)
	})

	t.Run("should return error when handling like but liked photo does not exist", func(t *testing.T) {
		//given
		input := "1\nAlice follows Billy\n2\nAlice likes Billy photo\n5\n"
		isMyself := false
		buf := strings.NewReader(input)
		scanner := bufio.NewScanner(buf)

		//when
		err := app.NewApp().RunApp(scanner)

		//then
		assert.Equal(t, err, customerror.ErrPhotoDoesntExist("Billy", isMyself))
	})

	t.Run("should return error when handling like their own photo but photo does not exist", func(t *testing.T) {
		//given
		input := "1\nTiara follows Bob\n2\nTiara likes Tiara photo\n5\n"
		isMyself := true
		buf := strings.NewReader(input)
		scanner := bufio.NewScanner(buf)

		//when
		err := app.NewApp().RunApp(scanner)

		//then
		assert.Equal(t, err, customerror.ErrPhotoDoesntExist("Tiara", isMyself))
	})

	t.Run("should return error when handling like but user has not followed", func(t *testing.T) {
		//given
		input := "1\nAlice follows Bob\n2\nAlice uploaded photo\n2\nBob likes Alice photo\n5\n"
		buf := strings.NewReader(input)
		scanner := bufio.NewScanner(buf)

		//when
		err := app.NewApp().RunApp(scanner)

		//then
		assert.Equal(t, err, customerror.ErrUnableToLike("Alice"))
	})

	t.Run("should return error when handling like but user already liked", func(t *testing.T) {
		//given
		input := "1\nAlice follows Bob\n2\nBob uploaded photo\n2\nAlice likes Bob photo\n2\nAlice likes Bob photo\n5\n"
		buf := strings.NewReader(input)
		scanner := bufio.NewScanner(buf)

		//when
		err := app.NewApp().RunApp(scanner)

		//then
		assert.ErrorIs(t, err, customerror.ErrAlreadyLiked)
	})

	t.Run("should return error invalid keyword when handling like but keyword is invalid", func(t *testing.T) {
		//given
		input := "1\nAlice follows Bob\n2\nBob uploaded photo\n2\nAlice liek Bob photo\n5\n"
		buf := strings.NewReader(input)
		scanner := bufio.NewScanner(buf)

		//when
		err := app.NewApp().RunApp(scanner)

		//then
		assert.ErrorIs(t, err, customerror.ErrInvalidKeyword)
	})

	t.Run("should return error unknown user when handling like but user does not exist", func(t *testing.T) {
		//given
		input := "1\nAlice follows Bob\n2\nBob uploaded photo\n2\nJohn likes Bob photo\n5\n"
		buf := strings.NewReader(input)
		scanner := bufio.NewScanner(buf)

		//when
		err := app.NewApp().RunApp(scanner)

		//then
		assert.Equal(t, err, customerror.ErrUnknownUser("John"))
	})

	t.Run("should return error unknown user when handling like but liked user does not exist", func(t *testing.T) {
		//given
		input := "1\nAlice follows Bob\n2\nBob uploaded photo\n2\nAlice likes John photo\n5\n"
		buf := strings.NewReader(input)
		scanner := bufio.NewScanner(buf)

		//when
		err := app.NewApp().RunApp(scanner)

		//then
		assert.Equal(t, err, customerror.ErrUnknownUser("John"))
	})

	t.Run("should be able to handle display when input is valid", func(t *testing.T) {
		//given
		input := "1\nAlice follows Bob\n2\nAlice uploaded photo\n3\nAlice\n5\n"
		buf := strings.NewReader(input)
		scanner := bufio.NewScanner(buf)

		//when
		err := app.NewApp().RunApp(scanner)

		//then
		assert.Nil(t, err)
	})

	t.Run("should return error unknown user when handling display but user does not exist", func(t *testing.T) {
		//given
		input := "3\nMai\n5\n"
		buf := strings.NewReader(input)
		scanner := bufio.NewScanner(buf)

		//when
		err := app.NewApp().RunApp(scanner)

		//then
		assert.Equal(t, err, customerror.ErrUnknownUser("Mai"))
	})

	t.Run("should be able to handle trending", func(t *testing.T) {
		//given
		input := "1\nRicky follows June\n2\nJune uploaded photo\n2\nRicky likes June photo\n1\nYun follows June\n2\nYun likes June photo\n4\n5\n"
		buf := strings.NewReader(input)
		scanner := bufio.NewScanner(buf)

		//when
		err := app.NewApp().RunApp(scanner)

		//then
		assert.Nil(t, err)
	})

	t.Run("should return invalid input menu when menu is invalid", func(t *testing.T) {
		//given
		input := "6\n5\n"
		buf := strings.NewReader(input)
		scanner := bufio.NewScanner(buf)

		//when
		err := app.NewApp().RunApp(scanner)

		//then
		assert.ErrorIs(t, err, customerror.ErrInvalidMenu)
	})

}
