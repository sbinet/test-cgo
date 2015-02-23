#include "lib.h"

Object object;
Object* Global;

void ObjectInit(void);
int Get(Object*);

void ObjectInit() {
  Global = &object;
  Global->i = 42;
}

int Get(Object* o) {
  if (o==0) {
	return Global->i;
  }
  return o->i;
}
