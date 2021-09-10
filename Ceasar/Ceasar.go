package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//alfEN :=[]rune("abcdefghijklmnopqrstuvwxyz")
	alfRU :=[]rune("абвгдеёжзийклмнопрстуфхцчшщъыьэюя")
	key:=[]rune("пок")
	message:= []rune("привет, как дела")
	ceasar(alfRU, key, message)
}


func ceasar(alf, key, message []rune){
	var newAlf []rune
	offset:= 4
	for _,w:= range alf {
		if strings.Contains(string(key),string(w)){
			continue
		}

		newAlf = append(newAlf,w)
	}
	off := newAlf[len(newAlf)-offset:]
	for _,w:=range key{
		off = append(off,w)
	}
	for i:=0;i<len(newAlf)-offset;i++{
		off = append(off, newAlf[i])
	}


	for i,w:=range message{
		if !unicode.IsLetter(w){continue}
		ind := strings.Index(string(alf),string(w))
		message[i] = off[ind/2]
	}
	fmt.Println(string(off))
	fmt.Println(string(message))
}
