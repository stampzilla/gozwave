Gozwave generators
=============

To convert Sigma Designs header file (ZW_classcmd.h) to go code we use a tool called castxml. It can parse c code and convert it into an XML document. That xml document is then used by the commandclasses.go generator.

```
castxml ZW_classcmd.h -std=c89 --castxml-gccxml -o ZW_classcmd.xml
```
