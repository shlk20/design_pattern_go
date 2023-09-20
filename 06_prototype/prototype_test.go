package prototype

import (
	"fmt"
	"testing"
)

func TestFolder(t *testing.T) {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		children: []INode{file1},
		name:     "folder1",
	}

	folder2 := &Folder{
		name: "folder2",
		children: []INode{
			folder1, file2, file3,
		},
	}

	fmt.Println("Printing hierarchy for Folder2")
	folder2.print("  ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")
}
