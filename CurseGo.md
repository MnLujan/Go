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

