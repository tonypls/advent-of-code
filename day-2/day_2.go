package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)

type CubeCount struct {
    red   int
    green int
    blue  int
}

func main() {
    // Open the first file
    file1, err := os.Open("test.txt")
    if err != nil {
        fmt.Println("Error opening file1:", err)
        return
    }
    defer file1.Close()

    // Open the second file
    file2, err := os.Open("input1.txt")
    if err != nil {
        fmt.Println("Error opening file2:", err)
        return
    }
    defer file2.Close()

    // Process each file, false indicates cumulative count
    // testSum := processFile(file1, false)
    // fmt.Println("Sum of IDs of possible games (cumulative count):", testSum)

    // Reset the counter for each part of the game
    testSumReset := processFile(file1, false)
    fmt.Println("Sum of IDs of possible games (reset count):", testSumReset)

	  // Process each file, false indicates cumulative count
	//   sumOfPossibleGameIDs :=  processFile(file2, false)
	//   fmt.Println("Sum of IDs of possible games (cumulative count):", sumOfPossibleGameIDs)
  
	  // Reset the counter for each part of the game
	  sumOfPossibleGameIDsReset :=  processFile(file2, true)
	  fmt.Println("Sum of IDs of possible games (reset count):", sumOfPossibleGameIDsReset)
}

func processFile(file *os.File, resetAfterEachPart bool) int {
    var sumOfPossibleGameIDs int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        game := scanner.Text()
        if canPlayGame(game, resetAfterEachPart) {
            // fmt.Println(game, getGameID(game), "is possible")
            sumOfPossibleGameIDs += getGameID(game)
        } else {
            // fmt.Println(game, "is not possible")
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading from file:", err)
    }

    return sumOfPossibleGameIDs
}

func canPlayGame(game string, resetAfterEachPart bool) bool {
    // Available cubes
    available := CubeCount{12, 13, 14}

    // Parse the game string
    parts := strings.Split(game, ";")
    totalRequired := CubeCount{}

    for _, part := range parts {
        partRequired := CubeCount{}

        // Process each part
        if !processPart(part, &partRequired) {
            return false
        }
		// fmt.Println("Part required:", available.red, available.green, available.blue)
        if resetAfterEachPart {
            // In reset mode, check if each part can be played independently
            if partRequired.red > available.red || partRequired.green > available.green || partRequired.blue > available.blue {
                return false
            }
        } else {
            // In cumulative mode, accumulate the counts and check if the total can be played
            totalRequired.red += partRequired.red
            totalRequired.green += partRequired.green
            totalRequired.blue += partRequired.blue

            if totalRequired.red > available.red || totalRequired.green > available.green || totalRequired.blue > available.blue {
                return false
            }
        }
    }

    // In reset mode, always return true as long as each part is individually valid
    return true
}


func processPart(part string, totalRequired *CubeCount) bool {
    // Regex to find cube counts and colors
    re := regexp.MustCompile(`(\d+)\s*(red|green|blue)`)
    matches := re.FindAllStringSubmatch(part, -1)
    for _, match := range matches {
        if len(match) != 3 {
            continue
        }
        count, _ := strconv.Atoi(match[1])
        color := match[2]
        switch color {
        case "red":
            totalRequired.red += count
        case "green":
            totalRequired.green += count
        case "blue":
            totalRequired.blue += count
        }
    }
    return true
}


func getGameID(game string) int {
    parts := strings.Split(game, ":")
    gameName := strings.Split(parts[0], " ")
    if len(gameName) > 1 {
        id, err := strconv.Atoi(gameName[1])
        if err == nil {
            return id
        }
    }
    return 0
}
