/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * importer.go
 *
 *  Created on Feb 27, 2017
 *      Author Massimiliano Ghilardi
 */

package genimport

import (
	"bytes"
	"fmt"
	"go/build"
	"go/types"
	"io/ioutil"
	"os"
	r "reflect"

	"github.com/cosmos72/gomacro/base/output"
	"github.com/cosmos72/gomacro/base/paths"
	"github.com/cosmos72/gomacro/base/reflect"
	"github.com/cosmos72/gomacro/imports"
)

type ImportMode int

const (
	// ImBuiltin import mechanism is:
	// 1. write a file $GOPATH/src/github.com/cosmos72/gomacro/imports/$PKGPATH.go containing a single func init()
	//    i.e. *inside* gomacro sources
	// 2. tell the user to recompile gomacro
	ImBuiltin ImportMode = iota

	// ImThirdParty import mechanism is the same as ImBuiltin, except that files are created in a thirdparty/ subdirectory:
	// 1. write a file $GOPATH/src/github.com/cosmos72/gomacro/imports/thirdparty/$PKGPATH.go containing a single func init()
	//    i.e. *inside* gomacro sources
	// 2. tell the user to recompile gomacro
	ImThirdParty

	// ImInception import mechanism is:
	// 1. write a file $GOPATH/src/$PKGPATH/x_package.go containing a single func init()
	//    i.e. *inside* the package to be imported
	// 2. tell the user to recompile $PKGPATH
	ImInception

	// ImPlugin import mechanism is:
	// 1. write a file $GOPATH/src/gomacro.imports/$PKGPATH/$PKGNAME.go containing a var Packages map[string]Package
	//    and a single func init() to populate it
	// 2. invoke "go build -buildmode=plugin" on the file to create a shared library
	// 3. load such shared library with plugin.Open().Lookup("Packages")
	ImPlugin
)

func (mode ImportMode) String() string {
	switch mode {
	case ImBuiltin:
		return "Builtin"
	case ImThirdParty:
		return "ThirdParty"
	case ImInception:
		return "Inception"
	case ImPlugin:
		return "Plugin"
	default:
		return fmt.Sprintf("ImportMode(%d)", int(mode))
	}
}

type PackageName = imports.PackageName // package default name, or package alias

type PackageRef struct {
	imports.Package
	Path string
}

func (ref *PackageRef) DefaultName() PackageName {
	return ref.Package.DefaultName(ref.Path)
}

func (ref *PackageRef) String() string {
	return fmt.Sprintf("{%s %q, %d binds, %d types}", ref.DefaultName(), ref.Path, len(ref.Binds), len(ref.Types))
}

type Importer struct {
	srcDir     string
	mode       types.ImportMode
	PluginOpen r.Value // = reflect.ValueOf(plugin.Open)
	output     *Output
}

func DefaultImporter(o *Output) *Importer {
	return &Importer{output: o}
}

func (imp *Importer) havePluginOpen() bool {
	if !imp.PluginOpen.IsValid() {
		imp.PluginOpen = imports.Packages["plugin"].Binds["Open"]
		if !imp.PluginOpen.IsValid() {
			imp.PluginOpen = reflect.NoneR // cache the failure
		}
	}
	return imp.PluginOpen != reflect.NoneR
}

// LookupPackage returns a package if already present in cache
func LookupPackage(alias PackageName, pkgpath string) *PackageRef {
	pkg, found := imports.Packages[pkgpath]
	if !found {
		return nil
	}
	if len(pkg.Name) == 0 {
		// missing pkg.Name, initialize it
		_ = pkg.DefaultName(pkgpath)
		imports.Packages[pkgpath] = pkg
	}
	if len(alias) == 0 {
		// import "foo" => get alias from package name
		alias = pkg.DefaultName(pkgpath)
	}
	return &PackageRef{Package: pkg, Path: pkgpath}
}

func (imp *Importer) wrapImportError(pkgpath string, enableModule bool, err error) output.RuntimeError {
	if rerr, ok := err.(output.RuntimeError); ok {
		return rerr
	}
	if enableModule {
		return imp.output.MakeRuntimeError("error loading package %q metadata: %v", pkgpath, err)
	}
	return imp.output.MakeRuntimeError(
		"error loading package %q metadata, maybe you need to download (go get), compile (go build) and install (go install) it? %v",
		pkgpath, err)
}

