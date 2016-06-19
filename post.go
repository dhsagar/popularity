package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Post struct {
	// TODO shift to status
	Error string `json:"error" csv:"-"`

	ID             string            `json:"id" csv:"id"`
	SortKey        string            `json:"sort_key" csv:"-"`
	Title          string            `json:"title" csv:"title" csv:"-"`
	TitleWordCount int               `json:"-" csv:"title_word_count" csv:"-"`
	TitleTag       interface{}       `json:"title_tag" csv:"-"`
	Images         map[string]string `json:"images" csv:"-"`
	ImageCount     int               `json:"-" csv:"image_count" csv:"-"`
	Topics         []string          `json:"topics" csv:"-"`
	MetaKeyWords   string            `json:"-" csv:"meta_key_word" csv:"-"`
	Channel        string            `json:"channel" csv:"news_category" csv:"-"`
	ChannelName    string            `json:"channel_name" csv:"-"`
	SubChannels    []string          `json:"subchannels" csv:"-"`
	Author         string            `json:"author" csv:"author" csv:"-"`
	AuthorID       string            `json:"author_id" csv:"-"`
	PostDate       time.Time         `json:"post_date" csv:"post_date"`
	PostDay        string            `json:"-" csv:"post_day"`
	PostDelta      int               `json:"-" csv:"post_delta"`
	PostDateRfc    string            `json:"post_date_rfc" csv:"-"`
	Link           string            `json:"link" csv:"link" csv:"-"`
	Content        struct {
		Full  string      `json:"full" csv:"-"`
		Intro interface{} `json:"intro" csv:"-"`
		Plain string      `json:"plain" csv:"-"`
	} `json:"content" csv:"-"`
	ContentPlainText string `json:"-" csv:"content"`
	HyperLinks int `json:"-" csv:"hyperlinks"`
	MashableLinks int `json:"-" csv:"mashable_links"`
	Shares           struct {
		Facebook   int `json:"facebook" csv:"-"`
		Twitter    int `json:"twitter" csv:"-"`
		GooglePlus int `json:"google_plus" csv:"-"`
		Pinterest  int `json:"pinterest" csv:"-"`
		LinkedIn   int `json:"linked_in" csv:"-"`
		Total      int `json:"total" csv:"-"`
	} `json:"shares" csv:"-"`
	TotalShare    int                    `json:"-" csv:"shares"`
	CommentsCount int                    `json:"comments_count" csv:"comment"`
	URL           string                 `json:"url" csv:"-"`
	Velocity      []int                  `json:"velocity" csv:"-"`
	ShortURL      string                 `json:"short_url" csv:"-"`
	Targeting     map[string]interface{} `json:"targeting" csv:"-"`
	ContentSource string                 `json:"content_source" csv:"-"`
	LeadType      string                 `json:"lead_type" csv:"-"`
	CommentsURL   string                 `json:"comments_url" csv:"-"`
	ShortcodeData struct {
		Gallery []interface{} `json:"gallery" csv:"-"`
	} `json:"shortcode_data"csv:"-"`
	Webview        bool   `json:"webview" csv:"-"`
	SponsoredBy    string `json:"sponsored_by" csv:"-"`
	SponsoredByURL string `json:"sponsored_by_url" csv:"-"`
	SeriesType     string `json:"series_type" csv:"series_type"`
	SeriesSlug     string `json:"series_slug" csv:"-"`
}

func GetPost(id string) *Post {
	url := "https://api.mashable.com/v1/posts/" + id
	log.Println("calling api for ", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	post := &Post{}

	log.Println("got response with", resp.Status)
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(responseBody, post)
		if err != nil {
			log.Fatal(err)
		}

		if post.Error != "" {
			log.Println("got api response status as", post.Error)
			log.Fatal()
		}

		post.TitleWordCount = len(strings.Split(post.Title, " "))
		post.ImageCount = len(post.Images)
		post.MetaKeyWords = strings.Join(post.Topics, " ")
		post.PostDay = post.PostDate.Weekday().String()
		post.PostDelta = int(time.Since(post.PostDate).Hours() / 24)
		post.ContentPlainText = post.Content.Plain
		post.TotalShare = post.Shares.Total
		post.HyperLinks = strings.Count(post.Content.Full, "<a href=")
		post.MashableLinks = strings.Count(post.Content.Full, `<a href="http://mashable.com`)
		return post
	} else {
		log.Println("response contains error, exiting...")
		log.Fatal()
	}
	return post
}
