#include "hello.h"

char *greet(const char *name, size_t size)
{
  size_t len = size + 9;
  char *msg = malloc(len);
  if (!msg) {
    return NULL;
  }
  snprintf(msg, len, "Hello, %s!", name);
  return msg;
}

// NOTE: Caller must take care of msgpack_sbuffer_init
// and msgpack_sbuffer_destroy
void greet_msgpack(const char *name, size_t size, msgpack_sbuffer *buf)
{
  size_t str_len = size + 21;
  char *msg = malloc(str_len + 1);
  if (!msg) {
    return;
  }
  snprintf(msg, str_len + 1, "Hello, %s! from msgpack", name);

  msgpack_packer pk;
  msgpack_packer_init(&pk, buf, msgpack_sbuffer_write);

  msgpack_pack_str(&pk, str_len);
  msgpack_pack_str_body(&pk, msg, str_len);

  free(msg);
}
