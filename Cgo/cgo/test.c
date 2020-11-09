#include <stdio.h>
#include "test.h"

void hello() {
    printf("Hello,World~\n");
}

#ifdef __TEST__

int main(){
    hello();
    return 0;
}

#endif