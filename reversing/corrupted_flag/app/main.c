#include <stdio.h>
#include <stdlib.h>
#include <fcntl.h>
#include <unistd.h>

#define NthBit(x, n) (((x) >> (n)) & 1)

int r;

int __corrupt_internal(char *bits) {
  char x;
  if (read(r, &x, sizeof(char)) != sizeof(char)) exit(1);
  if ((x & 0b11111) < 7) {
    bits[x & 0b11111] ^= 1;
  }
}

void corrupt(char *in, size_t ilen, char **out, size_t *olen) {
  char *bits = (char*)malloc(ilen * 14);
  *olen = ilen * 7 / 4 + 1;
  *out = (char*)malloc(*olen);

  size_t cnt = 0;
  for (size_t i = 0; i < ilen; i++) {
    for (int j = 0; j < 2; j++) {
      bits[cnt++] = NthBit(in[i], 4*j+0);
      bits[cnt++] = NthBit(in[i], 4*j+1);
      bits[cnt++] = NthBit(in[i], 4*j+2);
      bits[cnt++] = NthBit(in[i], 4*j+0) ^ NthBit(in[i], 4*j+1) ^ NthBit(in[i], 4*j+2);
      bits[cnt++] = NthBit(in[i], 4*j+3);
      bits[cnt++] = NthBit(in[i], 4*j+0) ^ NthBit(in[i], 4*j+1) ^ NthBit(in[i], 4*j+3);
      bits[cnt++] = NthBit(in[i], 4*j+0) ^ NthBit(in[i], 4*j+2) ^ NthBit(in[i], 4*j+3);
      __corrupt_internal(&bits[cnt-7]);
    }
  }

  for (size_t i = 0; i < cnt; i++) {
    (*out)[i / 8] |= bits[i] << (i % 8);
  }
}

int main() {
  char buf[0x100], *cor;
  size_t len, sz;
  FILE *fp;

  fp = fopen("./flag.txt", "rb");
  if (fp == NULL) return 1;
  len = fread(buf, sizeof(char), 0x100, fp);
  fclose(fp);

  corrupt(buf, len, &cor, &sz);
  fp = fopen("./flag.txt.enc", "wb");
  if (fp == NULL) return 1;
  fwrite(cor, sizeof(char), sz, fp);
  fclose(fp);
  free(cor);

  return 0;
}

__attribute__((constructor))
void setup() {
  r = open("/dev/urandom", O_RDONLY);
  if (r == -1) exit(1);
}

__attribute__((destructor))
void cleanup() {
  close(r);
}
