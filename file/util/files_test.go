package util

import "testing"

func TestGetAllFiles(t *testing.T) {
f:=new(File)

f2:=f.GetAllFiles("D:\\FTP\\online\\ABC")
for _, v := range f2.FilesList {
	t.Log(v)
}
t.Log("-------------------------")
f2=f2.GetFilterSuffix(".zip")
for _, v := range f2.FilesList {
	t.Log(v)
}
t.Log("-------------------------")
f2=f2.GetFilterZH()
for _, v := range f2.FilesList {
	t.Log(v)
}

}