func (imp *Importer) ImportPackage(alias PackageName, pkgpath string, enableModule bool) *PackageRef {
	refs, err := imp.ImportPackagesOrError(map[string]PackageName{pkgpath: alias}, enableModule)
	if err != nil {
		panic(err)
	}
	return refs[pkgpath]
}

func (imp *Importer) ImportPackagesOrError(pkgpaths map[string]PackageName, enableModule bool) (map[string]*PackageRef, error) {
	paths.GetImportsSrcDir() // warns if GOPATH or paths.ImportsDir may be wrong

	refs := make(map[string]*PackageRef)
	for pkgpath, alias := range pkgpaths {
		ref := LookupPackage(alias, pkgpath)
		if ref != nil {
			// found in package cache, no need to compile and load this package
			refs[pkgpath] = ref
			delete(pkgpaths, pkgpath)
		}
	}
	if len(pkgpaths) != 0 {
		mode := imp.importModeFor(pkgpaths)
		for pkgpath, alias := range pkgpaths {
			ref, err := imp.importPackageOrError(alias, pkgpath, mode, enableModule)
			if err != nil {
				return nil, err
			}
			refs[pkgpath] = ref
		}
	}
	return refs, nil
}

func (imp *Importer) importModeFor(pkgpaths map[string]PackageName) ImportMode {
	var curr ImportMode
	havemode := false
	for _, alias := range pkgpaths {
		var mode ImportMode
		switch alias {
		case "_b":
			mode = ImBuiltin
		case "_i":
			mode = ImInception
		case "_3":
			mode = ImThirdParty
		default:
			if imp.havePluginOpen() {
				mode = ImPlugin
			} else {
				mode = ImThirdParty
			}
		}
		highest := mode
		if havemode && mode != curr {
			if highest < curr {
				highest = curr
			}
			imp.output.Warnf("attempt to import multiple packages with conflicting modes %v and %v - numerically higher mode %v will be used",
				curr, mode, highest)
		}
		curr = highest
		havemode = true
	}
	return curr
}

func (imp *Importer) importPackageOrError(alias PackageName, pkgpath string, mode ImportMode, enableModule bool) (*PackageRef, error) {

	o := imp.output
	gpkg, err := imp.Load(pkgpath, enableModule) // loads names and types, not the values!
	if err != nil {
		return nil, imp.wrapImportError(pkgpath, enableModule, err)
	}
	if len(alias) == 0 {
		alias = PackageName(gpkg.Name())
	}
	file := createImportFile(imp.output, pkgpath, gpkg, mode, enableModule)
	ref := &PackageRef{Path: pkgpath}
	if len(file) == 0 || mode != ImPlugin {
		// either the package exports nothing, or user must rebuild gomacro.
		// in both cases, still cache it to avoid recreating the file.
		imports.Packages[pkgpath] = ref.Package
		return ref, nil
	}
	soname := compilePlugin(o, file, enableModule, o.Stdout, o.Stderr)
	ipkgs := imp.loadPluginSymbol(soname, "Packages")
	pkgs := *ipkgs.(*map[string]imports.PackageUnderlying)

	// cache *all* packages found for future use
	imports.Packages.Merge(pkgs)

	// but return only requested one
	pkg, found := imports.Packages[pkgpath]
	if !found {
		return nil, imp.output.MakeRuntimeError(
			"error loading package %q: the compiled plugin %q does not contain it! internal error? %v",
			pkgpath, soname)
	}
	ref.Package = pkg
	return ref, nil
}

func createImportFile(o *Output, pkgpath string, pkg *types.Package, mode ImportMode, enableModule bool) string {
	dir := computeImportDir(o, pkgpath, mode)
	if mode == ImPlugin {
		createDir(o, dir)
		removeAllFilesInDirExcept(o, dir, []string{"go.mod", "go.sum"})
	}
	filepath := computeImportFilename(o, pkgpath, mode)
	filepath = paths.Subdir(dir, filepath)

	buf := bytes.Buffer{}
	isEmpty := writeImportFile(o, &buf, pkgpath, pkg, mode)
	if isEmpty {
		o.Warnf("package %q exports zero constants, functions, types and variables", pkgpath)
		return ""
	}

	err := ioutil.WriteFile(filepath, buf.Bytes(), os.FileMode(0644))
	if err != nil {
		o.Errorf("error writing file %q: %v", filepath, err)
	}
	switch mode {
	case ImBuiltin, ImThirdParty:
		o.Warnf("created file %q, recompile gomacro to use it", filepath)
	case ImInception:
		o.Warnf("created file %q, recompile %s to use it", filepath, pkgpath)
	case ImPlugin:
		// if needed, go.mod file was created already by Importer.Load()
		env := environForCompiler(enableModule)
		runGoModTidyIfNeeded(o, pkgpath, dir, env)
	}
	return filepath
}

