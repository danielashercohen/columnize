package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
)

func main() {
    var columnWidth map[int]int   
    columnWidth = make (map[int]int)
    
    content, err := ioutil.ReadFile("./test/test")
    if err == nil {
        lines := strings.Split(string(content), "\n")
        // process each line looking at each column and keeping a record of the 
        // widest value seen in each column
        for _, element := range lines {
        	columns := strings.Fields(string(element))
        	for columnIndex, columnValue := range columns {
        		if column, ok := columnWidth[columnIndex]; ok {
        			if column < len(columnValue) {
        				columnWidth[columnIndex] = len(columnValue)
        			}
        		} else {
        			columnWidth[columnIndex] = len(columnValue)
        		}
        	}
        }
        // Loop through the lines again,  printing each column entry with trailing spaces
        // to pad the column to the widest value seen in that column (plus a space)
        for _, element := range lines {
        	columns := strings.Fields(string(element))
        	for columnIndex, columnValue := range columns {
        		// print each column value on this line followed by the appropriate number of 
        		// trailing spaces
        	    fmt.Print(columnValue, strings.Repeat(" ", columnWidth[columnIndex]-len(columnValue)+1)) 
        	}
        	fmt.Println();
        }
    } else {
        log.Fatal(err)
    } 
  

}



