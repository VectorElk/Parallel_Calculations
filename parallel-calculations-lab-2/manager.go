package main

import (
  "strings"
  "os"
  "encoding/csv"
)

type UglyDB struct{
    db map[string]string
    Path string
}

func (u *UglyDB) Act(command string) (result string){
  l := strings.Split(command, " ")

  for i := 0; i < len(l); i++ {
    l[i] = strings.Trim(l[i], "\n\r")
  }

  if len(l) == 0 {
    return "wrong command"
  }

  if l[0] == "Read"{
    if len(l) != 2 {
      return "wrong number of arguments"
    }
    value, e := u.Read(l[1])
    if e != 0{
      return "The key does not exist"
    }else{
      return value
    }
  }

  if l[0] == "Update"{
    if len(l) != 3 {
      return "wrong number of arguments"
    }
    e := u.Update(l[1], l[2])
    if e == 0 {
      return "Updated successfully"
    } else {
      return "The key does not exist"
    }
  }

  if l[0] == "Append"{
    if len(l) != 3 {
      return "wrong number of arguments"
    }
    e := u.Append(l[1], l[2])
    if e == 0 {
      return "Created new entry"
    } else {
      return "The key already exists"
    }
  }

  if l[0] == "Delete"{
    if len(l) != 2 {
      return "wrong number of arguments"
    }
    e := u.Delete(l[1])
    if e == 0 {
      return "Deleted successfully"
    } else {
      return "The key does not exist"
    }
  }

  return "wrong command"
}

func (u *UglyDB) Read(key string) (string, int){
  val,exists := u.db[key]
  if exists{
    return val,0
  }else{
    return val,1
  }
}

func (u * UglyDB) Update(key string, value string)(int){
  _,exists := u.db[key]
  if exists{
    u.db[key] = value
    return 0
  }else{
    return 1
  }
}

func (u *UglyDB) Append(key string, value string) (int){
  _,exists := u.db[key]
  if exists{
    return 1
  }else{
    u.db[key] = value
    return 0
  }
}

func (u *UglyDB) Delete(key string) (int){
  _,exists := u.db[key]
  if exists{
    delete(u.db, key)
    return 0
  }else{
    return 1
  }
}

func (u *UglyDB) Load() (int){
  file,_ := os.Open(u.Path)
  records, err := csv.NewReader(file).ReadAll()
  if err != nil {
    panic(err)
  }
  for _, row := range records {
    u.Append(row[0], row[1])
    }
  return 0
}

func (u *UglyDB) Save(){
  file,err := os.Create(u.Path)
  checkCreate("Cannot create file", err)
  writer := csv.NewWriter(file)
  for key, value := range u.db{
        err := writer.Write([]string{key, value})
        checkCreate("Cannot write to file", err)
    }
  writer.Flush()
  file.Close()
}

func (u *UglyDB) Init(path string){
  u.Path = path
  u.db = map[string]string{
    "key":"value",
  }
}

func checkCreate(message string, err error) {
    if err != nil {
        panic(err)
    }
}
