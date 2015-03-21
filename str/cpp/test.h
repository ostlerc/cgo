#ifndef TEST_H
#define TEST_H

#include <string.h>

#ifdef __cplusplus
extern "C" {
#endif

struct mystring {
    void* str;
    int length;
    const char* data;
};

void del(struct mystring s);
struct mystring HelloStr();
#ifdef __cplusplus
}
#endif

#endif //TEST_H
