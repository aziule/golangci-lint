//args: -Efilebuildtag
//config_path: testdata/configs/filebuildtag.yml
package testdata

// +build foo // ERROR "build comment must appear before package clause and be followed by a blank line$"
