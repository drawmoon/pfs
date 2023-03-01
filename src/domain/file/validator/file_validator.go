package validator

import "github.com/PowerReport/pfs/src/domain/file/model"

// 校验文件添加
func ValidateForAdding(file *model.File) {
	// TODO: Method not implemented.
}

// 校验文件修改
func ValidateForUpdating(file *model.File) {
	// TODO: Method not implemented.
}

// 校验文件移动
func ValidateForMove(file *model.File) {
	// TODO: Method not implemented.
}

// 校验文件名称是否合法
func ValidateName(file *model.File) {
	// TODO: Method not implemented.
}

// 校验文件对于当前用户是否有写（更新）的权限
func ValidatePermissionForWriting(file *model.File) {
	// TODO: Method not implemented.
}

// 校验文件对于当前用户是否有删除的权限
func ValidatePermissionForDeleting(file *model.File) {
	// TODO: Method not implemented.
}
