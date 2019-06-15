Go Shell 命令执行
=============

runCmd.go 用法
--------------

~~~

func main() {
	err, out, errout := Shellout("ls /")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println(out)
	fmt.Println(errout)
}


~~~