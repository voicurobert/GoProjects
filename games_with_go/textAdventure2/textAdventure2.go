package main

// 1. NPC (non player characters) - talk to them, fight
// 2. NPC move around the grapth
// 3. items that can picked up or placed down
// 4. accept natural language as input (verb noun)

/*
type npc struct {
	desc string
	health int
}
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type choice struct {
	cmd         string
	description string
	nextNode    *storyNode
}

type storyNode struct {
	text    string
	choices []*choice
}

func (node *storyNode) addChoice(cmd, description string, nextNode *storyNode) {
	choice := &choice{cmd, description, nextNode}
	node.choices = append(node.choices, choice)
}

func (node *storyNode) render() {
	fmt.Println(node.text)
	if node.choices != nil {
		for _, choice := range node.choices {
			fmt.Println(choice.cmd, choice.cmd)
		}
	}
}

func (node *storyNode) executeCmd(cmd string) *storyNode {
	for _, choice := range node.choices {
		if strings.ToLower(choice.cmd) == strings.ToLower(cmd) {
			return choice.nextNode
		}
	}
	fmt.Println("Sorry I didn't understand that.")
	return node
}

var scanner *bufio.Scanner

func (node *storyNode) play() {
	node.render()
	if node.choices != nil {
		scanner.Scan()
		node.executeCmd(scanner.Text()).play()
	}
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)

	start := storyNode{text: `
		You are in a large chamber, deep underground.
		You see three passages leading out. A north passge leads into darkness.
		To the sout, a passage to head upward. The eastern passages appears
		flat and well traveled
	`}

	darkRoom := storyNode{text: "It is so dark. You cannot see a thing."}

	darkRoomLit := storyNode{text: "The dark passage is now lit by your lantern. You can continue north or head back south"}

	grue := storyNode{text: "While stumbling around in the darkness, you are eaten by a grue."}

	trap := storyNode{text: "You head down the well traveled path when suddenly a trap door opens and you fall into a pit."}

	treasure := storyNode{text: "You arrive at a small chamber, filled with treasure!"}

	start.addChoice("N", "Go north", &darkRoom)
	start.addChoice("S", "Go south", &darkRoom)
	start.addChoice("E", "Go east", &trap)

	darkRoom.addChoice("S", "Try to go back", &grue)
	darkRoom.addChoice("O", "Turn on lantern", &darkRoomLit)

	darkRoomLit.addChoice("N", "Go north", &treasure)
	darkRoomLit.addChoice("S", "Go south", &start)

	start.play()

	fmt.Println()
	fmt.Println("The end.")

}
