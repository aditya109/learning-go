# Dependency Injection

https://elliotchance.medium.com/a-new-simpler-way-to-do-dependency-injection-in-go-9e191bef50d5

https://levelup.gitconnected.com/dependency-injection-in-go-using-receiver-functions-d76b7e541ecd

## Overview and Terminology

DI literally means to inject your dependencies. A dependency can be anything that effects the behaviour or outcome of your logic. Some examples are:

1. Other services
2. Configuration
3. System or environment state
4. Stubs of externals APIs.

A *service* is an instance of a class. It is called a service because its often referred to by name rather than type.

A *container* is a collection of services. Services are lazy-loaded and only initialized when they are requested from the container.

A *singleton* is an instance that is initialised once, but can be reused many times.

## 