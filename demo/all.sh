set -ex
cd $(dirname $0)
( cd .. ; make )

../goterp hello_world.t 2>_world.err | grep '^Hello World$'

../goterp hello_web.t 2>_web.err &
P=$!
sleep 3
wget --output-document=/dev/stdout --server-response 'http://localhost:8080/' 2>&1 | fgrep '200 OK'
wget --output-document=/dev/stdout --server-response 'http://localhost:8080/' 2>&1 | fgrep 'Hello Web!'
kill $P

../goterp triang.t 2>_triang.err | grep '^The triangs under 100 are 0 1 3 6 10 15 21 28 36 45 55 66 78 91$'

../goterp wiki.t 2>_wiki.err &
P=$!
sleep 3
wget --output-document=/dev/stdout --server-response 'http://localhost:8080/' 2>&1 | fgrep '200 OK'
wget --output-document=/dev/stdout --server-response 'http://localhost:8080/' 2>&1 | fgrep 'This is the home page.'
wget --output-document=/dev/stdout --server-response 'http://localhost:8080/list' 2>&1 | fgrep '200 OK'
wget --output-document=/dev/stdout --server-response 'http://localhost:8080/list' 2>&1 | fgrep '<a href="/view?page=JoePage">"JoePage"</a>'
wget --output-document=/dev/stdout --server-response 'http://localhost:8080/view?page=JoePage' 2>&1 | fgrep '200 OK'
wget --output-document=/dev/stdout --server-response 'http://localhost:8080/view?page=JoePage' 2>&1 | fgrep "This is Joe's page."
kill $P

cat << END
@@@@@@@@@@@@@@@
@             @
@   SUCCESS   @
@             @
@@@@@@@@@@@@@@@
END
