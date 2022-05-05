package wrapper

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <taos.h>
*/
import "C"
import (
	"unsafe"

	"github.com/taosdata/driver-go/v2/wrapper/cgo"
)

//typedef void(tmq_commit_cb(tmq_t *, tmq_resp_err_t, tmq_topic_vgroup_list_t *, void *param));

//export TMQCommitCB
func TMQCommitCB(consumer unsafe.Pointer, resp C.enum_tmq_resp_err_t, offsets unsafe.Pointer, param unsafe.Pointer) {
	c := cgo.Handle(param).Value().(chan *TMQCommitCallbackResult)
	r := GetTMQCommitCallbackResult(int32(resp), consumer, offsets)
	c <- r
}
