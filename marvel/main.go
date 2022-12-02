package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(os.Getenv("MARVEL_PUBLIC_KEY"))

	publicKey := os.Getenv("MARVEL_PUBLIC_KEY")
	privateKey := os.Getenv("MARVEL_PRIVATE_KEY")

	client := marvelClient{
		baseURL:    "https://gateway.marvel.com:443/v1/public",
		publicKey:  publicKey,
		privateKey: privateKey,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
	event, err := client.getEvents()
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(event)
}

type marvelClient struct {
	baseURL    string
	publicKey  string
	privateKey string
	httpClient *http.Client
}

func (c *marvelClient) md5Hash(ts int64) string {
	tsForHash := strconv.Itoa(int(ts))
	hash := md5.Sum([]byte(tsForHash + c.privateKey + c.publicKey))
	return hex.EncodeToString(hash[:])
}
func (c *marvelClient) signUrl(url string) string {
	ts := time.Now().Unix()
	hash := c.md5Hash(ts)
	return fmt.Sprintf("%s?ts=%d&apikey=%s&hash=%s", url, ts, c.publicKey, hash)
}

func (c *marvelClient) getEvents() ([]Event, error) {
	url := c.baseURL + "/events?Limit=2"
	url = c.signUrl(url)
	spew.Dump(url)

	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	spew.Dump(res.Status, res.StatusCode)

	var eventResponse EventResponse
	if err := json.NewDecoder(res.Body).Decode(&eventResponse); err != nil {
		return nil, err
	}

	return eventResponse.Data.Results, nil
}
