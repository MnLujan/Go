#Documentacion Go

###*Variables*

Tenemos los mismos tipos de variables al igual que todos los leguajes, pero ¿como declararlos?

El lenguaje requiere la siguiente sintaxis: 

> **var** var_name **tipo**

+ Por ejemplo:
    > **var** x **int**
    > **var** cadena **string**
    > **var** boleano **bool**
    **var** numeros **[]int**

Cabe aclarar que los datos de tipo string, se inicializan como una cadena vacia, o los valores **bool** en **false**. Ademas, GO nos exigue que cada variable declarada, sea utilizada. Si lo anterior no se cumple, nuestro codigo no compilara.

Si queremos declarar e inicializar la variable con un cierto valor, tenemos una segunda sintaxis:

> var_name **:=** *val*

+ Por ejemplo:
    >nombre **:=** "Martin"

El compilador en tiempo de ejecucion detecta el tipo de variable que es y se lo asigna.

###*Casteo*

La conversion de tipo de datos es siempre necesaria, para esto nos ofrecen la libreria **strconv** la cual debemos importar a nuestro proyecto. Esta libreria contiene varias funciones asique siempre es recomendable leer la documentacion de la misma.
La sintaxis para la conversion de un dato es de las iguiente manera:

> edad **:=** 22
> edad_string **:=** strconv.Itoa(edad)

Si nosotros importamos una libreria y no es utiulizada, el compilador nos obligara a eliminarla.

Las funciones en GO pueden devolver mas de un valor, por lo que deberemos verificar esto y analizar si necesitamos este segundo valor que nos retorna. Si no lo necesitamos, debemos recurrir a una sintaxis, ya que recordemos que no podemos declara variables que no vamos a usar. Esta sintaxis es **"_"**, este operador indica que el valor a retornar no me interesa y puede ser descartado. 

+ Por ejemplo
> edad_int,_ **:=** strconv.Atoi(edad)

###*Lectura de datos*

El input y output en GO se encuentra manejado por varias librerias, veremos algunas de ellas.

+ **fmt**:
    Esta libreria utiliza funciones similares a C.
    >Por ejemplo:
    La funcion "Scanf" puede ser utilizada para leer la entrada por teclado. Esta funcion se debe aclara el tipo que esperar y podemos asignarle la **direccion de memoria** mediante **&** de una variable para almacenar el dato ingresado.
    *Sintaxis*: fmt.Scanf("%d", &variable)

+ **bufio**

    Este paquete utiliza un objeto de tipo "Reader", se necesitan un poco mas de instrucciones pero nos da mas opciones y chequeo de errores.
    >Ejemplo:
    reader **:=** bufio.NewReader(os.Stdin)

    Aqui aparece un nuevo paquete llamado "os" el cual nos sirve para indicarle a reader que tipo de lector debe utilizar. En el caso anterior la entrada utilizada es por teclado.

    >Continuando:
    ¿Como ejecutamos la lectura?
    reader nos retornara el valor leido y un valor adicional que indica la ausencia o presencia de error. Si la lectura es exitosa, retorna **nil**.
    *Sintaxis*:
    text,err **:=** reader.ReadString('\n') (El parametro es el stop de lectura)

###*Ciclos*

GO el unico metodo de iteracion que tiene es el **for** y tiene la misma sintaxis que en el resto de lenguajes.

>Ejemplo:
 - for i:=0; i<10; i++{ }
 - for{} *Este ciclo es infinito*
 - for i>10{i++} *Habiendo declarado antes i*
 - for index; content := range myArray{} *Este ultimo simula el foreach, si el index no nos interesa, podemos usar el operador **_***

###*Arreglos*

Al igual que en C/C++ podemos declarar arreglos de tamaño fijo pero no dinamico. Si no especificamos el valor de inicializacion este se inicializara con los valores predeterminados por Go. (para enteros el 0, string "", etc).
Podemos preguntar el tamaño de los arreglos mediante la funcion **len(array)**.

>Sintaxis:
**var** arreglo [10]**int**
arreglo **:=** [10]**int**{}

Podemos declarar el arreglo e inicializar algunas variables, el resto quedara con los valores predeterminados por GO.

>Ejemplo
arreglo := [3]int{1,2}
*La tercer posicion quedara con el valor 0*

###*Slices*

Estructura basada en arreglos que nos permite modificar su tamaño en tiempo de ejecucion y no es necesario definir su tamaño. La declaracion se realiza igual que un arreglo, pero sin especificar su tamaño, esto lo convierte en un **Slices**. Basicamente un Slices es una estructura en la que uno de sus datos es un puntero a un arreglo, en otro lugar de memoria.
Al declarar el un Slices a diferendia de los arreglos, el valor almacenado es de tipo nulo o **nil**.
Cuando se declara un Slices, declaramos un puntero al primer elemento del arreglo, al igual que C/C++. 

Otra utilidad es el llamado "Slicing" que nos permite partir un arreglo o un slice, creando otro nuevo.

>Por ejemplo:
arreglo **:=** [3]int{1,2,3}
slice **:=** arreglo[**0:2**]

