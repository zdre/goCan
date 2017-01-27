package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
	"strconv"
    "math"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convertTs(s string) int {
	time, err := strconv.ParseFloat(s,64)
    check(err)
    return int(time*1000)
}

func atoi (frame []string, pos int) int {
    i, err := strconv.Atoi(frame[pos])
    check(err)
    return i
}

func decn(num int, pos int, negate bool) string {
    divisor := int64(math.Pow10(pos))
    //fmt.Println("div:", divisor)
    return strconv.FormatInt(int64(num)/divisor,10) + "." + strconv.FormatInt(int64(num) % divisor,10)
}

func dec(num int, pos int) string {
    return decn(num,pos,false)
}

func main() {
	file, err := os.Open("/work/data.txt")
	check(err)
	defer file.Close()


    var nomEnergyRemain string
    var expEnergyRemain string
    var idealEnergyRemain string
    var nomPackFullEnergy string
    var energyTillChargeDone string
    var socUI string
    var battOdo string
    var packVolt string

	scanner := bufio.NewScanner(file)
	for i:=0; scanner.Scan() && i<100000; i++ {
		line := scanner.Text()
	    split := strings.Split(line,",")
        //fmt.Println(split)
        //timestamp := convertTs(split[0])
        //fmt.Println(timestamp)

        // code, err := strconv.ParseInt(split[1],16,16)
        // check(err)
        code := split[1]
        frame := split[3:]

        //fmt.Println(code)
        //fmt.Println(frame)

        switch code {
        case "382":
            nomEnergyRemain = dec((atoi(frame,1)>>2) + ((atoi(frame,2) & 0x0F) * 64),1)
            expEnergyRemain = dec((atoi(frame,2)>>4) + ((atoi(frame,3) & 0x3F) * 16),1)
            idealEnergyRemain = dec((atoi(frame,3)>>6) + ((atoi(frame,4) & 0xFF) * 4),1)
            nomPackFullEnergy = dec(atoi(frame,0) + ((atoi(frame,1) & 0x03)<<8),1)
            energyTillChargeDone = dec(atoi(frame,5) + ((atoi(frame,6) & 0x03)<<8),1)
        case "302":
            socUI = dec((atoi(frame,1)>>2) + ((atoi(frame,2) & 0xF)<<6),1)
        case "562": 
            battOdo = dec(atoi(frame,0) + (atoi(frame,1)<<8) + (atoi(frame,2)<<16) + (atoi(frame,3)<<24),3)
        case "102":
            packVolt = dec(atoi(frame,0) + atoi(frame,1)*256,2)
            f3:= atoi(frame,3)
            var exp = 1
            if f3 & 0x80 == 1 {
                exp = -1
            }
           fmt.Println("packCurr" , dec(atoi(frame,2) + (f3 & 0x7F)*256 - 10000,exp))
        }
	}
        fmt.Println("nomEnergyRemain: ", nomEnergyRemain)
        fmt.Println("expEnergyRemain: ", expEnergyRemain)
        fmt.Println("idealEnergyRemain: ", idealEnergyRemain)
        fmt.Println("nomPackFullEnergy: ", nomPackFullEnergy)
        fmt.Println("energyTillChargeDone: ", energyTillChargeDone)
        fmt.Println("socUI: ", socUI)
        fmt.Println("battOdo: ", battOdo)
        fmt.Println("packVolt: ", packVolt)

}
