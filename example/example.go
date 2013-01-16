package main

import (
	"fmt"
	"os/exec"
	"github.com/mattn/go-mruby"
)

func main() {
	mrb := mruby.New()
	defer mrb.Close()

	fmt.Printf("%#v\n", mrb.Eval(`
begin
  p Module.constants
rescue => ex
  p ex
end
  `))

  b, _ := exec.Command("ruby", "-v").Output()
  b = b[:len(b)-1]
  mrb.DefineConst(mrb.ClassGet("Kernel"), "RUBY_DESCRIPTION", mrb.NewString(string(b)))
  mrb.Eval(`p RUBY_DESCRIPTION`)
}
