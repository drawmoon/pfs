package obs

// 获取文件
func Get(name string) ([]byte, error) {
	// TODO: Method not implemented.
	return nil, nil
}

// 推送文件
func Put(filename string, data []byte, contentType string) (string, error) {
	// TODO: Method not implemented.
	return "", nil
}

// 拷贝文件
func Copy(filename, target string) (string, error) {
	data, err := Get(target)
	if err != nil {
		return "", nil
	}

	return Put(filename, data, "")
}

// 移除文件
func Remove(name string) (bool, error) {
	// TODO: Method not implemented.
	return true, nil
}
