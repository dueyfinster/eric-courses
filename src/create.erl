-module(create).
-export([create/1, reverse_create/1]).

create(N) ->
    create(1,N).

create(M,M) ->
    [M];
create(M,N) ->
    [M | create(M+1, N)].

reverse_create(N) ->
    reverse_create(1,N).
reverse_create(M,M) ->
    [M];
reverse_create(M,N) ->
    create(M-1, N).

