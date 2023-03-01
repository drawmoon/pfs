package validator

import "github.com/PowerReport/pfs/src/domain/directory/model"

// 校验目录添加
func ValidateForAdding(dir *model.Directory) {
	// TODO: Method not implemented.
}

// 校验目录修改
func ValidateForUpdating(dir *model.Directory) {
	// TODO: Method not implemented.
}

// 校验目录移动
func ValidateForMove(dir *model.Directory) {
	// TODO: Method not implemented.
}

// 校验目录名称是否合法
func ValidateName(dir *model.Directory) {
	// TODO: Method not implemented.
}

// 校验目录对于当前用户是否有写（更新）的权限
func ValidatePermissionForWriting(dir *model.Directory) {
	// TODO: Method not implemented.
}

// 校验目录对于当前用户是否有删除的权限
func ValidatePermissionForDeleting(dir *model.Directory) {
	// TODO: Method not implemented.
}
