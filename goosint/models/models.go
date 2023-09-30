package models

type InfoData struct {
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Birthday    string `json:"birthday"`
	Friends     int64  `json:"friends"`
	Subscribers int64  `json:"subscribers"`
	CreatedTime string `json:"created_time"`
}

type Facebook struct {
	UID  string   `json:"uid"`
	Data InfoData `json:"data"`
}

type GithubCommitSearch struct {
	TotalCount int `json:"total_count"`
	Items      []struct {
		Commit struct {
			Author struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"author"`
			Committer struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"committer"`
		} `json:"commit"`
		Repository struct {
			Name  string `json:"name"`
			Owner struct {
				Login string `json:"login"`
			} `json:"owner"`
			HTMLURL string `json:"html_url"`
		} `json:"repository"`
	} `json:"items"`
}

type DomainReverseList struct {
	DomainsCount int      `json:"domainsCount"`
	DomainsList  []string `json:"domainsList"`
}

type SeonEmail struct {
	Success bool `json:"success"`
	Data    struct {
		Email          string `json:"email"`
		AccountDetails struct {
			Facebook struct {
				Registered interface{} `json:"registered"`
			} `json:"facebook"`
			Github struct {
				Registered bool `json:"registered"`
			} `json:"github"`
			Google struct {
				Registered bool   `json:"registered"`
				Photo      string `json:"photo"`
			} `json:"google"`
			Gravatar struct {
				Registered bool        `json:"registered"`
				Location   interface{} `json:"location"`
				Name       interface{} `json:"name"`
				ProfileURL interface{} `json:"profile_url"`
				Username   interface{} `json:"username"`
			} `json:"gravatar"`
			Instagram struct {
				Registered bool `json:"registered"`
			} `json:"instagram"`
			Linkedin struct {
				Registered      interface{} `json:"registered"`
				URL             interface{} `json:"url"`
				Name            interface{} `json:"name"`
				Company         interface{} `json:"company"`
				Title           interface{} `json:"title"`
				Location        interface{} `json:"location"`
				Website         interface{} `json:"website"`
				Twitter         interface{} `json:"twitter"`
				Photo           interface{} `json:"photo"`
				ConnectionCount interface{} `json:"connection_count"`
			} `json:"linkedin"`
			Skype struct {
				Registered bool        `json:"registered"`
				Country    interface{} `json:"country"`
				City       interface{} `json:"city"`
				Gender     interface{} `json:"gender"`
				Name       string      `json:"name"`
				ID         string      `json:"id"`
				Handle     interface{} `json:"handle"`
				Bio        interface{} `json:"bio"`
				Age        interface{} `json:"age"`
				Language   interface{} `json:"language"`
				State      interface{} `json:"state"`
				Photo      string      `json:"photo"`
			} `json:"skype"`
			Twitter struct {
				Registered bool `json:"registered"`
			} `json:"twitter"`
		} `json:"account_details"`
		BreachDetails struct {
			NumberOfBreaches int `json:"number_of_breaches"`
			Breaches         []struct {
				Name   string `json:"name"`
				Domain string `json:"domain"`
				Date   string `json:"date"`
			} `json:"breaches"`
		} `json:"breach_details"`
	} `json:"data"`
}

type SeonPhone struct {
	Success bool `json:"success"`
	Data    struct {
		Number         int64  `json:"number"`
		Valid          bool   `json:"valid"`
		Type           string `json:"type"`
		Country        string `json:"country"`
		Carrier        string `json:"carrier"`
		AccountDetails struct {
			Facebook struct {
				Registered bool `json:"registered"`
			} `json:"facebook"`
			Google struct {
				Registered bool        `json:"registered"`
				AccountID  interface{} `json:"account_id"`
				FullName   interface{} `json:"full_name"`
			} `json:"google"`
			Twitter struct {
				Registered bool `json:"registered"`
			} `json:"twitter"`
			Instagram struct {
				Registered interface{} `json:"registered"`
			} `json:"instagram"`
			Skype struct {
				Registered bool        `json:"registered"`
				Age        interface{} `json:"age"`
				City       interface{} `json:"city"`
				Bio        interface{} `json:"bio"`
				Country    interface{} `json:"country"`
				Gender     interface{} `json:"gender"`
				Language   interface{} `json:"language"`
				Name       interface{} `json:"name"`
				Handle     interface{} `json:"handle"`
				ID         interface{} `json:"id"`
				Photo      interface{} `json:"photo"`
				State      interface{} `json:"state"`
			} `json:"skype"`
			Telegram struct {
				Registered bool        `json:"registered"`
				Photo      interface{} `json:"photo"`
				LastSeen   interface{} `json:"last_seen"`
			} `json:"telegram"`
			Viber struct {
				Registered bool        `json:"registered"`
				Photo      interface{} `json:"photo"`
				LastSeen   interface{} `json:"last_seen"`
				Name       interface{} `json:"name"`
			} `json:"viber"`
			Zalo struct {
				Registered  bool        `json:"registered"`
				DateOfBirth interface{} `json:"date_of_birth"`
				Name        interface{} `json:"name"`
				UID         interface{} `json:"uid"`
			} `json:"zalo"`
		} `json:"account_details"`
	} `json:"data"`
}
