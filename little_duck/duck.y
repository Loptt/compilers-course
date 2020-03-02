%{
    #include <iostream>
%}

%token IF 
%token ELSE 
%token VAR
%token INTTYPE
%token FLOATTYPE  
%token PRINT
%token PLUS
%token MINUS
%token MULT
%token DIV
%token RELOP
%token ID
%token INTEGER
%token FLOAT
%token STRING

%%

programa : ID ':' vars bloque   
         | ID ':' bloque       

bloque : '{' estatutos '}' 

estatutos : estatuto estatutos 
          | estatuto 
          |

estatuto : asignacion
         | condicion
         | escritura

asignacion : ID '=' expresion ';'

expresion : exp
          | exp RELOP exp

exp : termino
    | termino PLUS termino
    | termino MINUS termino

termino : factor
        | factor MULT factor
        | factor DIV factor

factor : '(' expresion ')'
       | PLUS varcte
       | MINUS varcte
       | varcte

varcte : ID
       | INTEGER
       | FLOAT

condicion : IF '(' expresion ')' bloque ';'
          | IF '(' expresion ')' bloque ELSE bloque ';'

escritura : PRINT '(' escrituraarg ')' ';'

escrituraarg : expresion ',' escrituraarg
             | STRING ',' escrituraarg
             | expresion
             | STRING

vars : VAR varargs

varargs : ids ':' tipo ';' varargs
        | ids ':' tipo ';'

ids : ID ',' ids
    | ID

tipo : INTTYPE
     | FLOATTYPE

%%

