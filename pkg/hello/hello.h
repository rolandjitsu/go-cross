#ifndef HELLO_H_
#define HELLO_H_

#include <stdio.h>
#include <stdlib.h>
#include <msgpack.h>

char *greet(const char *name, size_t size);
void greet_msgpack(const char *name, size_t size, msgpack_sbuffer *buf);

#endif
