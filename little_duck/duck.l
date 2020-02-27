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

"if"          { std::cout << "Found IF" << std::endl; }
"else"        { std::cout << "Found ELSE" << std::endl; }
"var"         { std::cout << "Found VAR" << std::endl; }
"int"         { std::cout << "Found INT TYPE" << std::endl; }
"float"       { std::cout << "Found FLOAT TYPE" << std::endl; }
"{"           { std::cout << "Found OPEN BREACKET" << std::endl; }
"}"           { std::cout << "Found CLOSE BREACKET" << std::endl; }
"("           { std::cout << "Found OPEN PAREN" << std::endl; }
")"           { std::cout << "Found CLOSE PAREN" << std::endl; }
";"           { std::cout << "Found SEMICOLON" << std::endl; }
"+"           { std::cout << "Found PLUS" << std::endl; }
"-"           { std::cout << "Found MINUS" << std::endl; }
"*"           { std::cout << "Found MULTI" << std::endl; }
"/"           { std::cout << "Found DIV" << std::endl; }

{integer}   { std::cout << "Found: INTEGER" << std::endl;}
{float}     { std::cout << "Found: FLOAT" << std::endl;}
{id}        { std::cout << "Found: ID" << std::endl; }

%%

int main(int, char**) {
    while(yylex());
}