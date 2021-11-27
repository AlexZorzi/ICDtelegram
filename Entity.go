package main

import "strings"

type Entity struct {
	UrlID string `json:"@id"`
	Title struct {
		Language string `json:"@language"`
		Value    string `json:"@value"`
	} `json:"title"`
	Definition struct {
		Language string `json:"@language"`
		Value    string `json:"@value"`
	} `json:"definition"`

	Parent []string `json:"parent"`
	Child  []string `json:"child"`

	LongDefinition struct {
		Language string `json:"@language"`
		Value    string `json:"@value"`
	} `json:"longDefinition"`
}

func (entity Entity) ID() string {
	return GetIDfromUrl(entity.UrlID)
}
func (entity Entity) GetChildren() []Entity {
	var listchild []Entity
	for _, child := range entity.Child {
		splits := strings.Split(child, "/")
		ID := splits[len(splits)-1]
		childEntity, err := GetICDFoundationByID(ID)
		if err != nil {
			panic(err)
		}
		listchild = append(listchild, childEntity)
	}
	return listchild
}
func (entity Entity) GetParent() []Entity {
	var listparent []Entity
	for _, parent := range entity.Parent {
		splits := strings.Split(parent, "/")
		ID := splits[len(splits)-1]
		parentEntity, err := GetICDFoundationByID(ID)
		if err != nil {
			panic(err)
		}
		listparent = append(listparent, parentEntity)
	}
	return listparent
}
