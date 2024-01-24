package activityreporter

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
}

func (s *SocialGraph) IsUserExist(username string) (*User, bool) {
	val, ok := s.userList[username]
	return val, ok
}

func (s *SocialGraph) AddToTrending(user *User) {
	s.trending = append(s.trending, user)
}

func (s *SocialGraph) UpdateTrending(user *User) {
	for i, v := range s.trending {
		if v.LikesCount() < user.LikesCount() {
			s.trending = append(s.trending[:i+1], s.trending[i:len(s.trending)-1]...)
			s.trending[i] = user
			break
		}
	}
}

func (s *SocialGraph) Trending() []*User {

	if len(s.trending) > 3 {
		return s.trending[:3]
	}

	return s.trending
}
