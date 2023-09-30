package fbinfo

import (
	"goosint/models"
	"io"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetFacebookInfo(uid string) models.Facebook {
	client := &http.Client{}
	name := make(chan string)
	avatar := make(chan string)
	birthday := make(chan string)
	friends := make(chan int64)
	subscribers := make(chan int64)
	cretime := make(chan string)

	go getAvatar(client, uid, name, avatar)
	go getBirthday(client, uid, birthday)
	go getFriends(client, uid, friends)
	go getSubscribers(client, uid, subscribers)
	go getCreatedTime(client, uid, cretime)

	return models.Facebook{
		UID: uid,
		Data: models.InfoData{
			Name:        <-name,
			Avatar:      <-avatar,
			Birthday:    <-birthday,
			Friends:     <-friends,
			Subscribers: <-subscribers,
			CreatedTime: <-cretime,
		},
	}
}

func getAvatar(client *http.Client, uid string, name chan<- string, avatar chan<- string) {
	urlAvatar := `https://www.facebook.com/api/graphql?doc_id=5341536295888250&variables=%7B%22height%22:500,%22scale%22:1,%22userID%22:%22` + uid + `%22,%22width%22:500%7D`

	req, err := http.NewRequest("POST", urlAvatar, nil)

	check(err)
	res, err := client.Do(req)
	check(err)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	check(err)

	nameValue := gjson.Get(string(body), "data.profile.name")
	avatarValue := gjson.Get(string(body), "data.profile.profile_picture.uri")
	name <- nameValue.String()
	avatar <- avatarValue.String()
}

func getBirthday(client *http.Client, uid string, birthday chan<- string) {
	urlBirthday := `https://www.facebook.com/api/graphql/?q=node(` + uid + `)%7Bbirthday%7D`

	req, err := http.NewRequest("POST", urlBirthday, nil)

	check(err)
	res, err := client.Do(req)
	check(err)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	check(err)

	value := gjson.Get(string(body), uid+".birthday")
	birthday <- value.String()
}

func getSubscribers(client *http.Client, uid string, subcribers chan<- int64) {
	urlSubscribers := `https://www.facebook.com/api/graphql/?q=node(` + uid + `)%7Bsubscribers%7Bcount%7D%7D`

	req, err := http.NewRequest("POST", urlSubscribers, nil)

	check(err)
	res, err := client.Do(req)
	check(err)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	check(err)
	value := gjson.Get(string(body), uid+".subscribers.count")
	subcribers <- value.Int()
}

func getFriends(client *http.Client, uid string, friends chan<- int64) {
	urlFriends := `https://www.facebook.com/api/graphql/?q=node(` + uid + `)%7Bfriends%7Bcount%7D%7D`

	req, err := http.NewRequest("POST", urlFriends, nil)

	check(err)
	res, err := client.Do(req)
	check(err)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	check(err)
	value := gjson.Get(string(body), uid+".friends.count")
	friends <- value.Int()
}

func getCreatedTime(client *http.Client, uid string, cretime chan<- string) {
	urlCreatedTime := `https://www.facebook.com/api/graphql/?q=node(` + uid + `)%7Bcreated_time%7D`

	req, err := http.NewRequest("POST", urlCreatedTime, nil)

	check(err)
	res, err := client.Do(req)
	check(err)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	check(err)
	value := gjson.Get(string(body), uid+".created_time")
	cretime <- time.Unix(value.Int(), 0).String()
}
