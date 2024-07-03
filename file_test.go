package iutils

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestDirExists(t *testing.T) {

	tempDirPath, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %s", err)
	}
	defer os.RemoveAll(tempDirPath)

	if !DirExists(tempDirPath) {
		t.Errorf("Expected directory to exist at path: %s", tempDirPath)
	}

	if DirExists("/path/to/nonexistent/directory") {
		t.Errorf("Expected directory to not exist at path: /path/to/nonexistent/directory")
	}
}

func TestFileExists(t *testing.T) {

	tmpFile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %s", err)
	}
	defer os.Remove(tmpFile.Name())

	exists := FileExists(tmpFile.Name())
	if !exists {
		t.Errorf("FileExists() = %v for file %s, want %v", exists, tmpFile.Name(), true)
	}

	notExists := FileExists("nonexistentfile.txt")
	if notExists {
		t.Errorf("FileExists() = %v for file nonexistentfile.txt, want %v", notExists, false)
	}
}

func TestGetFileSize(t *testing.T) {

	existingFile := "testfile.txt"
	f, _ := os.Create(existingFile)
	defer f.Close()
	defer os.Remove(existingFile)

	testData := []byte("Hello, World!")
	f.Write(testData)

	expectedSize := int64(len(testData))
	actualSize, err := GetFileSize(existingFile)
	if err != nil {
		t.Errorf("GetFileSize returned an error for an existing file: %v", err)
	}
	if actualSize != expectedSize {
		t.Errorf("GetFileSize = %d, want %d for an existing file", actualSize, expectedSize)
	}

	nonExistingFile := "nonexistingfile.txt"
	_, err = GetFileSize(nonExistingFile)
	if err == nil {
		t.Errorf("GetFileSize did not return an error for a non-existing file")
	}
}

func TestReadFile(t *testing.T) {

	content := []byte("Hello, World!")
	filename := "testfile.txt"
	err := os.WriteFile(filename, content, 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(filename)

	result, err := ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if !bytes.Equal(result, content) {
		t.Errorf("File content does not match. Expected: %s, Got: %s", string(content), string(result))
	}
}

func TestReadFileParts(t *testing.T) {

	tempFile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	content := []byte("Hello, World!")
	if _, err := tempFile.Write(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	tests := []struct {
		filename    string
		fileoffset  int64
		length      int64
		expected    []byte
		expectedErr bool
	}{
		{
			filename:    tempFile.Name(),
			fileoffset:  0,
			length:      int64(len(content)),
			expected:    content,
			expectedErr: false,
		},
		{
			filename:    tempFile.Name(),
			fileoffset:  0,
			length:      int64(len(content)) + 10,
			expected:    content,
			expectedErr: false,
		},
		{
			filename:    tempFile.Name(),
			fileoffset:  int64(len(content)) + 10,
			length:      int64(len(content)),
			expected:    nil,
			expectedErr: false,
		},
		{
			filename:    "nonexistent.txt",
			fileoffset:  0,
			length:      int64(len(content)),
			expected:    nil,
			expectedErr: true,
		},
	}

	for _, test := range tests {
		result, err := ReadFileParts(test.filename, test.fileoffset, test.length)
		if (err != nil) != test.expectedErr {
			t.Errorf("ReadFileParts(%s, %d, %d) unexpected error result: %v",
				test.filename, test.fileoffset, test.length, err)
		}
		if !test.expectedErr && string(result) != string(test.expected) {
			t.Errorf("ReadFileParts(%s, %d, %d) expected %s, got %s",
				test.filename, test.fileoffset, test.length, string(test.expected), string(result))
		}
	}
}

func TestWriteFile(t *testing.T) {

	tests := []struct {
		filename string
		data     []byte
		wantErr  bool
	}{

		{"test.txt", []byte("Hello, World!"), false},
		{"", []byte("Hello, World!"), true},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			err := WriteFile(tt.filename, tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteFile() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err == nil {
				file, err := os.Open(tt.filename)
				if err != nil {
					t.Fatal(err)
				}
				defer file.Close()

				buf := make([]byte, len(tt.data))
				n, err := file.Read(buf)
				if err != nil {
					t.Fatal(err)
				}

				if n != len(tt.data) {
					t.Errorf("only read %d bytes out of %d", n, len(tt.data))
				}

				if string(buf) != string(tt.data) {
					t.Errorf("file content = %s, want %s", string(buf), string(tt.data))
				}
			}
		})
	}
}

func TestWriteFile2(t *testing.T) {
	tests := []struct {
		filename string
		data     []byte
		wantErr  bool
	}{
		{"test.txt", []byte("hello world"), false},
		{"dir/test.txt", []byte("hello world"), false},
		{"dir/subdir/test.txt", []byte("hello world"), false},
		{"", []byte("hello world"), true},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {

			err := WriteFile2(tt.filename, tt.data)

			if (err != nil) != tt.wantErr {
				t.Errorf("WriteFile2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			// 检查文件是否存在
			if _, err := os.Stat(tt.filename); os.IsNotExist(err) {
				t.Errorf("WriteFile2() failed to create file %v", tt.filename)
			}

			// 检查文件内容
			if err == nil {
				content, _ := os.ReadFile(tt.filename)
				if string(content) != string(tt.data) {
					t.Errorf("WriteFile2() failed to write data to file %v", tt.filename)
				}
			}

			// 清理文件和目录
			if err == nil {
				os.Remove(tt.filename)
				dirName := filepath.Dir(tt.filename)
				if dirName != "." {
					os.RemoveAll(dirName)
				}
			}
		})
	}
}

func TestWriteFile3(t *testing.T) {
	// Define a test case struct
	type testCase struct {
		filename string
		data     []byte
		wantErr  bool
	}

	// Create a list of test cases
	tests := []testCase{
		{"testfile.txt", []byte("hello world"), false},
		{"", []byte("hello world"), true},
		{"testfile.txt", nil, false},
	}

	// Iterate through the test cases
	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			// Call the function under test
			err := WriteFile3(tt.filename, tt.data)

			// Check for expected errors
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteFile3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Check if the file was created successfully
			if err == nil {
				if _, err := os.Stat(tt.filename); os.IsNotExist(err) {
					t.Errorf("WriteFile3() failed to create file %v", tt.filename)
				}
			}

			// Cleanup the file after the test
			if err == nil {
				_ = os.Remove(tt.filename)
			}
		})
	}
}

