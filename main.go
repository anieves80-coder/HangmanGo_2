package main

import (
	"fmt"
	"os"
	"os/exec"
	"bufio"	
	"strings"
)
var inWord = word{}
var guessWord string
var wordCnt int
var err int
var verifyErr bool

func main(){	
	guessWord = inWord.getInput()	
	showWord()
	getChars()
}

func getChars() {
	fmt.Println("Type in a single letter to guess!")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err == nil {
		setChar(string(char))
	} else {
		fmt.Println("try again ")
		getChars()
	}
}

func setChar(char string){	

	clearTerminal()
	
	verifyErr = true
	
	for i, lt := range guessWord{
		if char == string(lt) {
			inWord[i] = char
			wordCnt++
			verifyErr = false
		}
	}		
	
	if verifyErr {
		err++
	}
	
	showWord()
	
	if err >= 7 {
		fmt.Println("You LOSE!!!\n", "Word was",guessWord)
	} else if wordCnt < len(inWord){
		getChars()
	} else {
		fmt.Println("You WIN!!!", "Word was",guessWord)
	}
	
}

func showWord() {
	fmt.Println(strings.Join(inWord, " "))
	fmt.Println(" ")
	fmt.Println("Tries: ", err, "out of 7")
	fmt.Println(" ")
}

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
