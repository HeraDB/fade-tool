/* Author : Akita
 * Date   : 17 January 2021
 */
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

//Types
 type Color struct{
   R,G,B   int
 }
 
// Lazy implementation
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

 
 // https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
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
   if err!=nil{
	 fmt.Printf("Error: %v\n",err)
	 os.Exit(1)
   }
 }

 
 func main(){
   if len(os.Args)!=4{
	 fmt.Printf("Usage: %s <color1> <color2> <banner file>\nNote: colors must be in hex, ex.: ff00ff\n",os.Args[0])
	 os.Exit(1)
   }
 
   file,err := ioutil.ReadFile(os.Args[3])
   HandleErr(err)
 
   lines := strings.SplitAfter(string(file),"\n")
   llen := []int{}
   for _,v := range lines{
	 llen = append(llen, utf8.RuneCountInString(v))
   }

   LOL, _ := os.Open(os.Args[3])

   lol := bufio.NewScanner(LOL)

   
   for lol.Scan() {
     count:=0
   }

   log.Println(len(llen),"lol", )

   for _, i := range llen {


    log.Println(utf8.RuneCountInString(strconv.Itoa(i)))
   }
   //log.Println(llen[len(llen)-1], llen[utf8.RuneCountInString(llen)-1])

   // llen[utf8.RuneCountInString(llen)-1]
   sort.Ints(llen[:])
 
   c1,err := ToRGB(os.Args[1])
   HandleErr(err)
   c2,err := ToRGB(os.Args[2])
   HandleErr(err)
   log.Println(llen[len(llen)-1])
   lol := utf8.RuneCountInString(strings.Join(lines, "")) - 0
   r,g,b := Gradient(c1,c2, lol)
 

    for _,line := range lines{
		for i,v := range line {

		fmt.Print(Colorize(string(v),r[i],g[i],b[i]))
	}
}
 

 }