-module(boolean).
-export([b_not/1, b_and/2, b_or/2]).

b_not(B) ->
    not B.

b_and(B, C) ->
    B and C.

b_or(B, C) ->
    B or C.
