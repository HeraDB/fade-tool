package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"log"
	"unicode/utf8"
)

type Color struct{
	R,G,B   int
}

func ToRGB(h string)(c Color,err error){
	switch len(h) {
	case 6:
		_,err = fmt.Sscanf(h,"%02x%02x%02x",&c.R,&c.G,&c.B)
     case 3:
		_,err = fmt.Sscanf(h,"%1x%1x%1x",&c.R,&c.G,&c.B)
	  	c.R *= 17
	  	c.G *= 17
	  	c.B *= 17
     default:
		err = fmt.Errorf("Invalid hex color")
	}
	return
}

func Bresenham(s,e float64,steps int)[]int{
	delta   := (e-s)/(float64(steps)-1)
	colors  := []int{int(s)}
	err     := 0.0
	for i:=0;i<steps-1;i++{
		n   := float64(colors[i])+delta
		err  = err+(n-float64(int(n)))
		if err>=0.5{
			n   = n+1.0
			err = err-1.0
		}
		colors = append(colors,int(n))
	}
	return colors
}
 
func Gradient(c1,c2 Color,n int)([]int,[]int,[]int){
	if n<3{
		r := []int{c1.R,c2.R}
	   	g := []int{c1.G,c2.G}
	   	b := []int{c1.B,c2.B}
		return r,g,b
	}
   
	R := Bresenham(float64(c1.R),float64(c2.R),n)
	G := Bresenham(float64(c1.G),float64(c2.G),n)
	B := Bresenham(float64(c1.B),float64(c2.B),n)
	return R,G,B
}

func Colorize(text string,r,g,b int)string{
	fg := fmt.Sprintf("\x1b[38;2;%d;%d;%dm",r,g,b)
	return fg+text+"\x1b[0m"
}

func HandleErr(err error){
	if err != nil{
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Ussage correction: %s <colour one> <colour two> <banner file>", os.Args[0])
		return
	}


	File, error := ioutil.ReadFile(os.Args[3]); if error != nil {
		fmt.Printf("Failed to read file correctly")
		return
	}

	Files := strings.SplitAfter(string(File), "\n")

	LenDeep := []int{}
	Count := 0

	for _, Line := range Files {
		LenDeep = append(LenDeep, utf8.RuneCountInString(Line))
		Count++
	}

	sort.Ints(LenDeep[:])

	ColorHexOne, error := ToRGB(os.Args[1]); if error != nil {
		fmt.Println(error.Error())
	}

	ColorHexTwo, error := ToRGB(os.Args[2]); if error != nil {
		fmt.Println(error.Error())
	}
	log.Println(LenDeep[utf8.RuneCountInString(strings.Join(Files, " "))], utf8.RuneCountInString(strings.Join(Files, " ")))
	R, G, B := Gradient(ColorHexOne, ColorHexTwo, LenDeep[utf8.RuneCountInString(strings.Join(Files, " "))-1])

	out := []string{}
	count := 0
	for _, Line := range Files {
		for _, V := range Line {

			out = append(out, Colorize(string(V), R[count], G[count], B[count]))
			count++
		}
	}

	for _, V := range out {
		fmt.Printf("%s", V)
	}
	fmt.Println("")


}