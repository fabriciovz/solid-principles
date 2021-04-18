Liskov Substitution Principle
Coined by Barbara Liskov, the Liskov substitution principle states, roughly, that two types are substitutable if they exhibit behaviour such that the caller is unable to tell the difference.

In a class based language, Liskov’s substitution principle is commonly interpreted as a specification for an abstract base class with various concrete subtypes. But Go does not have classes, or inheritance, so substitution cannot be implemented in terms of an abstract class hierarchy.

https://dave.cheney.net/2016/08/20/solid-go-design