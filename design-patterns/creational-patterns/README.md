# Creational Patterns

They are design patterns that deal with object creation mechanisms, tring to create objects in a manner suitable to the situation. The basic form of object creation could result in design problems or added complexity to the design. Creational patterns solve this problem by somehow controlling this object creation.

## Rules of thumb

1. Sometimes creational patterns are competitors; these are cases when either `Prototype` or `Abstract Factory` could be used profitably. At other times they are complementary; `Abstract Factory` might store a set of `Prototypes` from which to clone and return product objects, `Builder` can use one of the other patterns to implement which comp