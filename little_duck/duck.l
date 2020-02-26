%{
    #include <iostream>
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
{integer}   { std::cout << "Found: INTEGER" << std::endl;}
{float}     { std::cout << "Found: FLOAT" << std::endl;}
{id}        { std::cout << "Found: ID" << std::endl; }


%%
int main(int, char**) {
    while(yylex());
}