#ifndef DATABASE_H
#define DATABASE_H

#define KEY_SIZE 64
#define VALUE_SIZE 256

typedef struct {
    char key[KEY_SIZE];
    char value[VALUE_SIZE];
    int is_deleted;
} Record;

#endif