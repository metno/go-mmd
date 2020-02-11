package mmd

import "encoding/xml"

// MetadataMMD is struct defining the MMD format.
// Use for marshalling. For unmarshalling we need separate structs to handle namespaces properly.
type MetadataMMD struct {
	XMLName                 xml.Name              `xml:"mmd:mmd"`
	Title                   MMDTitle              `json:"title" xml:"mmd:title"`
	Abstract                MMDAbstract           `json:"abstract" xml:"mmd:abstract"`
	LastMetadataUpdate      string                `json:"last_metadata_update" xml:"mmd:last_metadata_update"`
	Keywords                MMDKeywords           `json:"keywords" xml:"mmd:keywords"`
	MetadataIdentifier      string                `json:"metadata_identifier" xml:"mmd:metadata_identifier"`
	Collections             []string              `xml:"mmd:collection"`
	MetadataStatus          string                `xml:"mmd:metadata_status"`
	TemporalExtent          TemporalExtent        `xml:"mmd:temporal_extent"`
	GeographicExtent        GeographicExtent      `xml:"mmd:geographic_extent"`
	DatasetProductionStatus string                `xml:"mmd:dataset_production_status"`
	DatasetLanguage         string                `xml:"mmd:dataset_language"`
	OperationalStatus       string                `xml:"mmd:operational_status"`
	AccessConstraint        string                `xml:"mmd:access_constraint"`
	DataCenter              MMDDataCenter         `xml:"mmd:data_center"`
	DataAccess              MMDDataAccess         `xml:"mmd:data_access"`
	RelatedDataset          MMDRelatedDataset     `xml:"mmd:related_dataset"`
	RelatedInformation      MMDRelatedInformation `xml:"mmd:related_information"`
	ISOTopicCategories      []string              `xml:"mmd:mmd>mmd:iso_topic_category"`
	Project                 MMDShortLongName      `xml:"mmd:project"`
	Platform                MMDShortLongName      `xml:"mmd:platform"`
	Instrument              MMDShortLongName      `xml:"mmd:instrument"`
	ActivityType            []string              `xml:"mmd:mmd>mmd:activity_type"`
}

type MMDKeywords struct {
	Vocabulary string   `xml:"vocabulary,attr"`
	Keyword    []string `xml:"mmd:keyword"`
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
	SrsName string  `xml:"srsName,attr"`
	North   float32 `xml:"mmd:north"`
	South   float32 `xml:"mmd:south"`
	West    float32 `xml:"mmd:west"`
	East    float32 `xml:"mmd:east"`
}

type MMDDataCenter struct {
	DataCenterName MMDShortLongName `xml:"mmd:data_center_name"`
	DataCenterURL  string           `xml:"mmd:data_center_url"`
	DatasetID      string           `xml:"mmd:dataset_id"`
}

type MMDDataAccess struct {
	Type        string   `xml:"mmd:type"`
	Resource    string   `xml:"mmd:resource"`
	Description string   `xml:"mmd:description"`
	WMSLayers   []string `xml:"mmd:mmd>mmd:wms_layers>mmd:wms_layer"`
}

type MMDRelatedDataset struct {
	RelatedDataset string `xml:",chardata"`
	RelationType   string `xml:"relation_type,attr"`
}

type MMDRelatedInformation struct {
	Type     string `xml:"type"`
	Resource string `xml:"resource"`
}

type MMDShortLongName struct {
	ShortName string `xml:"short_name"`
	LongName  string `xml:"long_name"`
}
