was
===

Its a new verb.  "Was it"


Description:
============

Stupid simple but useful tool to move a file/files or directory and
move it/them back later.  Once you have this, you'll use it often,
and you'll be annoyed if you don't have it installed somewhere.

Was moves a list of files to files with a .was extension, and/or
moves them back if they already have a .was extension.  It moves
directories, and handles conflicts in a sensible way. It sort of
toggles the name between whatever and whatever.was.

Maybe some file's in your way and you'll probably want it again in
like 5 minutes after you're done goofing around?  Was it.

You want to mess with some code or a config file and are pretty
sure you'll screw it up?  Was it first.

You've installed a database, and want to move it aside and test a
clean one?  Was the entire thing.

You've cloned you favorite project and want to clone some other
guy's fork for a quick peek?  Was yours then clone his.

It sounds stupid but its very handy.

It leaves everything right in the same dir so you can see what
you've been doing, and clean up old ones very easily.

Examples
========

    was thisFile -> thisFile.was
    was thisFile.was -> thisFile
    was thisFile thatFile.was -> thisFile.was thatFile
    was -c thisFile -> thisFile thisFile.was

    was filename1 [filename2 filename3 ...]

    was --help
      .
      .
      .
      -c=false: copy instead of move
      -f=false: clobber any conflicting files
      -v=false: verbose output



WIP

Make it return non-zero if there were any errors

Let user choose the extension.

Read file list from STDIN

