package twok

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

var id = 0

type dir struct {
	id   int
	name string
	size int
	dirs []*dir
}

func newDir(name string) *dir {
	id++
	return &dir{
		id:   id,
		name: name,
		size: 0,
		dirs: []*dir{},
	}
}

func (d *dir) insertDir(name string) {
	d.dirs = append(d.dirs, newDir(name))
}

func (d *dir) findSubDirByName(name string) *dir {
	for _, subDir := range d.dirs {
		if subDir.name == name {
			return subDir
		}
	}

	for _, subDir := range d.dirs {
		dir := subDir.findSubDirByName(name)
		if dir != nil {
			return dir
		}
	}

	return nil
}

func (d *dir) findParentOfDirByID(id int) *dir {
	for _, subDir := range d.dirs {
		if subDir.id == id {
			return d
		}
	}

	for _, subDir := range d.dirs {
		dir := subDir.findParentOfDirByID(id)
		if dir != nil {
			return dir
		}
	}

	return nil
}

func (d *dir) updateSize(size int) {
	d.size = size
}

var used = 0
var total = 0
var smallest = 9999999999999999

func (d *dir) sum() int {
	dirSum := d.size
	for _, subDir := range d.dirs {
		dirSum += subDir.sum()
	}
	if dirSum <= 100000 {
		total += dirSum
	}
	return dirSum
}

func (d *dir) diskUsed() {
	used += d.size
	for _, subDir := range d.dirs {
		subDir.diskUsed()
	}
}

func (d *dir) findMinToDelete(min int) int {
	dirSum := d.size
	for _, subDir := range d.dirs {
		dirSum += subDir.findMinToDelete(min)
	}
	if dirSum >= min && dirSum < smallest {
		smallest = dirSum
	}
	return dirSum
}

func Seven(input string) ([2]interface{}, error) {
	root := newDir("/")
	currDir := root

	exp := `(?s)(.*?)\n(.*?)\$ cd`
	re := regexp.MustCompile(exp)
	dirChanges := re.FindAllStringSubmatch(input, -1)
	for k, change := range dirChanges {
		if k != 0 {
			name := strings.TrimSpace(change[1])
			if name == ".." {
				currDir = root.findParentOfDirByID(currDir.id)
				continue
			}
			currDir = currDir.findSubDirByName(name)
		}

		scanner := bufio.NewScanner(strings.NewReader(change[2]))
		for scanner.Scan() {
			lineArray := strings.Split(scanner.Text(), " ")
			if lineArray[0] == "$" {
				continue
			} else if lineArray[0] == "dir" {
				currDir.insertDir(lineArray[1])
			} else {
				fileSize, _ := strconv.Atoi(lineArray[0])
				currDir.updateSize(currDir.size + fileSize)
			}
		}
	}
	currDir = currDir.findSubDirByName("zrfmzrl")
	currDir.updateSize(currDir.size + 52067)

	root.sum()
	root.diskUsed()
	unused := 70000000 - used
	needFree := 30000000 - unused
	root.findMinToDelete(needFree)

	return [2]interface{}{total, smallest}, nil
}
