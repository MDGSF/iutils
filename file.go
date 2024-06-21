package iutils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func DirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func FileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func GetFileSize(filename string) (int64, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	fileSize := fileInfo.Size()
	return fileSize, nil
}

func ReadFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return io.ReadAll(file)
}

func ReadFileParts(filename string, fileoffset int64, length int64) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = file.Seek(fileoffset, io.SeekStart)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(io.LimitReader(file, length))
}

func WriteFile(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close() // 关闭文件，defer 语句确保在函数退出前执行

	// 使用 Write 方法将 []byte 写入文件
	n, err := file.Write(data)
	if err != nil {
		return err
	}

	if n != len(data) {
		return fmt.Errorf("only wrote %d bytes out of %d", n, len(data))
	}

	return nil
}

// writeFile 自动创建目录并写入文件
func WriteFile2(filename string, data []byte) error {
	// 分离文件名和路径
	dirName := filepath.Dir(filename)

	// 检查目录是否存在，不存在则创建
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		// 使用 MkdirAll 创建缺失的多级目录
		if err := os.MkdirAll(dirName, 0755); err != nil {
			return err
		}
	}

	// 写入文件
	if err := os.WriteFile(filename, data, 0644); err != nil {
		return err
	}

	return nil
}

// WriteFile writes data to the file specified by filename.
// It uses a temporary file for safety and ensures data is flushed to disk.
func WriteFile3(filename string, data []byte) error {
	// Create a temporary file
	tempFile, err := os.CreateTemp(os.TempDir(), ".tmp-"+filepath.Base(filename))
	if err != nil {
		return err
	}
	defer func(name string) {
		_ = os.Remove(name) // Attempt to remove the temp file if it still exists.
	}(tempFile.Name())

	// Write data to the temporary file.
	if _, err := tempFile.Write(data); err != nil {
		return err
	}

	// Close the temporary file before moving it to avoid file locking issues.
	if err := tempFile.Close(); err != nil {
		return err
	}

	// Move the temporary file to the destination, overwriting it if it exists.
	if err := os.Rename(tempFile.Name(), filename); err != nil {
		return err
	}

	// Sync the directory to ensure the rename operation is flushed to disk.
	if err := syncDir(filepath.Dir(filename)); err != nil {
		return err
	}

	return nil
}

// syncDir forces a synchronization of the file system metadata and any delayed writes to disk for the given directory.
func syncDir(dirName string) error {
	dir, err := os.Open(dirName)
	if err != nil {
		return err
	}
	defer func(dir *os.File) {
		_ = dir.Close()
	}(dir)

	return dir.Sync()
}

func WriteJson(filename string, data interface{}) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return WriteFile3(filename, bytes)
}

func WriteFileParts(
	filename string,
	fileoffset int64,
	data []byte,
) error {
	// 分离文件名和路径
	dirName := filepath.Dir(filename)

	// 检查目录是否存在，不存在则创建
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		// 使用 MkdirAll 创建缺失的多级目录
		if err := os.MkdirAll(dirName, 0755); err != nil {
			return err
		}
	}

	// 写入文件
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Seek(fileoffset, io.SeekStart)
	if err != nil {
		return err
	}

	n, err := file.Write(data)
	if err != nil {
		return err
	}

	if n != len(data) {
		return fmt.Errorf("only wrote %d bytes out of %d", n, len(data))
	}

	return nil
}

func WriteScript(filename string, data []byte) error {
	// 分离文件名和路径
	dirName := filepath.Dir(filename)

	// 检查目录是否存在，不存在则创建
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		// 使用 MkdirAll 创建缺失的多级目录
		if err := os.MkdirAll(dirName, 0755); err != nil {
			return err
		}
	}

	// 写入文件
	if err := os.WriteFile(filename, data, 0755); err != nil {
		return err
	}

	if err := os.Chmod(filename, 0755); err != nil {
		return err
	}

	return nil
}

// CalcFileMD5 计算并返回指定文件的MD5哈希值
func CalcFileMD5(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := md5.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	return hashString, nil
}