La expresion **[0:2]** indica que corta al arreglo entre la posicion 0 a 2 (sin incluir el valor que hay en la posicion 2) y se la asigna al slice, es decir, ahora slice sera un arreglo con 2 lugares y que contendra los valores que habia en esas posiciones de la variable *arreglo*. Por lo tanto, si se imprimiera los valores dentro de slice tendremos:

>[1,2]

#####Make y Append

Otra manera de construir Slices es mediante la funcion **make(-,-,-)** la cual recibe 2 o 3 parametros, estos son:

+ Tipo de Slice a crear
+ Longuitud del mismo
+ Capacidad del Slice

Basicamente este tercer parametro nos permite extender el arreglo en caso de necesitarlo, es reservar memoria de mas para optimizar. Utilizando la funcion **append(slice, val)** podemos agregar un valor al final del arreglo utilizando este espacio "extra" declarado anteriormente. Aqui tenemos 2 datos para una misma variable, su longuitud que es el tamaño del arreglo y su capacidad, que es la cantidad de datos capaz de almacenar.
En caso de no declarar la capacidad, podemos utilizar igual la funcion **append()** pero es menos eficiente, debido a que se deberia alocar memoria en tiempo de ejecucion. 

>Sintaxis:
slice := make([]int, 3, 5)
slice = append(slice, 6)

Aqui estamos diciendo que en la posicion final del slice, agrege el valor 6. Entonces en la posicion 4 del arreglo se almacenara este valor, por mas que el largo del arreglo sea de 3.

#####Copy

Esta es una funcion que nos permite copiar un arreglo en otro. Recive 2 parametros dicha funcion, el destino de la copia y la fuente de onde copiar los datos. Dicha funcion copia el minimo de elementos que se termiten en el arreglo de destino. Es decir, si queremos copiar desde una fuente de largo 5 a un destino de largo 2, la funcion solo copiara 2 elementos. *Recordar que se puede jugar con la capacidad del arreglo para estos casos*.
+ Dato: La funcion **copy()** solo se puede utilizar con Slices, no con arreglos de tamaño fijo. 

>Sintaxis:
    copy(destino, fuente)

###*Punteros*

La definicion de punteros en GO es identica a la de C/C++, por lo tanto no es necesario aclarar para que sirven y como funcionan. La sintaxis de la declaracion de un puntero es la siguiente:

>Por ejemplo:
var x ***int**
>>*Para agregarle un valor:*
entero := 5
x = &entero
*Si no se agrega un valor, si valor predeterminado sera **nil***


###Estructuras

Las estructuras no sirven para definir un tipo de dato espcifico que dentro de este, se encuentren otros datos de tipo primitivos del lenguaje (int, string, bool, etc) o variables de otro tipo definido anteriormente. 

Go nos permite definir una estructura de 2 formas:

+ Mediante el operador **var**.
+ Mediante el operador **new()**

Este ultimo nos devuelve un puntero a la estructura.

>Sintaxis de la definicion de la estructura
**type** name_struct **struct**{
    }
    *Dentro de la misma se pueden poner los campos que se requieran*
>>Sintaxis de la declaracion:
    **var** name_var **name_struct**
    *Aqui todos los campos se inicializaran con el valor por defecto de GO*
    **var** name_var **name_struct**{var1, var2, ...}
    *Aqui podemos darle valores de inicializacion a la estructura, deben estar en el mismo orden que en la definicion de la **Struct***
    name_var **:=** **name_struc**{var1, var2, ...}
    name_var **:=** **new(name_struct)**
    *name_var es ahora un puntero a la estructura*


######Metodos

Go no es un lenguaje orientado a objetos pero nos permite agregar funciones a las estructuras. Tenemos 2 maneras de declarar las funciones y hay que tener en claro para que las utilizaremos. Estas son:

>Sintaxis 1:
**func(this name_struc)** name_func () tipe_return{}

Esta sintaxis tiene la particularidad de que la estructura pasada como parametro, es una copia exacta de quien la llama, por lo tanto cualquier cambie que se realice en las variables dentro de la funcion, no se vera reflejada, solo sera modficiada en la copia.

>Sintaxis 2:
**func**(this *name_struc) name_func () tipe_return{}

En este caso la funcion recibe un puntero a la estructura, por lo tanto los cambios que se realicen se veran reflejados en la estructura que llama.
Esto se debe tener en cuenta dependiendo de que queremos hacer y si queremos ahorra memoria, ya que enviar una copia exacta de la estructura es costoso.

######Campos anonimos

Como bien dijimos anteriormente, GO no es un lenguaje orientado a objetos pero existen estas opciones porque nos permiten replicar la **herencia** de estos lenguajes. Primero definimos una estructura para ejemplificar:

>type Human struct{
    name string
    }

A continuacion definimos un campo anonimo:

>type Tutor struct{
    Human
    }
    >>*Se declara como*
    tutor **:=** Human{"Martin"}
    *Si no se le pasa los valores la estructura inicializa con los valores por defecto*

