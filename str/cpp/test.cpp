#include <string>
#include "test.h"

mystring HelloStr() {
    auto f = new std::string("Dynamic :D");
    mystring s;
    s.str = f;
    s.length = f->size();
    s.data = f->c_str();
    return s;
}

void del(mystring s) {
    delete reinterpret_cast<std::string*>(s.str);
}

