# ParrilladaGolang


Trabajo práctico realizado en GO. 

Algunas aclaraciones importantes:

Por motivos de tiempo no pude dejarlo UP en alguna instancia cloud como AWS. 

Esto no impide que el proyecto se pueda levantar localmente 
(con Go instalado o dockerizado) y los tests puedan correrse tambien localmente. 

De todos modos, estoy mas que dispuesto a defenderlo
localmente levantandolo, debatiendo deudas técnicas, etc.

Por otro lado, al querer adaptar este trabajo a una API, 
me tome algunas licencias con los requests y responses, 
asi como también con algunas decisiones de lógica no expresas. 
También, por tiempo, no pude documentar esto para que sea mas facil el 
entendimiento. De nuevo, quedo totalmente a disposicion para explicarlo.

Por último y antes de dejar una copia del README original, quiero aclarar algunas
carácteristicas no menores.

En un principio comencé el desarrollo con Java (lenguaje que no trabajo hace mucho)
y que al avanzar me comenzó a presentar dudas en cuanto a prácticas en la implementación.

Por esto, decidi hacer un remake contrarreloj en Go. Lamentablemente,
el lenguaje no tiene una orientación/soporte NATIVO para manejar herencia, polimorfismo, etc,
por lo que debí (mediante prácticas) simular estos recursos que por ej en Java son más directos.

De todos modos, se puede observar que mediante por ej el patron de diseño Composite y el uso 
de interfaces, 
se logro polimorfismo y herencia. Incluso en algunos lugares se intento replicar clases y métodos abstractos.

Por lo anterior, también el código puede no ser el más limpio ya que estas adaptaciones impactan directamente sobre 
las buenas practicas y/o estilos.

De todos modos, quedo totalmente abierto a consultas, preguntas o lo que se desee.

Soy netamente consciente de la deuda técnica en algunos puntos así como también la necesidad de subir el coverage de tests.

Desde ya muchas gracias, 

Martín Ezequiel Abogado

####### Enunciado

<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Strict//EN" "http://www.w3.org/TR/html4/strict.dtd">
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8"> 





</head>
<body class="markdown-body">
<h1> <a id="la-parrilla-de-rosendo" class="anchor" href="#la-parrilla-de-rosendo" aria-hidden="true"><span aria-hidden="true" class="octicon octicon-link"></span></a>La Parrilla de Rosendo</h1> 
<p>Rosendo tiene una parrilla y nos pidi&oacute; un sistema para administrar su men&uacute;.</p> 
<p>Original: https://github.com/obj1-unahur/parrillada_2020s2</p>
<h2> <a id="1-comidas" class="anchor" href="#1-comidas" aria-hidden="true"><span aria-hidden="true" class="octicon octicon-link"></span></a>1. Comidas</h2> 
<p>Nos piden modelar los distintos platos que ofrece la parrilla. De cada <strong>Plato</strong> nos interesa conocer:</p> 
<ul> 
 <li>su <strong>peso</strong>, medido en gramos;</li> 
 <li>si es <strong>apto vegetariano</strong>;</li> 
 <li>su <strong>valoraci&oacute;n</strong>, un n&uacute;mero que indica qu&eacute; tan bueno es el plato;</li> 
 <li>si <strong>es abundante</strong>, lo cual es cierto cuando su peso es mayor a 250 gramos.</li> 
</ul> 
<p>Consideraremos inicialmente estos platos:</p> 
<h3> <a id="provoleta" class="anchor" href="#provoleta" aria-hidden="true"><span aria-hidden="true" class="octicon octicon-link"></span></a>Provoleta</h3> 
<p>Cada provoleta tiene un peso diferente y pueden tener o no especias. Las provoletas son <em>apto vegetariano</em> si no tiene especias. Tambi&eacute;n decimos que una provoleta <strong>es especial</strong> cuando se cumple alguna de estas condiciones:</p> 
<ul> 
 <li><em>es abundante</em></li> 
 <li>&oacute; <em>tiene especias</em>;</li> 
</ul> 
<p>Su <em>valoraci&oacute;n</em> es de 120 puntos si es especial, y de 80 en caso contrario.</p> 
<h3> <a id="hamburguesas-de-carne" class="anchor" href="#hamburguesas-de-carne" aria-hidden="true"><span aria-hidden="true" class="octicon octicon-link"></span></a>Hamburguesas de carne</h3> 
<p>Su <em>peso</em> es siempre de 250 gramos y l&oacute;gicamente no son <em>apto para vegetariano</em>. A cada hamburguesa se le configura su pan, y la <em>valoraci&oacute;n</em> del plato se calcula como <code>60 + valoraci&oacute;n del pan</code>.</p> 
<p>Los tres &uacute;nicos panes posibles en nuestro modelo son:</p> 
<ul> 
 <li> <strong>industrial</strong> otorga 0 puntos;</li> 
 <li> <strong>casero</strong> otorga 20 puntos;</li> 
 <li> <strong>de masa madre</strong> otorga 45 puntos.</li> 
