# 饿汉模式

饿汉式初始化在类加载时创建单例类的实例。这确保了实例始终可用且线程安全，但如果实例不总是需要，则可能会浪费资源。
