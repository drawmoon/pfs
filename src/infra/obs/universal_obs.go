package obs

// 通用的对象存储服务
type UniversalObs struct {
}

func NewUniversalObs() *UniversalObs {
	return &UniversalObs{}
}

// 获取文件
func (obs *UniversalObs) get(name string) ([]byte, error) {
	// TODO: Method not implemented.
	return nil, nil
}

// 推送文件
func (obs *UniversalObs) put(filename string, data []byte, contentType string) (string, error) {
	// TODO: Method not implemented.
	return "", nil
}

// 拷贝文件
func (obs *UniversalObs) copy(filename, target string) (string, error) {
	data, err := obs.get(target)
	if err != nil {
		return "", nil
	}

	return obs.put(filename, data, "")
}

// 移除文件
func (obs *UniversalObs) remove(name string) (bool, error) {
	// TODO: Method not implemented.
	return true, nil
}
