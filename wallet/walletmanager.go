//walletmanager.go

package wallet

import (
	"bytes"
	"encoding/gob"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"qanly_chain/constcoe"
	"qanly_chain/utils"
	"strings"
)

type RefList map[string]string //key值为钱包地址，value为钱包的别名

func (r *RefList) Save() {
	filename := constcoe.WalletsRefList + "ref_list.data"
	var content bytes.Buffer
	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(r)
	utils.Handle(err)
	err = ioutil.WriteFile(filename, content.Bytes(), 0644)
	utils.Handle(err)
}

func (r *RefList) Update() {
	err := filepath.Walk(constcoe.Wallets, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		fileName := f.Name()
		if strings.Compare(fileName[len(fileName)-4:], ".wlt") == 0 {
			_, ok := (*r)[fileName[:len(fileName)-4]]
			if !ok {
				(*r)[fileName[:len(fileName)-4]] = ""
			}
		}
		return nil
	})
	utils.Handle(err)
}

func LoadRefList() *RefList {
	filename := constcoe.WalletsRefList + "ref_list.data"
	var reflist RefList
	if utils.FileExists(filename) {
		fileContent, err := ioutil.ReadFile(filename)
		utils.Handle(err)
		decoder := gob.NewDecoder(bytes.NewBuffer(fileContent))
		err = decoder.Decode(&reflist)
		utils.Handle(err)
	} else {
		reflist = make(RefList)
		reflist.Update()
	}
	return &reflist
}

func (r *RefList) BindRef(address, refname string) { //使用别名方便调试
	(*r)[address] = refname
}

func (r *RefList) FindRef(refname string) (string, error) {
	temp := ""
	for key, val := range *r {
		if val == refname {
			temp = key
			break
		}
	}
	if temp == "" {
		err := errors.New("the refname is not found")
		return temp, err
	}
	return temp, nil
}
