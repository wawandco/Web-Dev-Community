### Acerca de
Este repositorio es una extensión y continuación del taller [Pet Store Loneliness 2000](https://github.com/wawandco/Web-Dev-Community/blob/main/petstore/README.md) donde hubo una exploración y aplicación del paquete de buffalo [Pop v6.0](https://github.com/gobuffalo/pop). Con la intención de darle un enfoque al presente taller y seguir explorando las herramientas que usamos en nuestro día a día, no será necesario completar el taller anterior para desarrollar este. Los invitamos a leer en primer lugar las siguientes [instrucciones](https://github.com/wawandco/Web-Dev-Community/blob/main/petstorehttp/INSTRUCTIONS.md) que serán el punto de partida para el desarrollo del taller. Muchas gracias por participar y les deseamos el mejor de los éxitos.

### Petstore loneliness 2000 API
En la actualidad Petstore loneliness 2000 continúa en una constante búsqueda por optimizar sus procesos. Una de sus mayores necesidades en estos momentos es la de habilitar herramientas que les permitan revisar el estado actual del negocio en tiempo real. Para llevar a cabo la implementación, Petstore loneliness 2000 y su equipo de desarrollo han decidido crear un prototipo web en forma de servicio API que les facilite integración futura con otras herramientas para analisis de datos.

Dentro de los requerimientos formulados se encuentran los siguientes:

1. Crear un endpoint que permita obtener la lista de las mascotas registradas en la tienda.
    - Solo debe recibir peticiones GET
    - La respuesta exitosa, debe responder con su respectivo código HTTP y debe proveer un JSON con una lista que incluya toda la información de las mascotas.

2. Crear un endpoint que permita obtener la información completa de una mascota específica.
    -  Solo debe recibir peticiones GET
    -  La respuesta exitosa, debe responder con su respectivo código HTTP y debe proveer un JSON con la información de la mascota asociada al ID provisto en la petición.
    -  Si no existe una mascota asociada al ID provisto, la respuesta debe presentarse con su respectivo código de HTTP y un JSON con un mensaje de error que indique que no se encontró la mascota requerida. 

3. Crear un endpoint que permita eliminar la información completa de una mascota específico.
    -  Solo debe recibir peticiones DELETE
    -  La respuesta exitosa, debe responder con su respectivo código HTTP y debe proveer un JSON con la información de la mascota que fue eliminada.
    -  Si no existe una mascota asociada al ID provisto, la respuesta debe presentarse con su respectivo código de HTTP y un JSON con un mensaje de error que indique que no se encontró la mascota requerida. 

4. Crear un endpoint que permita obtener la lista de clientes registrados en la tienda.
    -  Solo debe recibir peticiones GET
    -  La respuesta exitosa, debe responder con su respectivo código HTTP y debe proveer un JSON con la lista que contenga la información de todos los clientes.


5. Crear un endpoint que permita obtener la información completa de un cliente específico.
    -  Solo debe recibir peticiones GET
    -  La respuesta exitosa, debe responder con su respectivo código HTTP y debe proveer un JSON con la información del cliente asociado al ID provisto en la petición.
    -  Si no existe un cliente asociado al ID provisto, la respuesta debe presentarse con su respectivo código de HTTP y un JSON con un mensaje de error que indique que no se encontró el cliente. 


6. Crear un endpoint que permita eliminar la información completa de un cliente específico.
    -  Solo debe recibir peticiones DELETE
    -  La respuesta exitosa, debe responder con su respectivo código HTTP y debe proveer un JSON con la información del cliente que fue eliminado.
    -  Si no existe un cliente asociado al ID provisto, la respuesta debe presentarse con su respectivo código de HTTP y un JSON con un mensaje de error que indique que no se encontró el cliente requerido. 

7. Crear un endpoint que permita registrar una nueva venta para una mascota específica.
    -  Solo debe recibir peticiones POST
    -  La respuesta exitosa, debe responder con su respectivo código HTTP y debe proveer un JSON con la información de la venta realizada.
    -  Si no existe el cliente o la mascota provista en la petición, la respuesta debe presentarse con su respectivo código de HTTP y un JSON con un mensaje de error que indique la entidad que no se encontró.

8. Crear un endpoint que permita obtener las mascotas compradas por un cliente específico.
    -  Solo debe recibir peticiones GET
    -  La respuesta exitosa, debe responder con su respectivo código HTTP y debe proveer un JSON con la información de la venta asociada al ID provisto en la petición.
    -  Si no existe una venta asociada al ID provisto, la respuesta debe presentarse con su respectivo código de HTTP y un JSON con un mensaje de error que indique que no se encontró dicha venta. 

