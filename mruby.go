package mruby

/*
#include <mruby.h>
#include <mruby/class.h>
#include <mruby/data.h>
#include <mruby/compile.h>
#include <mruby/string.h>

#cgo CFLAGS: -Imruby/include
#cgo linux LDFLAGS: -L. libmruby.so -lm
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type MRuby struct {
	state *C.mrb_state
}

func New() *MRuby {
	return &MRuby{C.mrb_open()}
}

type Symbol C.mrb_sym

func (m *MRuby) Intern(s string) Symbol {
	p := C.mrb_intern(m.state, C.CString(s))
	return Symbol(p)
}

func (m *MRuby) P(obj C.mrb_value) {
	obj = C.mrb_funcall_argv(m.state, obj, C.mrb_sym(m.Intern("inspect")), 0, nil)
	fmt.Println(C.GoString(C.mrb_string_value_ptr(m.state, obj)))
}

func (m *MRuby) DefineConst(class *C.struct_RClass, name string, value C.mrb_value){
	C.mrb_define_const(m.state, class, C.CString(name), value)
}

func (m *MRuby) ClassGet(name string)*C.struct_RClass{
	return C.mrb_class_get(m.state, C.CString(name))
}

func (m *MRuby) NewString(str string) C.mrb_value {
	return C.mrb_str_new(m.state, C.CString(str), C.int(len(str)))
}

func (m *MRuby) Eval(code string, args ...interface{}) C.mrb_value {
	c := C.CString(code)
	defer C.free(unsafe.Pointer(c))

	value := C.mrb_load_string(m.state, c)
	mtype := C.struct_mrb_data_type{}
	C.mrb_get_datatype(m.state, value, &mtype)

	if m.state.exc != nil {
		exception := C.mrb_obj_value(unsafe.Pointer(m.state.exc))
		fmt.Printf("exception: %#v\n", exception)
	}

	return value
}

func (m *MRuby) Close() {
	if m.state != nil {
		C.mrb_close(m.state)
	}
}
