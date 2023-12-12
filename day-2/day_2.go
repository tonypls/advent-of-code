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
	  // Open the file
	  file1, err := os.Open("test.txt") // Replace 'games.txt' with your actual file name
	  if err != nil {
		  fmt.Println("Error opening file:", err)
		  return
	  }
	  defer file.Close()
    // Open the file
    file2, err := os.Open("input1.txt") // Replace 'games.txt' with your actual file name
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    var sumOfPossibleGameIDs int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        game := scanner.Text()
        if canPlayGame(game) {
            fmt.Println(game,getGameID(game), "is possible")
            gameID := getGameID(game)
            sumOfPossibleGameIDs += gameID
        } else {
            fmt.Println(game, "is not possible")
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading from file:", err)
    }

    fmt.Println("Sum of IDs of possible games:", sumOfPossibleGameIDs)
}

func canPlayGame(game string) bool {
    // Available cubes
    available := CubeCount{12, 13, 14}

    // Parse the game string
    parts := strings.Split(game, ";")
    totalRequired := CubeCount{}

    // Regex to find cube counts and colors
    re := regexp.MustCompile(`(\d+)\s*(red|green|blue)`)

    for _, part := range parts {
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
    }

    // Check if the game can be played
    return totalRequired.red <= available.red && totalRequired.green <= available.green && totalRequired.blue <= available.blue
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
