package util

import "testing"

func TestGetAllFiles(t *testing.T) {
f:=new(File)

f1:=f.GetAllFiles("D:\\FTP\\online\\ABC").GetFilterSuffix(".zip").GetFilterZH()
for _, v := range f1.FilesList {
	t.Log(v)
}
// t.Log("-------------------------")
// f1=f1.GetFilterSuffix(".zip")
// for _, v := range f.FilesList {
// 	t.Log(v)
// }
// t.Log("-------------------------")
// f1=f1.GetFilterZH()
// for _, v := range f.FilesList {
// 	t.Log(v)
// }

}
