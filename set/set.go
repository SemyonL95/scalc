package set

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Set map[int]struct{}

func (s Set) Display() {
	var arr []int
	for num := range s {
		arr = append(arr, num)
	}

	sort.Ints(arr)

	for _, num := range arr {
		fmt.Println(num)
	}
}

// ReadSet from file
func ReadSet(path string) (Set, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening file, %s: %w", path, err)
	}
	defer file.Close()

	var set = make(Set)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("conversion to int: %w", err)
		}

		set[num] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanning file, %s: %w", path, err)
	}

	return set, nil
}

func Union(sets ...Set) Set {
	result := NewSet()
	for _, set := range sets {
		for num := range set {
			result[num] = struct{}{}
		}
	}

	return result
}

func Intersect(sets ...Set) Set {
	result := NewSet()

	if len(sets) == 0 {
		return result
	}

EXTERNAL:
	for num := range sets[0] {
		for _, set := range sets {
			_, ok := set[num]
			if !ok {
				continue EXTERNAL
			}
		}
		result[num] = struct{}{}
	}

	return result
}

// Diff returns difference of first set with sets
func Diff(sets ...Set) Set {
	if len(sets) == 0 {
		return NewSet()
	}

	result := sets[0]

	for num, set := range sets {
		if num == 0 {
			continue
		}

		intersect := Intersect(result, set)

		for key := range intersect {
			delete(result, key)
		}
	}

	return result
}

func NewSet(ints ...int) Set {
	set := make(Set)
	for _, num := range ints {
		set[num] = struct{}{}
	}

	return set
}

// ParseExpression returns result of expression if it is valid
func ParseExpression(tokens []string) (Set, error) {

	// Expression must at least contain "[", "]", "SUM|DIF|INT" and set which is in turn at least one token
	if len(tokens) < 4 {
		return nil, fmt.Errorf("invalid minimal length of expresion")
	}

	if tokens[0] != "[" || tokens[len(tokens)-1] != "]" {
		return nil, fmt.Errorf("invalid open and close delimeters of expresions")
	}

	setsTokens := tokens[2 : len(tokens)-1]

	sets, err := ParseSets(setsTokens)
	if err != nil {
		return nil, err
	}

	switch tokens[1] {
	case "SUM":
		return Union(sets...), nil
	case "DIF":
		return Diff(sets...), nil
	case "INT":
		return Intersect(sets...), nil
	default:
		return nil, fmt.Errorf("Invalid operator value: %s", tokens[1])
	}
}

func ParseSets(tokens []string) ([]Set, error) {
	var result []Set

	if len(tokens) < 1 {
		return result, nil
	}

	if tokens[0] != "[" {
		set, err := ReadSet(tokens[0])

		if err != nil {
			return nil, err
		}

		result = append(result, set)

		sets, err := ParseSets(tokens[1:])

		if err != nil {
			return nil, err
		}

		result = append(result, sets...)

		return result, nil
	}

	if tokens[0] == "[" {
		index, err := FindEndOfSet(tokens)

		if err != nil {
			return nil, err
		}

		set, err := ParseExpression(tokens[0 : index+1])

		if err != nil {
			return nil, err
		}

		result = append(result, set)

		sets, err := ParseSets(tokens[index+1:])

		if err != nil {
			return nil, err
		}

		result = append(result, sets...)

		return result, nil
	}

	return nil, fmt.Errorf("invalid open delimeter: %s", tokens[0])
}

func FindEndOfSet(tokens []string) (int, error) {
	counter := 0

	for key, val := range tokens {

		if val == "[" {
			counter++
		}

		if val == "]" {
			counter--
		}

		if counter == 0 {
			return key, nil
		}
	}

	return 0, fmt.Errorf("error ivalid expresion: %q", tokens)
}
