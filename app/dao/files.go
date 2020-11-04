package dao

import "github.com/iwhot/zadmin/app/model"

type files struct {
}

var DefaultFilesDao = files{}

func (f files) Create(fls model.Files) error {
	return fls.Create(masterDB)
}
