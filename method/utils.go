package method

import (
	"archive/zip"
	"io"
	"os"
	"os/exec"

	"github.com/google/uuid"
)

func GitCloneReturnZipPath(git_url string) (string, error) {
	uuid := uuid.New()

	// 执行 git clone
	cmd := exec.Command("git", "clone", git_url, "tmp/"+uuid.String())

	// 执行并返回结果
	_, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	// fmt.Println(string(out))

	path := "tmp/" + uuid.String() + ".zip"

	// 压缩 tmp/uuid
	err = CompressPathToZip("tmp/"+uuid.String(), path)
	if err != nil {
		return "", err
	}

	// 删除 tmp/uuid
	err = os.RemoveAll("tmp/" + uuid.String())
	if err != nil {
		return "", err
	}

	return path, nil
}

// 压缩文件夹
func CompressPathToZip(path, targetFile string) error {
	d, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()

	f, err := os.Open(path)
	if err != nil {
		return err
	}

	err = compress(f, "", w)

	return err
}

// 压缩文件
func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
