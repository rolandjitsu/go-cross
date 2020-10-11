#include "hello.h"

char *greet(const char *name, size_t size)
{
  char *msg = malloc(size + 8);
  sprintf(msg, "Hello, %s!", name);
  return msg;
}