func TestSyncDir(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Test syncing a valid directory
	err = syncDir(tempDir)
	if err != nil {
		t.Errorf("syncDir returned an error: %v", err)
	}

	// Test syncing a non-existent directory
	nonExistentDir := tempDir + "/non-existent"
	err = syncDir(nonExistentDir)
	if err == nil {
		t.Errorf("syncDir did not return an error for non-existent directory")
	}
}

func TestWriteJson(t *testing.T) {
	// 测试数据准备
	type TestData struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	testData := TestData{
		Name: "Alice",
		Age:  25,
	}

	// 执行函数
	err := WriteJson("test.json", testData)
	if err != nil {
		t.Errorf("WriteJson() failed, error: %v", err)
	}
	defer os.Remove("test.json")
}

func TestWriteFileParts(t *testing.T) {
	// 定义测试用例
	tests := []struct {
		filename   string
		fileoffset int64
		data       []byte
		wantErr    bool
	}{
		// 添加测试用例
		{"testfile.txt", 0, []byte("hello"), false},
		{"testfile.txt", 5, []byte("world"), false},
		{"nonexistent_dir/testfile.txt", 0, []byte("hello"), false},
		{"testfile.txt", -1, []byte("invalid offset"), true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("filename=%s,offset=%d,data=%s", tt.filename, tt.fileoffset, tt.data), func(t *testing.T) {
			// 清理工作：删除已存在的文件或目录
			defer os.RemoveAll(tt.filename)

			err := WriteFileParts(tt.filename, tt.fileoffset, tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteFileParts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			// 验证文件内容
			content, err := ReadFileParts(tt.filename, tt.fileoffset, int64(len(tt.data)))
			if err != nil {
				t.Errorf("Failed to read file: %v", err)
				return
			}

			if !tt.wantErr && string(content) != string(tt.data) {
				t.Errorf("WriteFileParts() wrote incorrect content: got %v, want %v", string(content), string(tt.data))
			}
		})
	}
}

func TestWriteScript(t *testing.T) {
	// 定义测试用例
	tests := []struct {
		name     string
		filename string
		data     []byte
		wantErr  bool
	}{
		{
			name:     "TestValidFile",
			filename: "./testdata/testscript.sh",
			data:     []byte("#!/bin/bash\necho 'Hello, World!'\n"),
			wantErr:  false,
		},
	}

	// 遍历测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 调用函数
			err := WriteScript(tt.filename, tt.data)

			// 检查错误
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteScript() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// 检查文件是否存在
			if _, err := os.Stat(tt.filename); os.IsNotExist(err) {
				t.Errorf("WriteScript() failed to create file %v", tt.filename)
			}

			// 清理文件和目录
			defer os.Remove(tt.filename)
			defer os.RemoveAll(filepath.Dir(tt.filename))
		})
	}
}
