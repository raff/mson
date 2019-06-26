# mson
A parser for MongoDB shell "dumps"

Convert MongoDB "extended" JSON / BSON to Go maps or JSON.

This is mostly a JSON parser modified to support MongoDB shell BSON objects (BinData, NumberLong, ISODate).
It also support parsing multiple consecutive object (i.e. a stream of objects not wrapped into an array).

The parser returns a Go map (map[string]interface{}) or array ([]interface{}).

A command line tool is available in cmd/mson that converts the input MongoDB file to JSON.
The tool can also skip initial "garbage" that may have been written by MongoDB shell before the real data (see --prologue option).
