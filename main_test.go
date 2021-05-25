package main

func TestBookController() {
}

/**
 * UTILITY FUNCTIONS FOR TESTING.
 */

func ReadBooks() {
	resp, err := http.Get(fmt.Sprintf("%s/book", ts.URL))
}

func CreateBook(body) {
	resp, err := http.Post(fmt.Sprintf("%s/book", ts.URL), "application/json", body)
}

func UpdateBook(body) {
	resp, err := http.Put(fmt.Sprintf("%s/book", ts.URL), "application/json", body)
}

func DeleteBook(body) {
	resp, err := http.Delete(fmt.Sprintf("%s/book", ts.URL), "application/json", body)
}

func CreateBookA() {
	body, _ := json.Marshal(map[string]string{
		"title":  "Harry Potter and The Philosopher's Stone",
		"author": "J.K. Rowling",
	})
	CreateBook(body)
}

func CreateBookB() {
	body, _ := json.Marshal(map[string]string{
		"title":  "Harry Potter and The Chamber of Secrets",
		"author": "J.K. Rowling",
	})
	CreateBook(body)
}

func CreateBookC() {
	body, _ := json.Marshal(map[string]string{
		"title":  "Harry Potter and The Prisoner of Azkaban",
		"author": "J.K. Rowling",
	})
	CreateBook(body)
}

func CreateBookD() {
	body, _ := json.Marshal(map[string]string{
		"title":  "Harry Potter and The Goblet of Fire",
		"author": "J.K. Rowling",
	})
	CreateBook(body)
}

func UpdateBookA() {
	body, _ := json.Marshal(map[string]string{
		"id":  0,
		"title": "Harry Potter and The Sorcerer's Stone"
	})
	UpdateBook(body)
}

func UpdateBookB() {
	body, _ := json.Marshal(map[string]string{
		"id":  1,
		"author": "Joanne Kathleen Rowling",
	})
	UpdateBook(body)
}

func DeleteBookA() {
	body, _ := json.Marshal(map[string]string{
		"id":  0
	})
	DeleteBook(body)
}
