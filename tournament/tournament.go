package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

type matchMap map[string]int

func contains(list []string, element string) bool {
	for _, value := range list {
		if value == element {
			return true
		}
	}
	return false
}

func PadRight(str, pad string, lenght int) string {
	for {
		str += pad
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}

func sorted(standings []string) {
	for _, standing := range standings {
		teamName := strings.TrimSpace(strings.Split(standing, "|")[0])
		teamPoints := strings.TrimSpace(strings.Split(standing, "|")[5])
		combined := teamPoints + teamName
	}

}

func Tally(input io.Reader, output io.Writer) error {

	wins, draws, losses, points := make(matchMap), make(matchMap), make(matchMap), make(matchMap)
	var teams []string
	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)
	for scanner.Scan() {
		match := scanner.Text()
		if strings.HasPrefix(match, "#") || match == "\n" || match == "" {
			continue
		}
		if len(strings.Split(match, ";")) != 3 {
			return errors.New("Line needs to have 3 parts")
		}
		parts := strings.Split(match, ";")
		teamA, teamB, result := parts[0], parts[1], parts[2]
		for _, team := range parts[0:1] {
			if !contains(teams, team) {
				teams = append(teams, team)
			}
		}
		switch result {
		case "win":
			wins[teamA]++
			losses[teamB]++
		case "loss":
			losses[teamA]++
			wins[teamB]++
		case "draw":
			draws[teamA]++
			draws[teamB]++
		default:
			return errors.New("Invalid result")
		}
	}
	var standings []string
	for _, team := range teams {
		total := wins[team] + losses[team] + draws[team]
		points[team] = 3*wins[team] + 0*losses[team] + 1*draws[team]
		fmtString := fmt.Sprintf("%s|  %d |  %d |  %d |  %d|  %d\n", PadRight(team, " ", 31), total, wins[team], draws[team], losses[team], points)
		standings = append(standings, fmtString)
	}
	standings = sorted(standings)
	//if err := writer.Flush(); err != nil {
	//	return errors.New("Unable to write to output writer")
	//}
	return nil
}
