
<-------- Laboratorio 2 Sistemas Distribuidos -------->

Integrantes:
    + Joaquín Castillo Tapia - 201773520-1
    + María Paz Morales - 201773505-8

-------------------------------------------------------

Dependencias instaladas:
    + dist05 - DataNode y Client:
        * git
        * grpc
    + dist06 - DataNode y CLient:
        * git
        * grpc
    + dist07 - DataNode y CLient:
        * git
        * grpc
    + dist08 - NameNode y Client:
        * git
        * grpc

-------------------------------------------------------

Ejecución:
    ** Se recomienda seguir el siguiente orden:

    + dist08 - NameNode:
        1. Dirigirse al la carpeta nameNode mediante:
                $ cd LAB2
                $ cd NameNode

        2. Ejecutar el programa mediante:
                $ make

    + dist05 - DataNode:
        1. Asegurase que el programa en NameNode está ejecutándose.

        2. Dirigirse a la carpeta DataNode mediante:
                $ cd LAB2
                $ cd DataNode

        3. Ejecutar el programa mediante:
                $ make

    + dist06 - DataNode:
        1. Asegurase que el programa en NameNode está ejecutándose.

        2. Dirigirse a la carpeta DataNode mediante:
                $ cd LAB2
                $ cd DataNode

        3. Ejecutar el programa mediante:
                $ make

    + dist07 - DataNode:
        1. Asegurase que el programa en NameNode está ejecutándose.

        2. Dirigirse a la carpeta DataNode mediante:
                $ cd LAB2
                $ cd DataNode

        3. Ejecutar el programa mediante:
                $ make

    + dist05 o dist06 o dist07 o dist08 - Client:
        1. Asegurase que el programa en NameNode y DataNode están ejecutándose.

        2. Dirigirse a la carpeta Client mediante:
                $ cd LAB2
                $ cd Client

        3. Ejecutar el programa mediante:
                $ make
        --> Si se desea simular más de un cliente a la vez se debe abrir otra terminal y seguir los pasos 1 al 3.

	--> Para poder simular los clientes de manera ordenada, decidimos darle una sola funcionalidad a cada ejecución del Client. 
	    Entonces, primero se elige si se quiere simular tipo Upload o Download. Si se elige Upload se solicita el nombre del archivo a subir,
        este archivo debe estar en la misma carpeta que el código de Client. En la máquina hay varios archivos de extención .epub y .pdf 
        de ejemplo y que fueron utilizados para probar el programa. Si se elige Download, se desplegaran los libros disponible, se debe elegir uno
        y este se guardara en la misma carpeta donde se encuentra el código de Client.


- Para finalizar la ejecucion se debe cada ejecucion por separado.
- Registros generados:

    + dist08 - NameNode:
        * Log.txt : Contiene el detalle de todos los libors subidos y las ip de dónde se encuentran sus partes.
    
    + dist05 o dist06 o dist07 - DataNode: 
        * Libros/ : carpeta que almacena los chunks de los libros en cada máquina según le corresponda
    
    + ddist05 o dist06 o dist07 o dist08 - Client:
        * Libros : libros descargados según solcitudes hechas

-------------------------------------------------------

Consideraciones:
- La única forma de que se rechace una propuesta de distribucion es que el DataNode esté caído (no se esté ejecutando)
- Después de verificar qué DataNodes están caídos no habrán nuevas caídas
- La extension de los libors puede ser cualquiera. Las extensiones probadas son .epub y .pdf
- Los nombres de los libors a subir no pueden contener espacios, se recomienda reemplazarlos por guiones bajos. 