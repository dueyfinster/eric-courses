-module(echo).
-export([start/0, stop/0, listen/0, print/1]).

start() ->
    register(echo, spawn(echo, listen, [])),
    ok.

stop() ->
    echo ! stop,
    ok.

listen() ->
    receive 
        {print, Message} ->
		 io:format("~p~n", [Message]),
		 listen();
	stop ->
            true
    end.

print(Message) ->
    echo ! {print, Message},
    ok.