</ul> 
<h3> <a id="hamburguesas-vegetarianas" class="anchor" href="#hamburguesas-vegetarianas" aria-hidden="true"><span aria-hidden="true" class="octicon octicon-link"></span></a>Hamburguesas vegetarianas</h3> 
<p>Se comportan igual que las hamburguesas de carne, con tres diferencias:</p> 
<ul> 
 <li>son <em>apto para vegetarianos</em>;</li> 
 <li>para cada hamburguesa, se informa mediante un string de qu&eacute; legumbre est&aacute; hecha (por ejemplo: <code>&quot;lentejas&quot;</code> o <code>&quot;garbanzos&quot;</code>);</li> 
 <li>a la valoraci&oacute;n se le suma otro plus que, como valor m&aacute;ximo, puede ser <em><strong>17</strong></em>, y se calcula como <code>2 * cantidad de letras del nombre de la legumbre</code>. Por ejemplo, si es de lentejas (que tiene 8 letras) el plus ser&aacute; de 16, pero si fuera garbanzos (que tiene 9 letras) ser&aacute; 17 (el m&aacute;ximo permitido).</li> 
</ul> 
<p><strong>Tip:</strong> para sacar la cantidad de letras de una palabra podes utilizar el mensaje <code>size()</code> a un string.</p> 
<h3> <a id="parrillada" class="anchor" href="#parrillada" aria-hidden="true"><span aria-hidden="true" class="octicon octicon-link"></span></a>Parrillada</h3> 
<p>Para cada parrillada se nos indica los cortes de carne pedidos. De cada corte se conoce su nombre, <em>calidad</em> (un n&uacute;mero) y su <em>peso</em>.</p> 
<p>El <em>peso</em> de la parrillada es la suma de los pesos de sus cortes. No es <em>apto vegetariano</em>. La <em>valoraci&oacute;n</em> se calcula como <code>15 * m&aacute;xima calidad de sus cortes - cantidad de cortes</code>, y no puede dar un resultado negativo.</p> 
<p>Por ejemplo, si una parillada incluye los tres siguientes cortes</p> 
<ul> 
 <li>asado, de calidad 2 y peso 800 gramos,</li> 
 <li>vacio, de calidad 3 y peso 1200 gramos,</li> 
 <li>matambre de cerdo, de calidad 1 y peso 1350 gramos,</li> 
</ul> 
<p>El peso de la parrillada ser&iacute;a de <strong>3350 gramos</strong> y la valoracion ser&iacute;a de (15 * 3) - 3 = <strong>42</strong></p> 
<h2> <a id="2-comensales" class="anchor" href="#2-comensales" aria-hidden="true"><span aria-hidden="true" class="octicon octicon-link"></span></a>2. Comensales</h2> 
<p>Ya tenemos la comida, ahora nos faltan los comensales. 
 <g-emoji class="g-emoji" alias="fork_and_knife" fallback-src="https://github.githubassets.com/images/icons/emoji/unicode/1f374.png">
  🍴
 </g-emoji></p> 
<p>De cada comensal nos interesa saber:</p> 
<ul> 
 <li>su <strong>peso</strong>, medido en gramos;</li> 
 <li>si <strong>le agrada o no una comida</strong>, lo cual depender&aacute; de tipo de comensal. Implementar el m&eacute;todo <code>leAgrada(unaComida)</code> a cada tipo de comensal.</li> 
 <li>las <strong>comidas que comi&oacute;</strong>, una lista de todo lo que haya ingerido. Implementar un m&eacute;todo <code>comer(unaComida)</code> que la agregue a la lista;</li> 
 <li>y, si est&aacute; <strong>satisfecho o no</strong>, lo cual explicaremos a continuaci&oacute;n.</li> 
