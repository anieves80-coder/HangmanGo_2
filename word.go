package main

import (
	"fmt"
	"net/http"	
	"io/ioutil"
	"encoding/json"
)

type word []string

// Word struct for JSON thats going to be received from an API GET request
type Word struct {
	ID int `json:"id"`
	APIWord string `json:"word"`
}

func (w *word) getInput() (string, int) {
	url := "https://api.wordnik.com/v4/words.json/randomWord?&minLength=5&maxLength=-1&api_key=48dd829661f515d5abc0d03197a00582e888cc7da2484d5c7"
	resp, err := http.Get(url)	
	if err != nil{
		fmt.Println(err)
	}
	defer resp.Body.Close() //defer keyword waits for execution... in this case waits for the API call
	body, _ := ioutil.ReadAll(resp.Body)
	
	var getWord Word // creates struct object // sets the variable getWord as a type Word struct
	json.Unmarshal(body, &getWord) // sets the json as the struct. NOTE: pointer
    c := w.initDashes(len(getWord.APIWord),getWord.APIWord)
	return getWord.APIWord, c
}

func (w *word) initDashes(cnt int, APIWord string) int {	
	dashCnt := 0		
	for i := 0; i < cnt; i++ {
		if string(APIWord[i]) == "-"{
			dashCnt++
			*w = append (*w, "-")
		} else {
			*w = append (*w, "_")
		}
	}

	return dashCnt

}
