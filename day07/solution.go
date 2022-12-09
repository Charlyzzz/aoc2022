package day07

import (
	"aoc2022/challenge"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type c struct{}

func (c) Part1() interface{} {
	input := challenge.Input()
	return sumDirectoriesWithSizeUnder(input)
}

func (c) Part2() interface{} {
	input := challenge.Input()

	return smallestDirectorySizeToDelete(input)
}

func Challenge() challenge.Challenge {
	return c{}
}

type directory struct {
	name        string
	parent      *directory
	directories map[string]*directory
	files       int
	size        int
}

func (d *directory) cd(dir string) *directory {
	if dir == ".." {
		return d.mustGetParent()
	}
	return d.directories[dir]
}

func (d *directory) mustGetParent() *directory {
	if d.parent == nil {
		log.Fatal("parent dir was nil")
	}
	return d.parent
}

func (d *directory) addDir(dirName string) {
	dir := newDirectory(d, dirName)
	d.directories[dirName] = dir
}

func (d *directory) addFile(size int) {
	d.files += size
}

func (d *directory) getSize() int {
	if d.hasSize() {
		return d.size
	}
	dirsSize := 0
	for _, dir := range d.directories {
		dirsSize += dir.getSize()
	}
	d.size = d.files + dirsSize
	return d.size
}

func (d *directory) sumSizesUnder(sizeUnder int) int {
	sum := 0
	for _, dir := range d.directories {
		sum += dir.sumSizesUnder(sizeUnder)
		size := dir.getSize()
		if size <= sizeUnder {
			sum += size
		}
	}
	return sum
}

func (d *directory) closestDirToSize(sizeToDelete int) *directory {
	var bestChildren *directory
	for _, dir := range d.directories {
		size := dir.getSize()
		if size >= sizeToDelete {
			closestChildren := dir.closestDirToSize(sizeToDelete)
			bestChildren = closerChildren(sizeToDelete, bestChildren, closestChildren)
		}
	}
	return closerChildren(sizeToDelete, d, bestChildren)
}

func (d *directory) hasSize() bool {
	return d.size != -1
}

func closerChildren(sizeToDelete int, children1 *directory, children2 *directory) *directory {
	if children1 == nil && children2 == nil {
		return nil
	}
	if children1 == nil {
		return children2
	}
	if children2 == nil {
		return children1
	}
	if (children1.getSize() - sizeToDelete) > (children2.getSize() - sizeToDelete) {
		return children2
	}
	return children1
}

func sumDirectoriesWithSizeUnder(input string) int {
	sizeUnder := 100_000
	root := parseFilesystem(input)
	return root.sumSizesUnder(sizeUnder)
}

func smallestDirectorySizeToDelete(input string) int {
	totalSpace := 70_000_000
	requiredSpace := 30_000_000
	root := parseFilesystem(input)
	unusedSpace := totalSpace - root.getSize()
	deleteSpace := requiredSpace - unusedSpace
	dir := root.closestDirToSize(deleteSpace)
	return dir.getSize()
}

func newDirectory(parent *directory, name string) *directory {
	return &directory{
		name:        name,
		parent:      parent,
		directories: make(map[string]*directory),
		files:       0,
		size:        -1,
	}
}

func parseFilesystem(input string) *directory {
	root := newDirectory(nil, "/")
	current := root
	for _, command := range strings.Split(input, "\n") {
		if isRoot(command) || isLs(command) {
			continue
		}
		if ok, dir := isCdInto(command); ok {
			current = current.cd(dir)
			continue
		}
		if ok, name := isDir(command); ok {
			current.addDir(name)
			continue
		}
		if ok, size := isFile(command); ok {
			current.addFile(size)
			continue
		}
		log.Fatalf("unknown command: %s", command)
	}
	return root
}

var fileRegexp = regexp.MustCompile(`(\d+)`)

func isFile(command string) (bool, int) {
	match := fileRegexp.FindStringSubmatch(command)
	if match == nil {
		return false, 0
	}
	size, err := strconv.Atoi(match[1])
	if err != nil {
		log.Fatalf("error parsing [%s] %v", command, err)
	}
	return true, size
}

var dirRegexp = regexp.MustCompile(`dir (\w+)`)

func isDir(command string) (bool, string) {
	match := dirRegexp.FindStringSubmatch(command)
	if match == nil {
		return false, ""
	}
	return true, match[1]
}

var cdRegexp = regexp.MustCompile(`cd (\w+|\.\.)`)

func isCdInto(command string) (bool, string) {
	match := cdRegexp.FindStringSubmatch(command)
	if match == nil {
		return false, ""
	}
	return true, match[1]
}

func isLs(command string) bool {
	return command == "$ ls"
}

func isRoot(command string) bool {
	return command == "$ cd /"
}
