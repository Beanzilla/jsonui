# JSONUI

**Currently in developmental state**

## Node structure

Based on my [gonode](https://git.red-green.com/david/gonode).

Changes:

- Data of type `string` not of type `any`.
- No Excess Methods (keep it simple, avoid many functions that aren't needed, make things public/exported for quick access instead)
- Added Collapsed `bool` to track if we should iterate over this Node or ignore it (this will allow collapsed nodes to not be seen in the Text side)
- Depth is predefined (currently is off by one, users should subtract 1 to get the correct depth, needs fix, see issue #1)

## The Project History

This Project is a remake of [JSONUI](https://github.com/gulyasm/jsonui) by [gulyasm](https://github.com/gulyasm).

Major changes:

- Possible Performance gain
- Graphical fixes
- Complete rebuild
