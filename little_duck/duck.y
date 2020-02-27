%{
    #include <iostream>
%}

%token IF 
%token ELSE 
%token VAR
%token INTTYPE
%token FLOATTYPE  
%token PLUS
%token MINUS
%token MULT
%token DIV
%token ID

%%

programa : ID ':' vars bloque   { std::cout << "PROGRMA" << std::endl; }
         | ID ':' bloque

%%

