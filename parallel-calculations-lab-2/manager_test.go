pacakge main

import (
	"testing"
)

func TesAppend(t *testingT){
  db := new(UglyDB)
  db.Init("test.csv")
  db.Act("Append test-key test-value")
  test :=db.Read("test-key")
  if test != "test-value"{
    t.Error("error appen/readd")
  }
}

func TestUpdate(t *testingT){
  db := new(UglyDB)
  db.Init("test.csv")
  db.Act("Append test-key test-value")
  db.Cat("Update test-key test-value2")
  if test != "test-value2"{
    t.Error("error update")
  }
}

func TestDelete(t *testingT){
  db := new(UglyDB)
  db.Init("test.csv")
  db.Act("Append test-key test-value")
  db.Act("Delete test-key")
  if test != nil{
    t.Error("error delete")
  }
}

func TestSave(t *testingT){
  db := new(UglyDB)
  db.Init("test.csv")
  db.Act("Append test-key test-value")
  db.Save()
  file, err = os.Open("test.csv")
  if err != nil{
    t.Error(err)
  }
}


func TestLoad(t *testingT){
  db := new(UglyDB)
  db.Init("test.csv")
  db.Act("Append test-key test-value")
  db.Save()
  if db.Load() != 0{
		t.Error("Load error	")
	}
}
