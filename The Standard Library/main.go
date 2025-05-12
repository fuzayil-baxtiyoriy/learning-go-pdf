package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// io
func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

// encoding/json
// {
//     "id":"12345",
//     "date_ordered":"2020-05-01T13:01:02Z",
//     "customer_id":"3",
//     "items":[{"id":"xyz123","name":"Thing 1"},{"id":"abc789","name":"Thing 2"}]
// }

type Order struct {
	ID          string    `json:"id"`
	DateOrdered time.Time `json:"date_ordered"`
	CustomerID  string    `json:"customer_id"`
	Items       []Item    `json:"items"`
}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// net/http

func jsonPlaceHolder() {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequestWithContext(context.Background(),
		http.MethodGet, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("X-My-Client", "Learning Go")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status: got %v", res.Status))
	}
	fmt.Println(res.Header.Get("Content-Type"))
	var data struct {
		UserID    int    `json:"userId"`
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", data)
}

func main() {
	// io examples
	// s := "The quick brown fox jupmed over the lazy dog"
	// sr := strings.NewReader(s)
	// counts, err := countLetters(sr)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(counts)

	// time examples
	// t, err := time.Parse("2006-01-02 15:04:05 -0700", "2023-03-13 00:00:00 +0000")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println(t.Format("January 2, 2006 at 3:04:05PM MST"))

	// net/http
	jsonPlaceHolder()

	// http.Server
	s := http.Server{
		Addr: ":8080",
		ReadTimeout: 30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout: 120 * time.Second,
	}
	defer s.Close()

}
