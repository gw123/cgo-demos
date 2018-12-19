package main

/*
#include <stdio.h>
#include "cgo_wrap.h"

// The gateway function
int callOnMeGo_cgo(int in)
{
	printf("C.callOnMeGo_cgo(): called with arg = %d\n", in);
	int callOnMeGo(int);
	return callOnMeGo(in);
};

int testCallback(callback_fcn f)
{
	printf("C.testCallback(): called with arg = 1\n");
	f(1);
}

*/
import "C"