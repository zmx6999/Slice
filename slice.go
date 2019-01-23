package main

/*
#include <stdlib.h>
 */
import "C"
import (
	"unsafe"
	"fmt"
	)

type Slice struct {
	Data unsafe.Pointer
	Len int
	Cap int
}

var (
	ptr uintptr
	intLen int
)

func init()  {
	x:=10
	ptr=unsafe.Sizeof(x)
	intLen=int(ptr)
}

func main()  {
	var s Slice
	s.Create(5,10,1,2,3,4,5)
	s.Print()
	s.Append(6,7,8)
	s.Print()
	fmt.Println(s.Len)
	fmt.Println(s.Cap)
	s.Append(9,10,11)
	s.Print()
	fmt.Println(s.Len)
	fmt.Println(s.Cap)
	fmt.Println(s.GetData(3))
	fmt.Println(s.Search(10))
	s.Insert(3,12)
	s.Print()
	fmt.Println(s.Len)
	fmt.Println(s.Cap)
	s.Delete(3)
	s.Print()
	fmt.Println(s.Len)
	fmt.Println(s.Cap)
	s.Destroy()
	fmt.Print(s)
	/*var x []int
	fmt.Println(unsafe.Sizeof(x))*/
	/*var i interface{}
	i=10
	v,ok:=i.(int)
	if ok {
		fmt.Printf("%T\n",v)
		fmt.Println(v)
	}
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.ValueOf(i))
	i1:=10
	i2:=20
	fmt.Print(reflect.TypeOf(i1)==reflect.TypeOf(i2))*/
}

func (s *Slice) Create(l int,c int,data ...int)  {
	if len(data)==0 {
		return
	}
	if len(data)>l || l>c {
		return
	}
	s.Data=C.malloc(C.ulong(intLen)*C.ulong(c))
	s.Len=l
	s.Cap=c
	p:=uintptr(s.Data)
	for _,v:=range data{
		*(*int)(unsafe.Pointer(p))=v
		p+=ptr
	}
}

func (s *Slice) Print()  {
	if s==nil || s.Data==nil {
		return
	}
	p:=uintptr(s.Data)
	for i:=0; i<s.Len; i++ {
		fmt.Print(*(*int)(unsafe.Pointer(p))," ")
		p+=ptr
	}
	fmt.Println()
}

func (s *Slice) GetData(i int) int {
	if s==nil || s.Data==nil {
		return 0
	}
	if i<0 || i>=s.Len {
		return 0
	}
	p:=uintptr(s.Data)
	for j:=0; j<i; j++ {
		p+=ptr
	}
	return *(*int)(unsafe.Pointer(p))
}

func (s *Slice) Search(x int) int {
	if s==nil || s.Data==nil {
		return -1
	}
	p:=uintptr(s.Data)
	for i:=0; i<s.Len; i++ {
		if *(*int)(unsafe.Pointer(p))==x {
			return i
		}
		p+=ptr
	}
	return -1
}

func (s *Slice) extendCapacity(data ...int)  {
	if s==nil || s.Data==nil {
		return
	}
	if s.Len+len(data)>s.Cap {
		s.Data=C.realloc(s.Data,C.ulong(intLen)*C.ulong(s.Cap)*2)
		s.Cap*=2
	}
}

func (s *Slice) Append(data ...int)  {
	s.extendCapacity(data...)
	p:=uintptr(s.Data)
	for i:=0; i<s.Len; i++ {
		p+=ptr
	}
	for _,v:=range data{
		*(*int)(unsafe.Pointer(p))=v
		p+=ptr
	}
	s.Len+=len(data)
}

func (s *Slice) Insert(i int,data int)  {
	if s==nil || s.Data==nil {
		return
	}
	if i<0 || i>=s.Len {
		return
	}
	s.extendCapacity(data)
	p:=uintptr(s.Data)
	tp:=p
	j:=0
	for ; j<s.Len; j++ {
		if j<i {
			p+=ptr
		}
		tp+=ptr
	}
	for ; j>i; j-- {
		*(*int)(unsafe.Pointer(tp))=*(*int)(unsafe.Pointer(tp-ptr))
		tp-=ptr
	}
	*(*int)(unsafe.Pointer(p))=data
	s.Len++
}

func (s *Slice) Delete(i int)  {
	if s==nil || s.Data==nil {
		return
	}
	if i<0 || i>=s.Len {
		return
	}
	p:=uintptr(s.Data)
	for j:=0; j<i; j++ {
		p+=ptr
	}
	for j:=i; j<s.Len-1; j++ {
		*(*int)(unsafe.Pointer(p))=*(*int)(unsafe.Pointer(p+ptr))
		p+=ptr
	}
	*(*int)(unsafe.Pointer(p))=0
	s.Len--
}

func (s *Slice) Destroy()  {
	C.free(s.Data)
	s.Data=nil
	s.Len=0
	s.Cap=0
	s=nil
}
