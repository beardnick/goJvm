# constantPool

## constantInfo接口：抽象出所有14种常量的基本内存布局

```java
cp_info {
    u1 tag;
    u1 info[];
}
```

即带一个tag外加若干byte的其它信息的数据

## 14 种常量

```java
CONSTANT_Class                  
CONSTANT_Fieldref               
CONSTANT_Methodref              
CONSTANT_InterfaceMethodref     
CONSTANT_String                 
CONSTANT_Integer                
CONSTANT_Float                  
CONSTANT_Long                   
CONSTANT_Double                 
CONSTANT_NameAndType            
CONSTANT_Utf8                   
CONSTANT_MethodHandle           
CONSTANT_MethodType             
CONSTANT_InvokeDynamic          
```
