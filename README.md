# Go Design Pattern

implementing Design Patterns in golang.

source: [design patterns](https://refactoring.guru/design-patterns)

|  **Category**   |             **Design Patterns（Github）**              |
| :------------: | :----------------------------------------------------: |
| **Creational** |        [Factory Method (工厂模式)](01_factory)         |
|                | [Abstract Factory (抽象工厂模式)](02_abstract_factory) |
|                |           [Builder (建造者模式)](03_builder)           |
|                |          [Prototype (原型模式)](04_prototype)          |
|                |          [Singleton (单例模式)](05_singleton)          |
| **Structural** |           [Adapter (适配器模式)](06_adapter)           |
|                |             [Bridge (桥接模式)](07_bridge)             |
|                |          [Composite (组合模式)](08_composite)          |
|                |               [Decorator (装饰器模式)]()               |
|                |                 [Facade (门面模式)]()                  |
|                |                [Flyweight (享元模式)]()                |
|                |                  [Proxy (代理模式)]()                  |
| **Behavioral** |        [Chain of Responsibility (职责链模式)]()        |
|                |                 [Command (命令模式)]()                 |
|                |               [Iterator (迭代器模式)]()                |
|                |                [Mediator (中介模式)]()                 |
|                |                [Memento (备忘录模式)]()                |
|                |               [Observer (观察者模式)]()                |
|                |                  [State (状态模式)]()                  |
|                |                [Strategy (策略模式)]()                 |
|                |             [Template Method (模板模式)]()             |
|                |                [Visitor (访问者模式)]()                |

## Six principles of design pattern

### 1.Single Responsibility Principle (单一职责原则)

单一责任原则也被称为单一功能原则，即导致类变更的原因不应该超过一个。用外行的话说，一个类只负责一个功能。

### 2.Liskov Substitution Principle LSP (里氏代换原则)

里氏代换原则是面向对象设计的基本原则之一。 里氏代换原则中说，任何基类可以出现的地方，子类一定可以出现。LSP 是继承复用的基石，只有当派生类可以替换掉基类，且软件单位的功能不受到影响时，基类才能真正被复用，而派生类也能够在基类的基础上增加新的行为。里氏代换原则是对开闭原则的补充。实现开闭原则的关键步骤就是抽象化，而基类与子类的继承关系就是抽象化的具体实现，所以里氏代换原则是对实现抽象化的具体步骤的规范。

### 3.Dependency Inversion Principle (依赖倒转原则)

这个原则是开闭原则的基础，具体内容：针对接口编程，依赖于抽象而不依赖于具体。

### 4.Interface Isolation Principle (接口隔离原则)

这个原则的意思是：使用多个隔离的接口，比使用单个接口要好。它还有另外一个意思是：降低类之间的耦合度。由此可见，其实设计模式就是从大型软件架构出发、便于升级和维护的软件设计思想，它强调降低依赖，降低耦合。

### 5.Law of Demeter (迪米特法则，又称最少知道原则)

最少知道原则是指：一个实体应当尽量少地与其他实体之间发生相互作用，使得系统功能模块相对独立。

### 6.Open-Closed Principle (开闭原则)

开闭原则的意思是：对扩展开放，对修改关闭。在程序需要进行拓展的时候，不能去修改原有的代码，实现一个热插拔的效果。简言之，是为了使程序的扩展性好，易于维护和升级。想要达到这样的效果，我们需要使用接口和抽象类，后面的具体设计中我们会提到这点。
