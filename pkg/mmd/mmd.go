package mmd

import "encoding/xml"

// MetadataMMD is struct defining the MMD format.
// Use for marshalling. For unmarshalling we need separate structs to handle namespaces properly.
type MetadataMMD struct {
	XMLName                 xml.Name              `xml:"mmd"`
	Title                   MMDTitle              `xml:"title"`
	Abstract                MMDAbstract           `xml:"abstract"`
	LastMetadataUpdate      string                `xml:"last_metadata_update"`
	Keywords                []MMDKeywords         `xml:"keywords"`
	MetadataIdentifier      string                `xml:"metadata_identifier"`
	Collections             []string              `xml:"collection"`
	MetadataStatus          string                `xml:"metadata_status"`
	TemporalExtent          TemporalExtent        `xml:"temporal_extent"`
	GeographicExtent        GeographicExtent      `xml:"geographic_extent"`
	DatasetProductionStatus string                `xml:"dataset_production_status"`
	DatasetLanguage         string                `xml:"dataset_language"`
	OperationalStatus       string                `xml:"operational_status"`
	AccessConstraint        string                `xml:"access_constraint"`
	DataCenter              MMDDataCenter         `xml:"data_center"`
	DataAccess              []MMDDataAccess       `xml:"data_access"`
	RelatedDataset          MMDRelatedDataset     `xml:"related_dataset"`
	RelatedInformation      MMDRelatedInformation `xml:"related_information"`
	ISOTopicCategories      []string              `xml:"iso_topic_category"`
	Project                 MMDShortLongName      `xml:"project"`
	Platform                MMDShortLongName      `xml:"platform"`
	Instrument              MMDShortLongName      `xml:"instrument"`
	ActivityType            []string              `xml:"activity_type"`
	StorageInformation      StorageInformation    `xml:"storage_information"`
	DatasetCitation         []DatasetCitation     `xml:"dataset_citation"`
	UseConstraint           UseConstraint         `xml:"use_constraint"`
	SpatialRepresentation   string                `xml:"spatial_representation"`
}

type UseConstraint struct {
	Identifier string `xml:"identifier"`
	Resource   string `xml:"resource"`
}

type Personnel struct {
	Role         string `xml:"role"`
	Name         string `xml:"name"`
	Email        string `xml:"email"`
	Organisation string `xml:"organisation"`
}

type DatasetCitation struct {
	Author          string `xml:"auth"`
	PublicationDate string `xml:"publication_date"`
	Title           string `xml:"title"`
}

type StorageInformation struct {
	FileName     string  `xml:"file_name"`
	FileLocation string  `xml:"file_location"`
	FileFormat   string  `xml:"file_format"`
	FileSize     float64 `xml:"file_size"`
}

type MMDKeywords struct {
	Vocabulary string   `xml:"vocabulary,attr"`
	Keyword    []string `xml:"keyword"`
	Resource   string   `xml:"resource"`
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
	StartDate string `xml:"start_date"`
	EndDate   string `xml:"end_date"`
}

type GeographicExtent struct {
	Rectangle Rectangle `xml:"rectangle"`
}
type Rectangle struct {
	SrsName string  `xml:"srsName,attr"`
	North   float32 `xml:"north"`
	South   float32 `xml:"south"`
	West    float32 `xml:"west"`
	East    float32 `xml:"east"`
}

type MMDDataCenter struct {
	DataCenterName MMDShortLongName `xml:"data_center_name"`
	DataCenterURL  string           `xml:"data_center_url"`
	DatasetID      string           `xml:"dataset_id"`
}

type MMDDataAccess struct {
	Type        string `xml:"type"`
	Resource    string `xml:"resource"`
	Description string `xml:"description"`
}

type MMDRelatedDataset struct {
	RelatedDataset string `xml:",chardata"`
	RelationType   string `xml:"relation_type,attr"`
}

type MMDRelatedInformation struct {
	Type        string `xml:"type"`
	Description string `xml:"description"`
	Resource    string `xml:"resource"`
}

type MMDShortLongName struct {
	ShortName string `xml:"short_name"`
	LongName  string `xml:"long_name"`
}
