
#include <stdio.h>
#include "database.h"
 

int main(void) {
    Record record;
   
    printf("ViralDB initialized!\n");

    printf("Size of Record: %zu bytes\n", sizeof(Record));
    printf("Size of key: %zu bytes\n", sizeof(record.key));
    printf("Size of value: %zu bytes\n", sizeof(record.value));
    printf("Size of is_deleted: %zu bytes\n", sizeof(record.is_deleted));
}