</ul> 
<p>Para que un comensal est&eacute; satisfecho, se tiene que cumplir que el peso de las comidas ingeridas sea mayor o igual al 1% de su propio peso <em>y adem&aacute;s</em> una condici&oacute;n que define cada comensal (para que quede claro: se tienen que cumplir <em>ambas</em> condiciones).</p> 
<p>Consideraremos los siguientes tipos de comensales:</p> 
<h3> <a id="vegetarianos" class="anchor" href="#vegetarianos" aria-hidden="true"><span aria-hidden="true" class="octicon octicon-link"></span></a>Vegetarianos</h3> 
<p>Les agradan las comidas que son apto vegetariano y tienen una valoraci&oacute;n mayor a 85. La condici&oacute;n adicional para estar satisfechos es que ninguna comida de las ingeridas sea abundante.</p> 
<h3> <a id="hambre-popular" class="anchor" href="#hambre-popular" aria-hidden="true"><span aria-hidden="true" class="octicon octicon-link"></span></a>Hambre popular</h3> 
<p>Simplemente les agradan las comidas abundantes. No tienen ninguna condici&oacute;n adicional para estar satisfechos.</p> 
<h3> <a id="de-paladar-fino" class="anchor" href="#de-paladar-fino" aria-hidden="true"><span aria-hidden="true" class="octicon octicon-link"></span></a>De paladar fino</h3> 
<p>Les agradan las comidas que pesan entre 150 y 300 gramos, y adem&aacute;s tienen una valoraci&oacute;n mayor a 100. La condici&oacute;n adicional para satisfacerse es que la cantidad de comidas ingeridas sea par.</p> 
<h2> <a id="3-cocina" class="anchor" href="#3-cocina" aria-hidden="true"><span aria-hidden="true" class="octicon octicon-link"></span></a>3. Cocina</h2> 
<p>Agregar al modelo la cocina, que tiene <em>todas las comidas</em> que la parrilla tiene preparadas.</p> 
<p>Se quiere poder consultar:</p> 
<ul> 
 <li>si <strong>tiene buena oferta vegetariana</strong>: esto es as&iacute; si la diferencia entre comidas vegetarianas y no vegetarianas es de a lo sumo 2. Por ejemplo: si hay 10 carn&iacute;voras y 8 vegetarianas s&iacute; tiene buena oferta, pero si hay 11 carn&iacute;voras no (porque la diferencia es mayor a 2). Otro ejemplo, si hay 11 vegetarianas y 9 carn&iacute;voras, tiene buena oferta vegetariana.</li> 
 <li> <strong>el plato fuerte carn&iacute;voro</strong>: el de m&aacute;xima valoraci&oacute;n entre los no apto vegetariano;</li> 
 <li>dado un comensal, la lista de <strong>comidas que le gustan</strong>.</li> 
</ul> 
<p>Tambi&eacute;n se pide <strong>poder elegir un plato</strong> para un comensal - por ahora es cualquier plato que le guste. Si no le gusta ning&uacute;n plato, lanzar un error. Si el plato existe, sacarlo de la cocina y hacer que el comensal lo coma.</p> 
<h2> <a id="4-test" class="anchor" href="#4-test" aria-hidden="true"><span aria-hidden="true" class="octicon octicon-link"></span></a>4. Test</h2> 
<p>Realizar al menos los siguientes tests.</p> 
<h3> <a id="41-comidas" class="anchor" href="#41-comidas" aria-hidden="true"><span aria-hidden="true" class="octicon octicon-link"></span></a>4.1 Comidas</h3> 
<ul> 
 <li>Valoraci&oacute;n de la Hambruguesa de Carne con pan casero. El resultado ser&aacute;: <strong>80</strong> (60 + 20)</li> 
 <li>Valoraci&oacute;n de la Hambruguesa Vegetariana, con pan de masaMadre y de garbanzos como legumbre. El resultado ser&aacute;: <strong>122</strong> (60 + 45 + 17)</li> 
 <li>Provoleta, de peso 190 y con especias. El resultado de ser especial es: <strong>true</strong> </li> 
 <li>Parillada, que incluya un asado, de calidad 10 y peso 250 gramos, una entra&ntilde;a, de calidad 7 y peso 200 gramos y un chorizo, de calidad 8 y peso 50 gramos. El resultado de la valoracion ser&aacute; de <strong>147</strong> (15 * 10) - 3.</li> 
</ul> 
<h3> <a id="42-comensales" class="anchor" href="#42-comensales" aria-hidden="true"><span aria-hidden="true" class="octicon octicon-link"></span></a>4.2 Comensales</h3> 
<ul> 
 <li>Un vegetariano de peso 68500 gramos come una provoleta de 190 gramos con especias y dos Hamburguesas Vegetarianas (250 grs cada una) con pan de masaMadre y de garbanzos como legumbre. El resultado si esta satisfecho es: <strong>True</strong> porque (190 + 250 + 250) &gt;= (68500 * 0.01) y adem&aacute;s ninguna de las tres comidas es abundante, es decir, ninguna supera los 250 gramos.</li> 
 <li>Un comensal popular de peso 85000 gramos come una parrillada que incluye un asado, de calidad 10 y peso 250 gramos, una entra&ntilde;a, de calidad 7 y peso 200 gramos y un chorizo, de calidad 8 y peso 50 gramos. El resultado si esta satisfecho es: <strong>False</strong> porque 500 &lt; 850</li> 
 <li>Un comensal de paladar fino le agrada comer una Hamburguesa de carne de 250 gramos con pan de masa madre.</li> 
 <li>Un comensal de paladar fino NO agrada comer una Hamburguesa de carne de 250 gramos con pan casero.</li> 
</ul>
</body>
</html>


