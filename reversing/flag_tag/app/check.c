#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <emscripten.h>

typedef unsigned char u8;

EM_JS(void, setupTag, (), {
    let convertToCArray = (s) => {
      let ptr = Module.allocate(Module.intArrayFromString(s),
                                Module.ALLOC_STACK);
      return ptr;
    };
    let flagElements = document.getElementsByTagName("flag");
    for (let element of flagElements) {
      let enc = element.innerHTML;
      let key = element.attributes.key;
      let placeholder = element.attributes.placeholder;
      if (enc && key) {
        element.innerHTML = "";
        let button = document.createElement('button');
        let input = document.createElement('input');
        input.setAttribute("placeholder", placeholder
                           ? placeholder.nodeValue
                           : "");
        input.setAttribute("enc", enc);
        input.setAttribute("key", key.nodeValue);
        input.setAttribute("onerror", element.attributes.onerror
                           ? element.attributes.onerror.nodeValue
                           : "");
        input.setAttribute("onsuccess", element.attributes.onsuccess
                           ? element.attributes.onsuccess.nodeValue
                           : "");
        button.innerHTML = "Check";
        button.onclick = (event) => {
            let input = event.target.previousSibling;
            let stack = Module.stackSave();
            let flag = convertToCArray(input.value);
            let enc = convertToCArray(input.attributes.enc.nodeValue.trim());
            let key = convertToCArray(input.attributes.key.nodeValue);
            if (Module._check(flag, key, enc)) {
              eval(input.attributes.onerror.nodeValue);
            } else {
              eval(input.attributes.onsuccess.nodeValue);
            }
            Module.stackRestore(stack);
        };
        element.appendChild(input);
        element.appendChild(button);
      }
    }
  });

void EMSCRIPTEN_KEEPALIVE init()
{
  setupTag();
}

#define ROTL(a,b) (((a) << (b)) | ((a) >> (8 - (b))))

void QR(u8 *s, int a, int b, int c, int d) {
  s[b] ^= ROTL((s[a] + s[d]) & 0xff, 1);
  s[c] ^= ROTL((s[b] + s[a]) & 0xff, 2);
  s[d] ^= ROTL((s[c] + s[b]) & 0xff, 3);
  s[a] ^= ROTL((s[d] + s[c]) & 0xff, 4);
}

int EMSCRIPTEN_KEEPALIVE check(const char* cflag,
                               const char* key,
                               const char* enc)
{
  const char table[] = {'0', '1', '2', '3', '4', '5', '6', '7',
                        '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'};
  int flaglen = strlen(cflag);
  int keylen = strlen(key);
  int enclen = strlen(enc);
  if (keylen < 8)
    return -1;
  if ((flaglen * 4 <= enclen - 32) || (flaglen * 4 > enclen))
    return -1;

  int padlen = (flaglen + 15) & ~15;
  char *flag = (char*)malloc(padlen);
  memset(flag, 0, padlen);
  memcpy(flag, cflag, flaglen);

  u8 s[16];
  for (int i = 0; i < 8; i++)
    s[i] = key[i];

  int correct = 0;
  for (int i = 0; i < flaglen; i += 8) {
    for (int j = 0; j < 8; j++)
      s[8+j] = flag[i+j];

    for (int j = 0; j < 128; j++) {
      // odd
      QR(s, 0, 4, 8, 12);
      QR(s, 5, 9, 13, 1);
      QR(s, 10, 14, 2, 6);
      QR(s, 15, 3, 7, 11);
      // even
      QR(s, 0, 1, 2, 3);
      QR(s, 5, 6, 7, 4);
      QR(s, 10, 11, 8, 9);
      QR(s, 15, 12, 13, 14);
    }

    for (int j = 0; j < 16; j++) {
      correct |= (enc[i*4+j*2+0] != table[s[j] / 0x10]);
      correct |= (enc[i*4+j*2+1] != table[s[j] % 0x10]);
    }
  }

  free(flag);
  return correct;
}
