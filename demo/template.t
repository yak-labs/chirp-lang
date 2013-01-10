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

  {% set b 1 %}
  <p>
    The variable $b is {{set b}}
  </p>

  {% if $b =}
    <p>
      The variable $b is true.
    </p>
  {%}
  {#
  {= else =}
    <p>
      The variable $b is false.
    </p>
  {%}
  #}
</body>
</html>
}]
