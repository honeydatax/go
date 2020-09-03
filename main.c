#include <stdio.h>
#define WASM_EXPORT __attribute__((visibility("default")))

WASM_EXPORT
int main() {
  int n=21+45;
  return n;
}
