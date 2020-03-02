%{
    #include <iostream>
    #include "duck.tab.h"
    extern int yylex();
%}

%option noyywrap

delim   [ \t\n]
ws      {delim}+
alfa    [A-Za-z]
digit   [0-9]
id      {alfa}({alfa}|{digit})*
integer {digit}+
float   {digit}+\.({digit})+

%%

{ws}        {}

"if"          { return (IF); }
"else"        { return (ELSE); }
"var"         { return (VAR);}
"int"         { return (INTTYPE);}
"float"       { return (FLOATTYPE); }
"print"       { return (PRINT); }
"+"           { return (PLUS); }
"-"           { return (MINUS); }
"*"           { return (MULT); }
"/"           { return (DIV); }
">"           { return (RELOP); }
"<"           { return (RELOP); }
"<>"          { return (RELOP); }

{integer}   { return (INT);}
{float}     { return (FLOAT);}
{id}        { return (ID); }

%%

int main(int, char**) {
    while(yylex());
}