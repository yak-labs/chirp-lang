set h [hash]
hset $h title Demo
hset $h blob {The quick brown fox jumps over the lazy dog.}
puts [template {<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>{{hget $h title}}</test>
</head>
<body>
  {{hget $h blob}}
  {#
    This comment should not appear in the output.
  #}
</body>
</html>
}]
