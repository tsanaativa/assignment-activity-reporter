package customerror

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidMenu    = errors.New("invalid menu")
	ErrInvalidInput   = errors.New("invalid input")
	ErrInvalidKeyword = errors.New("invalid keyword")

	ErrAlreadyFollowed  = errors.New("you already followed the user")
	ErrAlreadyLiked     = errors.New("you already liked the photo")
	ErrFollowThemselves = errors.New("a user cannot follow themselves")
	ErrAlreadyUploaded  = errors.New("you cannot upload more than once")
)

func ErrUnknownUser(username string) error {
	return fmt.Errorf("unknown user %s", username)
}

func ErrPhotoDoesntExist(username string) error {
	return fmt.Errorf("%s doesn't have a photo", username)
}

func ErrUnableToLike(username string) error {
	return fmt.Errorf("unable to like %s's photo", username)
}
