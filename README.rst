Chirp
================================================================================

Chirp is a dynamic language inspired by Tcl_ and written in Go_.

.. image:: https://drone.io/github.com/chirp-lang/chirp/status.png
   :alt: Build Status
   :target: https://drone.io/github.com/chirp-lang/chirp/latest
   :width: 80
   :height: 18

Building Chirp
--------------------------------------------------------------------------------

First, you'll need to download and install Git_.  Once you have Git installed,
make sure that the ``git`` command is available from the command line.  If not,
you'll need to add it to your ``PATH`` environment variable.

Next, you'll need to download and install Go_.  Once it's installed, you'll
need to create a workspace and set your ``GOPATH`` environment variable to your
workspace folder.  The Go documentation contains a great page explaining the
details on setting up your workspace titled, `How to Write Go Code`_.  There is
also a nice `screencast`_ available as well.

Once you've setup your workspace, open up a command line to your workspace
folder and run the following::

    go get github.com/yak-people/chirp-lang
    cd src/github.com/yak-people/chirp-lang/chirp
    git checkout master
    go install

This will fetch the Chirp source code, ensure you're on the ``master`` branch,
and build the command line executable to run Chirp code.  If you've added ``GOPATH/bin`` to your ``PATH``, you should now be able to run ``chrip`` from your command line.

.. _Tcl: http://tcl.tk/
.. _Go: http://golang.org/
.. _Git: http://git-scm.com/
.. _How to Write Go Code: http://golang.org/doc/code.html
.. _screencast: http://youtu.be/XCsL89YtqCs
