#include <stdio.h>
#include <stdlib.h>

int main(int argc, const char *argv[]) {
    int num = atoi(argv[1]);
    printf("%d %d\n", num, num << 2);

    return 0;
}
