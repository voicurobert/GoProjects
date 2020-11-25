package main

// add a function that will insert a new page after a given page
// add a function that will delete a page

import (
	"fmt"
)

// linked list
type storyPage struct {
	text     string
	nextPage *storyPage
}

func addPage(page, nextPage *storyPage) {
	if page.nextPage == nil {
		page.nextPage = nextPage
	}
}

func removePage(page, pageToRemove *storyPage) {
	if page.nextPage == pageToRemove {
		page.nextPage = nil
	}
}

func playStory(page *storyPage) {
	if page == nil {
		return
	}
	fmt.Println(page.text)
	playStory(page.nextPage)
}

func main() {
	//scanner := bufio.NewScanner(os.Stdin)

	page1 := storyPage{"It was a dark and stormy night.", nil}
	page2 := storyPage{"You are alone, and you need to find the sacred helmet before the bad guys do", nil}
	page3 := storyPage{"You see a troll ahead", nil}

	page1.nextPage = &page2
	page2.nextPage = &page3
	playStory(&page1)
}
