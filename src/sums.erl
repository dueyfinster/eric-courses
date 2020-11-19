-module(sums).
-export([sum/1, sum_interval/2]).


sum(0) ->
    0;
sum(N) ->
    N + sum(N-1).

sum_interval(Max, Max) ->
    Max;
sum_interval(Min, Max) when Min =< Max ->
    Min + sum_interval(Min + 1, Max).

