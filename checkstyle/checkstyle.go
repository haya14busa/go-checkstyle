// Package checkstyle provides checkstyle utility.
// http://checkstyle.sourceforge.net/
package checkstyle

import "encoding/xml"

// Result represents checkstyle XML result.
// <?xml version="1.0" encoding="utf-8"?><checkstyle version="4.3"><file ...></file>...</checkstyle>
//
// References:
//   - http://checkstyle.sourceforge.net/
//   - http://eslint.org/docs/user-guide/formatters/#checkstyle
type Result struct {
	XMLName xml.Name `xml:"checkstyle"`
	Version string   `xml:"version,attr"`
	Files   []*File  `xml:"file,omitempty"`
}

// File represents <file name="fname"><error ... />...</file>
type File struct {
	Name   string   `xml:"name,attr"`
	Errors []*Error `xml:"error"`
}

// Error represents <error line="1" column="10" severity="error" message="msg" source="src" />
type Error struct {
	Column   int    `xml:"column,attr,omitempty"`
	Line     int    `xml:"line,attr"`
	Message  string `xml:"message,attr"`
	Severity string `xml:"severity,attr,omitempty"`
	Source   string `xml:"source,attr,omitempty"`
}
