docker run -it --rm --name erlang-inst1 -v "$PWD/src":/usr/src/myapp -w /usr/src/myapp erlang:18 escript sums.erl
