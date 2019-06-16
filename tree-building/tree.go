package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}
type Node struct {
	ID       int
	Children []*Node
}

func areUniqueRecords(records []Record) bool {
	f := make(map[int]int)
	for _, record := range records {
		f[record.ID]++
	}
	for _, v := range f {
		if v > 1 {
			return false
		}
	}
	return true
}

func Build(records []Record) (*Node, error) {
	lenght := len(records)
	if lenght == 0 {
		return nil, nil
	}
	if !areUniqueRecords(records) {
		return nil, errors.New("Found duplicates")
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	if records[0].ID != 0 || records[0].Parent != 0 {
		return nil, errors.New("Please check parent node.")
	}

	nodes := make([]*Node, lenght)
	for index, record := range records {
		if record.Parent > record.ID {
			return nil, errors.New("Higher id parent of lower id")
		}
		if record.ID != index {
			return nil, errors.New("Non-contiguous input")
		}
		node := Node{
			ID: record.ID,
		}
		nodes[index] = &node

	}
	for index, record := range records {
		current := nodes[record.ID]
		parent := nodes[record.Parent]
		if record.ID == record.Parent {
			if index == 0 {
				continue
			}
			return nil, errors.New("Cycle detected")
		}
		children := append(parent.Children, current)
		parent.Children = children

	}
	return nodes[0], nil
}
