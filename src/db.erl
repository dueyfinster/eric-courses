-module(db).
-export([new/0, write/3, read/0, delete/2]).

new() ->
    [].

write(Key, Value, []) ->
    [Key, Value].

read() ->
    [].

delete(Key, []) ->
    [].
