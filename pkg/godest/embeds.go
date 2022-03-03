package godest

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"html/template"
	"io/fs"
	"path"

	"github.com/benbjohnson/hashfs"
	"github.com/pkg/errors"

	"github.com/sargassum-world/fluitans/pkg/godest/fingerprint"
	"github.com/sargassum-world/fluitans/pkg/godest/fsutil"
	tp "github.com/sargassum-world/fluitans/pkg/godest/template"
)

func ComputeCSPHash(resource []byte) string {
	rawHash := sha512.Sum512(resource)
	encodedHash := base64.StdEncoding.EncodeToString(rawHash[:])
	return fmt.Sprintf("'sha512-%s'", encodedHash)
}

func identifyModuleNonpageFiles(templates fs.FS) (map[string][]string, error) {
	modules, err := fsutil.ListDirectories(templates, tp.FilterModule)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't list template modules")
	}

	moduleFiles := make(map[string][]string)
	for _, module := range modules {
		var subfs fs.FS
		if module == "" {
			subfs = templates
		} else {
			subfs, err = fs.Sub(templates, module)
		}
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("couldn't list template module %s", module))
		}
		moduleSubfiles, err := fsutil.ListFiles(subfs, tp.FilterNonpage)
		moduleFiles[module] = make([]string, len(moduleSubfiles))
		for i, subfile := range moduleSubfiles {
			if module == "" {
				moduleFiles[module][i] = subfile
			} else {
				moduleFiles[module][i] = fmt.Sprintf("%s/%s", module, subfile)
			}
		}
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("couldn't list template files in module %s", module))
		}
	}
	return moduleFiles, nil
}

// Etag pre-computation

func computeAppFingerprint(
	appAssets, sharedTemplates []string, templates, app fs.FS,
) (string, error) {
	appConcatenated, err := fsutil.ReadConcatenated(appAssets, app)
	if err != nil {
		return "", errors.Wrap(err, "couldn't load all app assets together for fingerprinting")
	}
	sharedConcatenated, err := fsutil.ReadConcatenated(sharedTemplates, templates)
	if err != nil {
		return "", errors.Wrap(err, "couldn't load all shared templates together for fingerprinting")
	}

	return fingerprint.Compute(append(appConcatenated, sharedConcatenated...)), nil
}

func computePageFingerprints(
	moduleNonpageFiles map[string][]string, pageFiles []string, templates fs.FS,
) (map[string]string, error) {
	moduleNonpages := make(map[string][]byte)
	for module, files := range moduleNonpageFiles {
		loadedNonpages, err := fsutil.ReadConcatenated(files, templates)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf(
				"couldn't load non-page template files in template module %s for fingerprinting", module,
			))
		}
		moduleNonpages[module] = loadedNonpages
	}

	pages := make(map[string][]byte)
	for _, file := range pageFiles {
		loadedPage, err := fsutil.ReadFile(file, templates)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf(
				"couldn't load page template %s for fingerprinting", file,
			))
		}
		pages[file] = loadedPage
	}

	pageFingerprints := make(map[string]string)
	for _, pageFile := range pageFiles {
		module := path.Dir(pageFile)
		pageFingerprints[pageFile] = fingerprint.Compute(append(
			// Each page's fingerprint is computed from the page template itself as well as any non-page
			// files (e.g. partials) within its module, recursively including all non-page files in
			// submodules (i.e. subdirectories)
			pages[pageFile], moduleNonpages[module]...,
		))
	}
	return pageFingerprints, nil
}

// Embedded filesystem template support

type Embeds struct {
	StaticFS    fs.FS
	StaticHFS   *hashfs.FS
	TemplatesFS fs.FS
	AppFS       fs.FS
	AppHFS      *hashfs.FS
	FontsFS     fs.FS
}

func (e Embeds) ComputeAppFingerprint() (fingerprint string, err error) {
	appAssetFiles, err := fsutil.ListFiles(e.AppFS, tp.FilterAsset)
	if err != nil {
		err = errors.Wrap(err, "couldn't list app assets")
		return
	}
	sharedFiles, err := fsutil.ListFiles(e.TemplatesFS, tp.FilterShared)
	if err != nil {
		err = errors.Wrap(err, "couldn't list shared templates")
		return
	}
	fingerprint, err = computeAppFingerprint(appAssetFiles, sharedFiles, e.TemplatesFS, e.AppFS)
	if err != nil {
		err = errors.Wrap(err, "couldn't compute fingerprint for app")
		return
	}
	return
}

func (e Embeds) ComputePageFingerprints() (fingerprints map[string]string, err error) {
	moduleNonpageFiles, err := identifyModuleNonpageFiles(e.TemplatesFS)
	if err != nil {
		err = errors.Wrap(err, "couldn't list template module non-page template files")
		return
	}
	pageFiles, err := fsutil.ListFiles(e.TemplatesFS, tp.FilterPage)
	if err != nil {
		err = errors.Wrap(err, "couldn't list template pages")
		return
	}

	fingerprints, err = computePageFingerprints(moduleNonpageFiles, pageFiles, e.TemplatesFS)
	if err != nil {
		err = errors.Wrap(err, "couldn't compute fingerprint for page/module templates")
		return
	}
	return
}

func (e Embeds) NewTemplates(funcs ...template.FuncMap) (r tp.Templates, err error) {
	pageFiles, err := fsutil.ListFiles(e.TemplatesFS, tp.FilterPage)
	if err != nil {
		err = errors.Wrap(err, "couldn't list template pages")
	}
	r = tp.NewTemplates(e.TemplatesFS, pageFiles, funcs...)
	return
}

func (e Embeds) GetAppHashedNamer(urlPrefix string) func(string) string {
	return func(unhashedFilename string) string {
		return fmt.Sprintf(urlPrefix + e.AppHFS.HashName(unhashedFilename))
	}
}

func (e Embeds) GetStaticHashedNamer(urlPrefix string) func(string) string {
	return func(unhashedFilename string) string {
		return fmt.Sprintf(urlPrefix + e.StaticHFS.HashName(unhashedFilename))
	}
}

// Inline snippet support

type Inlines struct {
	CSS      map[string]template.CSS
	JS       map[string]template.JS
	JSStr    map[string]template.JSStr
	HTML     map[string]template.HTML
	HTMLAttr map[string]template.HTMLAttr
	SrcSet   map[string]template.Srcset
}

func (i Inlines) ComputeCSSHashesForCSP() (hashes []string) {
	hashes = make([]string, 0, len(i.CSS))
	for _, inline := range i.CSS {
		hashes = append(hashes, ComputeCSPHash([]byte(inline)))
	}
	return
}

func (i Inlines) ComputeJSHashesForCSP() (hashes []string) {
	hashes = make([]string, 0, len(i.CSS))
	for _, inline := range i.JS {
		hashes = append(hashes, ComputeCSPHash([]byte(inline)))
	}
	return
}