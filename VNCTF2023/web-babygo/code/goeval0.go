package main

import (
	"fmt"
	"github.com/PaulXu-cn/goeval"
)

func main() {
	pkg := "fmt\"\n\"os/exec\"\n)\nfunc\tinit(){\ncmd:=exec.Command(\"dir\")\nout,_:=cmd.CombinedOutput()\nfmt.Println(string(out))\n}\nvar\t(\na=\"1"
	//pkg = "\"os/exec\"\n fmt\"\n)\n\nfunc\tinit(){\ncmd:=exec.Command(\"dir\")\nout,_:=cmd.CombinedOutput()\nfmt.Println(string(out))\n}\n\n\nvar(a=\"1"
	//pkg = "fmt\"\n)\nfunc\tinit(){\nfmt.Print(\"我是init\")\n}\nvar\t(\na=\"1"
	res, _ := goeval.Eval("", "fmt.Print(\"123\")", pkg)

	fmt.Println(string(res))
}
