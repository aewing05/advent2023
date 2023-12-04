package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "strconv"
)

func parseColorResult(input string, color string) []int {
    if color != "blue" && color != "green" && color != "red" {
        return []int{}
    }


    pattern := regexp.MustCompile(`(\d+)\s+` + regexp.QuoteMeta(color))
    matches := pattern.FindAllStringSubmatch(input, -1)
    var result []int
    for _, match := range matches {
        num, err := strconv.Atoi(match[1])
        if err == nil {
            result = append(result, num)
        }
    }

    return result
}

func parseGameID(input string) (int, bool) {
    gamePattern := regexp.MustCompile(`Game+\s(\d+)`)
    matches := gamePattern.FindStringSubmatch(input)
    if len(matches) > 0 {
        num, err := strconv.Atoi(matches[1])
        if err == nil {
           return num, true 
        }
    } 

    return 0, false
}

func part1 () {
    file, err := os.Open("day2.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)

    // For each game (line), figure out the max number of each color
    // and increment where the maxes are less than 12 red cubes,
    // 13 green cubes, and 14 blue cubes
   
    const blueMax = 14
    const greenMax = 13
    const redMax = 12
    possible_game_sum := 0

    for scanner.Scan() {
        input := scanner.Text()

        blueResult := parseColorResult(input, "blue")
        greenResult := parseColorResult(input, "green")
        redResult := parseColorResult(input, "red")

        blueEligible := true
        greenEligible := true
        redEligible := true

        for _, result := range blueResult {
            if result > blueMax {
                blueEligible = false
                break
            }
        }

        for _, result := range greenResult {
            if result > greenMax {
                greenEligible = false
                break
            }
        }

        for _, result := range redResult {
            if result > redMax {
                redEligible = false
                break
            }
        }
   

        if blueEligible && greenEligible && redEligible {
            if num, ok := parseGameID(input); ok {
                possible_game_sum += num
            }

        }
    }
    fmt.Println("Possible Games: ", possible_game_sum)
}

func getMaxPositive (arr []int) int {
    max := 0
    for _, num := range arr {
        if num > max {
            max = num
        }
    }
    return max
}

func part2 () {
    file, err := os.Open("day2.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    //
    scanner := bufio.NewScanner(file)
    
    sumOfPowers := 0 
    for scanner.Scan() {
        input := scanner.Text()
        fmt.Println(input)

        blueResult := parseColorResult(input, "blue")
        greenResult := parseColorResult(input, "green")
        redResult := parseColorResult(input, "red")
 
        blueMax := getMaxPositive(blueResult)    
        greenMax := getMaxPositive(greenResult)    
        redMax := getMaxPositive(redResult)    
        
        sumOfPowers += blueMax * greenMax * redMax

    }
    fmt.Println("Total Sum of Powers: ", sumOfPowers) 
}


func main () {
//    part1()
    part2()
}
