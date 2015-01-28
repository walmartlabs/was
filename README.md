was
===


Description:
============

Stupid simple but useful tool to move a file or directory and move it back later.
Was moves a list of files to files with a .was extension, and/or moves them back if they already have a .was extension.
Moves directories, and handles conflicts in a sensible way.

Some file's in your way and you'll probably want it again later?  Was it.

You want to mess with some code and are pretty sure you'll screw it up?  Was it first.

You've installed a database, and want to move it aside and test a clean one?  Was it.

You've cloned you favorite project and want to clone some other guy's fork for a quick peek?  Was yours then clone his.

It sounds stupid but its very handy.

It leaves them all right in the same dir so you can see what you've been doing.

Examples
========

    was thisFile -> thisFile.was
    was thisFile.was -> thisFile
    was thisFile thatFile.was -> thisFile.was thatFile

    was filename1 [filename2 filename3 ...]

WIP

Make it return non-zero if there were any errors

Let user choose the extension.

Read file list from STDIN