Podemos acceder a los valores de **Human** y a las funciones que esten asociadas a esta estructura desde **Tutor**. 

###*Interfaces*

Las interfaces son estructuras de datos que definen metodos vacios. Estos a su vez sin un tipo de dato que podemos pasar entre funciones. Esto nos permite que varias estructuras diferentes implementen la misma interfaz y a su vez, estas estructuras utilizar funciones que reciban a la interfaz como parametro, ahorrando codigo. 
La declaraion de interfaces tiene la siguiente sintaxis:

>Sintaxis:
**Type** User **interface**{
    *name_func1* **tipe**
    *name_func2* **tipe**
    *name_func3* **tipe**
    }

Podemos sobreescribir cada metodo para una estructura diferente.

>Por ejemplo:
**func** (this name_struc) **name_func1 tipe**{
    }
    *Para cada estructura el metodo hace algo distinto*

Ademas, podemos implementar una funcion que reciba a la interfaz como parametro, esto nos permite que cualquer estructura que implemente la interfaz, podra hacer uso de dicha funcion.

>Por ejemplo:
**func** auth(*var* **name_interface**) **tipe**{
    }

###*GOroutines*

Entramos en la seccion de concurrencia dentro del lenguaje. Go no nos permite crear hilos como en java, ejecutamos "hilos" implementados a nivel de software, lo cual los hace mas livianos que un hilo comun. Go cuenta con un balanceador de cargas lo cual administra las distintas **Go routines** que asignemos. Se comportan de tal manera que pareciera que fueran hilos independientes, pero a bajo nivel no lo son. 

Si realziamos la llamada de una funcion y anteponemos la instruccion **go** bastara con ejecutar esa parte del codigo en concurrencia con el hilo Main.

>Por ejemplo:
**go** name_func()

Algo que se debe aclarar es que el hilo padre no esperara a que el hijo termine su ejecucion, por lo que si no se le aclara, el Main finalizara su ejecucion matando a cualquier hilo que se encuentra ejecutando en ese momento. 

Otra manera de crear concurrencia es en lugar de llamar una funcion anteponiendo la instruccion **go** podemos ejecutar un bloque de codigo de manera concurrente. 

>Por ejemplo:
**go func()**{
    *Cualquier codigo aqui dentro sera ejecutado en concurrencia*
    }()

######Channels

Los **Channels** o canales nos sirven para comunicar tareas que se esten realizando concurrentemente. Mediante estos canales podemos enviar numeros, cadenas, estructuras, etc.
Para la creacion de un canal debemos utilizar la funcion **make()** antes vista.

>Sintaxis:
channel **:=** make(**chan** *tipe*)

A continuacion debemos indicarle a la **go routine** que utilizaremos un canal, esto es de la siguiente manera:

>Sintaxis:
**go func**(*name_channel* **chan** *tipe*){
    *Codigo a ejecutar*
    }(*name_channel*)

Para ingresar o extraer datos desde el **channel** se utiliza el operador **<-**.

>Sintaxis:
*name_channel* **<-** *var*
*var* **<-** *name_channel*

En el repositorio se adjunta un ejercicio **ej_channel/channels.go** ejemplificando mejor un codigo que haga uso de esta herramienta.

###*Lectura de archivos*

######Version 1:

Implementaremos una nueva libreria llamada **io/ioutil** la cual contiene las funciones necesarias para la lectura de archivos.

La siguiente manera es la mas sencilla pero tiene una ventaja que veremos acontinuacion:

>Sintaxis:
data_file,err **:=** ioutil.ReadFile("*./ruta*")
*El valor entregado por la funcion es en Bytes*
**string**(data_file)
*Ejemplo de como convertirlo en una cadena*


- La ruta se la podemos especificar con "./MiRuta." Con el ./ indicamos que el archivo se eucneutra en la misma carpeta que el binario. 
- Si el archivo no esta en la misma carpeta se debera indicar la ruta absoluta mediante **/**,  por ejemplo: "/MiRuta"

Go para leer el archvio de esa manera lo cargara en su totalidad en memoria, lo cual si nuestro archivo es grande podria traernos problemas. Ademas de que si quisieramos imprimir, nos imprime la totalidad del mismo.

######Version 2:

Para la siguiente version necesitamos dos librerias antes vistas **os** y **bufio**. Estas librerias nos permiten leer el archivo linea por linea.
Lo primero que debemos hacer es abrir el archivo:

>Sintaxis:
*name_var*,*err* **:=** os.**Open**("./MiRuta")
*La funcion nos retorna el puntero a la primera linea del archivo y err es la por si se genera un error al intentar abrirlo*

Paso siguiente deberemos generar un **Scanner** para ir leyendo las lineas. 

>Sintaxis:
*name_var* **:=** bufio.**NewScanner**(*name_file*)

Una vez creado el Scanner podremos ir linea por linea leyendo el archivo de diferentes maneras.
Dentro del ejemplo en el repositorio se encuentra un ejemplo detallado. Buscar en **ej_Readfile/readfile.go**.





