#ifndef CLIBRARY_H
#define CLIBRARY_H
typedef int (*t_handel)(char *event);
int call_handel(t_handel fn,char *event);
#endif