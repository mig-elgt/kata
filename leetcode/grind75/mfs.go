package grind75

// 588. Design In-Memory File System

// Design a data structure that simulates an in-memory file system.

// Implement the FileSystem class:

// FileSystem() Initializes the object of the system.
// List<String> ls(String path)
// If path is a file path, returns a list that only contains this file's name.
// If path is a directory path, returns the list of file and directory names in this directory.
// The answer should in lexicographic order.
// void mkdir(String path) Makes a new directory according to the given path. The given directory path does not exist. If the middle directories in the path do not exist, you should create them as well.
// void addContentToFile(String filePath, String content)
// If filePath does not exist, creates that file containing given content.
// If filePath already exists, appends the given content to original content.
// String readContentFromFile(String filePath) Returns the content in the file at filePath.

// Input
// ["FileSystem", "ls", "mkdir", "addContentToFile", "ls", "readContentFromFile"]
// [[], ["/"], ["/a/b/c"], ["/a/b/c/d", "hello"], ["/"], ["/a/b/c/d"]]
// Output
// [null, [], null, null, ["a"], "hello"]

// Explanation
// FileSystem fileSystem = new FileSystem();
// fileSystem.ls("/");                         // return []
// fileSystem.mkdir("/a/b/c");
// fileSystem.addContentToFile("/a/b/c/d", "hello");
// fileSystem.ls("/");                         // return ["a"]
// fileSystem.readContentFromFile("/a/b/c/d"); // return "hello"

// Constraints:

// 1 <= path.length, filePath.length <= 100
// path and filePath are absolute paths which begin with '/' and do not end with '/' except that the path is just "/".
// You can assume that all directory names and file names only contain lowercase letters, and the same names will not exist in the same directory.
// You can assume that all operations will be passed valid parameters, and users will not attempt to retrieve file content or list a directory or file that does not exist.
// 1 <= content.length <= 50
// At most 300 calls will be made to ls, mkdir, addContentToFile, and readContentFromFile.

import (
	"sort"
	"strings"
)

type FileSystemNode struct {
	files map[string]string
	dirs  map[string]*FileSystemNode
}

type FileSystem struct {
	root *FileSystemNode
}

func NewFileSystem() FileSystem {
	return FileSystem{
		root: &FileSystemNode{
			files: map[string]string{},
			dirs:  map[string]*FileSystemNode{},
		},
	}
}

func (this *FileSystem) Ls(path string) []string {
	if path == "/" {
		res := []string{}
		for fileName := range this.root.files {
			res = append(res, fileName)
		}
		for dirName := range this.root.dirs {
			res = append(res, dirName)
		}
		sort.Strings(res)
		return res
	}
	var items, dirs []string
	items = strings.Split(path[1:], "/")
	dirs = items[:len(items)-1]
	currNode := this.root
	for _, dir := range dirs {
		if _, ok := currNode.dirs[dir]; !ok {
			return []string{}
		}
		currNode = currNode.dirs[dir]
	}
	lastItem := items[len(items)-1]
	if _, ok := currNode.files[lastItem]; ok {
		return []string{lastItem}
	}
	var res []string
	if _, ok := currNode.dirs[lastItem]; ok {
		node := currNode.dirs[lastItem]
		for fileName := range node.files {
			res = append(res, fileName)
		}
		for dirName := range node.dirs {
			res = append(res, dirName)
		}
	}
	sort.Strings(res)
	return res
}

func (this *FileSystem) Mkdir(path string) {
	currNode := this.root
	dirs := strings.Split(path[1:], "/")
	for _, dir := range dirs {
		if _, found := currNode.dirs[dir]; !found {
			currNode.dirs[dir] = &FileSystemNode{
				files: map[string]string{},
				dirs:  map[string]*FileSystemNode{},
			}
		}
		currNode = currNode.dirs[dir]
	}
}

func (this *FileSystem) AddContentToFile(filePath string, content string) {
	path := strings.Split(filePath[1:], "/")
	dirs := path[:len(path)-1]
	fileName := path[len(path)-1]
	currNode := this.root
	for _, dir := range dirs {
		if _, found := currNode.dirs[dir]; !found {
			return
		}
		currNode = currNode.dirs[dir]
	}
	currNode.files[fileName] += content
}

func (this *FileSystem) ReadContentFromFile(filePath string) string {
	path := strings.Split(filePath[1:], "/")
	dirs := path[:len(path)-1]
	fileName := path[len(path)-1]
	currNode := this.root
	for _, dir := range dirs {
		if _, found := currNode.dirs[dir]; !found {
			return ""
		}
		currNode = currNode.dirs[dir]
	}
	return currNode.files[fileName]
}
