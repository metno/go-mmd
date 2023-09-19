package mmd

import "encoding/xml"

// DecodeMetadataMMD is struct defining the MMD format.
// Use these for unmarshalling/decoding only, as we need separate structs for decoding and encoding to handle namespaces properly.
type DecodeMetadataMMD struct {
	XMLName                 xml.Name                    `xml:"mmd"`
	Title                   DecodeMMDTitle              `xml:"title"`
	Abstract                DecodeMMDAbstract           `xml:"abstract"`
	LastMetadataUpdate      string                      `xml:"last_metadata_update"`
	Keywords                []DecodeMMDKeywords         `xml:"keywords"`
	MetadataIdentifier      string                      `xml:"metadata_identifier"`
	Collections             []string                    `xml:"collection"`
	MetadataStatus          string                      `xml:"metadata_status"`
	TemporalExtent          DecodeTemporalExtent        `xml:"temporal_extent"`
	GeographicExtent        DecodeGeographicExtent      `xml:"geographic_extent"`
	DatasetProductionStatus string                      `xml:"dataset_production_status"`
	DatasetLanguage         string                      `xml:"dataset_language"`
	OperationalStatus       string                      `xml:"operational_status"`
	AccessConstraint        string                      `xml:"access_constraint"`
	DataCenter              DecodeMMDDataCenter         `xml:"data_center"`
	DataAccess              []DecodeMMDDataAccess       `xml:"data_access"`
	RelatedDataset          DecodeMMDRelatedDataset     `xml:"related_dataset"`
	RelatedInformation      DecodeMMDRelatedInformation `xml:"related_information"`
	ISOTopicCategories      []string                    `xml:"iso_topic_category"`
	Project                 DecodeMMDShortLongName      `xml:"project"`
	Platform                DecodeMMDShortLongName      `xml:"platform"`
	Instrument              DecodeMMDShortLongName      `xml:"instrument"`
	ActivityType            []string                    `xml:"activity_type"`
	StorageInformation      DecodeStorageInformation    `xml:"storage_information"`
	DatasetCitation         []DecodeDatasetCitation     `xml:"dataset_citation"`
	UseConstraint           DecodeUseConstraint         `xml:"use_constraint"`
	SpatialRepresentation   string                      `xml:"spatial_representation"`
}

type DecodeUseConstraint struct {
	Identifier string `xml:"identifier"`
	Resource   string `xml:"resource"`
}

type DecodePersonnel struct {
	Role         string `xml:"role"`
	Name         string `xml:"name"`
	Email        string `xml:"email"`
	Organisation string `xml:"organisation"`
}

type DecodeDatasetCitation struct {
	Author          string `xml:"auth"`
	PublicationDate string `xml:"publication_date"`
	Title           string `xml:"title"`
}

type DecodeStorageInformation struct {
	FileName     string  `xml:"file_name"`
	FileLocation string  `xml:"file_location"`
	FileFormat   string  `xml:"file_format"`
	FileSize     float64 `xml:"file_size"`
}

type DecodeMMDKeywords struct {
	Vocabulary string   `xml:"vocabulary,attr"`
	Keyword    []string `xml:"keyword"`
	Resource   string   `xml:"resource"`
}
type DecodeMMDTitle struct {
	Language string `xml:"xml:lang,attr"`
	Title    string `xml:",chardata"`
}
type DecodeMMDAbstract struct {
	Language string `xml:"xml:lang,attr"`
	Abstract string `xml:",chardata"`
}

type DecodeTemporalExtent struct {
	StartDate string `xml:"start_date"`
	EndDate   string `xml:"end_date"`
}

type DecodeGeographicExtent struct {
	Rectangle DecodeRectangle `xml:"rectangle"`
}
type DecodeRectangle struct {
	SrsName string  `xml:"srsName,attr"`
	North   float32 `xml:"north"`
	South   float32 `xml:"south"`
	West    float32 `xml:"west"`
	East    float32 `xml:"east"`
}

type DecodeMMDDataCenter struct {
	DataCenterName DecodeMMDShortLongName `xml:"data_center_name"`
	DataCenterURL  string                 `xml:"data_center_url"`
	DatasetID      string                 `xml:"dataset_id"`
}

type DecodeMMDDataAccess struct {
	Type        string `xml:"type"`
	Resource    string `xml:"resource"`
	Description string `xml:"description"`
}

type DecodeMMDRelatedDataset struct {
	RelatedDataset string `xml:",chardata"`
	RelationType   string `xml:"relation_type,attr"`
}

type DecodeMMDRelatedInformation struct {
	Type        string `xml:"type"`
	Description string `xml:"description"`
	Resource    string `xml:"resource"`
}

type DecodeMMDShortLongName struct {
	ShortName string `xml:"short_name"`
	LongName  string `xml:"long_name"`
}
