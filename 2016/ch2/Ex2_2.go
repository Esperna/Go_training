package main

import(
	"fmt"
	"os"
	"strconv"
	"gopl.io/ch2/tempconv_Ex2_2"
)

func main(){
	if len(os.Args) > 1{
		for _, arg := range os.Args[1:]{
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v",err)
				os.Exit(1)
			}
			printParameter(t)
		}
	}else {
		var t float64
		fmt.Scan(&t)
		printParameter(t)
	}
}

func printParameter(t float64){
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	m := tempconv.Meter(t)
	ft := tempconv.Feet(t)
	pd := tempconv.Pond(t)
	kg := tempconv.Kilogram(t)
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	fmt.Printf("%s = %s, %s = %s\n", m, tempconv.MToFT(m), ft, tempconv.FTToMT(ft))
	fmt.Printf("%s = %s, %s = %s\n", pd, tempconv.PDToKG(pd), kg, tempconv.KGToPD(kg))
}