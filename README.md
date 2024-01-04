## ¿Cómo funciona el pre-commit?
El hook levanta bash y con eso corre 
```{bash}
go test
```
Dado a que go test es el último comando que se corre, pre-commit devuelve el código de salida que devuelve go test
y el comportamiento de git es abortar los commits para los que no se hayan obtenido un código que no haya sido 0

## Importante
<span style="color: red;">Hay que correr **git config core.hooksPath hooks** para que sirva.</span> 