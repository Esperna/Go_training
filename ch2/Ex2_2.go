package main

import(
	"fmt"
	"os"
	"strconv"
	"gopl.io/ch2/tempconv"
)

func main(){
	if len(os.Args) > 1{
		for _, arg := range os.Args[1:]{
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v",err)
				os.Exit(1)
			}
			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)
			//m := tempconv.meter(t)
			//ft := tempconv.feet(t)
			//pd := tempconv.pond(t)
			//kg := tempconv.kilogram(t)
			fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
			//fmt.Printf("%s = %s, %s = %s\n", m, tempconv.MToFT(m), ft, tempconv.FTtoMT(ft))
			//fmt.Printf("%s = %s, %s = %s\n", pd, tempconv.PDToKG(pd), kg, tempconv.KGToPD(kg))
		}
	}else {
		var t float64
		fmt.Scan(&t)
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
