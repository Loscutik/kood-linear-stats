package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"01.kood.tech/git/obudarah/linear-stats/statistics"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Wrong numbers of arguments. A name of file with values of data is needed")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	Y, err := ReadNumbers(file)
	if err != nil {
		log.Fatal(err)
	}
	X:=make([]float64,len(Y))
	for i := 0; i < len(Y); i++ {
		X[i]=float64(i)
	}
	
	// sums needed for calculating the coefficients of the Linear Regression
	n:=float64(len(Y))
	Sx:=float64(n*(n-1))/2.0 // the sum of Xi={0,1,...,n-1}
	Sy, Sxy, Sx2:=0.0,0.0,0.0
	for i := 0; i < len(Y); i++ {
		Sy+=Y[i]
		Sxy+=float64(i)*Y[i]
		Sx2+=float64(i)*float64(i)
	}

	// the coefficients of the Linear Regression
	k1:=(n*Sxy-Sx*Sy)/(n*Sx2-Sx*Sx)
	k0:=(Sy-k1*Sx)/n
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n",k1,k0)
	
	PearsonCorrelationCoefficent:=statistics.Covariance(X,Y)/(statistics.StandardDeviation(X)*statistics.StandardDeviation(Y))
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n",PearsonCorrelationCoefficent)
	
}

/*
reads file line by line, convert numbers to int and returns slice of the numbers
*/
func ReadNumbers(file *os.File) (data []float64, err error) {
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	
	for fileScanner.Scan() {
		num,err:= strconv.ParseFloat(fileScanner.Text(),64)
		if err != nil {
			return nil, err
		}
		data = append(data, num)
	}
	err = fileScanner.Err()

	return
}

