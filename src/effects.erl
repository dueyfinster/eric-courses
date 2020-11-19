-module(effects).
-export([print/1, even_print/1]).

print(0) ->
    ok;
print(N) ->
    print(N-1),
    io:format("Number:~p~n", [N]).

even_print(0) ->
    ok;
even_print(N) when N rem 2 == 0 ->
    even_print(N-1),
    io:format("Number:~p~n", [N]);
even_print(N) ->
    even_print(N-1).
