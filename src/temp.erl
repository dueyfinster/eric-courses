-module(temp).
-export([convert/1]).

-spec f2c(number()) -> number().

f2c(F) ->
    5*(F-32)/9.

-spec c2f(number()) -> number().

c2f(C) ->
    9*C/5+32.

-spec convert({'c',number()}) -> {'f',number()};
             ({'f',number()}) -> {'c',number()}.

convert({c, C}) ->
    {f,c2f(C)};
convert({f, F}) ->
    {c,f2c(F)}.  