func createDir(o *Output, dir string) {
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		o.Errorf("error creating directory %q: %v", dir, err)
	}
}

func listDir(o *Output, dir string) []os.FileInfo {
	d, err := os.Open(dir)
	if err != nil {
		o.Errorf("error opening directory %q: %v", dir, err)
	}
	defer d.Close()
	infos, err := d.Readdir(0)
	if err != nil {
		o.Errorf("error listing directory %q: %v", dir, err)
	}
	return infos
}

func removeAllFilesInDir(o *Output, dir string) {
	removeAllFilesInDirExcept(o, dir, nil)
}

func removeAllFilesInDirExcept(o *Output, dir string, except_list []string) {
	for _, info := range listDir(o, dir) {
		if info.IsDir() {
			continue
		}
		name := info.Name()
		for _, except_name := range except_list {
			if name == except_name {
				name = ""
				break
			}
		}
		if name == "" {
			continue
		}
		filepath := paths.Subdir(dir, name)
		if err := os.Remove(filepath); err != nil {
			o.Errorf("error removing file %q: %v", filepath, err)
		}
	}
}

func createPluginGoModFile(o *Output, pkgpath string, dir string) string {
	gomod := paths.Subdir(dir, "go.mod")
	err := ioutil.WriteFile(gomod, []byte("module gomacro.imports/"+pkgpath+"\n"), os.FileMode(0644))
	if err != nil {
		o.Errorf("error writing file %q: %v", gomod, err)
	}
	return gomod
}

func packageSanitizedName(pkgpath string) string {
	return sanitizeIdent(paths.FileName(pkgpath))
}

func sanitizeIdent(str string) string {
	return sanitizeIdent2(str, '_')
}

func sanitizeIdent2(str string, replacement rune) string {
	runes := []rune(str)
	for i, ch := range runes {
		if ch >= 'a' && ch <= 'z' {
			continue
		} else if ch >= 'A' && ch <= 'Z' {
			if i == 0 {
				// first rune must be lowercase to avoid conflict
				// with Packages, ValueOf, TypeOf
				runes[i] = ch - 'A' + 'a'
			}
			continue
		} else if i > 0 && (ch == '_' || (ch >= '0' && ch <= '9')) {
			continue
		}
		runes[i] = replacement
	}
	str = string(runes)
	if isReservedKeyword(str) {
		runes = append(runes, '_')
		str = string(runes)
	}
	return str
}

func computeImportDir(o *Output, pkgpath string, mode ImportMode) string {
	switch mode {
	case ImBuiltin:
		// user will need to recompile gomacro
		return paths.GetImportsSrcDir()
	case ImThirdParty:
		// either plugin.Open is not available, or user explicitly requested import _3 "package".
		// In both cases, user will need to recompile gomacro
		return paths.Subdir(paths.GetImportsSrcDir(), "thirdparty")
	case ImInception:
		// user will need to recompile the package being imported
		for _, srcdir := range paths.GoSrcDirs {
			dir := paths.Subdir(srcdir, pkgpath)
			if _, err := os.Stat(dir); err == nil {
				return paths.Subdir(srcdir, pkgpath)
			}
		}
		o.Errorf("unable to locate package %q in $GOPATH/src ($GOPATH=%s)",
			pkgpath, build.Default.GOPATH)
	case ImPlugin:
		return paths.Subdir(paths.GoSrcDir, "gomacro.imports", pkgpath)
	default:
		o.Errorf("unknown import mode: %v", mode)
	}
	return ""
}

func computeImportFilename(o *Output, pkgpath string, mode ImportMode) string {
	switch mode {
	case ImBuiltin:
		// user will need to recompile gomacro
		return sanitizeIdent(pkgpath) + ".go"
	case ImThirdParty:
		// either plugin.Open is not available, or user explicitly requested import _3 "package".
		// In both cases, user will need to recompile gomacro
		return sanitizeIdent(pkgpath) + ".go"
	case ImInception:
		// user will need to recompile package being imported
		return "x_package.go"
	case ImPlugin:
		return sanitizeIdent(paths.FileName(pkgpath)) + ".go"
	default:
		o.Errorf("unknown import mode: %v", mode)
		return ""
	}
}
