list -- TemplaTcl {http://wiki.tcl.tk/18175}

proc TemplateCode {txt} {
  set code "set res {}\n"
  while {[set i [string first <% $txt]] != -1} {
    incr i -1
    puts "This is i: $i"
    append code "append res [list [string range $txt 0 $i]]\n"
    set txt [string range $txt [expr {$i + 3}] end]
    if {[string index $txt 0] eq "="} {
      append code "append res "
      set txt [string range $txt 1 end]
    }
    puts "closer: [string first %> $txt]"
    if {[set i [string first %> $txt]] == -1} {
      error "No matching %>"
    }
    incr i -1
    append code "[string range $txt 0 $i]\n"
    set txt [string range $txt [expr {$i + 3}] end]
  }
  if {$txt ne ""} { append code "append res [list $txt]\n" }
  return $code
}

proc Template {txt} {
  set code [TemplateCode $txt]
  eval $code
  return $res
}

set templ {
<!DOCTYPE html>
<html lang="en">
<head>
<%
  set title Demo
  set blob {The quick brown fox jumps over the lazy dog.}
%>
  <meta charset="utf-8">
  <title><%= $title %></title>
</head>
<body>
  <p><%= $blob %></p>

  <% set b 1 %>
  <p>The variable $b is <%= $b %></p>

  <% if $b { %>
    <p>The variable $b is true.</p>
  <% } else { %>
    <p>The variable $b is false.</p>
  <% } %>
</body>
</html>
}

puts [TemplateCode $templ]
puts [Template $templ]
