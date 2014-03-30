flowette
========

`flowette` is the future backend server for [http://flostatus.neocities.org](http://flostatus.neocities.org).

It's unlikely that this project is useful to you and It's far from being finished.

Instructions
------------

After installing, you should use the `flowette` command to create a new database file

    $ flowette new flowette.db
    Database successfully initialised!

You can insert new data now

    $ flowette records -a -d 2014-03-30 -s true
    Record added

You can also list all entries in the database

    $ flowette records
    Date        Status
    2014-03-30  true

You might also want to start the server

    $ flowette serve


