package activityreporter

import (
	"sort"
)

type SocialGraph struct {
	userList map[string]*User
	trending []*User
}

func NewSocialGraph() SocialGraph {
	return SocialGraph{
		userList: make(map[string]*User),
	}
}

func (s *SocialGraph) AddNewUser(user *User) {
	s.userList[user.Username] = user
	s.trending = append(s.trending, user)
}

func (s *SocialGraph) IsUserExist(username string) (*User, bool) {
	val, ok := s.userList[username]
	return val, ok
}

func (s *SocialGraph) Trending() []*User {

	sort.Slice(s.trending[:], func(i, j int) bool {
		return s.trending[i].LikesCount() > s.trending[j].LikesCount()
	})

	if len(s.trending) > 3 {
		return s.trending[:3]
	}

	return s.trending
}
