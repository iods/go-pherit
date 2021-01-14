package Page

import "io/ioutil"

// Page Struct defining the Title and Body of the wiki page.
type Page struct {
	Title string
	Body  []byte
}

// save Takes the post body and saves it to a file, naming it the title. The
// signature is a method that takes a receiver p, a pointer to Page.
func (p Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// LoadPage First, the function constructs the title from the param. Second, it reads
// the contents into a new variable, and returns a pointer to the Page literal.
func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}