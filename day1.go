package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "strconv"
    "unicode"
)

var conversion_map = map[string]string{
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
    "zero": "0",
}

var reversed_conversion_map = map[string]string{
    "eno": "1",
    "owt": "2",
    "eerht": "3",
    "ruof": "4",
    "evif": "5",
    "xis": "6",
    "neves": "7",
    "thgie": "8",
    "enin": "9",
    "orez": "0",
}

func reverseString(input string) string {
    runes := []rune(input)

    // Use two cursors and swap
    for i, j := 0, len(runes) - 1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }

    return string(runes)
}

// func replaceFirstInstanceString (input string, replacement_map map[string]string) string {
    // TODO: Take string and replacement map,
    // and return the string with the first 
    // match replaced, based on the replacement
    // map

// }

func firstDigitInString (input string) int {
    for _, char := range input {
        if unicode.IsDigit(char) {
            num, err := strconv.Atoi(string(char))
            if err != nil {
                log.Fatal("Fatal Error:", err)
            }

            return num
        }
    }

    // Return 0 if no digit is found
    return 0
}

func part1 () {
    file, err := os.Open("day1.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    total_sum := 0
    for scanner.Scan() {
        calibration_string := scanner.Text()
        num := firstDigitInString(calibration_string) 
        total_sum += (10 * num)

        rev_calibration_string := reverseString(calibration_string)  
        rev_num := firstDigitInString(rev_calibration_string)
        total_sum += rev_num
    }

    fmt.Println(total_sum)  

    if err:= scanner.Err(); err != nil {
        log.Fatal(err)
    }


}

func part2() {
// Going to copy code from part 1, but as a first step do a regex replace
    file, err := os.Open("day1.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    total_sum := 0

 pattern := regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine|zero)")
 reverse_pattern := regexp.MustCompile("(eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|orez)")
 
    for scanner.Scan() {
        calibration_string := scanner.Text()
        rev_calibration_string := reverseString(calibration_string) 

        // Update strings


        updated_string := pattern.ReplaceAllStringFunc(
            calibration_string,
            func(match string) string {return conversion_map[match]})

        reverse_updated_string := reverse_pattern.ReplaceAllStringFunc(
            rev_calibration_string,
            func(match string) string {return reversed_conversion_map[match]})


        num := firstDigitInString(updated_string) 
        total_sum += (10 * num)

        rev_num := firstDigitInString(reverse_updated_string)
        total_sum += rev_num
    }


    if err:= scanner.Err(); err != nil {
        log.Fatal(err)
    }


    fmt.Println(total_sum)  

}

func main () {
//    part1()
    part2()
}
