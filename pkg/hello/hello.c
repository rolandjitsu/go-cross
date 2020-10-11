#include "hello.h"

char *greet(const char *name, size_t size)
{
  char *msg = malloc(size + 8);
  sprintf(msg, "Hello, %s!", name);
  return msg;
}

// NOTE: Caller must take care of msgpack_sbuffer_init
// and msgpack_sbuffer_destroy
void greet_msgpack(const char *name, size_t size, msgpack_sbuffer *buf)
{
  int sz = size + 21;
  char *msg = malloc(sz);
  sprintf(msg, "Hello, %s! from msgpack", name);

  msgpack_packer pk;
  msgpack_packer_init(&pk, buf, msgpack_sbuffer_write);

  msgpack_pack_str(&pk, sz);
  msgpack_pack_str_body(&pk, msg, sz);
}
