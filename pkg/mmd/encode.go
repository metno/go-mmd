package mmd

import "encoding/xml"

// MetadataMMD is struct defining the MMD format.
// Use these for marshalling/encoding, as we need separate structs for decoding and encoding to handle namespaces properly
type MetadataMMD struct {
	XMLName                 xml.Name               `xml:"mmd:mmd"`
	Title                   *MMDTitle              `xml:"mmd:title"`
	Abstract                *MMDAbstract           `xml:"mmd:abstract,omitempty"`
	LastMetadataUpdate      string                 `xml:"mmd:last_metadata_update,omitempty"`
	Keywords                []*MMDKeywords         `xml:"mmd:keywords,omitempty"`
	MetadataIdentifier      string                 `xml:"mmd:metadata_identifier,omitempty"`
	Collections             []string               `xml:"mmd:collection,omitempty"`
	MetadataStatus          string                 `xml:"mmd:metadata_status,omitempty"`
	TemporalExtent          *TemporalExtent        `xml:"mmd:temporal_extent,omitempty"`
	GeographicExtent        *GeographicExtent      `xml:"mmd:geographic_extent,omitempty"`
	DatasetProductionStatus string                 `xml:"mmd:dataset_production_status,omitempty"`
	DatasetLanguage         string                 `xml:"mmd:dataset_language,omitempty"`
	OperationalStatus       string                 `xml:"mmd:operational_status,omitempty"`
	AccessConstraint        string                 `xml:"mmd:access_constraint,omitempty"`
	DataCenter              *MMDDataCenter         `xml:"mmd:data_center,omitempty"`
	DataAccess              []*MMDDataAccess       `xml:"mmd:data_access,omitempty"`
	RelatedDataset          *MMDRelatedDataset     `xml:"mmd:related_dataset,omitempty"`
	RelatedInformation      *MMDRelatedInformation `xml:"mmd:related_information,omitempty"`
	ISOTopicCategories      []string               `xml:"mmd:iso_topic_category,omitempty"`
	Project                 *MMDShortLongName      `xml:"mmd:project,omitempty"`
	Platform                *MMDShortLongName      `xml:"mmd:platform,omitempty"`
	Instrument              *MMDShortLongName      `xml:"mmd:instrument,omitempty"`
	ActivityType            []string               `xml:"mmd:activity_type,omitempty"`
	StorageInformation      *StorageInformation    `xml:"mmd:storage_information,omitempty"`
	DatasetCitation         []*DatasetCitation     `xml:"mmd:dataset_citation,omitempty"`
	UseConstraint           *UseConstraint         `xml:"mmd:use_constraint,omitempty"`
	SpatialRepresentation   string                 `xml:"mmd:spatial_representation,omitempty"`
}

type UseConstraint struct {
	Identifier string `xml:"mmd:identifier"`
	Resource   string `xml:"mmd:resource"`
}

type Personnel struct {
	Role         string `xml:"mmd:role"`
	Name         string `xml:"mmd:name"`
	Email        string `xml:"mmd:email"`
	Organisation string `xml:"mmd:organisation"`
}

type DatasetCitation struct {
	Author          string `xml:"mmd:auth"`
	PublicationDate string `xml:"mmd:publication_date"`
	Title           string `xml:"mmd:title"`
}

type StorageInformation struct {
	FileName     string  `xml:"mmd:file_name"`
	FileLocation string  `xml:"mmd:file_location"`
	FileFormat   string  `xml:"mmd:file_format"`
	FileSize     float64 `xml:"mmd:file_size"`
}

type MMDKeywords struct {
	Vocabulary string   `xml:"mmd:vocabulary,attr"`
	Keyword    []string `xml:"mmd:keyword"`
	Resource   string   `xml:"mmd:resource"`
}
type MMDTitle struct {
	Language string `xml:"xml:lang,attr"`
	Title    string `xml:",chardata"`
}
type MMDAbstract struct {
	Language string `xml:"xml:lang,attr"`
	Abstract string `xml:",chardata"`
}

type TemporalExtent struct {
	StartDate string `xml:"mmd:start_date"`
	EndDate   string `xml:"mmd:end_date"`
}

type GeographicExtent struct {
	Rectangle Rectangle `xml:"mmd:rectangle"`
}
type Rectangle struct {
	SrsName string  `xml:"mmd:srsName,attr"`
	North   float32 `xml:"mmd:north"`
	South   float32 `xml:"mmd:south"`
	West    float32 `xml:"mmd:west"`
	East    float32 `xml:"mmd:east"`
}

type MMDDataCenter struct {
	DataCenterName MMDShortLongName `xml:"mmd:data_center_name"`
	DataCenterURL  string           `xml:"mmd:data_center_url"`
}

type MMDDataAccess struct {
	Type        string `xml:"mmd:type"`
	Resource    string `xml:"mmd:resource"`
	Description string `xml:"mmd:description"`
}

type MMDRelatedDataset struct {
	RelatedDataset string `xml:",chardata"`
	RelationType   string `xml:"mmd:relation_type,attr"`
}

type MMDRelatedInformation struct {
	Type        string `xml:"mmd:type"`
	Description string `xml:"mmd:description"`
	Resource    string `xml:"mmd:resource"`
}

type MMDShortLongName struct {
	ShortName string `xml:"mmd:short_name"`
	LongName  string `xml:"mmd:long_name"`
}
