package main

/*
#include <stdio.h>
#include "cgo_wrap.h"
// The gateway function
int call_handel(t_handel fn,char *event)
{
	printf("C.call_handel(): event %s\n",event);
	fn(event);
}
*/
import "C"
