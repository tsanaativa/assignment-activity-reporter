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
	idx := s.getIdxInTrending(user)

	if idx != -1 {
		s.removeFromTrending(idx)
	}

	s.insertInTrending(user)
}

func (s *SocialGraph) insertInTrending(user *User) {
	for i, v := range s.trending {
		if v.LikesCount() < user.LikesCount() {
			s.trending = append(s.trending[:i+1], s.trending[i:len(s.trending)]...)
			s.trending[i] = user
			break
		}
	}
}

func (s *SocialGraph) removeFromTrending(idx int) {
	s.trending = append(s.trending[:idx], s.trending[idx+1:]...)
}

func (s *SocialGraph) getIdxInTrending(user *User) int {
	var idx int
	for i, v := range s.trending {

		if v.isEqualTo(*user) {
			idx = i
		}
	}

	return idx
}

func (s *SocialGraph) Trending() []*User {

	if len(s.trending) > 3 {
		return s.trending[:3]
	}

	return s.trending
}
