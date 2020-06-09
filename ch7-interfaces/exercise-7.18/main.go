package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// TODO: not sure the tree initialization is correct

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

type scanner struct {
	*xml.Decoder
	tok xml.Token
	err error
}

func (s *scanner) scan() bool {
	if s.err != nil {
		return false
	}
	s.tok, s.err = s.Token()
	return s.err == nil
}

func Parse(doc *os.File) (*Element, error) {
	d := xml.NewDecoder(doc)
	scanner := scanner{Decoder: d}
	root := new(Element)

	if scanner.err != nil {
		return nil, scanner.err
	}
	if err := root.parse(&scanner); err != nil {
		return nil, err
	}
	return root, nil
}

func (el *Element) parse(scanner *scanner) error {

	for scanner.scan() {
		switch tok := scanner.tok.(type) {
		case xml.StartElement:
			child := Element{tok.Name, tok.Attr, []Node{}}
			if err := child.parse(scanner); err != nil {
				return err
			}
			el.Children = append(el.Children, child)
		case xml.EndElement:
			break
		}
	}
	return scanner.err
}

func main() {
	root, _ := Parse(os.Stdin)
	fmt.Printf("root: %v\n", root)
}
