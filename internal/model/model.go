package model

import (
	"encoding/xml"
	"log"
)

type Response struct {
	User    UserType     `xml:"user"`
	Book    BookType     `xml:"book"`
	Reviews []ReviewType `xml:"reviews>review"`
}

/* The actual types: */
type BookType struct {
	ID       string       `xml:"id"`
	Title    string       `xml:"title"`
	Link     string       `xml:"link"`
	ImageURL string       `xml:"image_url"`
	NumPages string       `xml:"num_pages"`
	Format   string       `xml:"format"`
	Authors  []AuthorType `xml:"authors>author"`
	ISBN     string       `xml:"isbn"`
}

type AuthorType struct {
	ID   string `xml:"id"`
	Name string `xml:"name"`
	Link string `xml:"link"`
}

type UserType struct {
	ID            string           `xml:"id"`
	Name          string           `xml:"name"`
	About         string           `xml:"about"`
	Link          string           `xml:"link"`
	ImageURL      string           `xml:"image_url"`
	SmallImageURL string           `xml:"small_image_url"`
	Location      string           `xml:"location"`
	LastActive    string           `xml:"last_active"`
	ReviewCount   int              `xml:"reviews_count"`
	Statuses      []UserStatusType `xml:"user_statuses>user_status"`
	Shelves       []ShelfType      `xml:"user_shelves>user_shelf"`
	LastRead      []ReviewType
}

type ReviewType struct {
	Book   Book   `xml:"book"`
	Rating int    `xml:"rating"`
	ReadAt string `xml:"read_at"`
	Link   string `xml:"link"`
}

type ShelfType struct {
	ID        string `xml:"id"`
	BookCount string `xml:"book_count"`
	Name      string `xml:"name"`
}

type UserStatusType struct {
	Page    int    `xml:"page"`
	Percent int    `xml:"percent"`
	Updated string `xml:"updated_at"`
	Book    Book   `xml:"book"`
}

// func getData(uri string, i interface{}) {
// 	data := getRequest(uri)
// 	xmlUnmarshal(data, i)
// }
//
// func getRequest(uri string) []byte {
// 	res, err := http.Get(uri)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	body, err := ioutil.ReadAll(res.Body)
// 	res.Body.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	return body
// }

func xmlUnmarshal(b []byte, i interface{}) {
	err := xml.Unmarshal(b, i)
	if err != nil {
		log.Fatal(err)
	}
}

// func parseDate(s string) (time.Time, error) {
// 	date, err := time.Parse(time.RFC3339, s)
// 	if err != nil {
// 		date, err = time.Parse(time.RubyDate, s)
// 		if err != nil {
// 			return time.Time{}, err
// 		}
// 	}
//
// 	return date, nil
// }
//
// func relativeDate(d string) string {
// 	date, err := parseDate(d)
// 	if err != nil {
// 		fmt.Println(err)
// 		return ""
// 	}
//
// 	s := time.Now().Sub(date)
//
// 	days := int(s / (24 * time.Hour))
// 	if days > 1 {
// 		return fmt.Sprintf("%v days ago", days)
// 	} else if days == 1 {
// 		return fmt.Sprintf("%v day ago", days)
// 	}
//
// 	hours := int(s / time.Hour)
// 	if hours > 1 {
// 		return fmt.Sprintf("%v hours ago", hours)
// 	}
//
// 	minutes := int(s / time.Minute)
// 	if minutes > 2 {
// 		return fmt.Sprintf("%v minutes ago", minutes)
// 	} else {
// 		return "Just now"
// 	}
// }
