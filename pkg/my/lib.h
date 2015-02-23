//#define API_DATA(RTYPE) extern RTYPE
#define API_DATA(RTYPE) RTYPE

typedef struct { int i; } Object;

API_DATA(Object*) Global;

void ObjectInit(void);
int Get(Object*);
