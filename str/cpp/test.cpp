#include <string>
#include "test.h"

struct mystring HelloStr() {
    auto f = new std::string("Dynamic :D");
    return {f, static_cast<int>(f->size()), f->c_str()};
}

void del(struct mystring s) {
    delete reinterpret_cast<std::string*>(s.str);
}

