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
    // Open and process 'test.txt' in both modes
    sumOfTestGameIDsCumulative, sumOfTestGameIDsReset := processBothModes("test.txt")

    // Open and process 'input1.txt' in both modes
    sumOfInput1GameIDsCumulative, sumOfInput1GameIDsReset := processBothModes("input1.txt")

    // Print results
    fmt.Println("Sum of IDs of possible games in 'test.txt' (cumulative count):", sumOfTestGameIDsCumulative)
    fmt.Println("Sum of IDs of possible games in 'test.txt' (reset count):", sumOfTestGameIDsReset)
    fmt.Println("Sum of IDs of possible games in 'input1.txt' (cumulative count):", sumOfInput1GameIDsCumulative)
    fmt.Println("Sum of IDs of possible games in 'input1.txt' (reset count):", sumOfInput1GameIDsReset)

    // Open and process the file for the second part of the puzzle
    sumOfPowers := processFileForMinSet("input1.txt")
    fmt.Println("Sum of the power of the minimum sets:", sumOfPowers)
}

func processFileForMinSet(filename string) int {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return 0
    }
    defer file.Close()

    var sumOfPowers int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        game := scanner.Text()
        power := calculateMinSetPower(game)
        sumOfPowers += power
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading from file:", err)
    }

    return sumOfPowers
}

func calculateMinSetPower(game string) int {
    parts := strings.Split(game, ";")
    minSet := CubeCount{}

    for _, part := range parts {
        partRequired := CubeCount{}
        processPart(part, &partRequired)
        minSet.red = max(minSet.red, partRequired.red)
        minSet.green = max(minSet.green, partRequired.green)
        minSet.blue = max(minSet.blue, partRequired.blue)
    }

    return minSet.red * minSet.green * minSet.blue
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func processBothModes(filename string) (int, int) {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return 0, 0
    }
    defer file.Close()

    // Process file in cumulative mode
    sumOfGameIDsCumulative := processFile(file, false)

    // Re-open file for reset mode processing
    file, err = os.Open(filename)
    if err != nil {
        fmt.Println("Error re-opening file:", err)
        return sumOfGameIDsCumulative, 0
    }
    defer file.Close()

    // Process file in reset mode
    sumOfGameIDsReset := processFile(file, true)

    return sumOfGameIDsCumulative, sumOfGameIDsReset
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
