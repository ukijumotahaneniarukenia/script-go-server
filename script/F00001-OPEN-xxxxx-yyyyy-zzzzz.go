package go_script

var Message string = "hello world"

var NNN string = nnn

func Mmm(a int,b int)int{//外部参照させるためには先頭を大文字に
  return a+b
}

func Show_Message() string {//外部参照させるためには先頭を大文字に
  return Message + NNN
}
